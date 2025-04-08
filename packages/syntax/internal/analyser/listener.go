package analyser

import (
	"figsyntax/internal/debugger"
	"figsyntax/internal/parser"
	"log/slog"

	"github.com/antlr4-go/antlr/v4"
)

type ComplexityListener struct {
	parser.JavaScriptParserListener
	logger  *slog.Logger
	scopes  []*Scope
	metrics Metrics
	path    string
}

func NewComplexityListener(logger *slog.Logger) *ComplexityListener {
	debugger.Trace()
	return &ComplexityListener{
		JavaScriptParserListener: new(parser.BaseJavaScriptParserListener),
		logger:                   logger,
		scopes:                   make([]*Scope, 0),
	}
}

func (l *ComplexityListener) GetMetrics() Metrics {
	debugger.Trace()
	return l.metrics
}

func (l *ComplexityListener) WithFile(path string) *ComplexityListener {
	debugger.Trace()
	l.scopes = make([]*Scope, 0)
	l.metrics = NewMetrics(path)
	l.path = path
	return l
}

func (l *ComplexityListener) incrementLloc(ctxs ...antlr.TerminalNode) *ComplexityListener {
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
func (l *ComplexityListener) incrementParams() *ComplexityListener {
	l.scopes[0].params++
	return l
}

func (l *ComplexityListener) incrementComplexity(ctxs ...antlr.TerminalNode) *ComplexityListener {
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

func (l *ComplexityListener) setClass(ctx antlr.ParserRuleContext) *ComplexityListener {
	if ctx != nil {
		l.scopes[0].class = ctx
	} else {
		l.scopes[0].class = l.scopes[0].assignee
		l.scopes[0].assignee = nil
	}
	return l
}
func (l *ComplexityListener) resetClass() *ComplexityListener {
	l.scopes[0].class = nil
	return l
}

func (l *ComplexityListener) setAssignee(ctx antlr.ParserRuleContext) *ComplexityListener {
	l.scopes[0].assignee = ctx
	return l
}

func (l *ComplexityListener) addOperators(operators ...antlr.TerminalNode) *ComplexityListener {
	debugger.Trace()
	for _, operator := range operators {
		l.scopes[0].addOperator(operator)
	}
	return l
}

func (l *ComplexityListener) addOperands(operands ...antlr.TerminalNode) *ComplexityListener {
	debugger.Trace()
	for _, operand := range operands {
		l.scopes[0].addOperand(operand)
	}
	return l
}

func (l *ComplexityListener) getScope() (*Scope, *Scope) {
	debugger.Trace()
	switch len(l.scopes) {
	case 0:
		return nil, nil
	case 1:
		return l.scopes[0], nil
	default:
		return l.scopes[0], l.scopes[1]
	}
}

func (l *ComplexityListener) pushProgram(ctx *parser.ProgramContext) *ComplexityListener {
	debugger.Trace()
	l.scopes = append([]*Scope{NewScope(l.path, "program", ctx)}, l.scopes...)
	return l
}

func (l *ComplexityListener) pushScope(identifier antlr.ParserRuleContext, ctx antlr.ParserRuleContext) *ComplexityListener {
	debugger.Trace()

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

func (l *ComplexityListener) popScope(c antlr.ParserRuleContext) *ComplexityListener {
	debugger.Trace()
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

func (l *ComplexityListener) EnterProgram(c *parser.ProgramContext) {
	l.pushProgram(c)
}
func (l *ComplexityListener) EnterStatement(c *parser.StatementContext) {
	switch {
	case c.Block() != nil:
	case c.VariableStatement() != nil:
	default:
		l.incrementLloc()
	}
}
func (l *ComplexityListener) EnterBlock(c *parser.BlockContext) {
	l.addOperators(c.OpenBrace())
}
func (l *ComplexityListener) EnterVariableDeclarationList(c *parser.VariableDeclarationListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *ComplexityListener) EnterVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.incrementLloc().
		setAssignee(c.Assignable()).
		addOperators(c.Assign())
}
func (l *ComplexityListener) EnterEmptyStatement_(c *parser.EmptyStatement_Context) {
	l.addOperators(c.SemiColon())
}
func (l *ComplexityListener) EnterIfStatement(c *parser.IfStatementContext) {
	l.incrementLloc(c.Else()).
		incrementComplexity().
		addOperators(c.If(), c.Else())
}
func (l *ComplexityListener) EnterDoStatement(c *parser.DoStatementContext) {
	l.incrementLloc().
		incrementComplexity().
		addOperators(c.Do(), c.While())
}
func (l *ComplexityListener) EnterWhileStatement(c *parser.WhileStatementContext) {
	l.incrementComplexity().
		addOperators(c.While())
}
func (l *ComplexityListener) EnterForStatement(c *parser.ForStatementContext) {
	l.incrementComplexity().
		addOperators(c.For())
}
func (l *ComplexityListener) EnterForInStatement(c *parser.ForInStatementContext) {
	l.incrementComplexity().
		addOperators(c.For(), c.In())
}
func (l *ComplexityListener) EnterForOfStatement(c *parser.ForOfStatementContext) {
	l.incrementComplexity().
		addOperators(c.For(), c.Await(), c.Of())
}
func (l *ComplexityListener) EnterVarModifier(c *parser.VarModifierContext) {
	l.addOperators(c.Var(), c.Const())
}
func (l *ComplexityListener) EnterContinueStatement(c *parser.ContinueStatementContext) {
	l.addOperators(c.Continue())
}
func (l *ComplexityListener) EnterBreakStatement(c *parser.BreakStatementContext) {
	l.addOperators(c.Break())
}
func (l *ComplexityListener) EnterReturnStatement(c *parser.ReturnStatementContext) {
	l.addOperators(c.Return())
}
func (l *ComplexityListener) EnterYieldStatement(c *parser.YieldStatementContext) {
	l.addOperators(c.Yield(), c.YieldStar())
}
func (l *ComplexityListener) EnterWithStatement(c *parser.WithStatementContext) {
	l.addOperators(c.With())
}
func (l *ComplexityListener) EnterSwitchStatement(c *parser.SwitchStatementContext) {
	l.addOperators(c.Switch())
}
func (l *ComplexityListener) EnterCaseBlock(c *parser.CaseBlockContext) {
	l.addOperators(c.OpenBrace())
}
func (l *ComplexityListener) EnterCaseClause(c *parser.CaseClauseContext) {
	l.incrementComplexity().addOperators(c.Case(), c.Colon())
}
func (l *ComplexityListener) EnterDefaultClause(c *parser.DefaultClauseContext) {
	l.incrementComplexity().addOperators(c.Default(), c.Colon())
}
func (l *ComplexityListener) EnterLabelledStatement(c *parser.LabelledStatementContext) {
	l.addOperators(c.Colon())
}
func (l *ComplexityListener) EnterThrowStatement(c *parser.ThrowStatementContext) {
	l.addOperators(c.Throw())
}
func (l *ComplexityListener) EnterTryStatement(c *parser.TryStatementContext) {
	l.addOperators(c.Try())
}
func (l *ComplexityListener) EnterCatchProduction(c *parser.CatchProductionContext) {
	l.incrementComplexity().
		addOperators(c.Catch())
}
func (l *ComplexityListener) EnterFinallyProduction(c *parser.FinallyProductionContext) {
	l.addOperators(c.Finally())
}
func (l *ComplexityListener) EnterFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.pushScope(c.Identifier(), c).
		addOperators(c.Async(), c.Function_(), c.Multiply())
}
func (l *ComplexityListener) EnterClassDeclaration(c *parser.ClassDeclarationContext) {
	l.setClass(c.Identifier()).
		addOperators(c.Class())
}
func (l *ComplexityListener) EnterClassTail(c *parser.ClassTailContext) {
	l.addOperators(c.Extends(), c.OpenBrace())
}
func (l *ComplexityListener) EnterClassElement(c *parser.ClassElementContext) {
	l.addOperators(c.Static())
}
func (l *ComplexityListener) EnterMethodDefinition(c *parser.MethodDefinitionContext) {
	l.pushScope(c.ClassElementName(), c).
		addOperators(c.Async(), c.Multiply())
}
func (l *ComplexityListener) EnterPrivateIdentifier(c *parser.PrivateIdentifierContext) {
	l.addOperators(c.Hashtag())
}
func (l *ComplexityListener) EnterFormalParameterList(c *parser.FormalParameterListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *ComplexityListener) EnterFormalParameterArg(c *parser.FormalParameterArgContext) {
	l.incrementParams().
		addOperators(c.Assign())
}
func (l *ComplexityListener) EnterLastFormalParameterArg(c *parser.LastFormalParameterArgContext) {
	l.incrementParams().
		addOperators(c.Ellipsis())
}
func (l *ComplexityListener) EnterFunctionBody(c *parser.FunctionBodyContext) {
	l.addOperators(c.OpenBrace())
}
func (l *ComplexityListener) EnterArrayLiteral(c *parser.ArrayLiteralContext) {
	l.addOperators(c.OpenBracket())
}
func (l *ComplexityListener) EnterElementList(c *parser.ElementListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *ComplexityListener) EnterArrayElement(c *parser.ArrayElementContext) {
	l.addOperators(c.Ellipsis())
}
func (l *ComplexityListener) EnterPropertyExpressionAssignment(c *parser.PropertyExpressionAssignmentContext) {
	l.addOperators(c.Colon())
}
func (l *ComplexityListener) EnterComputedPropertyExpressionAssignment(c *parser.ComputedPropertyExpressionAssignmentContext) {
	l.addOperators(c.OpenBracket(), c.Colon())
}
func (l *ComplexityListener) EnterFunctionProperty(c *parser.FunctionPropertyContext) {
	l.addOperators(c.Async(), c.Multiply())
}
func (l *ComplexityListener) EnterPropertyShorthand(c *parser.PropertyShorthandContext) {
	l.addOperators(c.Ellipsis())
}
func (l *ComplexityListener) EnterPropertyName(c *parser.PropertyNameContext) {
	l.addOperators(c.OpenBracket())
}
func (l *ComplexityListener) EnterArguments(c *parser.ArgumentsContext) {
	l.addOperators(c.OpenParen()).
		addOperators(c.AllComma()...)
}
func (l *ComplexityListener) EnterArgument(c *parser.ArgumentContext) {
	l.addOperators(c.Ellipsis())
}
func (l *ComplexityListener) EnterExpressionSequence(c *parser.ExpressionSequenceContext) {
	l.addOperators(c.AllComma()...)
}
func (l *ComplexityListener) EnterClassExpression(c *parser.ClassExpressionContext) {
	l.setClass(c.Identifier()).
		addOperators(c.Class())
}
func (l *ComplexityListener) EnterOptionalChainExpression(c *parser.OptionalChainExpressionContext) {
	l.incrementComplexity().
		addOperators(c.QuestionMarkDot())
}
func (l *ComplexityListener) EnterMemberIndexExpression(c *parser.MemberIndexExpressionContext) {
	l.incrementComplexity(c.QuestionMarkDot()).
		addOperators(c.OpenBracket(), c.QuestionMarkDot())
}
func (l *ComplexityListener) EnterMemberDotExpression(c *parser.MemberDotExpressionContext) {
	l.incrementComplexity(c.QuestionMark()).
		addOperators(c.QuestionMark(), c.Dot(), c.Hashtag())
}
func (l *ComplexityListener) EnterNewExpression(c *parser.NewExpressionContext) {
	l.addOperators(c.New())
}
func (l *ComplexityListener) EnterMetaExpression(c *parser.MetaExpressionContext) {
	l.addOperands(c.New()).addOperators(c.Dot())
}
func (l *ComplexityListener) EnterPostIncrementExpression(c *parser.PostIncrementExpressionContext) {
	l.addOperators(c.PlusPlus())
}
func (l *ComplexityListener) EnterPostDecreaseExpression(c *parser.PostDecreaseExpressionContext) {
	l.addOperators(c.MinusMinus())
}
func (l *ComplexityListener) EnterDeleteExpression(c *parser.DeleteExpressionContext) {
	l.addOperators(c.Delete())
}
func (l *ComplexityListener) EnterVoidExpression(c *parser.VoidExpressionContext) {
	l.addOperators(c.Void())
}
func (l *ComplexityListener) EnterTypeofExpression(c *parser.TypeofExpressionContext) {
	l.addOperators(c.Typeof())
}
func (l *ComplexityListener) EnterPreIncrementExpression(c *parser.PreIncrementExpressionContext) {
	l.addOperators(c.PlusPlus())
}
func (l *ComplexityListener) EnterPreDecreaseExpression(c *parser.PreDecreaseExpressionContext) {
	l.addOperators(c.MinusMinus())
}
func (l *ComplexityListener) EnterUnaryPlusExpression(c *parser.UnaryPlusExpressionContext) {
	l.addOperators(c.Plus())
}
func (l *ComplexityListener) EnterUnaryMinusExpression(c *parser.UnaryMinusExpressionContext) {
	l.addOperators(c.Minus())
}
func (l *ComplexityListener) EnterBitNotExpression(c *parser.BitNotExpressionContext) {
	l.addOperators(c.BitNot())
}
func (l *ComplexityListener) EnterNotExpression(c *parser.NotExpressionContext) {
	l.addOperators(c.Not())
}
func (l *ComplexityListener) EnterAwaitExpression(c *parser.AwaitExpressionContext) {
	l.addOperators(c.Await())
}
func (l *ComplexityListener) EnterPowerExpression(c *parser.PowerExpressionContext) {
	l.addOperators(c.Power())
}
func (l *ComplexityListener) EnterMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	l.addOperators(c.Divide(), c.Multiply(), c.Modulus())
}
func (l *ComplexityListener) EnterAdditiveExpression(c *parser.AdditiveExpressionContext) {
	l.addOperators(c.Plus(), c.Minus())
}
func (l *ComplexityListener) EnterCoalesceExpression(c *parser.CoalesceExpressionContext) {
	l.incrementComplexity().
		addOperators(c.NullCoalesce())
}
func (l *ComplexityListener) EnterBitShiftExpression(c *parser.BitShiftExpressionContext) {
	l.addOperators(c.LeftShiftArithmetic(), c.RightShiftArithmetic(), c.RightShiftLogical())
}
func (l *ComplexityListener) EnterRelationalExpression(c *parser.RelationalExpressionContext) {
	l.addOperators(c.LessThan(), c.MoreThan(), c.LessThanEquals(), c.GreaterThanEquals())
}
func (l *ComplexityListener) EnterInstanceofExpression(c *parser.InstanceofExpressionContext) {
	l.addOperators(c.Instanceof())
}
func (l *ComplexityListener) EnterInExpression(c *parser.InExpressionContext) {
	l.addOperators(c.In())
}
func (l *ComplexityListener) EnterEqualityExpression(c *parser.EqualityExpressionContext) {
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
func (l *ComplexityListener) EnterBitAndExpression(c *parser.BitAndExpressionContext) {
	l.addOperands(c.BitAnd())
}
func (l *ComplexityListener) EnterBitXOrExpression(c *parser.BitXOrExpressionContext) {
	l.addOperators(c.BitXOr())
}
func (l *ComplexityListener) EnterBitOrExpression(c *parser.BitOrExpressionContext) {
	l.addOperands(c.BitOr())
}
func (l *ComplexityListener) EnterLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	l.addOperators(c.And())
}
func (l *ComplexityListener) EnterLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	l.incrementComplexity().
		addOperators(c.Or())
}
func (l *ComplexityListener) EnterTernaryExpression(c *parser.TernaryExpressionContext) {
	l.incrementComplexity().
		addOperators(c.QuestionMark(), c.Colon())
}
func (l *ComplexityListener) EnterAssignmentExpression(c *parser.AssignmentExpressionContext) {
	l.setAssignee(c.SingleExpression(0)).addOperators(c.Assign())
}
func (l *ComplexityListener) EnterAssignmentOperator(c *parser.AssignmentOperatorContext) {
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
func (l *ComplexityListener) EnterImportExpression(c *parser.ImportExpressionContext) {
	l.addOperators(c.Import())
}
func (l *ComplexityListener) EnterThisExpression(c *parser.ThisExpressionContext) {
	l.addOperands(c.This())
}
func (l *ComplexityListener) EnterSuperExpression(c *parser.SuperExpressionContext) {
	l.addOperators(c.Super())
}
func (l *ComplexityListener) EnterParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {
	l.addOperators(c.OpenParen())
}
func (l *ComplexityListener) EnterInitializer(c *parser.InitializerContext) {
	l.addOperators(c.Assign())
}
func (l *ComplexityListener) EnterObjectLiteral(c *parser.ObjectLiteralContext) {
	l.addOperators(c.AllComma()...).
		addOperators(c.OpenBrace())
}
func (l *ComplexityListener) EnterAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.pushScope(l.scopes[0].assignee, c).
		addOperators(c.Async(), c.Function_(), c.Multiply())
}
func (l *ComplexityListener) EnterArrowFunction(c *parser.ArrowFunctionContext) {
	if c.ArrowFunctionBody().FunctionBody() != nil {
		l.pushScope(l.scopes[0].assignee, c).
			addOperators(c.Async(), c.ARROW())
	}
}
func (l *ComplexityListener) EnterLiteral(c *parser.LiteralContext) {
	l.addOperands(
		c.NullLiteral(),
		c.BooleanLiteral(),
		c.StringLiteral(),
		c.RegularExpressionLiteral(),
	)
}
func (l *ComplexityListener) EnterTemplateStringLiteral(c *parser.TemplateStringLiteralContext) {
	l.addOperands(c.AllBackTick()...)
}

