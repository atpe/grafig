package transpiler

import (
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"figsyntax/internal/validation"
	"log/slog"
	"path/filepath"
)

type Transpiler struct {
	logger                *slog.Logger
	errorListener         *validation.ValidationErrorListener
	transpilationListener *TranspilationListener
}

func NewTranspiler(logger *slog.Logger, errorListener *validation.ValidationErrorListener) *Transpiler {
	return &Transpiler{
		logger:        logger,
		errorListener: errorListener,
	}
}

func (t *Transpiler) TranspileInput(data string) (string, error) {
	logger.Trace()

	t.errorListener.Reset()

	parser, err := parser.New(
		t.logger.WithGroup("parser"),
		"figscript",
		parser.WithInputLexer(data),
		parser.WithErrorListener(t.errorListener),
	)
	if err != nil {
		return "", err
	}

	t.transpilationListener = NewTranspilationListener(t.logger.WithGroup("listener"), parser.GetHidden())

	if t.errorListener.HasError() {
		return "", t.errorListener.GetError()
	}

	parser.Walk(t.transpilationListener)

	return t.transpilationListener.GetText(), nil
}

func (t *Transpiler) TranspileFile(path string) (string, error) {
	logger.Trace()

	t.errorListener.Reset()

	parser, err := parser.New(
		t.logger.WithGroup("parser"),
		"figscript",
		parser.WithFileLexer(path),
		parser.WithErrorListener(t.errorListener),
	)
	if err != nil {
		return "", err
	}

	t.transpilationListener = NewTranspilationListener(t.logger.WithGroup("listener"), parser.GetHidden())

	if t.errorListener.HasError() {
		return "", t.errorListener.GetError()
	}

	parser.Walk(t.transpilationListener)

	return t.transpilationListener.GetText(), nil
}

func (t *Transpiler) TranspileGlob(glob string) (map[string]string, error) {
	paths, err := filepath.Glob(glob)
	if err != nil {
		t.logger.Error(err.Error())
		return nil, err
	}

	files := make(map[string]string, 0)
	for _, path := range paths {
		file, err := t.TranspileFile(path)
		if err != nil {
			t.logger.Error(err.Error(), slog.String("file", path))
			return nil, err
		}

		files[path] = file
	}

	return files, nil
}
