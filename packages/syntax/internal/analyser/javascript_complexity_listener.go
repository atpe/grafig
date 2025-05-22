package analyser

import (
	"figsyntax/internal/logger"
	"figsyntax/internal/parser"
	"log/slog"
)

type JavaScriptComplexityListener struct {
	parser.JavaScriptParserListener
	*Reporter
}

func NewJavaScriptComplexityListener(logger *slog.Logger) *JavaScriptComplexityListener {
	return &JavaScriptComplexityListener{
		new(parser.BaseJavaScriptParserListener),
		NewReporter(logger),
	}
}

func (l *JavaScriptComplexityListener) WithFile(path string) ComplexityListener {
	logger.Trace()
	l.scopes = make([]*Scope, 0)
	l.report = NewReport(path)
	l.path = path
	return l
}

func (l *JavaScriptComplexityListener) pushProgram(ctx *parser.ProgramContext) *JavaScriptComplexityListener {
	logger.Trace()
	l.scopes = append([]*Scope{NewScope(l.path, "program", ctx)}, l.scopes...)
	return l
}

func (l *JavaScriptComplexityListener) EnterProgram(c *parser.ProgramContext) {
	l.pushProgram(c)
}
func (l *JavaScriptComplexityListener) EnterStatement(c *parser.StatementContext) {
	switch {
	case c.Block() != nil:
	default:
		l.incrementLloc()
	}
}
func (l *JavaScriptComplexityListener) EnterBlock(c *parser.BlockContext) {
	l.addOperators(c.OpenBrace())
}
func (l *JavaScriptComplexityListener) EnterVariableDeclarationList(c *parser.VariableDeclarationListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *JavaScriptComplexityListener) EnterVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.setAssignee(c.Assignable()).
		addOperators(c.Assign())
}
func (l *JavaScriptComplexityListener) EnterEmptyStatement_(c *parser.EmptyStatement_Context) {
	l.addOperators(c.SemiColon())
}
func (l *JavaScriptComplexityListener) EnterIfStatement(c *parser.IfStatementContext) {
	l.incrementLloc(c.Else()).
		incrementComplexity().
		addOperators(c.If(), c.Else())
}
func (l *JavaScriptComplexityListener) EnterDoStatement(c *parser.DoStatementContext) {
	l.incrementComplexity().
		addOperators(c.Do(), c.While())
}
func (l *JavaScriptComplexityListener) EnterWhileStatement(c *parser.WhileStatementContext) {
	l.incrementComplexity().
		addOperators(c.While())
}
func (l *JavaScriptComplexityListener) EnterForStatement(c *parser.ForStatementContext) {
	l.incrementComplexity().
		addOperators(c.For())
}
func (l *JavaScriptComplexityListener) EnterForInStatement(c *parser.ForInStatementContext) {
	l.incrementComplexity().
		addOperators(c.For(), c.In())
}
func (l *JavaScriptComplexityListener) EnterForOfStatement(c *parser.ForOfStatementContext) {
	l.incrementComplexity().
		addOperators(c.For(), c.Await(), c.Of())
}
func (l *JavaScriptComplexityListener) EnterVarModifier(c *parser.VarModifierContext) {
	l.addOperators(c.Var(), c.Const())
}
func (l *JavaScriptComplexityListener) EnterContinueStatement(c *parser.ContinueStatementContext) {
	l.addOperators(c.Continue())
}
func (l *JavaScriptComplexityListener) EnterBreakStatement(c *parser.BreakStatementContext) {
	l.addOperators(c.Break())
}
func (l *JavaScriptComplexityListener) EnterReturnStatement(c *parser.ReturnStatementContext) {
	l.addOperators(c.Return())
}
func (l *JavaScriptComplexityListener) EnterYieldStatement(c *parser.YieldStatementContext) {
	l.addOperators(c.Yield(), c.YieldStar())
}
func (l *JavaScriptComplexityListener) EnterWithStatement(c *parser.WithStatementContext) {
	l.addOperators(c.With())
}
func (l *JavaScriptComplexityListener) EnterSwitchStatement(c *parser.SwitchStatementContext) {
	l.addOperators(c.Switch())
}
func (l *JavaScriptComplexityListener) EnterCaseBlock(c *parser.CaseBlockContext) {
	l.addOperators(c.OpenBrace())
}
func (l *JavaScriptComplexityListener) EnterCaseClause(c *parser.CaseClauseContext) {
	l.incrementComplexity().addOperators(c.Case(), c.Colon())
}
func (l *JavaScriptComplexityListener) EnterDefaultClause(c *parser.DefaultClauseContext) {
	l.incrementComplexity().addOperators(c.Default(), c.Colon())
}
func (l *JavaScriptComplexityListener) EnterLabelledStatement(c *parser.LabelledStatementContext) {
	l.addOperators(c.Colon())
}
func (l *JavaScriptComplexityListener) EnterThrowStatement(c *parser.ThrowStatementContext) {
	l.addOperators(c.Throw())
}
func (l *JavaScriptComplexityListener) EnterTryStatement(c *parser.TryStatementContext) {
	l.addOperators(c.Try())
}
func (l *JavaScriptComplexityListener) EnterCatchProduction(c *parser.CatchProductionContext) {
	l.incrementComplexity().
		addOperators(c.Catch())
}
func (l *JavaScriptComplexityListener) EnterFinallyProduction(c *parser.FinallyProductionContext) {
	l.addOperators(c.Finally())
}
func (l *JavaScriptComplexityListener) EnterFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.pushScope(c.Identifier(), c).
		addOperators(c.Async(), c.Function_(), c.Multiply())
}
func (l *JavaScriptComplexityListener) EnterClassDeclaration(c *parser.ClassDeclarationContext) {
	l.setClass(c.Identifier()).
		addOperators(c.Class())
}
func (l *JavaScriptComplexityListener) EnterClassTail(c *parser.ClassTailContext) {
	l.addOperators(c.Extends(), c.OpenBrace())
}
func (l *JavaScriptComplexityListener) EnterClassElement(c *parser.ClassElementContext) {
	l.addOperators(c.Static())
}
func (l *JavaScriptComplexityListener) EnterMethodDefinition(c *parser.MethodDefinitionContext) {
	l.pushScope(c.ClassElementName(), c).
		addOperators(c.Async(), c.Multiply())
}
func (l *JavaScriptComplexityListener) EnterPrivateIdentifier(c *parser.PrivateIdentifierContext) {
	l.addOperators(c.Hashtag())
}
func (l *JavaScriptComplexityListener) EnterFormalParameterList(c *parser.FormalParameterListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *JavaScriptComplexityListener) EnterFormalParameterArg(c *parser.FormalParameterArgContext) {
	l.incrementParams().
		addOperators(c.Assign())
}
func (l *JavaScriptComplexityListener) EnterLastFormalParameterArg(c *parser.LastFormalParameterArgContext) {
	l.incrementParams().
		addOperators(c.Ellipsis())
}
func (l *JavaScriptComplexityListener) EnterFunctionBody(c *parser.FunctionBodyContext) {
	l.addOperators(c.OpenBrace())
}
func (l *JavaScriptComplexityListener) EnterArrayLiteral(c *parser.ArrayLiteralContext) {
	l.addOperators(c.OpenBracket())
}
func (l *JavaScriptComplexityListener) EnterElementList(c *parser.ElementListContext) {
	l.addOperators(c.AllComma()...)
}
func (l *JavaScriptComplexityListener) EnterArrayElement(c *parser.ArrayElementContext) {
	l.addOperators(c.Ellipsis())
}
func (l *JavaScriptComplexityListener) EnterPropertyExpressionAssignment(c *parser.PropertyExpressionAssignmentContext) {
	l.addOperators(c.Colon())
}
func (l *JavaScriptComplexityListener) EnterComputedPropertyExpressionAssignment(c *parser.ComputedPropertyExpressionAssignmentContext) {
	l.addOperators(c.OpenBracket(), c.Colon())
}
func (l *JavaScriptComplexityListener) EnterFunctionProperty(c *parser.FunctionPropertyContext) {
	l.addOperators(c.Async(), c.Multiply())
}
func (l *JavaScriptComplexityListener) EnterPropertyShorthand(c *parser.PropertyShorthandContext) {
	l.addOperators(c.Ellipsis())
}
func (l *JavaScriptComplexityListener) EnterPropertyName(c *parser.PropertyNameContext) {
	l.addOperators(c.OpenBracket())
}
func (l *JavaScriptComplexityListener) EnterArguments(c *parser.ArgumentsContext) {
	l.addOperators(c.OpenParen()).
		addOperators(c.AllComma()...)
}
func (l *JavaScriptComplexityListener) EnterArgument(c *parser.ArgumentContext) {
	l.addOperators(c.Ellipsis())
}
func (l *JavaScriptComplexityListener) EnterExpressionSequence(c *parser.ExpressionSequenceContext) {
	l.addOperators(c.AllComma()...)
}
func (l *JavaScriptComplexityListener) EnterClassExpression(c *parser.ClassExpressionContext) {
	l.setClass(c.Identifier()).
		addOperators(c.Class())
}
func (l *JavaScriptComplexityListener) EnterOptionalChainExpression(c *parser.OptionalChainExpressionContext) {
	l.incrementComplexity().
		addOperators(c.QuestionMarkDot())
}
func (l *JavaScriptComplexityListener) EnterMemberIndexExpression(c *parser.MemberIndexExpressionContext) {
	l.incrementComplexity(c.QuestionMarkDot()).
		addOperators(c.OpenBracket(), c.QuestionMarkDot())
}
func (l *JavaScriptComplexityListener) EnterMemberDotExpression(c *parser.MemberDotExpressionContext) {
	l.incrementComplexity(c.QuestionMark()).
		addOperators(c.QuestionMark(), c.Dot(), c.Hashtag())
}
func (l *JavaScriptComplexityListener) EnterNewExpression(c *parser.NewExpressionContext) {
	l.addOperators(c.New())
}
func (l *JavaScriptComplexityListener) EnterMetaExpression(c *parser.MetaExpressionContext) {
	l.addOperands(c.New()).addOperators(c.Dot())
}
func (l *JavaScriptComplexityListener) EnterPostIncrementExpression(c *parser.PostIncrementExpressionContext) {
	l.addOperators(c.PlusPlus())
}
func (l *JavaScriptComplexityListener) EnterPostDecreaseExpression(c *parser.PostDecreaseExpressionContext) {
	l.addOperators(c.MinusMinus())
}
func (l *JavaScriptComplexityListener) EnterDeleteExpression(c *parser.DeleteExpressionContext) {
	l.addOperators(c.Delete())
}
func (l *JavaScriptComplexityListener) EnterVoidExpression(c *parser.VoidExpressionContext) {
	l.addOperators(c.Void())
}
func (l *JavaScriptComplexityListener) EnterTypeofExpression(c *parser.TypeofExpressionContext) {
	l.addOperators(c.Typeof())
}
func (l *JavaScriptComplexityListener) EnterPreIncrementExpression(c *parser.PreIncrementExpressionContext) {
	l.addOperators(c.PlusPlus())
}
func (l *JavaScriptComplexityListener) EnterPreDecreaseExpression(c *parser.PreDecreaseExpressionContext) {
	l.addOperators(c.MinusMinus())
}
func (l *JavaScriptComplexityListener) EnterUnaryPlusExpression(c *parser.UnaryPlusExpressionContext) {
	l.addOperators(c.Plus())
}
func (l *JavaScriptComplexityListener) EnterUnaryMinusExpression(c *parser.UnaryMinusExpressionContext) {
	l.addOperators(c.Minus())
}
func (l *JavaScriptComplexityListener) EnterBitNotExpression(c *parser.BitNotExpressionContext) {
	l.addOperators(c.BitNot())
}
func (l *JavaScriptComplexityListener) EnterNotExpression(c *parser.NotExpressionContext) {
	l.addOperators(c.Not())
}
func (l *JavaScriptComplexityListener) EnterAwaitExpression(c *parser.AwaitExpressionContext) {
	l.addOperators(c.Await())
}
func (l *JavaScriptComplexityListener) EnterPowerExpression(c *parser.PowerExpressionContext) {
	l.addOperators(c.Power())
}
func (l *JavaScriptComplexityListener) EnterMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	l.addOperators(c.Divide(), c.Multiply(), c.Modulus())
}
func (l *JavaScriptComplexityListener) EnterAdditiveExpression(c *parser.AdditiveExpressionContext) {
	l.addOperators(c.Plus(), c.Minus())
}
func (l *JavaScriptComplexityListener) EnterCoalesceExpression(c *parser.CoalesceExpressionContext) {
	l.incrementComplexity().
		addOperators(c.NullCoalesce())
}
func (l *JavaScriptComplexityListener) EnterBitShiftExpression(c *parser.BitShiftExpressionContext) {
	l.addOperators(c.LeftShiftArithmetic(), c.RightShiftArithmetic(), c.RightShiftLogical())
}
func (l *JavaScriptComplexityListener) EnterRelationalExpression(c *parser.RelationalExpressionContext) {
	l.addOperators(c.LessThan(), c.MoreThan(), c.LessThanEquals(), c.GreaterThanEquals())
}
func (l *JavaScriptComplexityListener) EnterInstanceofExpression(c *parser.InstanceofExpressionContext) {
	l.addOperators(c.Instanceof())
}
func (l *JavaScriptComplexityListener) EnterInExpression(c *parser.InExpressionContext) {
	l.addOperators(c.In())
}
func (l *JavaScriptComplexityListener) EnterEqualityExpression(c *parser.EqualityExpressionContext) {
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
func (l *JavaScriptComplexityListener) EnterBitAndExpression(c *parser.BitAndExpressionContext) {
	l.addOperands(c.BitAnd())
}
func (l *JavaScriptComplexityListener) EnterBitXOrExpression(c *parser.BitXOrExpressionContext) {
	l.addOperators(c.BitXOr())
}
func (l *JavaScriptComplexityListener) EnterBitOrExpression(c *parser.BitOrExpressionContext) {
	l.addOperands(c.BitOr())
}
func (l *JavaScriptComplexityListener) EnterLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	l.addOperators(c.And())
}
func (l *JavaScriptComplexityListener) EnterLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	l.incrementComplexity().
		addOperators(c.Or())
}
func (l *JavaScriptComplexityListener) EnterTernaryExpression(c *parser.TernaryExpressionContext) {
	l.incrementComplexity().
		addOperators(c.QuestionMark(), c.Colon())
}
func (l *JavaScriptComplexityListener) EnterAssignmentExpression(c *parser.AssignmentExpressionContext) {
	l.setAssignee(c.SingleExpression(0)).addOperators(c.Assign())
}
func (l *JavaScriptComplexityListener) EnterAssignmentOperator(c *parser.AssignmentOperatorContext) {
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
func (l *JavaScriptComplexityListener) EnterImportExpression(c *parser.ImportExpressionContext) {
	l.addOperators(c.Import())
}
func (l *JavaScriptComplexityListener) EnterThisExpression(c *parser.ThisExpressionContext) {
	l.addOperands(c.This())
}
func (l *JavaScriptComplexityListener) EnterSuperExpression(c *parser.SuperExpressionContext) {
	l.addOperators(c.Super())
}
func (l *JavaScriptComplexityListener) EnterParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {
	l.addOperators(c.OpenParen())
}
func (l *JavaScriptComplexityListener) EnterInitializer(c *parser.InitializerContext) {
	l.addOperators(c.Assign())
}
func (l *JavaScriptComplexityListener) EnterObjectLiteral(c *parser.ObjectLiteralContext) {
	l.addOperators(c.AllComma()...).
		addOperators(c.OpenBrace())
}
func (l *JavaScriptComplexityListener) EnterAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.pushScope(l.scopes[0].assignee, c).
		addOperators(c.Async(), c.Function_(), c.Multiply())
}
func (l *JavaScriptComplexityListener) EnterArrowFunction(c *parser.ArrowFunctionContext) {
	if c.ArrowFunctionBody().FunctionBody() != nil {
		l.pushScope(l.scopes[0].assignee, c)
	}
	l.addOperators(c.Async(), c.ARROW())
}
func (l *JavaScriptComplexityListener) EnterLiteral(c *parser.LiteralContext) {
	l.addOperands(
		c.NullLiteral(),
		c.BooleanLiteral(),
		c.StringLiteral(),
		c.RegularExpressionLiteral(),
	)
}
func (l *JavaScriptComplexityListener) EnterTemplateStringLiteral(c *parser.TemplateStringLiteralContext) {
	l.addOperands(c.AllBackTick()...)
}

