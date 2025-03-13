package cmd

import (
	"figsyntax/internal/listeners"
	"figsyntax/internal/parser"
	"log/slog"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var parseCmd = &cobra.Command{
	Use:    "parse",
	Short:  "Parses a given file in a known language.",
	Long:   `TODO`,
	Args:   cobra.ExactArgs(1),
	PreRun: parsePreRun,
	Run:    parseRun,
}
var parseViper = viper.New()

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringP("out", "o", "./parse-tree.json", "The output file path")
	parseCmd.Flags().StringP("target", "t", "file", "The target language to parse")
	parseCmd.Flags().String("log-level", "error", "The target language to parse")
	parseCmd.Flags().Bool("javascript", false, "Sets the target language to javascript")
	parseCmd.Flags().BoolP("warn", "w", false, "Sets the logger level to warn")
	parseCmd.Flags().BoolP("info", "i", false, "Sets the logger level to info")
	parseCmd.Flags().BoolP("debug", "d", false, "Sets the logger level to debug")

	parseCmd.MarkFlagFilename("out")
	parseCmd.MarkFlagsMutuallyExclusive("target", "javascript")
	parseCmd.MarkFlagsMutuallyExclusive("warn", "info", "debug")

	parseViper.BindPFlag("out", parseCmd.Flags().Lookup("out"))
	parseViper.BindPFlag("target", parseCmd.Flags().Lookup("target"))
	parseViper.BindPFlag("javascript", parseCmd.Flags().Lookup("javascript"))
	parseViper.BindPFlag("log-level", parseCmd.Flags().Lookup("log-level"))
	parseViper.BindPFlag("warn", parseCmd.Flags().Lookup("warn"))
	parseViper.BindPFlag("info", parseCmd.Flags().Lookup("info"))
	parseViper.BindPFlag("debug", parseCmd.Flags().Lookup("debug"))
}

func parsePreRun(cmd *cobra.Command, args []string) {
	target := parseViper.GetString("target")
	if target == "file" {
		switch {
		case strings.HasSuffix(args[0], ".js"):
			parseViper.Set("target", "javascript")
		default:
			slog.Warn("cannot infer target language from file name", "path", args[0])
			parseViper.Set("target", "javascript")
		}
	}

	if parseViper.GetString("log-level") == "error" {
		switch {
		case parseViper.GetBool("info"):
			parseViper.Set("log-level", "info")
		case parseViper.GetBool("warn"):
			parseViper.Set("log-level", "warn")
		case parseViper.GetBool("debug"):
			parseViper.Set("log-level", "debug")
		}
	}

	switch parseViper.GetString("log-level") {
	case "warn":
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	default:
		slog.SetLogLoggerLevel(slog.LevelError)
	}

	slog.Info(
		"starting...",
		slog.String("path", args[0]),
		slog.String("target", parseViper.GetString("target")),
		slog.String("log-level", parseViper.GetString("log-level")),
	)
}

func parseRun(cmd *cobra.Command, args []string) {
	errorListener := listeners.NewSlogErrorListener(slog.Default())
	debugListener := listeners.NewDebugListener(slog.Default())

	slog.Debug("reading file...", "path", args[0])

	input, err := antlr.NewFileStream(args[0])
	if err != nil {
		panic(err)
	}

	slog.Debug("tokenizing input...")

	lexer := parser.NewJavaScriptLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)

	slog.Debug("generating parser...")

	parser := parser.NewJavaScriptParser(tokens)
	parser.AddErrorListener(errorListener)

	tree := parser.Program()

	if errorListener.HasError() {
		panic(errorListener.GetError())
	}

	antlr.ParseTreeWalkerDefault.Walk(debugListener, tree)

	slog.Info("done")
}
