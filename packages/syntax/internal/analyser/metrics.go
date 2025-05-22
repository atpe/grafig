package analyser

type Metrics struct {
	Sloc       float64           `json:"sloc"`
	Lloc       float64           `json:"lloc"`
	Params     float64           `json:"params"`
	Halstead   HalsteadMetrics   `json:"halstead"`
	Cyclomatic CyclomaticMetrics `json:"cyclomatic"`
}

type CyclomaticMetrics struct {
	Complexity float64 `json:"complexity"`
	Density    float64 `json:"density"`
}

type HalsteadMetrics struct {
	OperatorsUnique float64 `json:"operatorsUnique"`
	OperatorsTotal  float64 `json:"operatorsTotal"`
	OperandsUnique  float64 `json:"operandsUnique"`
	OperandsTotal   float64 `json:"operandsTotal"`
	Vocabulary      float64 `json:"vocabulary"`
	Length          float64 `json:"length"`
	Difficulty      float64 `json:"difficulty"`
	Volume          float64 `json:"volume"`
	Effort          float64 `json:"effort"`
	Time            float64 `json:"time"`
	Bugs            float64 `json:"bugs"`
}
