package cmd

import (
	"figsyntax/internal/debugger"
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"figsyntax/internal/validation"
	"log/slog"

	"github.com/spf13/cobra"
)

type ParseCommand struct {
	*CommonCommand
}

func NewParseCommand(logger *slog.Logger, options ...CommonCommandOption) *ParseCommand {
	debugger.Trace()

	c := &ParseCommand{
		NewCommonCommand(logger.WithGroup("parse")),
	}
	for _, o := range options {
		o(c.CommonCommand)
	}
	c.CommonCommand.PreRun = c.PreRun
	c.CommonCommand.Run = c.Run
	return c
}

func (c *ParseCommand) PreRun(cmd *cobra.Command, args []string) {
	debugger.Trace()

	c.ConfigureLoggerFlags()
	c.ConfigureTargetFlags(args[0])

	c.logger = c.logger.With(
		slog.String("path", args[0]),
		slog.String("target", c.config.GetString("target")),
		slog.String("log-level", c.config.GetString("log-level")),
	)
}

func (c *ParseCommand) Run(cmd *cobra.Command, args []string) {
	debugger.Trace()

	errorListener := validation.NewValidationErrorListener(c.logger.WithGroup("errors"))
	debugListener := logger.NewDebugListener(c.logger.WithGroup("debug"))

	parser, err := parser.NewCommonParser(
		c.logger.WithGroup("parser"),
		c.config.GetString("target"),
		parser.WithFileLexer(args[0]),
		parser.WithErrorListener(errorListener),
	)
	if err != nil {
		c.logger.Error(err.Error())
		return
	}

	if errorListener.HasError() {
		c.logger.Error(errorListener.GetError().Error())
		return
	}

	parser.Walk(debugListener)
}