func (l *JavaScriptComplexityListener) EnterTemplateStringAtom(c *parser.TemplateStringAtomContext) {
	l.addOperands(
		c.TemplateStringAtom(),
		c.TemplateStringStartExpression(),
		c.TemplateCloseBrace(),
	)
}
func (l *JavaScriptComplexityListener) EnterNumericLiteral(c *parser.NumericLiteralContext) {
	l.addOperands(
		c.DecimalLiteral(),
		c.HexIntegerLiteral(),
		c.OctalIntegerLiteral(),
		c.OctalIntegerLiteral2(),
		c.BinaryIntegerLiteral(),
	)
}
func (l *JavaScriptComplexityListener) EnterBigintLiteral(c *parser.BigintLiteralContext) {
	l.addOperands(
		c.BigDecimalIntegerLiteral(),
		c.BigHexIntegerLiteral(),
		c.BigOctalIntegerLiteral(),
		c.BigBinaryIntegerLiteral(),
	)
}
func (l *JavaScriptComplexityListener) EnterIdentifier(c *parser.IdentifierContext) {
	l.addOperands(c.Identifier())
}
func (l *JavaScriptComplexityListener) EnterLet_(c *parser.Let_Context) {
	l.addOperators(c.NonStrictLet(), c.StrictLet())
}
func (l *JavaScriptComplexityListener) EnterEos(c *parser.EosContext) {
	l.addOperators(c.SemiColon())
}
func (l *JavaScriptComplexityListener) ExitProgram(c *parser.ProgramContext) {
	l.popScope(c)
}
func (l *JavaScriptComplexityListener) ExitVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.setAssignee(nil)
}
func (l *JavaScriptComplexityListener) ExitFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.popScope(c)
}
func (l *JavaScriptComplexityListener) ExitClassTail(c *parser.ClassTailContext) {
	l.resetClass()
}
func (l *JavaScriptComplexityListener) ExitMethodDefinition(c *parser.MethodDefinitionContext) {
	l.popScope(c)
}
func (l *JavaScriptComplexityListener) ExitAssignmentExpression(c *parser.AssignmentExpressionContext) {
	l.setAssignee(nil)
}
func (l *JavaScriptComplexityListener) ExitAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.popScope(c)
}
func (l *JavaScriptComplexityListener) ExitArrowFunctionBody(c *parser.ArrowFunctionBodyContext) {
	if c.FunctionBody() != nil {
		l.popScope(c)
	}
}
