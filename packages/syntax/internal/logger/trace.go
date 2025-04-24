package logger

import (
	"log/slog"
	"runtime"
)

func Trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	slog.Debug(
		f.Name(),
		slog.String("file", file),
		slog.Int("line", line),
	)
}
