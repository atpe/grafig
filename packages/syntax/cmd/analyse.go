package cmd

import (
	"bufio"
	"encoding/json"
	"figsyntax/internal/analyser"
	"figsyntax/internal/debugger"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

type AnalyseCommand struct {
	*CommonCommand
}

func NewAnalyseCommand(logger *slog.Logger, options ...CommonCommandOption) *AnalyseCommand {
	debugger.Trace()

	c := &AnalyseCommand{
		CommonCommand: NewCommonCommand(logger.WithGroup("analyse")),
	}
	for _, o := range options {
		o(c.CommonCommand)
	}
	c.CommonCommand.PreRun = c.PreRun
	c.CommonCommand.Run = c.Run
	return c
}

func (c *AnalyseCommand) PreRun(cmd *cobra.Command, args []string) {
	debugger.Trace()

	c.ConfigureLoggerFlags()
	c.ConfigureTargetFlags(args[0])

	c.logger = c.logger.With(
		slog.String("path", args[0]),
		slog.String("output", c.config.GetString(OUTPUT)),
		slog.String("target", c.config.GetString(TARGET)),
		slog.String("log-level", c.config.GetString(LOG_LEVEL)),
	)
}

func (c *AnalyseCommand) Run(cmd *cobra.Command, args []string) {
	debugger.Trace()

	analyser := analyser.NewAnalyser(c.logger)
	analysed, err := analyser.AnalyseGlob(args[0], c.config.GetString(TARGET))
	if err != nil {
		c.logger.Error(err.Error())
		return
	}

	if !c.config.GetBool(SILENT) {
		for _, metrics := range analysed {
			metrics.Print()
		}
	}

	path := c.config.GetString(OUTPUT)
	if path != DEFAULT {
		file, err := os.Create(path)
		if err != nil {
			c.logger.Error(err.Error())
			return
		}
		defer func() {
			if err := file.Close(); err != nil {
				c.logger.Error(err.Error())
				return
			}
		}()

		bytes, err := json.MarshalIndent(analysed, "", "  ")
		if err != nil {
			c.logger.Error(err.Error())
			return
		}

		writer := bufio.NewWriter(file)

		_, err = writer.Write(bytes)
		if err != nil {
			c.logger.Error(err.Error())
			return
		}

		if err = writer.Flush(); err != nil {
			panic(err)
		}
	}
}
