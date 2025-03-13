package listeners

import (
	"figsyntax/internal/parser"
	"log/slog"
)

type SlocMetrics struct {
	Logical  int
	Physical int
}

type CyclomaticMetrics struct {
	Complexity int
	Density    int
}

type HalsteadMetrics struct {
	Vocabulary int
	Difficulty int
	Volume     int
	Effort     int
	Bugs       int
	Time       int
}

type FunctionMetrics struct {
	Name       string
	Line       int
	Sloc       SlocMetrics
	Halstead   HalsteadMetrics
	Cyclomatic CyclomaticMetrics
}

type Metrics struct {
	Average   FunctionMetrics
	Functions []FunctionMetrics
}

type Scope struct {
	statements int
	blocks     int
	imports    int
	variables  int
	ifs        int
	dos        int
	whiles     int
	fors       int
	switches   int
	cases      int
	trys       int
	catches    int
	finallys   int
	functions  int
	classes    int
	methods    int
	ternarys   int
}

type ComplexityListener struct {
	parser.JavaScriptParserListener
	logger  *slog.Logger
	metrics Metrics
	scopes  []*Scope
}

func NewComplexityListener(logger *slog.Logger) *ComplexityListener {
	listener := new(ComplexityListener)
	listener.JavaScriptParserListener = new(parser.BaseJavaScriptParserListener)
	listener.logger = logger
	return listener
}

func (l *ComplexityListener) currentScope() *Scope {
	return l.scopes[0]
}

func (l *ComplexityListener) pushScope() {
	l.scopes = append([]*Scope{new(Scope)}, l.scopes...)
}

func (l *ComplexityListener) popScope() *Scope {
	scope := l.scopes[0]
	l.scopes = l.scopes[1:]
	return scope
}

func (l *ComplexityListener) GetMetrics() Metrics {
	return l.metrics
}

func (l *ComplexityListener) EnterProgram(c *parser.ProgramContext) {
	l.scopes = []*Scope{new(Scope)}
}

// func (l *ComplexityListener) EnterSourceElement(c *parser.SourceElementContext) {}
func (l *ComplexityListener) EnterStatement(c *parser.StatementContext) {
	l.currentScope().statements++
}

func (l *ComplexityListener) EnterBlock(c *parser.BlockContext) {
	l.currentScope().blocks++
}

func (l *ComplexityListener) EnterStatementList(c *parser.StatementListContext) {}

func (l *ComplexityListener) EnterImportStatement(c *parser.ImportStatementContext) {
	l.currentScope().imports++
}

func (l *ComplexityListener) EnterImportFromBlock(c *parser.ImportFromBlockContext)     {}
func (l *ComplexityListener) EnterImportModuleItems(c *parser.ImportModuleItemsContext) {}
func (l *ComplexityListener) EnterImportAliasName(c *parser.ImportAliasNameContext)     {}
func (l *ComplexityListener) EnterModuleExportName(c *parser.ModuleExportNameContext)   {}
func (l *ComplexityListener) EnterImportedBinding(c *parser.ImportedBindingContext)     {}
func (l *ComplexityListener) EnterImportDefault(c *parser.ImportDefaultContext)         {}
func (l *ComplexityListener) EnterImportNamespace(c *parser.ImportNamespaceContext)     {}
func (l *ComplexityListener) EnterImportFrom(c *parser.ImportFromContext)               {}
func (l *ComplexityListener) EnterAliasName(c *parser.AliasNameContext)                 {}
func (l *ComplexityListener) EnterExportDeclaration(c *parser.ExportDeclarationContext) {}
func (l *ComplexityListener) EnterExportDefaultDeclaration(c *parser.ExportDefaultDeclarationContext) {
}
func (l *ComplexityListener) EnterExportFromBlock(c *parser.ExportFromBlockContext)                 {}
func (l *ComplexityListener) EnterExportModuleItems(c *parser.ExportModuleItemsContext)             {}
func (l *ComplexityListener) EnterExportAliasName(c *parser.ExportAliasNameContext)                 {}
func (l *ComplexityListener) EnterDeclaration(c *parser.DeclarationContext)                         {}
func (l *ComplexityListener) EnterVariableStatement(c *parser.VariableStatementContext)             {}
func (l *ComplexityListener) EnterVariableDeclarationList(c *parser.VariableDeclarationListContext) {}
func (l *ComplexityListener) EnterVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.currentScope().variables++
}
func (l *ComplexityListener) EnterEmptyStatement_(c *parser.EmptyStatement_Context) {}

func (l *ComplexityListener) EnterExpressionStatement(c *parser.ExpressionStatementContext) {}

func (l *ComplexityListener) EnterIfStatement(c *parser.IfStatementContext) {
	l.currentScope().ifs++
}

