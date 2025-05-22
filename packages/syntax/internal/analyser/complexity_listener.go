package analyser

import (
	"github.com/antlr4-go/antlr/v4"
)

type ComplexityListener interface {
	antlr.ParseTreeListener
	WithFile(path string) ComplexityListener
	GetMetrics() Report
}
