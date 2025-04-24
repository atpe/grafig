package parser

import (
	"figsyntax/internal/logger"
	"figsyntax/internal/parser/figscript"
	"fmt"
	"log/slog"

	"github.com/antlr4-go/antlr/v4"
)

type Token interface {
	GetText() string
}

type Lexer interface {
	GetTokens() *antlr.CommonTokenStream
}

type FileLexer struct {
	logger *slog.Logger
	target string
	tokens *antlr.CommonTokenStream
	input  antlr.CharStream
}

func NewFileLexer(logger *slog.Logger, path string, target string) (*FileLexer, error) {
	lexer := &FileLexer{logger: logger, target: target}
	err := lexer.readFileStream(path)
	if err != nil {
		return nil, err
	}

	err = lexer.tokeniseInput(target)
	if err != nil {
		return nil, err
	}

	return lexer, nil
}

func (l *FileLexer) GetTokens() *antlr.CommonTokenStream {
	logger.Trace()

	return l.tokens
}

func (l *FileLexer) tokeniseInput(target string) error {
	logger.Trace()

	switch target {
	case "javascript":
		lexer := NewJavaScriptLexer(l.input)
		l.tokens = antlr.NewCommonTokenStream(lexer, 0)
	case "figscript":
		lexer := figscript.NewFigScriptLexer(l.input)
		l.tokens = antlr.NewCommonTokenStream(lexer, 0)
	default:
		return fmt.Errorf("could not tokenise file - target '%v' is not a valid option", target)
	}

	return nil
}

func (l *FileLexer) readFileStream(path string) error {
	logger.Trace()

	input, err := antlr.NewFileStream(path)
	if err != nil {
		return err
	}

	l.input = input

	return nil
}
