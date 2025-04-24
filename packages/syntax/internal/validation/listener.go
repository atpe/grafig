package validation

import (
	"figsyntax/internal/logger"
	"fmt"
	"log/slog"

	"github.com/antlr4-go/antlr/v4"
)

type ValidationErrorListener struct {
	antlr.ErrorListener
	logger *slog.Logger
	errors []error
}

func NewValidationErrorListener(logger *slog.Logger) *ValidationErrorListener {
	return &ValidationErrorListener{
		antlr.NewDefaultErrorListener(),
		logger.WithGroup("errors"),
		make([]error, 0),
	}
}

func (c *ValidationErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	logger.Trace()
	c.logger.Error(msg, "line", line, "col", column)
	c.errors = append(c.errors, fmt.Errorf(msg))
}

func (c *ValidationErrorListener) HasError() bool {
	logger.Trace()
	return len(c.errors) > 0
}

func (c *ValidationErrorListener) GetError() error {
	logger.Trace()
	if len(c.errors) > 0 {
		return c.errors[0]
	}
	return nil
}

func (c *ValidationErrorListener) Reset() {
	logger.Trace()
	c.errors = make([]error, 0)
}