func (l *ComplexityListener) EnterTemplateStringAtom(c *parser.TemplateStringAtomContext) {
	l.addOperands(
		c.TemplateStringAtom(),
		c.TemplateStringStartExpression(),
		c.TemplateCloseBrace(),
	)
}
func (l *ComplexityListener) EnterNumericLiteral(c *parser.NumericLiteralContext) {
	l.addOperands(
		c.DecimalLiteral(),
		c.HexIntegerLiteral(),
		c.OctalIntegerLiteral(),
		c.OctalIntegerLiteral2(),
		c.BinaryIntegerLiteral(),
	)
}
func (l *ComplexityListener) EnterBigintLiteral(c *parser.BigintLiteralContext) {
	l.addOperands(
		c.BigDecimalIntegerLiteral(),
		c.BigHexIntegerLiteral(),
		c.BigOctalIntegerLiteral(),
		c.BigBinaryIntegerLiteral(),
	)
}
func (l *ComplexityListener) EnterIdentifier(c *parser.IdentifierContext) {
	l.addOperands(c.Identifier())
}
func (l *ComplexityListener) EnterLet_(c *parser.Let_Context) {
	l.addOperators(c.NonStrictLet(), c.StrictLet())
}
func (l *ComplexityListener) EnterEos(c *parser.EosContext) {
	l.addOperators(c.SemiColon())
}