func (l *ComplexityListener) EnterDoStatement(c *parser.DoStatementContext) {
	l.currentScope().dos++
}

func (l *ComplexityListener) EnterWhileStatement(c *parser.WhileStatementContext) {
	l.currentScope().whiles++
}
func (l *ComplexityListener) EnterForStatement(c *parser.ForStatementContext) {
	l.currentScope().fors++
}
func (l *ComplexityListener) EnterForInStatement(c *parser.ForInStatementContext) {
	l.currentScope().fors++
}
func (l *ComplexityListener) EnterForOfStatement(c *parser.ForOfStatementContext) {
	l.currentScope().fors++
}
func (l *ComplexityListener) EnterVarModifier(c *parser.VarModifierContext)             {}
func (l *ComplexityListener) EnterContinueStatement(c *parser.ContinueStatementContext) {}
func (l *ComplexityListener) EnterBreakStatement(c *parser.BreakStatementContext)       {}
func (l *ComplexityListener) EnterReturnStatement(c *parser.ReturnStatementContext)     {}
func (l *ComplexityListener) EnterYieldStatement(c *parser.YieldStatementContext)       {}
func (l *ComplexityListener) EnterWithStatement(c *parser.WithStatementContext)         {}
func (l *ComplexityListener) EnterSwitchStatement(c *parser.SwitchStatementContext) {
	l.currentScope().switches++
}
func (l *ComplexityListener) EnterCaseBlock(c *parser.CaseBlockContext)     {}
func (l *ComplexityListener) EnterCaseClauses(c *parser.CaseClausesContext) {}
func (l *ComplexityListener) EnterCaseClause(c *parser.CaseClauseContext) {
	l.currentScope().cases++
}
func (l *ComplexityListener) EnterDefaultClause(c *parser.DefaultClauseContext) {
	l.currentScope().cases++
}

func (l *ComplexityListener) EnterLabelledStatement(c *parser.LabelledStatementContext) {}
func (l *ComplexityListener) EnterThrowStatement(c *parser.ThrowStatementContext)       {}
func (l *ComplexityListener) EnterTryStatement(c *parser.TryStatementContext) {
	l.currentScope().trys++
}
func (l *ComplexityListener) EnterCatchProduction(c *parser.CatchProductionContext) {
	l.currentScope().catches++
}
func (l *ComplexityListener) EnterFinallyProduction(c *parser.FinallyProductionContext) {
	l.currentScope().finallys++
}

func (l *ComplexityListener) EnterDebuggerStatement(c *parser.DebuggerStatementContext) {}

func (l *ComplexityListener) EnterFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.currentScope().functions++
}

func (l *ComplexityListener) EnterClassDeclaration(c *parser.ClassDeclarationContext) {
	l.currentScope().classes++
}

func (l *ComplexityListener) EnterClassTail(c *parser.ClassTailContext)       {}
func (l *ComplexityListener) EnterClassElement(c *parser.ClassElementContext) {}
func (l *ComplexityListener) EnterMethodDefinition(c *parser.MethodDefinitionContext) {
	l.currentScope().methods++
}
func (l *ComplexityListener) EnterFieldDefinition(c *parser.FieldDefinitionContext)   {}
func (l *ComplexityListener) EnterClassElementName(c *parser.ClassElementNameContext) {}

func (l *ComplexityListener) EnterPrivateIdentifier(c *parser.PrivateIdentifierContext)           {}
func (l *ComplexityListener) EnterFormalParameterList(c *parser.FormalParameterListContext)       {}
func (l *ComplexityListener) EnterFormalParameterArg(c *parser.FormalParameterArgContext)         {}
func (l *ComplexityListener) EnterLastFormalParameterArg(c *parser.LastFormalParameterArgContext) {}
func (l *ComplexityListener) EnterFunctionBody(c *parser.FunctionBodyContext)                     {}
func (l *ComplexityListener) EnterSourceElements(c *parser.SourceElementsContext)                 {}
func (l *ComplexityListener) EnterArrayLiteral(c *parser.ArrayLiteralContext)                     {}
func (l *ComplexityListener) EnterElementList(c *parser.ElementListContext)                       {}
func (l *ComplexityListener) EnterArrayElement(c *parser.ArrayElementContext)                     {}
func (l *ComplexityListener) EnterPropertyExpressionAssignment(c *parser.PropertyExpressionAssignmentContext) {
}
func (l *ComplexityListener) EnterComputedPropertyExpressionAssignment(c *parser.ComputedPropertyExpressionAssignmentContext) {
}
func (l *ComplexityListener) EnterFunctionProperty(c *parser.FunctionPropertyContext)     {}
func (l *ComplexityListener) EnterPropertyGetter(c *parser.PropertyGetterContext)         {}
func (l *ComplexityListener) EnterPropertySetter(c *parser.PropertySetterContext)         {}
func (l *ComplexityListener) EnterPropertyShorthand(c *parser.PropertyShorthandContext)   {}
func (l *ComplexityListener) EnterPropertyName(c *parser.PropertyNameContext)             {}
func (l *ComplexityListener) EnterArguments(c *parser.ArgumentsContext)                   {}
func (l *ComplexityListener) EnterArgument(c *parser.ArgumentContext)                     {}
func (l *ComplexityListener) EnterExpressionSequence(c *parser.ExpressionSequenceContext) {}
func (l *ComplexityListener) EnterTemplateStringExpression(c *parser.TemplateStringExpressionContext) {
}
func (l *ComplexityListener) EnterTernaryExpression(c *parser.TernaryExpressionContext) {
	l.currentScope().ternarys++
}
func (l *ComplexityListener) EnterLogicalAndExpression(c *parser.LogicalAndExpressionContext) {}
func (l *ComplexityListener) EnterPowerExpression(c *parser.PowerExpressionContext)           {}

