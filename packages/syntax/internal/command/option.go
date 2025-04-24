package command

import (
	"figsyntax/internal/file"
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type CommonCommandOption func(*CommonCommand)

func WithUse(text string) CommonCommandOption {
	return func(c *CommonCommand) {
		c.Use = text
	}
}

func WithShort(text string) CommonCommandOption {
	return func(c *CommonCommand) {
		c.Short = text
	}
}

func WithLong(text string) CommonCommandOption {
	return func(c *CommonCommand) {
		c.Long = text
	}
}

func WithSinglePathArg() CommonCommandOption {
	return func(c *CommonCommand) {
		c.Args = file.SinglePathArg
	}
}

func WithLoggerFlags() CommonCommandOption {
	return func(c *CommonCommand) {
		flags := useFlags(c)
		flags.String(logger.Level, Undefined, "specify a logger severity level")
		c.config.BindPFlag(logger.Level, flags.Lookup(logger.Level))

		levels := []string{
			logger.LevelError,
			logger.LevelWarn,
			logger.LevelInfo,
			logger.LevelDebug,
		}

		for _, level := range levels {
			usage := fmt.Sprintf("use the %v logger level", strings.ToUpper(level))
			flags.Bool(level, false, usage)
			c.config.BindPFlag(level, flags.Lookup(level))
		}

		c.MarkFlagsMutuallyExclusive(append(levels, logger.Level)...)
	}
}

func WithTargetFlags() CommonCommandOption {
	return func(c *CommonCommand) {
		flags := useFlags(c)
		flags.StringP(parser.Target, "t", Undefined, "specify a language to target")
		c.config.BindPFlag(parser.Target, flags.Lookup(parser.Target))

		languages := []string{
			parser.JavaScript,
			parser.FigScript,
		}

		for _, language := range languages {
			usage := fmt.Sprintf("use the %v target", strings.ToUpper(language))
			flags.Bool(language, false, usage)
			c.config.BindPFlag(language, flags.Lookup(language))
		}

		c.MarkFlagsMutuallyExclusive(append(languages, parser.Target)...)
	}
}

func WithOutputFlag(dir bool) CommonCommandOption {
	return func(c *CommonCommand) {
		flags := useFlags(c)
		if dir {
			flags.StringP(file.Output, "o", Undefined, "specify an output dir")
			c.MarkFlagDirname(file.Output)
		} else {
			flags.StringP(file.Output, "o", Undefined, "specify an output file")
			c.MarkFlagFilename(file.Output)
		}
		c.config.BindPFlag(file.Output, flags.Lookup(file.Output))
	}
}

func WithSilentFlag() CommonCommandOption {
	return func(c *CommonCommand) {
		flags := useFlags(c)
		flags.BoolP(logger.Silent, "s", false, "don't write any output to the console")
		c.config.BindPFlag(logger.Silent, flags.Lookup(logger.Silent))
	}
}

func WithSubcommand(command *CommonCommand) CommonCommandOption {
	return func(c *CommonCommand) {
		c.Bind(command)
	}
}

func useFlags(command *CommonCommand) *pflag.FlagSet {
	return command.Flags()
}
