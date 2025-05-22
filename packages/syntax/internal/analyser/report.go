package analyser

import (
	"fmt"
	"strconv"
)

const (
	Format = "format"
	Flat   = "flat"
	JSON   = "json"
	CSV    = "csv"
)

type Report struct {
	Metrics
	Total     Metrics  `json:"total"`
	Average   Metrics  `json:"average"`
	File      string   `json:"file"`
	Function  string   `json:"function"`
	Line      int      `json:"line"`
	Functions []Report `json:"functions,omitempty"`
}

func NewReport(file string) Report {
	return Report{
		File:      file,
		Functions: make([]Report, 0),
	}
}

func (r Report) Flat() []Report {
	copy := r
	copy.Functions = nil
	metrics := []Report{copy}
	for _, child := range r.Functions {
		metrics = append(metrics, child.Flat()...)
	}
	return metrics
}

func (r Report) Print() {
	fmt.Printf("\n%v\n", r.File)
	for _, metric := range r.Flat() {
		fmt.Printf("\nFunction: %v", metric.Function)
		fmt.Printf("\n\tLine No.: %v", metric.Line)
		fmt.Printf("\n\tPhysical LOC: %v", metric.Sloc)
		fmt.Printf("\n\tLogical LOC: %v", metric.Lloc)
		fmt.Printf("\n\tParameter count: %v", metric.Params)
		fmt.Printf("\n\tCyclomatic complexity: %v", metric.Cyclomatic.Complexity)
		fmt.Printf("\n\tCyclomatic density: %v%%", metric.Cyclomatic.Density*100)
		fmt.Printf("\n\tHalstead difficulty: %v", metric.Halstead.Difficulty)
		fmt.Printf("\n\tHalstead volume: %v", metric.Halstead.Volume)
		fmt.Printf("\n\tHalstead effort: %v", metric.Halstead.Effort)
		fmt.Println()
	}
	fmt.Println()
}

func (r Report) ToCsv() [][]string {
	rows := make([][]string, 0)
	for _, report := range r.Flat() {
		rows = append(rows, report.toCsv())
	}
	return rows
}

func (r Report) toCsv() []string {
	return []string{
		r.File,
		r.Function,
		strconv.FormatInt(int64(r.Line), 10),
		strconv.FormatFloat(r.Sloc, 'f', -1, 64),
		strconv.FormatFloat(r.Lloc, 'f', -1, 64),
		strconv.FormatFloat(r.Params, 'f', -1, 64),
		strconv.FormatFloat(r.Cyclomatic.Complexity, 'f', -1, 64),
		strconv.FormatFloat(r.Cyclomatic.Density, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.OperatorsUnique, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.OperatorsTotal, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.OperandsUnique, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.OperandsTotal, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.Vocabulary, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.Length, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.Volume, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.Difficulty, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.Effort, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.Time, 'f', -1, 64),
		strconv.FormatFloat(r.Halstead.Bugs, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Sloc, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Lloc, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Params, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Cyclomatic.Complexity, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Cyclomatic.Density, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.OperatorsUnique, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.OperatorsTotal, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.OperandsUnique, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.OperandsTotal, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.Vocabulary, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.Length, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.Volume, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.Difficulty, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.Effort, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.Time, 'f', -1, 64),
		strconv.FormatFloat(r.Total.Halstead.Bugs, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Sloc, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Lloc, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Params, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Cyclomatic.Complexity, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Cyclomatic.Density, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.OperatorsUnique, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.OperatorsTotal, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.OperandsUnique, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.OperandsTotal, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.Vocabulary, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.Length, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.Volume, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.Difficulty, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.Effort, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.Time, 'f', -1, 64),
		strconv.FormatFloat(r.Average.Halstead.Bugs, 'f', -1, 64),
	}
}

var Headers = []string{
	"Path",
	"Name",
	"Line",
	"SLOC",
	"LLOC",
	"Params",
	"Cyclomatic Complexity",
	"Cyclomatic Density",
	"Halstead Operators Distinct",
	"Halstead Operators Total",
	"Halstead Operands Distinct",
	"Halstead Operands Total",
	"Halstead Vocabulary",
	"Halstead Length",
	"Halstead Volume",
	"Halstead Difficulty",
	"Halstead Effort",
	"Halstead Time",
	"Halstead Bugs",
	"SLOC (Total)",
	"LLOC (Total)",
	"Params (Total)",
	"Cyclomatic Complexity (Total)",
	"Cyclomatic Density (Total)",
	"Halstead Operators Distinct (Total)",
	"Halstead Operators Total (Total)",
	"Halstead Operands Distinct (Total)",
	"Halstead Operands Total (Total)",
	"Halstead Vocabulary (Total)",
	"Halstead Length (Total)",
	"Halstead Volume (Total)",
	"Halstead Difficulty (Total)",
	"Halstead Effort (Total)",
	"Halstead Time (Total)",
	"Halstead Bugs (Total)",
	"SLOC (Average)",
	"LLOC (Average)",
	"Params (Average)",
	"Cyclomatic Complexity (Average)",
	"Cyclomatic Density (Average)",
	"Halstead Operators Distinct (Average)",
	"Halstead Operators Average (Average)",
	"Halstead Operands Distinct (Average)",
	"Halstead Operands Average (Average)",
	"Halstead Vocabulary (Average)",
	"Halstead Length (Average)",
	"Halstead Volume (Average)",
	"Halstead Difficulty (Average)",
	"Halstead Effort (Average)",
	"Halstead Time (Average)",
	"Halstead Bugs (Average)",
}
