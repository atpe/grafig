package parser

import (
	"figsyntax/internal/logger"
	"figsyntax/internal/parser/figscript"
	"fmt"
	"log/slog"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorListener interface {
	antlr.ErrorListener
	HasError() bool
	GetError() error
}

type Parser interface {
	GetError() error
	Walk()
}

type CommonParserOption func(*CommonParser) error

type CommonParser struct {
	logger         *slog.Logger
	target         string
	lexer          Lexer
	tokens         *antlr.CommonTokenStream
	tree           antlr.ParseTree
	hidden         map[int]antlr.Token
	errorListeners []ErrorListener
}

func New(logger *slog.Logger, target string, options ...CommonParserOption) (*CommonParser, error) {
	p := &CommonParser{
		logger:         logger,
		target:         target,
		errorListeners: make([]ErrorListener, 0),
		hidden:         make(map[int]antlr.Token),
	}

	err := p.apply(options)
	if err != nil {
		return nil, err
	}
	err = p.parse()
	if err != nil {
		return nil, err
	}
	err = p.validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func WithFileLexer(path string) CommonParserOption {
	return func(p *CommonParser) error {
		lexer, err := NewFileLexer(p.logger.WithGroup("lexer"), path, p.target)
		if err != nil {
			return err
		}
		p.lexer = lexer
		return nil
	}
}

func WithErrorListener(listener ErrorListener) CommonParserOption {
	return func(p *CommonParser) error {
		p.errorListeners = append(p.errorListeners, listener)
		return nil
	}
}

func (p *CommonParser) GetTree() antlr.Tree {
	logger.Trace()
	return p.tree
}

func (p *CommonParser) GetHidden() map[int]antlr.Token {
	logger.Trace()
	return p.hidden
}

func (p *CommonParser) Walk(listener antlr.ParseTreeListener) {
	logger.Trace()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.tree)
}

func (p *CommonParser) parse() error {
	logger.Trace()

	p.tokens = p.lexer.GetTokens()

	switch p.target {
	case "javascript":
		p.parseJavaScript()
	case "figscript":
		p.parseFigScript()
	default:
		return fmt.Errorf("cannot parse for invalid target '%v'", p.target)
	}

	for _, token := range p.tokens.GetAllTokens() {
		if token.GetChannel() == antlr.TokenHiddenChannel {
			p.hidden[token.GetTokenIndex()] = token
		}
	}

	return nil
}

func (p *CommonParser) parseJavaScript() {
	logger.Trace()

	parser := NewJavaScriptParser(p.tokens)
	p.bind(parser)
	p.tree = parser.Program()
}

func (p *CommonParser) parseFigScript() {
	logger.Trace()

	parser := figscript.NewFigScriptParser(p.tokens)
	p.bind(parser)
	p.tree = parser.Program()
}

func (p *CommonParser) apply(options []CommonParserOption) error {
	logger.Trace()

	for _, o := range options {
		err := o(p)
		if err != nil {
			return err
		}
	}

	if p.lexer == nil {
		return fmt.Errorf("cannot create a p without a 'With*Lexer' option")
	}

	return nil
}

func (p *CommonParser) bind(parser antlr.Parser) {
	logger.Trace()
	for _, listener := range p.errorListeners {
		parser.AddErrorListener(listener)
	}
}

func (p *CommonParser) validate() error {
	for _, listener := range p.errorListeners {
		if listener.HasError() {
			return listener.GetError()
		}
	}
	return nil
}
