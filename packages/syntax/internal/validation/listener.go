package validation

import (
	"figsyntax/internal/debugger"
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
	debugger.Trace()

	return &ValidationErrorListener{
		antlr.NewDefaultErrorListener(),
		logger,
		make([]error, 0),
	}
}

func (c *ValidationErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	debugger.Trace()

	c.logger.Error(msg, "line", line, "col", column)
	c.errors = append(c.errors, fmt.Errorf(msg))
}

func (c *ValidationErrorListener) HasError() bool {
	debugger.Trace()

	return len(c.errors) > 0
}

func (c *ValidationErrorListener) GetError() error {
	debugger.Trace()

	return c.errors[0]
}

func (c *ValidationErrorListener) Reset() {
	debugger.Trace()

	c.errors = make([]error, 0)
}