// EXIT

func (l *ComplexityListener) ExitProgram(c *parser.ProgramContext) {
	l.popScope(c)
}

// func (l *ComplexityListener) ExitSourceElement(c *parser.SourceElementContext) {}
// func (l *ComplexityListener) ExitStatement(c *parser.StatementContext) {}
// func (l *ComplexityListener) ExitBlock(c *parser.BlockContext) {}
// func (l *ComplexityListener) ExitStatementList(c *parser.StatementListContext) {}
// func (l *ComplexityListener) ExitImportStatement(c *parser.ImportStatementContext) {}
// func (l *ComplexityListener) ExitImportFromBlock(c *parser.ImportFromBlockContext) {}
// func (l *ComplexityListener) ExitImportModuleItems(c *parser.ImportModuleItemsContext) {}
// func (l *ComplexityListener) ExitImportAliasName(c *parser.ImportAliasNameContext) {}
// func (l *ComplexityListener) ExitModuleExportName(c *parser.ModuleExportNameContext) {}
// func (l *ComplexityListener) ExitImportedBinding(c *parser.ImportedBindingContext) {}
// func (l *ComplexityListener) ExitImportDefault(c *parser.ImportDefaultContext) {}
// func (l *ComplexityListener) ExitImportNamespace(c *parser.ImportNamespaceContext) {}
// func (l *ComplexityListener) ExitImportFrom(c *parser.ImportFromContext) {}
// func (l *ComplexityListener) ExitAliasName(c *parser.AliasNameContext) {}
// func (l *ComplexityListener) ExitExportDeclaration(c *parser.ExportDeclarationContext) {}
// func (l *ComplexityListener) ExitExportDefaultDeclaration(c *parser.ExportDefaultDeclarationContext) {}
// func (l *ComplexityListener) ExitExportFromBlock(c *parser.ExportFromBlockContext) {}
// func (l *ComplexityListener) ExitExportModuleItems(c *parser.ExportModuleItemsContext) {}
// func (l *ComplexityListener) ExitExportAliasName(c *parser.ExportAliasNameContext) {}
// func (l *ComplexityListener) ExitDeclaration(c *parser.DeclarationContext) {}
// func (l *ComplexityListener) ExitVariableStatement(c *parser.VariableStatementContext) {}
// func (l *ComplexityListener) ExitVariableDeclarationList(c *parser.VariableDeclarationListContext) {}
func (l *ComplexityListener) ExitVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.setAssignee(nil)
}

