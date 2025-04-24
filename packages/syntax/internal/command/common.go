package command

import (
	"figsyntax/internal/file"
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"figsyntax/internal/validation"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Undefined = "Undefined"
)

type CommonCommand struct {
	*cobra.Command
	config *viper.Viper
	logger *slog.Logger
}

func New(logger *slog.Logger, options ...CommonCommandOption) *CommonCommand {
	c := &CommonCommand{
		Command: &cobra.Command{},
		logger:  logger,
		config:  viper.New(),
	}
	for _, o := range options {
		o(c)
	}
	return c
}
func Start(options ...CommonCommandOption) {
	logger.Trace()
	cmd := New(slog.Default(), options...)
	cmd.DeferErrorFunc(cmd.Execute)
}

func (c *CommonCommand) Bind(child *CommonCommand) *CommonCommand {
	logger.Trace()
	c.AddCommand(child.Command)
	return c
}

func (c *CommonCommand) ConfigureLoggerFlags() {
	logger.Trace()
	if c.GetConfigString(logger.Level) == Undefined {
		switch {
		case c.GetConfigBool(logger.LevelWarn):
			c.SetConfigValue(logger.Level, logger.LevelWarn)
		case c.GetConfigBool(logger.LevelInfo):
			c.SetConfigValue(logger.Level, logger.LevelInfo)
		case c.GetConfigBool(logger.LevelDebug):
			c.SetConfigValue(logger.Level, logger.LevelDebug)
		default:
			c.SetConfigValue(logger.Level, logger.LevelError)
		}
	}

	switch c.GetConfigString(logger.Level) {
	case logger.LevelError:
		slog.SetLogLoggerLevel(slog.LevelError)
	case logger.LevelWarn:
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case logger.LevelInfo:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case logger.LevelDebug:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

func (c *CommonCommand) ConfigureTargetFlags(path string) {
	logger.Trace()
	target := c.GetConfigString(parser.Target)
	if target == Undefined {
		switch {
		case c.GetConfigBool(parser.JavaScript):
			c.SetConfigValue(parser.Target, parser.JavaScript)
		case c.GetConfigBool(parser.FigScript):
			c.SetConfigValue(parser.Target, parser.FigScript)
		case strings.HasSuffix(path, file.JS):
			c.SetConfigValue(parser.Target, parser.JavaScript)
		case strings.HasSuffix(path, file.FS):
			c.SetConfigValue(parser.Target, parser.FigScript)
		default:
			defer c.logger.Warn("could not infer target language from file name, defaulting to javascript")
			c.SetConfigValue(parser.Target, parser.JavaScript)
		}
	}
}

func (c *CommonCommand) GetConfigValue(key string) any {
	logger.Trace()
	return c.config.Get(key)
}
func (c *CommonCommand) SetConfigValue(key string, value any) {
	logger.Trace()
	c.config.Set(key, value)
}
func (c *CommonCommand) GetConfigString(key string) string {
	logger.Trace()
	return c.config.GetString(key)
}
func (c *CommonCommand) GetConfigBool(key string) bool {
	logger.Trace()
	return c.config.GetBool(key)
}
func (c *CommonCommand) GetConfigInt(key string) int {
	logger.Trace()
	return c.config.GetInt(key)
}
func (c *CommonCommand) UseConfigAttr(key string) slog.Attr {
	logger.Trace()
	return slog.Any(key, c.GetConfigValue(key))
}
func (c *CommonCommand) GetLogger() *slog.Logger {
	logger.Trace()
	return c.logger
}
func (c *CommonCommand) LogDebug(msg string, args ...any) {
	logger.Trace()
	c.logger.Debug(msg, args...)
}
func (c *CommonCommand) LogInfo(msg string, args ...any) {
	logger.Trace()
	c.logger.Info(msg, args...)
}
func (c *CommonCommand) LogWarn(msg string, args ...any) {
	logger.Trace()
	c.logger.Warn(msg, args...)
}
func (c *CommonCommand) LogError(msg string, args ...any) {
	logger.Trace()
	c.logger.Error(msg, args...)
}
func (c *CommonCommand) LogWith(args ...any) {
	logger.Trace()
	c.logger = c.logger.With(args...)
}

func (c *CommonCommand) UseValidationErrorListener(args ...any) *validation.ValidationErrorListener {
	logger.Trace()
	return validation.NewValidationErrorListener(c.logger)
}

func (c *CommonCommand) HandleError(err error) {
	logger.Trace()
	if err != nil {
		c.logger.Error(err.Error())
		os.Exit(1)
	}
}
func (c *CommonCommand) DeferErrorFunc(getError func() error, conditions ...bool) {
	logger.Trace()
	for _, condition := range conditions {
		if !condition {
			return
		}
	}
	c.HandleError(getError())
}
