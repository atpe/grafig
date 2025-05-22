package cmd

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"figsyntax/internal/analyser"
	"figsyntax/internal/command"
	"figsyntax/internal/file"
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const analyseUse = "analyse <path> [flags]"
const analyseShort = `Analyses the complexity of a given figscript or javascript file.`
const analyseLong = `
Given a file containing valid javascript or figscript, the analyse command will
parse the code and evalutate the following metrics:

	- Source Lines of Code (LOC/SLOC)
	- Logical Lines of Code (LLOC)
	- Parameters
	- Cyclomatic Complexity
	- Cyclomatic Density
	- Halstead Vocabulary
	- Halstead Length
	- Halstead Volume
	- Halstead Difficulty
	- Halstead Effort
	- Halstead Time
	- Halstead Bugs

The output is generated in a similar format to the npm package 'complexity-report'.
	See: https://www.npmjs.com/package/complexity-report
`

type analyseCommand struct {
	*command.CommonCommand
}

func newAnalyseCommand(logger *slog.Logger, options ...command.CommonCommandOption) *analyseCommand {
	c := &analyseCommand{command.New(logger.WithGroup(analyseUse), options...)}

	c.CommonCommand.PreRun = c.PreRun
	c.CommonCommand.Run = c.Run

	return c
}

func (c *analyseCommand) PreRun(cmd *cobra.Command, args []string) {
	logger.Trace()

	c.ConfigureLoggerFlags()
	c.ConfigureTargetFlags(args[0])

	c.LogWith(
		slog.String(file.Path, args[0]),
		c.UseConfigAttr(file.Output),
		c.UseConfigAttr(parser.Target),
		c.UseConfigAttr(logger.Level),
	)
}

func (c *analyseCommand) Run(cmd *cobra.Command, args []string) {
	logger.Trace()

	target := c.GetConfigString(parser.Target)
	switch target {
	case parser.JavaScript:
		break
	case parser.FigScript:
		break
	}

	a := analyser.New(c.GetLogger(), c.GetConfigString(parser.Target))
	analysed, err := a.AnalyseGlob(args[0])

	c.HandleError(err)

	if !c.GetConfigBool(logger.Silent) {
		for _, metrics := range analysed {
			metrics.Print()
		}
	}

	path := c.GetConfigString(file.Output)
	if path == command.Undefined {
		return
	}

	file, err := os.Create(path)
	c.HandleError(err)
	defer c.DeferErrorFunc(file.Close, err == nil)

	format, err := c.useFormat(path)
	c.HandleError(err)

	switch format {
	case analyser.JSON:
		data := analysed

		if c.GetConfigBool(analyser.Flat) {
			data = []analyser.Report{}
			for _, report := range analysed {
				data = append(data, report.Flat()...)
			}
		}

		bytes, err := json.MarshalIndent(data, "", "  ")
		c.HandleError(err)

		writer := bufio.NewWriter(file)
		_, err = writer.Write(bytes)
		c.HandleError(err)
		c.DeferErrorFunc(writer.Flush, err == nil)
	case analyser.CSV:
		data := [][]string{analyser.Headers}

		for _, report := range analysed {
			data = append(data, report.ToCsv()...)
		}

		writer := csv.NewWriter(file)
		err = writer.WriteAll(data)
		c.HandleError(err)
	}
}

func (c *analyseCommand) useFormat(path string) (string, error) {
	switch c.GetConfigString(analyser.Format) {
	case analyser.JSON:
		return analyser.JSON, nil
	case analyser.CSV:
		return analyser.CSV, nil
	default:
		break
	}

	switch {
	case strings.HasSuffix(path, file.JSON):
		return analyser.JSON, nil
	case strings.HasSuffix(path, file.CSV):
		return analyser.CSV, nil
	case strings.Contains(path, analyser.JSON):
		return analyser.JSON, nil
	case strings.Contains(path, analyser.CSV):
		return analyser.CSV, nil
	}

	return "", fmt.Errorf("cannot infer format from config or path")
}
