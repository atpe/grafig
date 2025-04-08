package parser

import (
	"figsyntax/internal/debugger"
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
	tree           antlr.ParseTree
	errorListeners []ErrorListener
}

func NewCommonParser(logger *slog.Logger, target string, options ...CommonParserOption) (*CommonParser, error) {
	debugger.Trace()

	p := &CommonParser{
		logger:         logger,
		target:         target,
		errorListeners: make([]ErrorListener, 0),
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

func (p *CommonParser) Walk(listener antlr.ParseTreeListener) {
	debugger.Trace()

	antlr.ParseTreeWalkerDefault.Walk(listener, p.tree)
}

func (p *CommonParser) parse() error {
	debugger.Trace()

	switch p.target {
	case "javascript":
		p.parseJavaScript()
	default:
		return fmt.Errorf("cannot parse for invalid target '%v'", p.target)
	}
	return nil
}

func (p *CommonParser) parseJavaScript() {
	debugger.Trace()

	parser := NewJavaScriptParser(p.lexer.GetTokens())
	p.bind(parser)
	p.tree = parser.Program()
}

func (p *CommonParser) apply(options []CommonParserOption) error {
	debugger.Trace()

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
	debugger.Trace()

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
