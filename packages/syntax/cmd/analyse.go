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

var analyseCmd = &cobra.Command{
	Use:    "analyse",
	Short:  "Analyses a given file in a known language.",
	Long:   `TODO`,
	Args:   cobra.ExactArgs(1),
	PreRun: analysePreRun,
	Run:    analyseRun,
}
var analyseViper = viper.New()

func init() {
	rootCmd.AddCommand(analyseCmd)

	analyseCmd.Flags().StringP("out", "o", "./metrics.json", "The output file path")
	analyseCmd.Flags().StringP("target", "t", "file", "The target language to analyse")
	analyseCmd.Flags().String("log-level", "error", "The target language to analyse")
	analyseCmd.Flags().Bool("javascript", false, "Sets the target language to javascript")
	analyseCmd.Flags().BoolP("warn", "w", false, "Sets the logger level to warn")
	analyseCmd.Flags().BoolP("info", "i", false, "Sets the logger level to info")
	analyseCmd.Flags().BoolP("debug", "d", false, "Sets the logger level to debug")

	analyseCmd.MarkFlagFilename("out")
	analyseCmd.MarkFlagsMutuallyExclusive("target", "javascript")
	analyseCmd.MarkFlagsMutuallyExclusive("warn", "info", "debug")

	analyseViper.BindPFlag("out", analyseCmd.Flags().Lookup("out"))
	analyseViper.BindPFlag("target", analyseCmd.Flags().Lookup("target"))
	analyseViper.BindPFlag("javascript", analyseCmd.Flags().Lookup("javascript"))
	analyseViper.BindPFlag("log-level", analyseCmd.Flags().Lookup("log-level"))
	analyseViper.BindPFlag("warn", analyseCmd.Flags().Lookup("warn"))
	analyseViper.BindPFlag("info", analyseCmd.Flags().Lookup("info"))
	analyseViper.BindPFlag("debug", analyseCmd.Flags().Lookup("debug"))
}

func analysePreRun(cmd *cobra.Command, args []string) {
	target := analyseViper.GetString("target")
	if target == "file" {
		switch {
		case strings.HasSuffix(args[0], ".js"):
			analyseViper.Set("target", "javascript")
		default:
			slog.Warn("cannot infer target language from file name", "path", args[0])
			analyseViper.Set("target", "javascript")
		}
	}

	if analyseViper.GetString("log-level") == "error" {
		switch {
		case analyseViper.GetBool("info"):
			analyseViper.Set("log-level", "info")
		case analyseViper.GetBool("warn"):
			analyseViper.Set("log-level", "warn")
		case analyseViper.GetBool("debug"):
			analyseViper.Set("log-level", "debug")
		}
	}

	switch analyseViper.GetString("log-level") {
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
		slog.String("target", analyseViper.GetString("target")),
		slog.String("log-level", analyseViper.GetString("log-level")),
	)
}

func analyseRun(cmd *cobra.Command, args []string) {
	errorListener := listeners.NewSlogErrorListener(slog.Default())
	complexityListener := listeners.NewComplexityListener(slog.Default())

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

	antlr.ParseTreeWalkerDefault.Walk(complexityListener, tree)

	slog.Info("done", "nodes", complexityListener.GetMetrics())
}
