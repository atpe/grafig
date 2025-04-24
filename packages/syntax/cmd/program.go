package cmd

import (
	"figsyntax/internal/command"
	"figsyntax/internal/logger"
	"log/slog"
)

const programUse = "figsyntax"
const programShort = `A parser, transpiler, and complexity analyser for the figscript syntax.`
const programLong = "Use a subcommand with --help to see more usage information."

func Execute() {
	logger.Trace()
	parse := newParseCommand(
		slog.Default(),
		command.WithUse(parseUse),
		command.WithShort(parseShort),
		command.WithLong(parseLong),
		command.WithSinglePathArg(),
		command.WithLoggerFlags(),
		command.WithTargetFlags(),
		command.WithSilentFlag(),
	)
	analyse := newAnalyseCommand(
		slog.Default(),
		command.WithUse(analyseUse),
		command.WithShort(analyseShort),
		command.WithLong(analyseLong),
		command.WithSinglePathArg(),
		command.WithLoggerFlags(),
		command.WithTargetFlags(),
		command.WithOutputFlag(false),
		command.WithSilentFlag(),
	)
	transpile := newTranspileCommand(
		slog.Default(),
		command.WithUse(transpileUse),
		command.WithShort(transpileShort),
		command.WithLong(transpileLong),
		command.WithSinglePathArg(),
		command.WithLoggerFlags(),
		command.WithOutputFlag(true),
		command.WithSilentFlag(),
	)
	command.Start(
		command.WithUse(programUse),
		command.WithShort(programShort),
		command.WithLong(programLong),
		command.WithSubcommand(parse.CommonCommand),
		command.WithSubcommand(analyse.CommonCommand),
		command.WithSubcommand(transpile.CommonCommand),
	)
}
