package listeners

import (
	"fmt"
	"log/slog"

	"github.com/antlr4-go/antlr/v4"
)

type SlogErrorListener struct {
	antlr.ErrorListener
	logger *slog.Logger
	errors []error
}

func NewSlogErrorListener(logger *slog.Logger) *SlogErrorListener {
	return &SlogErrorListener{
		antlr.NewDefaultErrorListener(),
		logger,
		[]error{},
	}
}

func (c *SlogErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	c.logger.Error(msg, "line", line, "col", column)
	c.errors = append(c.errors, fmt.Errorf(msg))
}

func (c *SlogErrorListener) HasError() bool {
	return len(c.errors) > 0
}

func (c *SlogErrorListener) GetError() error {
	return c.errors[0]
}
