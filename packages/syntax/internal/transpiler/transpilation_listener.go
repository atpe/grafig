package transpiler

import (
	"figsyntax/internal/logger"
	"figsyntax/internal/parser/figscript"
	"fmt"
	"log/slog"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

const FigScriptImport = "FigScript"

var overloads = map[string]string{
	"+": "add",
	"-": "sub",
	"*": "mult",
	"/": "div",
}

type TranspilationListener struct {
	figscript.FigScriptParserListener
	logger         *slog.Logger
	hidden         map[int]antlr.Token
	scopes         [][]antlr.Token
	previous       int
	component      int
	trailingLambda *figscript.TrailingLambdaExpressionContext
}

func NewTranspilationListener(l *slog.Logger, hidden map[int]antlr.Token) *TranspilationListener {
	logger.Trace()
	return &TranspilationListener{
		FigScriptParserListener: new(figscript.BaseFigScriptParserListener),
		logger:                  l,
		hidden:                  hidden,
		scopes:                  make([][]antlr.Token, 0),
	}
}

func (l *TranspilationListener) GetText() string {
	logger.Trace()
	text := ""
	for _, token := range l.scopes[0] {
		if token.GetTokenType() == figscript.FigScriptParserEOF {
			continue
		}

		if index := token.GetTokenIndex(); index >= 0 {
			text += l.getHiddenText(l.previous, index)
			l.previous = index
		}
		text += token.GetText()
	}

	return text + l.getHiddenText(l.previous, -1)
}

func (l *TranspilationListener) VisitTerminal(node antlr.TerminalNode) {
	logger.Trace()
	l.appendTokens(node.GetSymbol())
}

func (l *TranspilationListener) EnterEveryRule(node antlr.ParserRuleContext) {
	logger.Trace()

	rule := figscript.NewFigScriptParser(nil).RuleNames[node.GetRuleIndex()]
	l.logger.Info(
		fmt.Sprintf("parsing %v rule:\t", rule),
		slog.String("text", node.GetText()),
	)

	l.scopes = append([][]antlr.Token{make([]antlr.Token, 0)}, l.scopes...)
}

func (l *TranspilationListener) ExitEveryRule(node antlr.ParserRuleContext) {
	logger.Trace()
	if len(l.scopes) > 1 {
		scope := l.scopes[0]
		l.scopes = l.scopes[1:]
		l.appendTokens(scope...)
	}
}

func (l *TranspilationListener) ExitVectorLiteral(node *figscript.VectorLiteralContext) {
	logger.Trace()
	node.OpenParen().GetSymbol().SetText("{ ")
	node.CloseParen().GetSymbol().SetText(" }")
	l.component = 0
}

func (l *TranspilationListener) ExitComponent(node *figscript.ComponentContext) {
	logger.Trace()
	component := string(byte(l.component + 120))

	l.prependTokens(
		newToken(node, figscript.FigScriptLexerIdentifier, node.GetStart().GetTokenIndex(), component),
		newToken(node, figscript.FigScriptLexerColon, -1, ":"),
		newToken(node, figscript.FigScriptLexerWhiteSpaces, -1, " "),
	)

	l.component++
}

func (l *TranspilationListener) EnterTrailingLambdaExpression(node *figscript.TrailingLambdaExpressionContext) {
	logger.Trace()
	l.trailingLambda = node

	args := node.Arguments()
	if args != nil {
		args.RemoveLastChild()
	}
}

func (l *TranspilationListener) ExitTrailingLambdaExpression(node *figscript.TrailingLambdaExpressionContext) {
	logger.Trace()
	l.trailingLambda = nil
}

func (l *TranspilationListener) EnterTrailingLambda(node *figscript.TrailingLambdaContext) {
	logger.Trace()
	if l.trailingLambda.Arguments() == nil {
		l.appendTokens(
			newToken(node, figscript.FigScriptLexerOpenParen, -1, "("),
		)
	} else {
		l.appendTokens(
			newToken(node, figscript.FigScriptLexerComma, -1, ","),
			newToken(node, figscript.FigScriptLexerWhiteSpaces, -1, " "),
		)
	}

	l.appendTokens(
		newToken(node, figscript.FigScriptLexerOpenParen, -1, "("),
		newToken(node, figscript.FigScriptLexerIdentifier, -1, "it"),
		newToken(node, figscript.FigScriptLexerCloseParen, -1, ")"),
		newToken(node, figscript.FigScriptLexerWhiteSpaces, -1, " "),
		newToken(node, figscript.FigScriptLexerARROW, -1, "=>"),
	)
}

func (l *TranspilationListener) ExitTrailingLambda(node *figscript.TrailingLambdaContext) {
	logger.Trace()
	l.appendTokens(
		newToken(node, figscript.FigScriptLexerCloseParen, -1, ")"),
	)
}

func (l *TranspilationListener) EnterLambdaStatement(node *figscript.LambdaStatementContext) {
	logger.Trace()
	l.appendTokens(
		newToken(node, figscript.FigScriptLexerIdentifier, node.GetStart().GetTokenIndex(), "it"),
		newToken(node, figscript.FigScriptLexerDot, -1, "."),
	)
}

func (l *TranspilationListener) EnterSwizzleExpression(node *figscript.SwizzleExpressionContext) {
	logger.Trace()
	l.appendTokens(
		newToken(node, figscript.FigScriptLexerIdentifier, node.GetStart().GetTokenIndex(), "FigScript"),
		newToken(node, figscript.FigScriptLexerDot, -1, "."),
		newToken(node, figscript.FigScriptLexerIdentifier, -1, "from"),
		newToken(node, figscript.FigScriptLexerOpenParen, -1, "("),
	)
}

func (l *TranspilationListener) ExitSwizzleExpression(node *figscript.SwizzleExpressionContext) {
	logger.Trace()
	components := strings.Split(node.Swizzle().GetText(), "")
	l.scopes[0] = l.scopes[0][:len(l.scopes[0])-2]

	for _, component := range components {
		l.appendTokens(
			newToken(node, figscript.FigScriptLexerComma, -1, ","),
			newToken(node, figscript.FigScriptLexerWhiteSpaces, -1, " "),
			newToken(node, figscript.FigScriptLexerStringLiteral, -1, fmt.Sprintf("'%v'", component)),
		)
	}

	l.appendTokens(
		newToken(node, figscript.FigScriptLexerCloseParen, -1, ")"),
	)
}

func (l *TranspilationListener) EnterMultiplicativeExpression(node *figscript.MultiplicativeExpressionContext) {
	multiply := node.Multiply()
	divide := node.Divide()
	right := node.SingleExpression(1)

	var operator antlr.TerminalNode
	switch {
	case multiply != nil:
		operator = multiply
	case divide != nil:
		operator = divide
	default:
		return
	}

	l.handleInfixOperator(operator, right)(node)
}

// func (l *TranspilationListener) ExitMultiplicativeExpression(node *figscript.MultiplicativeExpressionContext) {
// 	l.appendTokens(
// 		newToken(node, figscript.FigScriptLexerOpenParen, -1, ")"),
// 	)
// }

func (l *TranspilationListener) EnterAdditiveExpression(node *figscript.AdditiveExpressionContext) {
	plus := node.Plus()
	minus := node.Minus()
	right := node.SingleExpression(1)

	var operator antlr.TerminalNode
	switch {
	case plus != nil:
		operator = plus
	case minus != nil:
		operator = minus
	default:
		return
	}

	l.handleInfixOperator(operator, right)(node)
}

// func (l *TranspilationListener) ExitAdditiveExpression(node *figscript.AdditiveExpressionContext) {
// 	l.appendTokens(
// 		newToken(node, figscript.FigScriptLexerOpenParen, -1, ")"),
// 	)
// }

func (l *TranspilationListener) appendTokens(tokens ...antlr.Token) {
	logger.Trace()
	l.scopes[0] = append(l.scopes[0], tokens...)
}
func (l *TranspilationListener) prependTokens(tokens ...antlr.Token) {
	logger.Trace()
	l.scopes[0] = append(tokens, l.scopes[0]...)
}

func (l TranspilationListener) getHiddenText(start int, stop int) string {
	logger.Trace()
	text := ""

	if stop < 0 {
		for i := start + 1; true; i += 1 {
			if hidden, ok := l.hidden[i]; ok {
				text += hidden.GetText()
			} else {
				break
			}
		}
		return text
	}

	for i := start + 1; i < stop; i += 1 {
		if hidden, ok := l.hidden[i]; ok {
			text += hidden.GetText()
		}
	}
	return text
}

func newToken(node antlr.ParserRuleContext, tokenType int, index int, text string) antlr.Token {
	logger.Trace()
	token := antlr.NewCommonToken(node.GetStart().GetSource(), tokenType, 0, -1, -1)
	token.SetTokenIndex(index)
	token.SetText(text)
	return token
}

func (l *TranspilationListener) handleInfixOperator(operator antlr.TerminalNode, right figscript.ISingleExpressionContext) func(antlr.ParserRuleContext) {
	identifier, ok := overloads[operator.GetText()]
	if !ok {
		return nil
	}

	index := right.GetStart().GetTokenIndex() - 1
	if hidden, ok := l.hidden[index]; ok {
		text := hidden.GetText()
		if text == " " {
			delete(l.hidden, index)
		}
		if strings.HasSuffix(text, " ") {
			text = text[:len(text)-1]
			hidden.SetText(text)
		}
	}

	return func(node antlr.ParserRuleContext) {
		node.RemoveLastChild()
		node.RemoveLastChild()
		node.AddTokenNode(newToken(node, figscript.FigScriptLexerComma, -1, ","))
		node.AddChild(right)
		node.AddTokenNode(newToken(node, figscript.FigScriptLexerCloseParen, -1, ")"))

		l.appendTokens(
			newToken(node, figscript.FigScriptLexerIdentifier, node.GetStart().GetTokenIndex(), "FigScript"),
			newToken(node, figscript.FigScriptLexerDot, -1, "."),
			newToken(node, figscript.FigScriptLexerIdentifier, -1, identifier),
			newToken(node, figscript.FigScriptLexerOpenParen, -1, "("),
		)
	}
}
