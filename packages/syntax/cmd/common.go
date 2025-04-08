package cmd

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	LOG_LEVEL         = "log-level"
	LOG_LEVEL_ERROR   = "error"
	LOG_LEVEL_WARN    = "warn"
	LOG_LEVEL_INFO    = "info"
	LOG_LEVEL_DEBUG   = "debug"
	TARGET            = "target"
	TARGET_JAVASCRIPT = "javascript"
	OUTPUT            = "output"
	SILENT            = "silent"
	DEFAULT           = "default"
)

type Command interface {
	Bind(*CommonCommand) *CommonCommand
}

type CommandFunc func(cmd *cobra.Command, args []string)
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

func WithExactArgs(n int) CommonCommandOption {
	return func(c *CommonCommand) {
		c.Args = cobra.ExactArgs(n)
	}
}

func WithLoggerFlags(persistent bool) CommonCommandOption {
	return func(c *CommonCommand) {
		name := LOG_LEVEL
		flags := useFlags(c, persistent)
		flags.String(name, LOG_LEVEL_ERROR, "specify a logger severity level")
		c.config.BindPFlag(name, flags.Lookup(name))

		levels := []string{
			LOG_LEVEL_ERROR,
			LOG_LEVEL_WARN,
			LOG_LEVEL_INFO,
			LOG_LEVEL_DEBUG,
		}

		for _, level := range levels {
			usage := fmt.Sprintf("use the %v logger level", strings.ToUpper(level))
			flags.Bool(level, false, usage)
			c.config.BindPFlag(level, flags.Lookup(level))
		}

		c.MarkFlagsMutuallyExclusive(append(levels, name)...)
	}
}

func WithTargetFlags(persistent bool) CommonCommandOption {
	return func(c *CommonCommand) {
		name := TARGET
		flags := useFlags(c, persistent)
		flags.StringP(name, "t", DEFAULT, "specify a language to target")
		c.config.BindPFlag(name, flags.Lookup(name))

		languages := []string{
			TARGET_JAVASCRIPT,
		}

		for _, language := range languages {
			usage := fmt.Sprintf("use the %v target", strings.ToUpper(language))
			flags.Bool(language, false, usage)
			c.config.BindPFlag(language, flags.Lookup(language))
		}

		c.MarkFlagsMutuallyExclusive(append(languages, name)...)
	}
}

func WithOutputFlag(persistent bool) CommonCommandOption {
	return func(c *CommonCommand) {
		flags := useFlags(c, persistent)
		flags.StringP(OUTPUT, "o", DEFAULT, "specify an output file")
		c.MarkFlagFilename(OUTPUT)
		c.config.BindPFlag(OUTPUT, flags.Lookup(OUTPUT))
	}
}

func WithSilentFlag(persistent bool) CommonCommandOption {
	return func(c *CommonCommand) {
		flags := useFlags(c, persistent)
		flags.BoolP(SILENT, "s", false, "don't write any output to the console")
		c.config.BindPFlag(SILENT, flags.Lookup(SILENT))
	}
}

type CommonCommand struct {
	*cobra.Command
	config *viper.Viper
	logger *slog.Logger
}

func NewCommonCommand(logger *slog.Logger) *CommonCommand {
	return &CommonCommand{
		Command: &cobra.Command{},
		logger:  logger,
		config:  viper.New(),
	}
}

func (c *CommonCommand) Bind(super *CommonCommand) *CommonCommand {
	super.AddCommand(c.Command)
	return c
}

func (c *CommonCommand) ConfigureLoggerFlags() {
	if c.config.GetString(LOG_LEVEL) == DEFAULT {
		switch {
		case c.config.GetBool(LOG_LEVEL_ERROR):
			c.config.Set(LOG_LEVEL, LOG_LEVEL_ERROR)
		case c.config.GetBool(LOG_LEVEL_WARN):
			c.config.Set(LOG_LEVEL, LOG_LEVEL_WARN)
		case c.config.GetBool(LOG_LEVEL_INFO):
			c.config.Set(LOG_LEVEL, LOG_LEVEL_INFO)
		case c.config.GetBool(LOG_LEVEL_DEBUG):
			c.config.Set(LOG_LEVEL, LOG_LEVEL_DEBUG)
		}
	}

	switch c.config.GetString(LOG_LEVEL) {
	case LOG_LEVEL_ERROR:
		slog.SetLogLoggerLevel(slog.LevelError)
	case LOG_LEVEL_WARN:
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case LOG_LEVEL_INFO:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case LOG_LEVEL_DEBUG:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

func (c *CommonCommand) ConfigureTargetFlags(path string) {
	target := c.config.GetString(TARGET)
	if target == DEFAULT {
		switch {
		case c.config.GetBool(TARGET_JAVASCRIPT):
			c.config.Set(TARGET, TARGET_JAVASCRIPT)
		case strings.HasSuffix(path, ".js"):
			c.config.Set(TARGET, TARGET_JAVASCRIPT)
		default:
			defer c.logger.Warn("could not infer target language from file name, defaulting to javascript")
			c.config.Set(TARGET, TARGET_JAVASCRIPT)
		}
	}
}

func useFlags(command *CommonCommand, persistent bool) *pflag.FlagSet {
	if persistent {
		return command.PersistentFlags()
	}
	return command.Flags()
}
