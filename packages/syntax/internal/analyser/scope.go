package analyser

import (
	"math"

	"github.com/antlr4-go/antlr/v4"
)

type Scope struct {
	file       string
	id         string
	class      antlr.ParserRuleContext
	assignee   antlr.ParserRuleContext
	line       int
	sloc       float64
	lloc       float64
	params     float64
	complexity float64
	operators  map[string]float64
	operands   map[string]float64
	reports    []Report
}

func NewScope(file string, id string, ctx antlr.ParserRuleContext) *Scope {
	line := ctx.GetStart().GetLine()
	sloc := (ctx.GetStop().GetLine() + 1) - line
	return &Scope{
		file:       file,
		id:         id,
		line:       line,
		sloc:       float64(sloc),
		complexity: 1,
		operators:  make(map[string]float64),
		operands:   make(map[string]float64),
		reports:    make([]Report, 0),
	}
}

func (s *Scope) addOperator(operator antlr.TerminalNode) *Scope {
	if operator != nil {
		text := operator.GetText()
		s.operators[text] = s.operators[text] + 1
	}
	return s
}

func (s *Scope) addOperand(operand antlr.TerminalNode) *Scope {
	if operand != nil {
		text := operand.GetText()
		s.operands[text] = s.operands[text] + 1
	}
	return s
}

func (s *Scope) countOperators() float64 {
	total := 0.
	for _, count := range s.operators {
		total += count
	}
	return float64(total)
}

func (s *Scope) countOperands() float64 {
	total := 0.
	for _, count := range s.operands {
		total += count
	}
	return float64(total)
}

func (s *Scope) aggregate(child *Scope) *Scope {
	s.reports = append(s.reports, child.getReport())
	return s
}

func (s *Scope) getReport() Report {
	return Report{
		Metrics:   s.getMetrics(),
		Total:     s.getTotal(),
		Average:   s.getAverage(),
		File:      s.file,
		Function:  s.id,
		Line:      s.line,
		Functions: s.reports,
	}
}

func (s *Scope) getMetrics() Metrics {
	return Metrics{
		s.sloc,
		s.lloc,
		s.params,
		s.getHalsteadMetrics(),
		s.getCyclomaticMetrics(),
	}
}

func (s *Scope) getTotal() Metrics {
	total := s.getMetrics()

	if len(s.reports) <= 0 {
		return total
	}

	for _, report := range s.reports {
		total.Lloc += report.Total.Lloc
		total.Params += report.Total.Params
		total.Halstead.OperatorsUnique += report.Total.Halstead.OperatorsUnique
		total.Halstead.OperandsUnique += report.Total.Halstead.OperandsUnique
		total.Halstead.OperatorsTotal += report.Total.Halstead.OperatorsTotal
		total.Halstead.OperandsTotal += report.Total.Halstead.OperandsTotal
		total.Halstead.Vocabulary += report.Total.Halstead.Vocabulary
		total.Halstead.Length += report.Total.Halstead.Length
		total.Halstead.Volume += report.Total.Halstead.Volume
		total.Halstead.Difficulty += report.Total.Halstead.Difficulty
		total.Halstead.Effort += report.Total.Halstead.Effort
		total.Halstead.Time += report.Total.Halstead.Time
		total.Halstead.Bugs += report.Total.Halstead.Bugs
		total.Cyclomatic.Complexity += report.Total.Cyclomatic.Complexity - 1
		total.Cyclomatic.Density += report.Total.Cyclomatic.Density
	}

	return total
}

func (s *Scope) getAverage() Metrics {
	average := s.getTotal()

	length := 1 + float64(len(s.reports))

	average.Sloc /= length
	average.Lloc /= length
	average.Params /= length
	average.Halstead.OperatorsUnique /= length
	average.Halstead.OperandsUnique /= length
	average.Halstead.OperatorsTotal /= length
	average.Halstead.OperandsTotal /= length
	average.Halstead.Vocabulary /= length
	average.Halstead.Length /= length
	average.Halstead.Volume /= length
	average.Halstead.Difficulty /= length
	average.Halstead.Effort /= length
	average.Halstead.Time /= length
	average.Halstead.Bugs /= length
	average.Cyclomatic.Complexity /= length
	average.Cyclomatic.Density /= length

	return average
}

func (s *Scope) getCyclomaticMetrics() CyclomaticMetrics {
	complexity := float64(s.complexity)
	density := 1.0
	if s.lloc != 0 {
		density = complexity / float64(s.lloc)
	}
	return CyclomaticMetrics{
		complexity,
		density,
	}
}

func (s *Scope) getHalsteadMetrics() HalsteadMetrics {
	operatorsUnique := float64(len(s.operators))
	operandsUnique := float64(len(s.operands))
	operatorsTotal := s.countOperators()
	operandsTotal := s.countOperands()

	if operandsTotal == 0 && operatorsTotal == 0 {
		return HalsteadMetrics{}
	}

	length := operatorsTotal + operandsTotal
	vocabulary := operatorsUnique + operandsUnique
	volume := float64(length) * math.Log2(float64(vocabulary))
	difficulty := (operatorsUnique / 2) * (operandsTotal / math.Max(operandsUnique, 1))
	effort := difficulty * volume
	time := effort / 18
	bugs := volume / 3000

	return HalsteadMetrics{
		operatorsUnique,
		operatorsTotal,
		operandsUnique,
		operandsTotal,
		vocabulary,
		length,
		difficulty,
		volume,
		effort,
		time,
		bugs,
	}
}
