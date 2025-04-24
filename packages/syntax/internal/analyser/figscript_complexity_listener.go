package analyser

import (
	"figsyntax/internal/logger"
	"figsyntax/internal/parser/figscript"
	"log/slog"

	"github.com/antlr4-go/antlr/v4"
)

type FigScriptComplexityListener struct {
	figscript.FigScriptParserListener
	logger  *slog.Logger
	scopes  []*Scope
	metrics Metrics
	path    string
}

func NewFigScriptComplexityListener(logger *slog.Logger) *FigScriptComplexityListener {
	return &FigScriptComplexityListener{
		FigScriptParserListener: new(figscript.BaseFigScriptParserListener),
		logger:                  logger,
		scopes:                  make([]*Scope, 0),
	}
}

func (l *FigScriptComplexityListener) GetMetrics() Metrics {
	logger.Trace()
	return l.metrics
}

func (l *FigScriptComplexityListener) WithFile(path string) ComplexityListener {
	logger.Trace()
	l.scopes = make([]*Scope, 0)
	l.metrics = NewMetrics(path)
	l.path = path
	return l
}

func (l *FigScriptComplexityListener) incrementLloc(ctxs ...antlr.TerminalNode) *FigScriptComplexityListener {
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
func (l *FigScriptComplexityListener) incrementParams() *FigScriptComplexityListener {
	l.scopes[0].params++
	return l
}

func (l *FigScriptComplexityListener) incrementComplexity(ctxs ...antlr.TerminalNode) *FigScriptComplexityListener {
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

func (l *FigScriptComplexityListener) setClass(ctx antlr.ParserRuleContext) *FigScriptComplexityListener {
	if ctx != nil {
		l.scopes[0].class = ctx
	} else {
		l.scopes[0].class = l.scopes[0].assignee
		l.scopes[0].assignee = nil
	}
	return l
}
func (l *FigScriptComplexityListener) resetClass() *FigScriptComplexityListener {
	l.scopes[0].class = nil
	return l
}

func (l *FigScriptComplexityListener) setAssignee(ctx antlr.ParserRuleContext) *FigScriptComplexityListener {
	l.scopes[0].assignee = ctx
	return l
}

func (l *FigScriptComplexityListener) addOperators(operators ...antlr.TerminalNode) *FigScriptComplexityListener {
	logger.Trace()
	for _, operator := range operators {
		l.scopes[0].addOperator(operator)
	}
	return l
}

func (l *FigScriptComplexityListener) addOperands(operands ...antlr.TerminalNode) *FigScriptComplexityListener {
	logger.Trace()
	for _, operand := range operands {
		l.scopes[0].addOperand(operand)
	}
	return l
}

func (l *FigScriptComplexityListener) getScope() (*Scope, *Scope) {
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

func (l *FigScriptComplexityListener) pushProgram(ctx *figscript.ProgramContext) *FigScriptComplexityListener {
	logger.Trace()
	l.scopes = append([]*Scope{NewScope(l.path, "program", ctx)}, l.scopes...)
	return l
}

func (l *FigScriptComplexityListener) pushScope(identifier antlr.ParserRuleContext, ctx antlr.ParserRuleContext) *FigScriptComplexityListener {
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

func (l *FigScriptComplexityListener) popScope(c antlr.ParserRuleContext) *FigScriptComplexityListener {
	logger.Trace()
	current, parent := l.getScope()
	if parent != nil {
		parent.aggregate(current)
		parent.assignee = nil
	} else {
		l.metrics = current.getMetrics()
	}
	l.scopes = l.scopes[1:]
	return l
}

// ENTER

func (l *FigScriptComplexityListener) EnterProgram(c *figscript.ProgramContext) {
	l.pushProgram(c)
}
func (l *FigScriptComplexityListener) EnterStatement(c *figscript.StatementContext) {
	switch {
	case c.Block() != nil:
	case c.VariableStatement() != nil:
	default:
		l.incrementLloc()
	}
}
func (l *FigScriptComplexityListener) EnterBlock(c *figscript.BlockContext) {
	l.addOperators(c.OpenBrace())
}
func (l *FigScriptComplexityListener) EnterVariableDeclarationList(c *figscript.VariableDeclarationListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *FigScriptComplexityListener) EnterVariableDeclaration(c *figscript.VariableDeclarationContext) {
	l.incrementLloc().
		setAssignee(c.Assignable()).
		addOperators(c.Assign())
}
func (l *FigScriptComplexityListener) EnterEmptyStatement_(c *figscript.EmptyStatement_Context) {
	l.addOperators(c.SemiColon())
}
func (l *FigScriptComplexityListener) EnterIfStatement(c *figscript.IfStatementContext) {
	l.incrementLloc(c.Else()).
		incrementComplexity().
		addOperators(c.If(), c.Else())
}
func (l *FigScriptComplexityListener) EnterDoStatement(c *figscript.DoStatementContext) {
	l.incrementLloc().
		incrementComplexity().
		addOperators(c.Do(), c.While())
}
func (l *FigScriptComplexityListener) EnterWhileStatement(c *figscript.WhileStatementContext) {
	l.incrementComplexity().
		addOperators(c.While())
}
func (l *FigScriptComplexityListener) EnterForStatement(c *figscript.ForStatementContext) {
	l.incrementComplexity().
		addOperators(c.For())
}
func (l *FigScriptComplexityListener) EnterForInStatement(c *figscript.ForInStatementContext) {
	l.incrementComplexity().
		addOperators(c.For(), c.In())
}
func (l *FigScriptComplexityListener) EnterForOfStatement(c *figscript.ForOfStatementContext) {
	l.incrementComplexity().
		addOperators(c.For(), c.Await(), c.Of())
}
func (l *FigScriptComplexityListener) EnterVarModifier(c *figscript.VarModifierContext) {
	l.addOperators(c.Var(), c.Const())
}
func (l *FigScriptComplexityListener) EnterContinueStatement(c *figscript.ContinueStatementContext) {
	l.addOperators(c.Continue())
}
func (l *FigScriptComplexityListener) EnterBreakStatement(c *figscript.BreakStatementContext) {
	l.addOperators(c.Break())
}
func (l *FigScriptComplexityListener) EnterReturnStatement(c *figscript.ReturnStatementContext) {
	l.addOperators(c.Return())
}
func (l *FigScriptComplexityListener) EnterYieldStatement(c *figscript.YieldStatementContext) {
	l.addOperators(c.Yield(), c.YieldStar())
}
func (l *FigScriptComplexityListener) EnterWithStatement(c *figscript.WithStatementContext) {
	l.addOperators(c.With())
}
func (l *FigScriptComplexityListener) EnterSwitchStatement(c *figscript.SwitchStatementContext) {
	l.addOperators(c.Switch())
}
func (l *FigScriptComplexityListener) EnterCaseBlock(c *figscript.CaseBlockContext) {
	l.addOperators(c.OpenBrace())
}
func (l *FigScriptComplexityListener) EnterCaseClause(c *figscript.CaseClauseContext) {
	l.incrementComplexity().addOperators(c.Case(), c.Colon())
}
func (l *FigScriptComplexityListener) EnterDefaultClause(c *figscript.DefaultClauseContext) {
	l.incrementComplexity().addOperators(c.Default(), c.Colon())
}
func (l *FigScriptComplexityListener) EnterLabelledStatement(c *figscript.LabelledStatementContext) {
	l.addOperators(c.Colon())
}
func (l *FigScriptComplexityListener) EnterThrowStatement(c *figscript.ThrowStatementContext) {
	l.addOperators(c.Throw())
}
func (l *FigScriptComplexityListener) EnterTryStatement(c *figscript.TryStatementContext) {
	l.addOperators(c.Try())
}
func (l *FigScriptComplexityListener) EnterCatchProduction(c *figscript.CatchProductionContext) {
	l.incrementComplexity().
		addOperators(c.Catch())
}
func (l *FigScriptComplexityListener) EnterFinallyProduction(c *figscript.FinallyProductionContext) {
	l.addOperators(c.Finally())
}
func (l *FigScriptComplexityListener) EnterFunctionDeclaration(c *figscript.FunctionDeclarationContext) {
	l.pushScope(c.Identifier(), c).
		addOperators(c.Async(), c.Function_(), c.Multiply())
}
func (l *FigScriptComplexityListener) EnterClassDeclaration(c *figscript.ClassDeclarationContext) {
	l.setClass(c.Identifier()).
		addOperators(c.Class())
}
func (l *FigScriptComplexityListener) EnterClassTail(c *figscript.ClassTailContext) {
	l.addOperators(c.Extends(), c.OpenBrace())
}
func (l *FigScriptComplexityListener) EnterClassElement(c *figscript.ClassElementContext) {
	l.addOperators(c.Static())
}
func (l *FigScriptComplexityListener) EnterMethodDefinition(c *figscript.MethodDefinitionContext) {
	l.pushScope(c.ClassElementName(), c).
		addOperators(c.Async(), c.Multiply())
}
func (l *FigScriptComplexityListener) EnterPrivateIdentifier(c *figscript.PrivateIdentifierContext) {
	l.addOperators(c.Hashtag())
}
func (l *FigScriptComplexityListener) EnterFormalParameterList(c *figscript.FormalParameterListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *FigScriptComplexityListener) EnterFormalParameterArg(c *figscript.FormalParameterArgContext) {
	l.incrementParams().
		addOperators(c.Assign())
}
func (l *FigScriptComplexityListener) EnterLastFormalParameterArg(c *figscript.LastFormalParameterArgContext) {
	l.incrementParams().
		addOperators(c.Ellipsis())
}
func (l *FigScriptComplexityListener) EnterFunctionBody(c *figscript.FunctionBodyContext) {
	l.addOperators(c.OpenBrace())
}
func (l *FigScriptComplexityListener) EnterArrayLiteral(c *figscript.ArrayLiteralContext) {
	l.addOperators(c.OpenBracket())
}
func (l *FigScriptComplexityListener) EnterVectorLiteral(c *figscript.VectorLiteralContext) {
	l.addOperators(c.OpenParen())
}
func (l *FigScriptComplexityListener) EnterComponentList(c *figscript.ComponentListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *FigScriptComplexityListener) EnterElementList(c *figscript.ElementListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *FigScriptComplexityListener) EnterArrayElement(c *figscript.ArrayElementContext) {
	l.addOperators(c.Ellipsis())
}
func (l *FigScriptComplexityListener) EnterPropertyExpressionAssignment(c *figscript.PropertyExpressionAssignmentContext) {
	l.addOperators(c.Colon())
}
func (l *FigScriptComplexityListener) EnterComputedPropertyExpressionAssignment(c *figscript.ComputedPropertyExpressionAssignmentContext) {
	l.addOperators(c.OpenBracket(), c.Colon())
}
func (l *FigScriptComplexityListener) EnterFunctionProperty(c *figscript.FunctionPropertyContext) {
	l.addOperators(c.Async(), c.Multiply())
}
func (l *FigScriptComplexityListener) EnterPropertyShorthand(c *figscript.PropertyShorthandContext) {
	l.addOperators(c.Ellipsis())
}
func (l *FigScriptComplexityListener) EnterPropertyName(c *figscript.PropertyNameContext) {
	l.addOperators(c.OpenBracket())
}
func (l *FigScriptComplexityListener) EnterArguments(c *figscript.ArgumentsContext) {
	l.addOperators(c.OpenParen()).
		addOperators(c.AllComma()...)
}
func (l *FigScriptComplexityListener) EnterArgument(c *figscript.ArgumentContext) {
	l.addOperators(c.Ellipsis())
}
func (l *FigScriptComplexityListener) EnterExpressionSequence(c *figscript.ExpressionSequenceContext) {
	l.addOperators(c.AllComma()...)
}
func (l *FigScriptComplexityListener) EnterClassExpression(c *figscript.ClassExpressionContext) {
	l.setClass(c.Identifier()).
		addOperators(c.Class())
}
func (l *FigScriptComplexityListener) EnterOptionalChainExpression(c *figscript.OptionalChainExpressionContext) {
	l.incrementComplexity().
		addOperators(c.QuestionMarkDot())
}
func (l *FigScriptComplexityListener) EnterMemberIndexExpression(c *figscript.MemberIndexExpressionContext) {
	l.incrementComplexity(c.QuestionMarkDot()).
		addOperators(c.OpenBracket(), c.QuestionMarkDot())
}
func (l *FigScriptComplexityListener) EnterMemberDotExpression(c *figscript.MemberDotExpressionContext) {
	l.incrementComplexity(c.QuestionMark()).
		addOperators(c.QuestionMark(), c.Dot(), c.Hashtag())
}
func (l *FigScriptComplexityListener) EnterNewExpression(c *figscript.NewExpressionContext) {
	l.addOperators(c.New())
}
func (l *FigScriptComplexityListener) EnterMetaExpression(c *figscript.MetaExpressionContext) {
	l.addOperands(c.New()).addOperators(c.Dot())
}
func (l *FigScriptComplexityListener) EnterPostIncrementExpression(c *figscript.PostIncrementExpressionContext) {
	l.addOperators(c.PlusPlus())
}
func (l *FigScriptComplexityListener) EnterPostDecreaseExpression(c *figscript.PostDecreaseExpressionContext) {
	l.addOperators(c.MinusMinus())
}
func (l *FigScriptComplexityListener) EnterDeleteExpression(c *figscript.DeleteExpressionContext) {
	l.addOperators(c.Delete())
}
func (l *FigScriptComplexityListener) EnterVoidExpression(c *figscript.VoidExpressionContext) {
	l.addOperators(c.Void())
}
func (l *FigScriptComplexityListener) EnterTypeofExpression(c *figscript.TypeofExpressionContext) {
	l.addOperators(c.Typeof())
}
func (l *FigScriptComplexityListener) EnterPreIncrementExpression(c *figscript.PreIncrementExpressionContext) {
	l.addOperators(c.PlusPlus())
}
func (l *FigScriptComplexityListener) EnterPreDecreaseExpression(c *figscript.PreDecreaseExpressionContext) {
	l.addOperators(c.MinusMinus())
}
func (l *FigScriptComplexityListener) EnterUnaryPlusExpression(c *figscript.UnaryPlusExpressionContext) {
	l.addOperators(c.Plus())
}
func (l *FigScriptComplexityListener) EnterUnaryMinusExpression(c *figscript.UnaryMinusExpressionContext) {
	l.addOperators(c.Minus())
}
func (l *FigScriptComplexityListener) EnterBitNotExpression(c *figscript.BitNotExpressionContext) {
	l.addOperators(c.BitNot())
}
func (l *FigScriptComplexityListener) EnterNotExpression(c *figscript.NotExpressionContext) {
	l.addOperators(c.Not())
}
func (l *FigScriptComplexityListener) EnterAwaitExpression(c *figscript.AwaitExpressionContext) {
	l.addOperators(c.Await())
}
func (l *FigScriptComplexityListener) EnterPowerExpression(c *figscript.PowerExpressionContext) {
	l.addOperators(c.Power())
}
func (l *FigScriptComplexityListener) EnterMultiplicativeExpression(c *figscript.MultiplicativeExpressionContext) {
	l.addOperators(c.Divide(), c.Multiply(), c.Modulus())
}
func (l *FigScriptComplexityListener) EnterAdditiveExpression(c *figscript.AdditiveExpressionContext) {
	l.addOperators(c.Plus(), c.Minus())
}
func (l *FigScriptComplexityListener) EnterCoalesceExpression(c *figscript.CoalesceExpressionContext) {
	l.incrementComplexity().
		addOperators(c.NullCoalesce())
}
func (l *FigScriptComplexityListener) EnterBitShiftExpression(c *figscript.BitShiftExpressionContext) {
	l.addOperators(c.LeftShiftArithmetic(), c.RightShiftArithmetic(), c.RightShiftLogical())
}
func (l *FigScriptComplexityListener) EnterRelationalExpression(c *figscript.RelationalExpressionContext) {
	l.addOperators(c.LessThan(), c.MoreThan(), c.LessThanEquals(), c.GreaterThanEquals())
}
func (l *FigScriptComplexityListener) EnterInstanceofExpression(c *figscript.InstanceofExpressionContext) {
	l.addOperators(c.Instanceof())
}
func (l *FigScriptComplexityListener) EnterInExpression(c *figscript.InExpressionContext) {
	l.addOperators(c.In())
}
func (l *FigScriptComplexityListener) EnterEqualityExpression(c *figscript.EqualityExpressionContext) {
	l.addOperators(
		c.Equals_(),
		c.NotEquals(),
		c.NotEquals(),
		c.IdentityEquals(),
		c.IdentityEquals(),
		c.IdentityNotEquals(),
		c.IdentityNotEquals(),
	)
}
func (l *FigScriptComplexityListener) EnterBitAndExpression(c *figscript.BitAndExpressionContext) {
	l.addOperands(c.BitAnd())
}
func (l *FigScriptComplexityListener) EnterBitXOrExpression(c *figscript.BitXOrExpressionContext) {
	l.addOperators(c.BitXOr())
}
func (l *FigScriptComplexityListener) EnterBitOrExpression(c *figscript.BitOrExpressionContext) {
	l.addOperands(c.BitOr())
}
func (l *FigScriptComplexityListener) EnterLogicalAndExpression(c *figscript.LogicalAndExpressionContext) {
	l.addOperators(c.And())
}
func (l *FigScriptComplexityListener) EnterLogicalOrExpression(c *figscript.LogicalOrExpressionContext) {
	l.incrementComplexity().
		addOperators(c.Or())
}
func (l *FigScriptComplexityListener) EnterTernaryExpression(c *figscript.TernaryExpressionContext) {
	l.incrementComplexity().
		addOperators(c.QuestionMark(), c.Colon())
}
func (l *FigScriptComplexityListener) EnterAssignmentExpression(c *figscript.AssignmentExpressionContext) {
	l.setAssignee(c.SingleExpression(0)).addOperators(c.Assign())
}
func (l *FigScriptComplexityListener) EnterAssignmentOperator(c *figscript.AssignmentOperatorContext) {
	l.incrementComplexity(c.NullishCoalescingAssign()).
		addOperators(
			c.MultiplyAssign(),
			c.DivideAssign(),
			c.ModulusAssign(),
			c.PlusAssign(),
			c.MinusAssign(),
			c.LeftShiftArithmeticAssign(),
			c.RightShiftArithmeticAssign(),
			c.RightShiftLogicalAssign(),
			c.BitAndAssign(),
			c.BitXorAssign(),
			c.BitOrAssign(),
			c.PowerAssign(),
			c.NullishCoalescingAssign(),
		)
}
func (l *FigScriptComplexityListener) EnterImportExpression(c *figscript.ImportExpressionContext) {
	l.addOperators(c.Import())
}
func (l *FigScriptComplexityListener) EnterThisExpression(c *figscript.ThisExpressionContext) {
	l.addOperands(c.This())
}
func (l *FigScriptComplexityListener) EnterSuperExpression(c *figscript.SuperExpressionContext) {
	l.addOperators(c.Super())
}
func (l *FigScriptComplexityListener) EnterParenthesizedExpression(c *figscript.ParenthesizedExpressionContext) {
	l.addOperators(c.OpenParen())
}
func (l *FigScriptComplexityListener) EnterInitializer(c *figscript.InitializerContext) {
	l.addOperators(c.Assign())
}
func (l *FigScriptComplexityListener) EnterObjectLiteral(c *figscript.ObjectLiteralContext) {
	l.addOperators(c.AllComma()...).
		addOperators(c.OpenBrace())
}
func (l *FigScriptComplexityListener) EnterAnonymousFunctionDecl(c *figscript.AnonymousFunctionDeclContext) {
	l.pushScope(l.scopes[0].assignee, c).
		addOperators(c.Async(), c.Function_(), c.Multiply())
}
func (l *FigScriptComplexityListener) EnterArrowFunction(c *figscript.ArrowFunctionContext) {
	if c.ArrowFunctionBody().FunctionBody() != nil {
		l.pushScope(l.scopes[0].assignee, c).
			addOperators(c.Async(), c.ARROW())
	}
}
func (l *FigScriptComplexityListener) EnterLiteral(c *figscript.LiteralContext) {
	l.addOperands(
		c.NullLiteral(),
		c.BooleanLiteral(),
		c.StringLiteral(),
		c.RegularExpressionLiteral(),
	)
}
func (l *FigScriptComplexityListener) EnterTemplateStringLiteral(c *figscript.TemplateStringLiteralContext) {
	l.addOperands(c.AllBackTick()...)
}

func (l *FigScriptComplexityListener) EnterTemplateStringAtom(c *figscript.TemplateStringAtomContext) {
	l.addOperands(
		c.TemplateStringAtom(),
		c.TemplateStringStartExpression(),
		c.TemplateCloseBrace(),
	)
}
func (l *FigScriptComplexityListener) EnterNumericLiteral(c *figscript.NumericLiteralContext) {
	l.addOperands(
		c.DecimalLiteral(),
		c.HexIntegerLiteral(),
		c.OctalIntegerLiteral(),
		c.OctalIntegerLiteral2(),
		c.BinaryIntegerLiteral(),
	)
}
func (l *FigScriptComplexityListener) EnterBigintLiteral(c *figscript.BigintLiteralContext) {
	l.addOperands(
		c.BigDecimalIntegerLiteral(),
		c.BigHexIntegerLiteral(),
		c.BigOctalIntegerLiteral(),
		c.BigBinaryIntegerLiteral(),
	)
}
func (l *FigScriptComplexityListener) EnterIdentifier(c *figscript.IdentifierContext) {
	l.addOperands(c.Identifier())
}
func (l *FigScriptComplexityListener) EnterLet_(c *figscript.Let_Context) {
	l.addOperators(c.NonStrictLet(), c.StrictLet())
}
func (l *FigScriptComplexityListener) EnterEos(c *figscript.EosContext) {
	l.addOperators(c.SemiColon())
}

// EXIT

func (l *FigScriptComplexityListener) ExitProgram(c *figscript.ProgramContext) {
	l.popScope(c)
}

// func (l *ComplexityListener) ExitSourceElement(c *figscript.SourceElementContext) {}
// func (l *ComplexityListener) ExitStatement(c *figscript.StatementContext) {}
// func (l *ComplexityListener) ExitBlock(c *figscript.BlockContext) {}
// func (l *ComplexityListener) ExitStatementList(c *figscript.StatementListContext) {}
// func (l *ComplexityListener) ExitImportStatement(c *figscript.ImportStatementContext) {}
// func (l *ComplexityListener) ExitImportFromBlock(c *figscript.ImportFromBlockContext) {}
// func (l *ComplexityListener) ExitImportModuleItems(c *figscript.ImportModuleItemsContext) {}
// func (l *ComplexityListener) ExitImportAliasName(c *figscript.ImportAliasNameContext) {}
// func (l *ComplexityListener) ExitModuleExportName(c *figscript.ModuleExportNameContext) {}
// func (l *ComplexityListener) ExitImportedBinding(c *figscript.ImportedBindingContext) {}
// func (l *ComplexityListener) ExitImportDefault(c *figscript.ImportDefaultContext) {}
// func (l *ComplexityListener) ExitImportNamespace(c *figscript.ImportNamespaceContext) {}
// func (l *ComplexityListener) ExitImportFrom(c *figscript.ImportFromContext) {}
// func (l *ComplexityListener) ExitAliasName(c *figscript.AliasNameContext) {}
// func (l *ComplexityListener) ExitExportDeclaration(c *figscript.ExportDeclarationContext) {}
// func (l *ComplexityListener) ExitExportDefaultDeclaration(c *figscript.ExportDefaultDeclarationContext) {}
// func (l *ComplexityListener) ExitExportFromBlock(c *figscript.ExportFromBlockContext) {}
// func (l *ComplexityListener) ExitExportModuleItems(c *figscript.ExportModuleItemsContext) {}
// func (l *ComplexityListener) ExitExportAliasName(c *figscript.ExportAliasNameContext) {}
// func (l *ComplexityListener) ExitDeclaration(c *figscript.DeclarationContext) {}
// func (l *ComplexityListener) ExitVariableStatement(c *figscript.VariableStatementContext) {}
// func (l *ComplexityListener) ExitVariableDeclarationList(c *figscript.VariableDeclarationListContext) {}
func (l *FigScriptComplexityListener) ExitVariableDeclaration(c *figscript.VariableDeclarationContext) {
	l.setAssignee(nil)
}

// func (l *ComplexityListener) ExitEmptyStatement_(c *figscript.EmptyStatement_Context) {}
// func (l *ComplexityListener) ExitExpressionStatement(c *figscript.ExpressionStatementContext) {}
// func (l *ComplexityListener) ExitIfStatement(c *figscript.IfStatementContext) {}
// func (l *ComplexityListener) ExitDoStatement(c *figscript.DoStatementContext) {}
// func (l *ComplexityListener) ExitWhileStatement(c *figscript.WhileStatementContext) {}
// func (l *ComplexityListener) ExitForStatement(c *figscript.ForStatementContext) {}
// func (l *ComplexityListener) ExitForInStatement(c *figscript.ForInStatementContext) {}
// func (l *ComplexityListener) ExitForOfStatement(c *figscript.ForOfStatementContext) {}
// func (l *ComplexityListener) ExitVarModifier(c *figscript.VarModifierContext) {}
// func (l *ComplexityListener) ExitContinueStatement(c *figscript.ContinueStatementContext) {}
// func (l *ComplexityListener) ExitBreakStatement(c *figscript.BreakStatementContext) {}
// func (l *ComplexityListener) ExitReturnStatement(c *figscript.ReturnStatementContext) {}
// func (l *ComplexityListener) ExitYieldStatement(c *figscript.YieldStatementContext) {}
// func (l *ComplexityListener) ExitWithStatement(c *figscript.WithStatementContext) {}
// func (l *ComplexityListener) ExitSwitchStatement(c *figscript.SwitchStatementContext) {}
// func (l *ComplexityListener) ExitCaseBlock(c *figscript.CaseBlockContext) {}
// func (l *ComplexityListener) ExitCaseClauses(c *figscript.CaseClausesContext) {}
// func (l *ComplexityListener) ExitCaseClause(c *figscript.CaseClauseContext) {}
// func (l *ComplexityListener) ExitDefaultClause(c *figscript.DefaultClauseContext) {}
// func (l *ComplexityListener) ExitLabelledStatement(c *figscript.LabelledStatementContext) {}
// func (l *ComplexityListener) ExitThrowStatement(c *figscript.ThrowStatementContext) {}
// func (l *ComplexityListener) ExitTryStatement(c *figscript.TryStatementContext) {}
// func (l *ComplexityListener) ExitCatchProduction(c *figscript.CatchProductionContext) {}
// func (l *ComplexityListener) ExitFinallyProduction(c *figscript.FinallyProductionContext) {}
// func (l *ComplexityListener) ExitDebuggerStatement(c *figscript.DebuggerStatementContext) {}
func (l *FigScriptComplexityListener) ExitFunctionDeclaration(c *figscript.FunctionDeclarationContext) {
	l.popScope(c)
}

// func (l *ComplexityListener) ExitClassDeclaration(c *figscript.ClassDeclarationContext) {}
func (l *FigScriptComplexityListener) ExitClassTail(c *figscript.ClassTailContext) {
	l.resetClass()
}

// func (l *ComplexityListener) ExitClassElement(c *figscript.ClassElementContext) {}
func (l *FigScriptComplexityListener) ExitMethodDefinition(c *figscript.MethodDefinitionContext) {
	l.popScope(c)
}

// func (l *ComplexityListener) ExitFieldDefinition(c *figscript.FieldDefinitionContext) {}
// func (l *ComplexityListener) ExitClassElementName(c *figscript.ClassElementNameContext) {}
// func (l *ComplexityListener) ExitPrivateIdentifier(c *figscript.PrivateIdentifierContext) {}
// func (l *ComplexityListener) ExitFormalParameterList(c *figscript.FormalParameterListContext) {}
// func (l *ComplexityListener) ExitFormalParameterArg(c *figscript.FormalParameterArgContext) {}
// func (l *ComplexityListener) ExitLastFormalParameterArg(c *figscript.LastFormalParameterArgContext) {}
// func (l *ComplexityListener) ExitFunctionBody(c *figscript.FunctionBodyContext) {}
// func (l *ComplexityListener) ExitSourceElements(c *figscript.SourceElementsContext) {}
// func (l *ComplexityListener) ExitArrayLiteral(c *figscript.ArrayLiteralContext) {}
// func (l *ComplexityListener) ExitElementList(c *figscript.ElementListContext) {}
// func (l *ComplexityListener) ExitArrayElement(c *figscript.ArrayElementContext) {}
// func (l *ComplexityListener) ExitPropertyExpressionAssignment(c *figscript.PropertyExpressionAssignmentContext) {}
// func (l *ComplexityListener) ExitComputedPropertyExpressionAssignment(c *figscript.ComputedPropertyExpressionAssignmentContext) {}
// func (l *ComplexityListener) ExitFunctionProperty(c *figscript.FunctionPropertyContext) {}
// func (l *ComplexityListener) ExitPropertyGetter(c *figscript.PropertyGetterContext) {}
// func (l *ComplexityListener) ExitPropertySetter(c *figscript.PropertySetterContext) {}
// func (l *ComplexityListener) ExitPropertyShorthand(c *figscript.PropertyShorthandContext) {}
// func (l *ComplexityListener) ExitPropertyName(c *figscript.PropertyNameContext) {}
// func (l *ComplexityListener) ExitArguments(c *figscript.ArgumentsContext) {}
// func (l *ComplexityListener) ExitArgument(c *figscript.ArgumentContext) {}
// func (l *ComplexityListener) ExitExpressionSequence(c *figscript.ExpressionSequenceContext) {}
// func (l *ComplexityListener) ExitTemplateStringExpression(c *figscript.TemplateStringExpressionContext) {}
// func (l *ComplexityListener) ExitTernaryExpression(c *figscript.TernaryExpressionContext) {}
// func (l *ComplexityListener) ExitLogicalAndExpression(c *figscript.LogicalAndExpressionContext) {}
// func (l *ComplexityListener) ExitPowerExpression(c *figscript.PowerExpressionContext) {}
// func (l *ComplexityListener) ExitPreIncrementExpression(c *figscript.PreIncrementExpressionContext) {}
// func (l *ComplexityListener) ExitObjectLiteralExpression(c *figscript.ObjectLiteralExpressionContext) {}
// func (l *ComplexityListener) ExitMetaExpression(c *figscript.MetaExpressionContext) {}
// func (l *ComplexityListener) ExitInExpression(c *figscript.InExpressionContext) {}
// func (l *ComplexityListener) ExitLogicalOrExpression(c *figscript.LogicalOrExpressionContext) {}
// func (l *ComplexityListener) ExitOptionalChainExpression(c *figscript.OptionalChainExpressionContext) {}
// func (l *ComplexityListener) ExitNotExpression(c *figscript.NotExpressionContext) {}
// func (l *ComplexityListener) ExitPreDecreaseExpression(c *figscript.PreDecreaseExpressionContext) {}
// func (l *ComplexityListener) ExitArgumentsExpression(c *figscript.ArgumentsExpressionContext) {}
// func (l *ComplexityListener) ExitAwaitExpression(c *figscript.AwaitExpressionContext) {}
// func (l *ComplexityListener) ExitThisExpression(c *figscript.ThisExpressionContext) {}
// func (l *ComplexityListener) ExitFunctionExpression(c *figscript.FunctionExpressionContext) {}
// func (l *ComplexityListener) ExitUnaryMinusExpression(c *figscript.UnaryMinusExpressionContext) {}
func (l *FigScriptComplexityListener) ExitAssignmentExpression(c *figscript.AssignmentExpressionContext) {
	l.setAssignee(nil)
}

// func (l *ComplexityListener) ExitPostDecreaseExpression(c *figscript.PostDecreaseExpressionContext) {}
// func (l *ComplexityListener) ExitTypeofExpression(c *figscript.TypeofExpressionContext) {}
// func (l *ComplexityListener) ExitInstanceofExpression(c *figscript.InstanceofExpressionContext) {}
// func (l *ComplexityListener) ExitUnaryPlusExpression(c *figscript.UnaryPlusExpressionContext) {}
// func (l *ComplexityListener) ExitDeleteExpression(c *figscript.DeleteExpressionContext) {}
// func (l *ComplexityListener) ExitImportExpression(c *figscript.ImportExpressionContext) {}
// func (l *ComplexityListener) ExitEqualityExpression(c *figscript.EqualityExpressionContext) {}
// func (l *ComplexityListener) ExitBitXOrExpression(c *figscript.BitXOrExpressionContext) {}
// func (l *ComplexityListener) ExitSuperExpression(c *figscript.SuperExpressionContext) {}
// func (l *ComplexityListener) ExitMultiplicativeExpression(c *figscript.MultiplicativeExpressionContext) {}
// func (l *ComplexityListener) ExitBitShiftExpression(c *figscript.BitShiftExpressionContext) {}
// func (l *ComplexityListener) ExitParenthesizedExpression(c *figscript.ParenthesizedExpressionContext) {}
// func (l *ComplexityListener) ExitAdditiveExpression(c *figscript.AdditiveExpressionContext) {}
// func (l *ComplexityListener) ExitRelationalExpression(c *figscript.RelationalExpressionContext) {}
// func (l *ComplexityListener) ExitPostIncrementExpression(c *figscript.PostIncrementExpressionContext) {}
// func (l *ComplexityListener) ExitYieldExpression(c *figscript.YieldExpressionContext) {}
// func (l *ComplexityListener) ExitBitNotExpression(c *figscript.BitNotExpressionContext) {}
// func (l *ComplexityListener) ExitNewExpression(c *figscript.NewExpressionContext) {}
// func (l *ComplexityListener) ExitLiteralExpression(c *figscript.LiteralExpressionContext) {}
// func (l *ComplexityListener) ExitArrayLiteralExpression(c *figscript.ArrayLiteralExpressionContext) {}
// func (l *ComplexityListener) ExitMemberDotExpression(c *figscript.MemberDotExpressionContext) {}
// func (l *ComplexityListener) ExitClassExpression(c *figscript.ClassExpressionContext) {}
// func (l *ComplexityListener) ExitMemberIndexExpression(c *figscript.MemberIndexExpressionContext) {}
// func (l *ComplexityListener) ExitIdentifierExpression(c *figscript.IdentifierExpressionContext) {}
// func (l *ComplexityListener) ExitBitAndExpression(c *figscript.BitAndExpressionContext) {}
// func (l *ComplexityListener) ExitBitOrExpression(c *figscript.BitOrExpressionContext) {}
// func (l *ComplexityListener) ExitAssignmentOperatorExpression(c *figscript.AssignmentOperatorExpressionContext) {}
// func (l *ComplexityListener) ExitVoidExpression(c *figscript.VoidExpressionContext) {}
// func (l *ComplexityListener) ExitCoalesceExpression(c *figscript.CoalesceExpressionContext) {}
// func (l *ComplexityListener) ExitInitializer(c *figscript.InitializerContext) {}
// func (l *ComplexityListener) ExitAssignable(c *figscript.AssignableContext) {}
// func (l *ComplexityListener) ExitObjectLiteral(c *figscript.ObjectLiteralContext) {}
// func (l *ComplexityListener) ExitNamedFunction(c *figscript.NamedFunctionContext) {}
func (l *FigScriptComplexityListener) ExitAnonymousFunctionDecl(c *figscript.AnonymousFunctionDeclContext) {
	l.popScope(c)
}

// func (l *ComplexityListener) ExitArrowFunction(c *figscript.ArrowFunctionContext) {}
// func (l *ComplexityListener) ExitArrowFunctionParameters(c *figscript.ArrowFunctionParametersContext) {}
func (l *FigScriptComplexityListener) ExitArrowFunctionBody(c *figscript.ArrowFunctionBodyContext) {
	if c.FunctionBody() != nil {
		l.popScope(c)
	}
}

// func (l *ComplexityListener) ExitAssignmentOperator(c *figscript.AssignmentOperatorContext) {}
// func (l *ComplexityListener) ExitLiteral(c *figscript.LiteralContext) {}
// func (l *ComplexityListener) ExitTemplateStringLiteral(c *figscript.TemplateStringLiteralContext) {}
// func (l *ComplexityListener) ExitTemplateStringAtom(c *figscript.TemplateStringAtomContext) {}
// func (l *ComplexityListener) ExitNumericLiteral(c *figscript.NumericLiteralContext) {}
// func (l *ComplexityListener) ExitBigintLiteral(c *figscript.BigintLiteralContext) {}
// func (l *ComplexityListener) ExitGetter(c *figscript.GetterContext) {}
// func (l *ComplexityListener) ExitSetter(c *figscript.SetterContext) {}
// func (l *ComplexityListener) ExitIdentifierName(c *figscript.IdentifierNameContext) {}
// func (l *ComplexityListener) ExitIdentifier(c *figscript.IdentifierContext) {}
// func (l *ComplexityListener) ExitReservedWord(c *figscript.ReservedWordContext) {}
// func (l *ComplexityListener) ExitKeyword(c *figscript.KeywordContext) {}
// func (l *ComplexityListener) ExitLet_(c *figscript.Let_Context) {}
// func (l *ComplexityListener) ExitEos(c *figscript.EosContext) {}