// func (l *ComplexityListener) ExitEmptyStatement_(c *parser.EmptyStatement_Context) {}
// func (l *ComplexityListener) ExitExpressionStatement(c *parser.ExpressionStatementContext) {}
// func (l *ComplexityListener) ExitIfStatement(c *parser.IfStatementContext) {}
// func (l *ComplexityListener) ExitDoStatement(c *parser.DoStatementContext) {}
// func (l *ComplexityListener) ExitWhileStatement(c *parser.WhileStatementContext) {}
// func (l *ComplexityListener) ExitForStatement(c *parser.ForStatementContext) {}
// func (l *ComplexityListener) ExitForInStatement(c *parser.ForInStatementContext) {}
// func (l *ComplexityListener) ExitForOfStatement(c *parser.ForOfStatementContext) {}
// func (l *ComplexityListener) ExitVarModifier(c *parser.VarModifierContext) {}
// func (l *ComplexityListener) ExitContinueStatement(c *parser.ContinueStatementContext) {}
// func (l *ComplexityListener) ExitBreakStatement(c *parser.BreakStatementContext) {}
// func (l *ComplexityListener) ExitReturnStatement(c *parser.ReturnStatementContext) {}
// func (l *ComplexityListener) ExitYieldStatement(c *parser.YieldStatementContext) {}
// func (l *ComplexityListener) ExitWithStatement(c *parser.WithStatementContext) {}
// func (l *ComplexityListener) ExitSwitchStatement(c *parser.SwitchStatementContext) {}
// func (l *ComplexityListener) ExitCaseBlock(c *parser.CaseBlockContext) {}
// func (l *ComplexityListener) ExitCaseClauses(c *parser.CaseClausesContext) {}
// func (l *ComplexityListener) ExitCaseClause(c *parser.CaseClauseContext) {}
// func (l *ComplexityListener) ExitDefaultClause(c *parser.DefaultClauseContext) {}
// func (l *ComplexityListener) ExitLabelledStatement(c *parser.LabelledStatementContext) {}
// func (l *ComplexityListener) ExitThrowStatement(c *parser.ThrowStatementContext) {}
// func (l *ComplexityListener) ExitTryStatement(c *parser.TryStatementContext) {}
// func (l *ComplexityListener) ExitCatchProduction(c *parser.CatchProductionContext) {}
// func (l *ComplexityListener) ExitFinallyProduction(c *parser.FinallyProductionContext) {}
// func (l *ComplexityListener) ExitDebuggerStatement(c *parser.DebuggerStatementContext) {}
func (l *ComplexityListener) ExitFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.popScope(c)
}

