package analyser

import (
	"figsyntax/internal/logger"
	"figsyntax/internal/parser/figscript"
	"log/slog"
)

type FigScriptComplexityListener struct {
	figscript.FigScriptParserListener
	*Reporter
}

func NewFigScriptComplexityListener(logger *slog.Logger) *FigScriptComplexityListener {
	return &FigScriptComplexityListener{
		new(figscript.BaseFigScriptParserListener),
		NewReporter(logger),
	}
}

func (l *FigScriptComplexityListener) WithFile(path string) ComplexityListener {
	logger.Trace()
	l.scopes = make([]*Scope, 0)
	l.report = NewReport(path)
	l.path = path
	return l
}

func (l *FigScriptComplexityListener) pushProgram(ctx *figscript.ProgramContext) *FigScriptComplexityListener {
	logger.Trace()
	l.scopes = append([]*Scope{NewScope(l.path, "program", ctx)}, l.scopes...)
	return l
}

func (l *FigScriptComplexityListener) EnterProgram(c *figscript.ProgramContext) {
	l.pushProgram(c)
}
func (l *FigScriptComplexityListener) EnterStatement(c *figscript.StatementContext) {
	switch {
	case c.Block() != nil:
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
		l.pushScope(l.scopes[0].assignee, c)
	}
	l.addOperators(c.Async(), c.ARROW())
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
func (l *FigScriptComplexityListener) ExitProgram(c *figscript.ProgramContext) {
	l.popScope(c)
}
func (l *FigScriptComplexityListener) ExitVariableDeclaration(c *figscript.VariableDeclarationContext) {
	l.setAssignee(nil)
}
func (l *FigScriptComplexityListener) ExitFunctionDeclaration(c *figscript.FunctionDeclarationContext) {
	l.popScope(c)
}
func (l *FigScriptComplexityListener) ExitClassTail(c *figscript.ClassTailContext) {
	l.resetClass()
}
func (l *FigScriptComplexityListener) ExitMethodDefinition(c *figscript.MethodDefinitionContext) {
	l.popScope(c)
}
func (l *FigScriptComplexityListener) ExitAssignmentExpression(c *figscript.AssignmentExpressionContext) {
	l.setAssignee(nil)
}
func (l *FigScriptComplexityListener) ExitAnonymousFunctionDecl(c *figscript.AnonymousFunctionDeclContext) {
	l.popScope(c)
}
func (l *FigScriptComplexityListener) ExitArrowFunctionBody(c *figscript.ArrowFunctionBodyContext) {
	if c.FunctionBody() != nil {
		l.popScope(c)
	}
}
