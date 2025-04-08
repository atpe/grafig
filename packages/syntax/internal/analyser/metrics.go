package analyser

import "fmt"

type Metrics struct {
	File       string            `json:"file"`
	Function   string            `json:"function"`
	Line       int               `json:"line"`
	Sloc       int               `json:"sloc"`
	Lloc       int               `json:"lloc"`
	Params     int               `json:"params"`
	Halstead   HalsteadMetrics   `json:"halstead"`
	Cyclomatic CyclomaticMetrics `json:"cyclomatic"`
	Functions  []Metrics         `json:"functions,omitempty"`
}

func NewMetrics(file string) Metrics {
	return Metrics{
		File:      file,
		Functions: make([]Metrics, 0),
	}
}

func (m Metrics) Flat() []Metrics {
	copy := m
	copy.Functions = nil
	metrics := []Metrics{copy}
	for _, child := range m.Functions {
		metrics = append(metrics, child.Flat()...)
	}
	return metrics
}

func (m Metrics) Print() {
	fmt.Printf("\n%v\n", m.File)
	for _, metric := range m.Flat() {
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

type CyclomaticMetrics struct {
	Complexity float64 `json:"complexity"`
	Density    float64 `json:"density"`
}

type HalsteadMetrics struct {
	OperatorsTotal  float64 `json:"operatorsTotal"`
	OperatorsUnique float64 `json:"operatorsUnique"`
	OperandsTotal   float64 `json:"operandsTotal"`
	OperandsUnique  float64 `json:"operandsUnique"`
	Vocabulary      float64 `json:"vocabulary"`
	Length          float64 `json:"length"`
	Difficulty      float64 `json:"difficulty"`
	Volume          float64 `json:"volume"`
	Effort          float64 `json:"effort"`
	Time            float64 `json:"time"`
	Bugs            float64 `json:"bugs"`
}
