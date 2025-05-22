package cmd

import (
	"bufio"
	"figsyntax/internal/command"
	"figsyntax/internal/file"
	"figsyntax/internal/logger"
	"figsyntax/internal/transpiler"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const logFormat = "\n%v:\n\n%v\n"
const fileFormat = "// Code generated from %v by @grafig/syntax. DO NOT EDIT.\n\n%v"

const transpileUse = "transpile <path> [flags]"
const transpileShort = `Transpiles a given figscript file.`
const transpileLong = `
Given a file containing valid figscript, the transpile command will parse the
code and generate the transpiled javascript:
`

type transpileCommand struct {
	*command.CommonCommand
}

func newTranspileCommand(logger *slog.Logger, options ...command.CommonCommandOption) *transpileCommand {
	c := &transpileCommand{command.New(logger.WithGroup(transpileUse), options...)}

	c.CommonCommand.PreRun = c.PreRun
	c.CommonCommand.Run = c.Run

	return c
}

func (c *transpileCommand) PreRun(cmd *cobra.Command, args []string) {
	logger.Trace()

	c.ConfigureLoggerFlags()
	c.LogWith(
		slog.String(file.Path, args[0]),
		c.UseConfigAttr(file.Output),
		c.UseConfigAttr(logger.Level),
	)
}

func (c *transpileCommand) Run(cmd *cobra.Command, args []string) {
	logger.Trace()

	errorListener := c.UseValidationErrorListener()

	transpiler := transpiler.NewTranspiler(c.GetLogger(), errorListener)

	transpiled := make(map[string]string, 0)

	if _, err := file.IsGlob(args[0]); err == nil {
		transpiled, err = transpiler.TranspileGlob(args[0])
		c.HandleError(err)
		c.HandleError(errorListener.GetError())
	} else {
		output, err := transpiler.TranspileInput(args[0])
		c.HandleError(err)
		c.HandleError(errorListener.GetError())
		transpiled["out"] = output
	}

	for path, content := range transpiled {
		if !c.GetConfigBool(logger.Silent) {
			fmt.Printf(logFormat, path, content)
		}
		if c.GetConfigString(file.Output) != command.Undefined {
			c.writeFile(path, content)
		}
	}

	_, err := os.Stdout.WriteString(transpiled["out"])
	c.HandleError(err)
}

func (c transpileCommand) writeFile(path string, content string) {
	logger.Trace()

	dir := c.GetConfigString(file.Output)
	if dir == command.Undefined {
		return
	}

	old := filepath.Base(path)
	new := strings.Replace(old, filepath.Ext(old), file.JS, 1)
	out := filepath.Join(dir, new)
	txt := fmt.Sprintf(fileFormat, path, content)

	file, err := os.Create(out)
	c.HandleError(err)

	defer c.DeferErrorFunc(file.Close, err == nil)

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte(txt))

	c.HandleError(err)
	c.DeferErrorFunc(writer.Flush, err == nil)
}