func (l *ComplexityListener) EnterPreIncrementExpression(c *parser.PreIncrementExpressionContext)   {}
func (l *ComplexityListener) EnterObjectLiteralExpression(c *parser.ObjectLiteralExpressionContext) {}
func (l *ComplexityListener) EnterMetaExpression(c *parser.MetaExpressionContext)                   {}
func (l *ComplexityListener) EnterInExpression(c *parser.InExpressionContext)                       {}
func (l *ComplexityListener) EnterLogicalOrExpression(c *parser.LogicalOrExpressionContext)         {}
func (l *ComplexityListener) EnterOptionalChainExpression(c *parser.OptionalChainExpressionContext) {}
func (l *ComplexityListener) EnterNotExpression(c *parser.NotExpressionContext)                     {}
func (l *ComplexityListener) EnterPreDecreaseExpression(c *parser.PreDecreaseExpressionContext)     {}
func (l *ComplexityListener) EnterArgumentsExpression(c *parser.ArgumentsExpressionContext)         {}
func (l *ComplexityListener) EnterAwaitExpression(c *parser.AwaitExpressionContext)                 {}
func (l *ComplexityListener) EnterThisExpression(c *parser.ThisExpressionContext)                   {}
func (l *ComplexityListener) EnterFunctionExpression(c *parser.FunctionExpressionContext)           {}
func (l *ComplexityListener) EnterUnaryMinusExpression(c *parser.UnaryMinusExpressionContext)       {}
func (l *ComplexityListener) EnterAssignmentExpression(c *parser.AssignmentExpressionContext)       {}
func (l *ComplexityListener) EnterPostDecreaseExpression(c *parser.PostDecreaseExpressionContext)   {}
func (l *ComplexityListener) EnterTypeofExpression(c *parser.TypeofExpressionContext)               {}
func (l *ComplexityListener) EnterInstanceofExpression(c *parser.InstanceofExpressionContext)       {}
func (l *ComplexityListener) EnterUnaryPlusExpression(c *parser.UnaryPlusExpressionContext)         {}
func (l *ComplexityListener) EnterDeleteExpression(c *parser.DeleteExpressionContext)               {}
func (l *ComplexityListener) EnterImportExpression(c *parser.ImportExpressionContext)               {}
func (l *ComplexityListener) EnterEqualityExpression(c *parser.EqualityExpressionContext)           {}
func (l *ComplexityListener) EnterBitXOrExpression(c *parser.BitXOrExpressionContext)               {}
func (l *ComplexityListener) EnterSuperExpression(c *parser.SuperExpressionContext)                 {}
func (l *ComplexityListener) EnterMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
}
func (l *ComplexityListener) EnterBitShiftExpression(c *parser.BitShiftExpressionContext)           {}
func (l *ComplexityListener) EnterParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {}
func (l *ComplexityListener) EnterAdditiveExpression(c *parser.AdditiveExpressionContext)           {}
func (l *ComplexityListener) EnterRelationalExpression(c *parser.RelationalExpressionContext)       {}
func (l *ComplexityListener) EnterPostIncrementExpression(c *parser.PostIncrementExpressionContext) {}
func (l *ComplexityListener) EnterYieldExpression(c *parser.YieldExpressionContext)                 {}
func (l *ComplexityListener) EnterBitNotExpression(c *parser.BitNotExpressionContext)               {}
func (l *ComplexityListener) EnterNewExpression(c *parser.NewExpressionContext)                     {}
func (l *ComplexityListener) EnterLiteralExpression(c *parser.LiteralExpressionContext)             {}
func (l *ComplexityListener) EnterArrayLiteralExpression(c *parser.ArrayLiteralExpressionContext)   {}
func (l *ComplexityListener) EnterMemberDotExpression(c *parser.MemberDotExpressionContext)         {}
func (l *ComplexityListener) EnterClassExpression(c *parser.ClassExpressionContext)                 {}
func (l *ComplexityListener) EnterMemberIndexExpression(c *parser.MemberIndexExpressionContext)     {}
func (l *ComplexityListener) EnterIdentifierExpression(c *parser.IdentifierExpressionContext)       {}
func (l *ComplexityListener) EnterBitAndExpression(c *parser.BitAndExpressionContext)               {}
func (l *ComplexityListener) EnterBitOrExpression(c *parser.BitOrExpressionContext)                 {}
func (l *ComplexityListener) EnterAssignmentOperatorExpression(c *parser.AssignmentOperatorExpressionContext) {
}
func (l *ComplexityListener) EnterVoidExpression(c *parser.VoidExpressionContext)         {}
func (l *ComplexityListener) EnterCoalesceExpression(c *parser.CoalesceExpressionContext) {}
func (l *ComplexityListener) EnterInitializer(c *parser.InitializerContext)               {}
func (l *ComplexityListener) EnterAssignable(c *parser.AssignableContext)                 {}
func (l *ComplexityListener) EnterObjectLiteral(c *parser.ObjectLiteralContext)           {}
func (l *ComplexityListener) EnterNamedFunction(c *parser.NamedFunctionContext)           {}
func (l *ComplexityListener) EnterAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.currentScope().functions++
}
func (l *ComplexityListener) EnterArrowFunction(c *parser.ArrowFunctionContext) {
	l.currentScope().functions++
}
func (l *ComplexityListener) EnterArrowFunctionParameters(c *parser.ArrowFunctionParametersContext) {}
func (l *ComplexityListener) EnterArrowFunctionBody(c *parser.ArrowFunctionBodyContext)             {}
func (l *ComplexityListener) EnterAssignmentOperator(c *parser.AssignmentOperatorContext)           {}
func (l *ComplexityListener) EnterLiteral(c *parser.LiteralContext)                                 {}
func (l *ComplexityListener) EnterTemplateStringLiteral(c *parser.TemplateStringLiteralContext)     {}
func (l *ComplexityListener) EnterTemplateStringAtom(c *parser.TemplateStringAtomContext)           {}
func (l *ComplexityListener) EnterNumericLiteral(c *parser.NumericLiteralContext)                   {}
func (l *ComplexityListener) EnterBigintLiteral(c *parser.BigintLiteralContext)                     {}
func (l *ComplexityListener) EnterGetter(c *parser.GetterContext)                                   {}
func (l *ComplexityListener) EnterSetter(c *parser.SetterContext)                                   {}
func (l *ComplexityListener) EnterIdentifierName(c *parser.IdentifierNameContext)                   {}
func (l *ComplexityListener) EnterIdentifier(c *parser.IdentifierContext)                           {}
func (l *ComplexityListener) EnterReservedWord(c *parser.ReservedWordContext)                       {}
func (l *ComplexityListener) EnterKeyword(c *parser.KeywordContext)                                 {}
func (l *ComplexityListener) EnterLet_(c *parser.Let_Context)                                       {}
func (l *ComplexityListener) EnterEos(c *parser.EosContext)                                         {}

