package analyser

import (
	"figsyntax/internal/debugger"
	"figsyntax/internal/parser"
	"figsyntax/internal/validation"
	"log/slog"
	"path/filepath"
)

type Analyser struct {
	logger             *slog.Logger
	errorListener      *validation.ValidationErrorListener
	complexityListener *ComplexityListener
}

func NewAnalyser(logger *slog.Logger) *Analyser {
	a := &Analyser{}
	a.logger = logger
	a.errorListener = validation.NewValidationErrorListener(a.logger.WithGroup("errors"))
	a.complexityListener = NewComplexityListener(a.logger.WithGroup("complexity"))
	return a
}

func (a *Analyser) AnalyseFile(path string, target string) (Metrics, error) {
	debugger.Trace()

	a.errorListener.Reset()

	parser, err := parser.NewCommonParser(
		a.logger.WithGroup("parser"),
		target,
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

func (a *Analyser) AnalyseGlob(glob string, target string) ([]Metrics, error) {
	paths, err := filepath.Glob(glob)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	metrics := make([]Metrics, 0)
	for _, path := range paths {
		m, err := a.AnalyseFile(path, target)
		if err != nil {
			a.logger.Error(err.Error(), slog.String("file", path))
			return nil, err
		}

		metrics = append(metrics, m)
	}

	return metrics, nil
}
