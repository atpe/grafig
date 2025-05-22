package analyser

import (
	"figsyntax/internal/logger"
	"log/slog"

	"github.com/antlr4-go/antlr/v4"
)

type Reporter struct {
	logger *slog.Logger
	scopes []*Scope
	report Report
	path   string
}

func NewReporter(logger *slog.Logger) *Reporter {
	return &Reporter{
		logger: logger,
		scopes: make([]*Scope, 0),
	}
}

func (l *Reporter) GetMetrics() Report {
	logger.Trace()
	return l.report
}

func (l *Reporter) incrementLloc(ctxs ...antlr.TerminalNode) *Reporter {
	if len(ctxs) == 0 {
		l.scopes[0].lloc++
		return l
	}

	for _, ctx := range ctxs {
		if ctx != nil {
			l.scopes[0].lloc++
		}
	}

	return l
}
func (l *Reporter) incrementParams() *Reporter {
	l.scopes[0].params++
	return l
}

func (l *Reporter) incrementComplexity(ctxs ...antlr.TerminalNode) *Reporter {
	if len(ctxs) == 0 {
		l.scopes[0].complexity++
		return l
	}

	for _, ctx := range ctxs {
		if ctx != nil {
			l.scopes[0].complexity++
		}
	}

	return l
}

func (l *Reporter) setClass(ctx antlr.ParserRuleContext) *Reporter {
	if ctx != nil {
		l.scopes[0].class = ctx
	} else {
		l.scopes[0].class = l.scopes[0].assignee
		l.scopes[0].assignee = nil
	}
	return l
}
func (l *Reporter) resetClass() *Reporter {
	l.scopes[0].class = nil
	return l
}

func (l *Reporter) setAssignee(ctx antlr.ParserRuleContext) *Reporter {
	l.scopes[0].assignee = ctx
	return l
}

func (l *Reporter) addOperators(operators ...antlr.TerminalNode) *Reporter {
	logger.Trace()
	for _, operator := range operators {
		l.scopes[0].addOperator(operator)
	}
	return l
}

func (l *Reporter) addOperands(operands ...antlr.TerminalNode) *Reporter {
	logger.Trace()
	for _, operand := range operands {
		l.scopes[0].addOperand(operand)
	}
	return l
}

func (l *Reporter) getScope() (*Scope, *Scope) {
	logger.Trace()
	switch len(l.scopes) {
	case 0:
		return nil, nil
	case 1:
		return l.scopes[0], nil
	default:
		return l.scopes[0], l.scopes[1]
	}
}

func (l *Reporter) pushScope(identifier antlr.ParserRuleContext, ctx antlr.ParserRuleContext) *Reporter {
	logger.Trace()

	id := ""
	if identifier == nil {
		id = "anonymous"
	} else {
		id = identifier.GetText()
	}

	if len(l.scopes) > 0 && l.scopes[0].class != nil {
		id = l.scopes[0].class.GetText() + "." + id
	}

	l.scopes = append([]*Scope{NewScope(l.path, id, ctx)}, l.scopes...)
	return l
}

func (l *Reporter) popScope(c antlr.ParserRuleContext) *Reporter {
	logger.Trace()
	current, parent := l.getScope()
	if parent != nil {
		parent.aggregate(current)
		parent.assignee = nil
	} else {
		l.report = current.getReport()
	}
	l.scopes = l.scopes[1:]
	return l
}
