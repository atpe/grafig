package cmd

import (
	"figsyntax/internal/debugger"
	"log/slog"
	"os"
)

type ProgramCommandOption func(*ProgramCommand)
type ProgramCommand struct {
	*CommonCommand
	subcommands []*CommonCommand
}

func NewProgram(options ...ProgramCommandOption) *ProgramCommand {
	debugger.Trace()

	c := &ProgramCommand{
		CommonCommand: NewCommonCommand(slog.Default()),
	}
	for _, o := range options {
		o(c)
	}
	return c
}

func WithCommandOptions(options ...CommonCommandOption) ProgramCommandOption {
	return func(p *ProgramCommand) {
		for _, o := range options {
			o(p.CommonCommand)
		}
	}
}

func WithSubcommands(commands ...*CommonCommand) ProgramCommandOption {
	return func(p *ProgramCommand) {
		for _, c := range commands {
			c.Bind(p.CommonCommand)
		}
		p.subcommands = append(p.subcommands, commands...)
	}
}

var parse = NewParseCommand(
	slog.Default(),
	WithUse("parse"),
	WithShort("Parses a given file in a known language."),
	WithLong("TODO"),
	WithExactArgs(1),
	WithLoggerFlags(false),
	WithTargetFlags(false),
	WithOutputFlag(false),
	WithSilentFlag(false),
)
var analyse = NewAnalyseCommand(
	slog.Default(),
	WithUse("analyse"),
	WithShort("Analyses a given file in a known language."),
	WithLong("TODO"),
	WithExactArgs(1),
	WithLoggerFlags(false),
	WithTargetFlags(false),
	WithOutputFlag(false),
	WithSilentFlag(false),
)

var program = NewProgram(
	WithCommandOptions(
		WithUse("figsyntax"),
		WithShort("A parser and transpiler for the figscript syntax."),
		WithLong("TODO"),
	),
	WithSubcommands(parse.CommonCommand, analyse.CommonCommand),
)

func Execute() {
	debugger.Trace()

	err := program.Execute()
	if err != nil {
		slog.Error(err.Error())

		os.Exit(1)
	}
}
