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
}

func NewInputLexer(logger *slog.Logger, data string, target string) (*FileLexer, error) {
	lexer := &FileLexer{logger: logger, target: target}

	input := antlr.NewInputStream(data)
	tokens, err := tokeniseInput(target, input)
	if err != nil {
		return nil, err
	}

	lexer.tokens = tokens

	return lexer, nil
}

func NewFileLexer(logger *slog.Logger, path string, target string) (*FileLexer, error) {
	lexer := &FileLexer{logger: logger, target: target}

	input, err := antlr.NewFileStream(path)
	if err != nil {
		return nil, err
	}

	tokens, err := tokeniseInput(target, input)
	if err != nil {
		return nil, err
	}

	lexer.tokens = tokens

	return lexer, nil
}

func (l *FileLexer) GetTokens() *antlr.CommonTokenStream {
	logger.Trace()

	return l.tokens
}

func tokeniseInput(target string, input antlr.CharStream) (*antlr.CommonTokenStream, error) {
	logger.Trace()

	switch target {
	case "javascript":
		lexer := NewJavaScriptLexer(input)
		return antlr.NewCommonTokenStream(lexer, 0), nil
	case "figscript":
		lexer := figscript.NewFigScriptLexer(input)
		return antlr.NewCommonTokenStream(lexer, 0), nil
	default:
		return nil, fmt.Errorf("could not tokenise file - target '%v' is not a valid option", target)
	}
}
