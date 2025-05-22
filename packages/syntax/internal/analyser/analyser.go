package analyser

import (
	"figsyntax/internal/file"
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"figsyntax/internal/validation"
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"
)

type CommonAnalyser struct {
	logger        *slog.Logger
	target        string
	errorListener *validation.ValidationErrorListener
}

func New(logger *slog.Logger, target string) *CommonAnalyser {
	return &CommonAnalyser{
		logger,
		target,
		validation.NewValidationErrorListener(logger),
	}
}

func (a *CommonAnalyser) AnalyseFile(path string) (Report, error) {
	logger.Trace()

	a.errorListener.Reset()

	target, err := a.useTarget(path)
	if err != nil {
		return Report{}, err
	}

	parser, err := parser.New(
		a.logger,
		target,
		parser.WithFileLexer(path),
		parser.WithErrorListener(a.errorListener),
	)
	if err != nil {
		return Report{}, err
	}

	if a.errorListener.HasError() {
		return Report{}, a.errorListener.GetError()
	}

	complexityListener, err := a.useComplexityListener(target)
	if err != nil {
		return Report{}, err
	}

	parser.Walk(complexityListener.WithFile(path))

	return complexityListener.GetMetrics(), nil
}

func (a *CommonAnalyser) AnalyseGlob(glob string) ([]Report, error) {
	paths, err := filepath.Glob(glob)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	metrics := make([]Report, 0)
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

func (a CommonAnalyser) useComplexityListener(target string) (ComplexityListener, error) {
	logger := a.logger.WithGroup("complexity")
	switch target {
	case parser.JavaScript:
		return NewJavaScriptComplexityListener(logger), nil
	case parser.FigScript:
		return NewFigScriptComplexityListener(logger), nil
	}
	return nil, fmt.Errorf("cannot create complexity listener - target '%v' is not a valid option", target)
}

func (a CommonAnalyser) useTarget(path string) (string, error) {
	switch a.target {
	case parser.JavaScript:
		return parser.JavaScript, nil
	case parser.FigScript:
		return parser.FigScript, nil
	}
	switch {
	case strings.HasSuffix(path, file.JS):
		return parser.JavaScript, nil
	case strings.HasSuffix(path, file.FS):
		return parser.FigScript, nil
	case strings.Contains(path, parser.JavaScript):
		return parser.JavaScript, nil
	case strings.Contains(path, parser.FigScript):
		return parser.FigScript, nil
	}
	return "", fmt.Errorf("cannot infer target from config or path")
}