// func (l *ComplexityListener) ExitClassDeclaration(c *parser.ClassDeclarationContext) {}
func (l *ComplexityListener) ExitClassTail(c *parser.ClassTailContext) {
	l.resetClass()
}

// func (l *ComplexityListener) ExitClassElement(c *parser.ClassElementContext) {}
func (l *ComplexityListener) ExitMethodDefinition(c *parser.MethodDefinitionContext) {
	l.popScope(c)
}

// func (l *ComplexityListener) ExitFieldDefinition(c *parser.FieldDefinitionContext) {}
// func (l *ComplexityListener) ExitClassElementName(c *parser.ClassElementNameContext) {}
// func (l *ComplexityListener) ExitPrivateIdentifier(c *parser.PrivateIdentifierContext) {}
// func (l *ComplexityListener) ExitFormalParameterList(c *parser.FormalParameterListContext) {}
// func (l *ComplexityListener) ExitFormalParameterArg(c *parser.FormalParameterArgContext) {}
// func (l *ComplexityListener) ExitLastFormalParameterArg(c *parser.LastFormalParameterArgContext) {}
// func (l *ComplexityListener) ExitFunctionBody(c *parser.FunctionBodyContext) {}
// func (l *ComplexityListener) ExitSourceElements(c *parser.SourceElementsContext) {}
// func (l *ComplexityListener) ExitArrayLiteral(c *parser.ArrayLiteralContext) {}
// func (l *ComplexityListener) ExitElementList(c *parser.ElementListContext) {}
// func (l *ComplexityListener) ExitArrayElement(c *parser.ArrayElementContext) {}
// func (l *ComplexityListener) ExitPropertyExpressionAssignment(c *parser.PropertyExpressionAssignmentContext) {}
// func (l *ComplexityListener) ExitComputedPropertyExpressionAssignment(c *parser.ComputedPropertyExpressionAssignmentContext) {}
// func (l *ComplexityListener) ExitFunctionProperty(c *parser.FunctionPropertyContext) {}
// func (l *ComplexityListener) ExitPropertyGetter(c *parser.PropertyGetterContext) {}
// func (l *ComplexityListener) ExitPropertySetter(c *parser.PropertySetterContext) {}
// func (l *ComplexityListener) ExitPropertyShorthand(c *parser.PropertyShorthandContext) {}
// func (l *ComplexityListener) ExitPropertyName(c *parser.PropertyNameContext) {}
// func (l *ComplexityListener) ExitArguments(c *parser.ArgumentsContext) {}
// func (l *ComplexityListener) ExitArgument(c *parser.ArgumentContext) {}
// func (l *ComplexityListener) ExitExpressionSequence(c *parser.ExpressionSequenceContext) {}
// func (l *ComplexityListener) ExitTemplateStringExpression(c *parser.TemplateStringExpressionContext) {}
// func (l *ComplexityListener) ExitTernaryExpression(c *parser.TernaryExpressionContext) {}
// func (l *ComplexityListener) ExitLogicalAndExpression(c *parser.LogicalAndExpressionContext) {}
// func (l *ComplexityListener) ExitPowerExpression(c *parser.PowerExpressionContext) {}
// func (l *ComplexityListener) ExitPreIncrementExpression(c *parser.PreIncrementExpressionContext) {}
// func (l *ComplexityListener) ExitObjectLiteralExpression(c *parser.ObjectLiteralExpressionContext) {}
// func (l *ComplexityListener) ExitMetaExpression(c *parser.MetaExpressionContext) {}
// func (l *ComplexityListener) ExitInExpression(c *parser.InExpressionContext) {}
// func (l *ComplexityListener) ExitLogicalOrExpression(c *parser.LogicalOrExpressionContext) {}
// func (l *ComplexityListener) ExitOptionalChainExpression(c *parser.OptionalChainExpressionContext) {}
// func (l *ComplexityListener) ExitNotExpression(c *parser.NotExpressionContext) {}
// func (l *ComplexityListener) ExitPreDecreaseExpression(c *parser.PreDecreaseExpressionContext) {}
// func (l *ComplexityListener) ExitArgumentsExpression(c *parser.ArgumentsExpressionContext) {}
// func (l *ComplexityListener) ExitAwaitExpression(c *parser.AwaitExpressionContext) {}
// func (l *ComplexityListener) ExitThisExpression(c *parser.ThisExpressionContext) {}
// func (l *ComplexityListener) ExitFunctionExpression(c *parser.FunctionExpressionContext) {}
// func (l *ComplexityListener) ExitUnaryMinusExpression(c *parser.UnaryMinusExpressionContext) {}
func (l *ComplexityListener) ExitAssignmentExpression(c *parser.AssignmentExpressionContext) {
	l.setAssignee(nil)
}

