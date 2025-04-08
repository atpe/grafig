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
	sloc       int
	lloc       int
	params     int
	complexity int
	operators  map[string]int
	operands   map[string]int
	metrics    []Metrics
}

func NewScope(file string, id string, ctx antlr.ParserRuleContext) *Scope {
	line := ctx.GetStart().GetLine()
	sloc := (ctx.GetStop().GetLine() + 1) - line
	return &Scope{
		file:       file,
		id:         id,
		line:       line,
		sloc:       sloc,
		complexity: 1,
		operators:  make(map[string]int),
		operands:   make(map[string]int),
		metrics:    make([]Metrics, 0),
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
	total := 0
	for _, count := range s.operators {
		total += count
	}
	return float64(total)
}

func (s *Scope) countOperands() float64 {
	total := 0
	for _, count := range s.operands {
		total += count
	}
	return float64(total)
}

func (s *Scope) aggregate(child *Scope) *Scope {
	s.metrics = append(s.metrics, child.getMetrics())
	return s
}

func (s *Scope) getMetrics() Metrics {
	return Metrics{
		File:       s.file,
		Function:   s.id,
		Line:       s.line,
		Sloc:       s.sloc,
		Lloc:       s.lloc,
		Params:     s.params,
		Halstead:   s.getHalsteadMetrics(),
		Cyclomatic: s.getCyclomaticMetrics(),
		Functions:  s.metrics,
	}
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
