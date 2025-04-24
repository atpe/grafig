package cmd

import (
	"figsyntax/internal/command"
	"figsyntax/internal/file"
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"log/slog"

	"github.com/spf13/cobra"
)

const parseUse = "parse <path> [flags]"
const parseShort = `Parses a given figscript or javascript file.`
const parseLong = `
Given a file containing valid javascript or figscript, the parse command will
output its resultant parse tree.
`

type parseCommand struct {
	*command.CommonCommand
}

func newParseCommand(logger *slog.Logger, options ...command.CommonCommandOption) *parseCommand {
	c := &parseCommand{command.New(logger.WithGroup(parseUse), options...)}

	c.CommonCommand.PreRun = c.PreRun
	c.CommonCommand.Run = c.Run

	return c
}

func (c *parseCommand) PreRun(cmd *cobra.Command, args []string) {
	logger.Trace()

	c.ConfigureLoggerFlags()
	c.ConfigureTargetFlags(args[0])

	c.LogWith(
		slog.String(file.Path, args[0]),
		c.UseConfigAttr(file.Output),
		c.UseConfigAttr(logger.Level),
	)
}

func (c *parseCommand) Run(cmd *cobra.Command, args []string) {
	logger.Trace()

	errorListener := c.UseValidationErrorListener()

	_, err := parser.New(
		c.GetLogger(),
		c.GetConfigString(parser.Target),
		parser.WithFileLexer(args[0]),
		parser.WithErrorListener(errorListener),
	)

	c.HandleError(err)
	c.HandleError(errorListener.GetError())
}