// func (l *ComplexityListener) ExitPostDecreaseExpression(c *parser.PostDecreaseExpressionContext) {}
// func (l *ComplexityListener) ExitTypeofExpression(c *parser.TypeofExpressionContext) {}
// func (l *ComplexityListener) ExitInstanceofExpression(c *parser.InstanceofExpressionContext) {}
// func (l *ComplexityListener) ExitUnaryPlusExpression(c *parser.UnaryPlusExpressionContext) {}
// func (l *ComplexityListener) ExitDeleteExpression(c *parser.DeleteExpressionContext) {}
// func (l *ComplexityListener) ExitImportExpression(c *parser.ImportExpressionContext) {}
// func (l *ComplexityListener) ExitEqualityExpression(c *parser.EqualityExpressionContext) {}
// func (l *ComplexityListener) ExitBitXOrExpression(c *parser.BitXOrExpressionContext) {}
// func (l *ComplexityListener) ExitSuperExpression(c *parser.SuperExpressionContext) {}
// func (l *ComplexityListener) ExitMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {}
// func (l *ComplexityListener) ExitBitShiftExpression(c *parser.BitShiftExpressionContext) {}
// func (l *ComplexityListener) ExitParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {}
// func (l *ComplexityListener) ExitAdditiveExpression(c *parser.AdditiveExpressionContext) {}
// func (l *ComplexityListener) ExitRelationalExpression(c *parser.RelationalExpressionContext) {}
// func (l *ComplexityListener) ExitPostIncrementExpression(c *parser.PostIncrementExpressionContext) {}
// func (l *ComplexityListener) ExitYieldExpression(c *parser.YieldExpressionContext) {}
// func (l *ComplexityListener) ExitBitNotExpression(c *parser.BitNotExpressionContext) {}
// func (l *ComplexityListener) ExitNewExpression(c *parser.NewExpressionContext) {}
// func (l *ComplexityListener) ExitLiteralExpression(c *parser.LiteralExpressionContext) {}
// func (l *ComplexityListener) ExitArrayLiteralExpression(c *parser.ArrayLiteralExpressionContext) {}
// func (l *ComplexityListener) ExitMemberDotExpression(c *parser.MemberDotExpressionContext) {}
// func (l *ComplexityListener) ExitClassExpression(c *parser.ClassExpressionContext) {}
// func (l *ComplexityListener) ExitMemberIndexExpression(c *parser.MemberIndexExpressionContext) {}
// func (l *ComplexityListener) ExitIdentifierExpression(c *parser.IdentifierExpressionContext) {}
// func (l *ComplexityListener) ExitBitAndExpression(c *parser.BitAndExpressionContext) {}
// func (l *ComplexityListener) ExitBitOrExpression(c *parser.BitOrExpressionContext) {}
// func (l *ComplexityListener) ExitAssignmentOperatorExpression(c *parser.AssignmentOperatorExpressionContext) {}
// func (l *ComplexityListener) ExitVoidExpression(c *parser.VoidExpressionContext) {}
// func (l *ComplexityListener) ExitCoalesceExpression(c *parser.CoalesceExpressionContext) {}
// func (l *ComplexityListener) ExitInitializer(c *parser.InitializerContext) {}
// func (l *ComplexityListener) ExitAssignable(c *parser.AssignableContext) {}
// func (l *ComplexityListener) ExitObjectLiteral(c *parser.ObjectLiteralContext) {}
// func (l *ComplexityListener) ExitNamedFunction(c *parser.NamedFunctionContext) {}
func (l *ComplexityListener) ExitAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.popScope(c)
}

