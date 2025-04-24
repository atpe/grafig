package analyser

import (
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"figsyntax/internal/validation"
	"log/slog"
	"path/filepath"
)

type CommonAnalyser struct {
	logger             *slog.Logger
	target             string
	errorListener      *validation.ValidationErrorListener
	complexityListener ComplexityListener
}

func NewCommonAnalyser(logger *slog.Logger, target string) *CommonAnalyser {
	errorListener := validation.NewValidationErrorListener(logger)
	switch target {
	case "javascript":
		return &CommonAnalyser{logger, target, errorListener, NewJavaScriptComplexityListener(logger.WithGroup("complexity"))}
	case "figscript":
		return &CommonAnalyser{logger, target, errorListener, NewFigScriptComplexityListener(logger.WithGroup("complexity"))}
	default:
		return &CommonAnalyser{logger, target, errorListener, NewJavaScriptComplexityListener(logger.WithGroup("complexity"))}
	}
}

func (a *CommonAnalyser) AnalyseFile(path string) (Metrics, error) {
	logger.Trace()

	a.errorListener.Reset()

	parser, err := parser.New(
		a.logger,
		a.target,
		parser.WithFileLexer(path),
		parser.WithErrorListener(a.errorListener),
	)
	if err != nil {
		return Metrics{}, err
	}

	if a.errorListener.HasError() {
		return Metrics{}, a.errorListener.GetError()
	}

	parser.Walk(a.complexityListener.WithFile(path))

	return a.complexityListener.GetMetrics(), nil
}

func (a *CommonAnalyser) AnalyseGlob(glob string) ([]Metrics, error) {
	paths, err := filepath.Glob(glob)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	metrics := make([]Metrics, 0)
	for _, path := range paths {
		m, err := a.AnalyseFile(path)
		if err != nil {
			a.logger.Error(err.Error(), slog.String("file", path))
			return nil, err
		}

		metrics = append(metrics, m)
	}

	return metrics, nil
}