func (l *ComplexityListener) ExitProgram(c *parser.ProgramContext) {
	l.logger.Info("finished walking", slog.Int("statements", l.scopes[0].statements))
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
// func (l *ComplexityListener) ExitVariableDeclaration(c *parser.VariableDeclarationContext) {}
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
// func (l *ComplexityListener) ExitFunctionDeclaration(c *parser.FunctionDeclarationContext) {}
// func (l *ComplexityListener) ExitClassDeclaration(c *parser.ClassDeclarationContext) {}
// func (l *ComplexityListener) ExitClassTail(c *parser.ClassTailContext) {}
// func (l *ComplexityListener) ExitClassElement(c *parser.ClassElementContext) {}
// func (l *ComplexityListener) ExitMethodDefinition(c *parser.MethodDefinitionContext) {}
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
// func (l *ComplexityListener) ExitAssignmentExpression(c *parser.AssignmentExpressionContext) {}
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
// func (l *ComplexityListener) ExitAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {}
// func (l *ComplexityListener) ExitArrowFunction(c *parser.ArrowFunctionContext) {}
// func (l *ComplexityListener) ExitArrowFunctionParameters(c *parser.ArrowFunctionParametersContext) {}
// func (l *ComplexityListener) ExitArrowFunctionBody(c *parser.ArrowFunctionBodyContext) {}
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