// func (l *ComplexityListener) ExitArrowFunction(c *parser.ArrowFunctionContext) {}
// func (l *ComplexityListener) ExitArrowFunctionParameters(c *parser.ArrowFunctionParametersContext) {}
func (l *ComplexityListener) ExitArrowFunctionBody(c *parser.ArrowFunctionBodyContext) {
	if c.FunctionBody() != nil {
		l.popScope(c)
	}
}

// func (l *ComplexityListener) ExitAssignmentOperator(c *parser.AssignmentOperatorContext) {}
// func (l *ComplexityListener) ExitLiteral(c *parser.LiteralContext) {}
// func (l *ComplexityListener) ExitTemplateStringLiteral(c *parser.TemplateStringLiteralContext) {}
// func (l *ComplexityListener) ExitTemplateStringAtom(c *parser.TemplateStringAtomContext) {}
// func (l *ComplexityListener) ExitNumericLiteral(c *parser.NumericLiteralContext) {}
// func (l *ComplexityListener) ExitBigintLiteral(c *parser.BigintLiteralContext) {}
// func (l *ComplexityListener) ExitGetter(c *parser.GetterContext) {}
// func (l *ComplexityListener) ExitSetter(c *parser.SetterContext) {}
// func (l *ComplexityListener) ExitIdentifierName(c *parser.IdentifierNameContext) {}
// func (l *ComplexityListener) ExitIdentifier(c *parser.IdentifierContext) {}
// func (l *ComplexityListener) ExitReservedWord(c *parser.ReservedWordContext) {}
// func (l *ComplexityListener) ExitKeyword(c *parser.KeywordContext) {}
// func (l *ComplexityListener) ExitLet_(c *parser.Let_Context) {}
// func (l *ComplexityListener) ExitEos(c *parser.EosContext) {}
