package logger

import (
	"figsyntax/internal/debugger"
	"figsyntax/internal/parser"
	"log/slog"
)

type DebugListener struct {
	parser.JavaScriptParserListener
	logger *slog.Logger
}

func NewDebugListener(logger *slog.Logger) *DebugListener {
	debugger.Trace()

	listener := new(DebugListener)
	listener.JavaScriptParserListener = new(parser.BaseJavaScriptParserListener)
	listener.logger = logger
	return listener
}

// EnterProgram is called when entering the program production.
func (l *DebugListener) EnterProgram(c *parser.ProgramContext) {
	l.logger.Info("EnterProgram called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSourceElement is called when entering the sourceElement production.
func (l *DebugListener) EnterSourceElement(c *parser.SourceElementContext) {
	l.logger.Info("EnterSourceElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterStatement is called when entering the statement production.
func (l *DebugListener) EnterStatement(c *parser.StatementContext) {
	l.logger.Info("EnterStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBlock is called when entering the block production.
func (l *DebugListener) EnterBlock(c *parser.BlockContext) {
	l.logger.Info("EnterBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterStatementList is called when entering the statementList production.
func (l *DebugListener) EnterStatementList(c *parser.StatementListContext) {
	l.logger.Info("EnterStatementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportStatement is called when entering the importStatement production.
func (l *DebugListener) EnterImportStatement(c *parser.ImportStatementContext) {
	l.logger.Info("EnterImportStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportFromBlock is called when entering the importFromBlock production.
func (l *DebugListener) EnterImportFromBlock(c *parser.ImportFromBlockContext) {
	l.logger.Info("EnterImportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportModuleItems is called when entering the importModuleItems production.
func (l *DebugListener) EnterImportModuleItems(c *parser.ImportModuleItemsContext) {
	l.logger.Info("EnterImportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportAliasName is called when entering the importAliasName production.
func (l *DebugListener) EnterImportAliasName(c *parser.ImportAliasNameContext) {
	l.logger.Info("EnterImportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterModuleExportName is called when entering the moduleExportName production.
func (l *DebugListener) EnterModuleExportName(c *parser.ModuleExportNameContext) {
	l.logger.Info("EnterModuleExportName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportedBinding is called when entering the importedBinding production.
func (l *DebugListener) EnterImportedBinding(c *parser.ImportedBindingContext) {
	l.logger.Info("EnterImportedBinding called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportDefault is called when entering the importDefault production.
func (l *DebugListener) EnterImportDefault(c *parser.ImportDefaultContext) {
	l.logger.Info("EnterImportDefault called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportNamespace is called when entering the importNamespace production.
func (l *DebugListener) EnterImportNamespace(c *parser.ImportNamespaceContext) {
	l.logger.Info("EnterImportNamespace called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportFrom is called when entering the importFrom production.
func (l *DebugListener) EnterImportFrom(c *parser.ImportFromContext) {
	l.logger.Info("EnterImportFrom called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAliasName is called when entering the aliasName production.
func (l *DebugListener) EnterAliasName(c *parser.AliasNameContext) {
	l.logger.Info("EnterAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportDeclaration is called when entering the ExportDeclaration production.
func (l *DebugListener) EnterExportDeclaration(c *parser.ExportDeclarationContext) {
	l.logger.Info("EnterExportDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportDefaultDeclaration is called when entering the ExportDefaultDeclaration production.
func (l *DebugListener) EnterExportDefaultDeclaration(c *parser.ExportDefaultDeclarationContext) {
	l.logger.Info("EnterExportDefaultDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportFromBlock is called when entering the exportFromBlock production.
func (l *DebugListener) EnterExportFromBlock(c *parser.ExportFromBlockContext) {
	l.logger.Info("EnterExportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportModuleItems is called when entering the exportModuleItems production.
func (l *DebugListener) EnterExportModuleItems(c *parser.ExportModuleItemsContext) {
	l.logger.Info("EnterExportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportAliasName is called when entering the exportAliasName production.
func (l *DebugListener) EnterExportAliasName(c *parser.ExportAliasNameContext) {
	l.logger.Info("EnterExportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDeclaration is called when entering the declaration production.
func (l *DebugListener) EnterDeclaration(c *parser.DeclarationContext) {
	l.logger.Info("EnterDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableStatement is called when entering the variableStatement production.
func (l *DebugListener) EnterVariableStatement(c *parser.VariableStatementContext) {
	l.logger.Info("EnterVariableStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
func (l *DebugListener) EnterVariableDeclarationList(c *parser.VariableDeclarationListContext) {
	l.logger.Info("EnterVariableDeclarationList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableDeclaration is called when entering the variableDeclaration production.
func (l *DebugListener) EnterVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.logger.Info("EnterVariableDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
func (l *DebugListener) EnterEmptyStatement_(c *parser.EmptyStatement_Context) {
	l.logger.Info("EnterEmptyStatement_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExpressionStatement is called when entering the expressionStatement production.
func (l *DebugListener) EnterExpressionStatement(c *parser.ExpressionStatementContext) {
	l.logger.Info("EnterExpressionStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIfStatement is called when entering the ifStatement production.
func (l *DebugListener) EnterIfStatement(c *parser.IfStatementContext) {
	l.logger.Info("EnterIfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDoStatement is called when entering the DoStatement production.
func (l *DebugListener) EnterDoStatement(c *parser.DoStatementContext) {
	l.logger.Info("EnterDoStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterWhileStatement is called when entering the WhileStatement production.
func (l *DebugListener) EnterWhileStatement(c *parser.WhileStatementContext) {
	l.logger.Info("EnterWhileStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForStatement is called when entering the ForStatement production.
func (l *DebugListener) EnterForStatement(c *parser.ForStatementContext) {
	l.logger.Info("EnterForStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForInStatement is called when entering the ForInStatement production.
func (l *DebugListener) EnterForInStatement(c *parser.ForInStatementContext) {
	l.logger.Info("EnterForInStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForOfStatement is called when entering the ForOfStatement production.
func (l *DebugListener) EnterForOfStatement(c *parser.ForOfStatementContext) {
	l.logger.Info("EnterForOfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVarModifier is called when entering the varModifier production.
func (l *DebugListener) EnterVarModifier(c *parser.VarModifierContext) {
	l.logger.Info("EnterVarModifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterContinueStatement is called when entering the continueStatement production.
func (l *DebugListener) EnterContinueStatement(c *parser.ContinueStatementContext) {
	l.logger.Info("EnterContinueStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBreakStatement is called when entering the breakStatement production.
func (l *DebugListener) EnterBreakStatement(c *parser.BreakStatementContext) {
	l.logger.Info("EnterBreakStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterReturnStatement is called when entering the returnStatement production.
func (l *DebugListener) EnterReturnStatement(c *parser.ReturnStatementContext) {
	l.logger.Info("EnterReturnStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterYieldStatement is called when entering the yieldStatement production.
func (l *DebugListener) EnterYieldStatement(c *parser.YieldStatementContext) {
	l.logger.Info("EnterYieldStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterWithStatement is called when entering the withStatement production.
func (l *DebugListener) EnterWithStatement(c *parser.WithStatementContext) {
	l.logger.Info("EnterWithStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSwitchStatement is called when entering the switchStatement production.
func (l *DebugListener) EnterSwitchStatement(c *parser.SwitchStatementContext) {
	l.logger.Info("EnterSwitchStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseBlock is called when entering the caseBlock production.
func (l *DebugListener) EnterCaseBlock(c *parser.CaseBlockContext) {
	l.logger.Info("EnterCaseBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseClauses is called when entering the caseClauses production.
func (l *DebugListener) EnterCaseClauses(c *parser.CaseClausesContext) {
	l.logger.Info("EnterCaseClauses called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseClause is called when entering the caseClause production.
func (l *DebugListener) EnterCaseClause(c *parser.CaseClauseContext) {
	l.logger.Info("EnterCaseClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDefaultClause is called when entering the defaultClause production.
func (l *DebugListener) EnterDefaultClause(c *parser.DefaultClauseContext) {
	l.logger.Info("EnterDefaultClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLabelledStatement is called when entering the labelledStatement production.
func (l *DebugListener) EnterLabelledStatement(c *parser.LabelledStatementContext) {
	l.logger.Info("EnterLabelledStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterThrowStatement is called when entering the throwStatement production.
func (l *DebugListener) EnterThrowStatement(c *parser.ThrowStatementContext) {
	l.logger.Info("EnterThrowStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTryStatement is called when entering the tryStatement production.
func (l *DebugListener) EnterTryStatement(c *parser.TryStatementContext) {
	l.logger.Info("EnterTryStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCatchProduction is called when entering the catchProduction production.
func (l *DebugListener) EnterCatchProduction(c *parser.CatchProductionContext) {
	l.logger.Info("EnterCatchProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFinallyProduction is called when entering the finallyProduction production.
func (l *DebugListener) EnterFinallyProduction(c *parser.FinallyProductionContext) {
	l.logger.Info("EnterFinallyProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDebuggerStatement is called when entering the debuggerStatement production.
func (l *DebugListener) EnterDebuggerStatement(c *parser.DebuggerStatementContext) {
	l.logger.Info("EnterDebuggerStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionDeclaration is called when entering the functionDeclaration production.
func (l *DebugListener) EnterFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.logger.Info("EnterFunctionDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassDeclaration is called when entering the classDeclaration production.
func (l *DebugListener) EnterClassDeclaration(c *parser.ClassDeclarationContext) {
	l.logger.Info("EnterClassDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassTail is called when entering the classTail production.
func (l *DebugListener) EnterClassTail(c *parser.ClassTailContext) {
	l.logger.Info("EnterClassTail called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassElement is called when entering the classElement production.
func (l *DebugListener) EnterClassElement(c *parser.ClassElementContext) {
	l.logger.Info("EnterClassElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMethodDefinition is called when entering the methodDefinition production.
func (l *DebugListener) EnterMethodDefinition(c *parser.MethodDefinitionContext) {
	l.logger.Info("EnterMethodDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFieldDefinition is called when entering the fieldDefinition production.
func (l *DebugListener) EnterFieldDefinition(c *parser.FieldDefinitionContext) {
	l.logger.Info("EnterFieldDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassElementName is called when entering the classElementName production.
func (l *DebugListener) EnterClassElementName(c *parser.ClassElementNameContext) {
	l.logger.Info("EnterClassElementName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPrivateIdentifier is called when entering the privateIdentifier production.
func (l *DebugListener) EnterPrivateIdentifier(c *parser.PrivateIdentifierContext) {
	l.logger.Info("EnterPrivateIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFormalParameterList is called when entering the formalParameterList production.
func (l *DebugListener) EnterFormalParameterList(c *parser.FormalParameterListContext) {
	l.logger.Info("EnterFormalParameterList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFormalParameterArg is called when entering the formalParameterArg production.
func (l *DebugListener) EnterFormalParameterArg(c *parser.FormalParameterArgContext) {
	l.logger.Info("EnterFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLastFormalParameterArg is called when entering the lastFormalParameterArg production.
func (l *DebugListener) EnterLastFormalParameterArg(c *parser.LastFormalParameterArgContext) {
	l.logger.Info("EnterLastFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionBody is called when entering the functionBody production.
func (l *DebugListener) EnterFunctionBody(c *parser.FunctionBodyContext) {
	l.logger.Info("EnterFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSourceElements is called when entering the sourceElements production.
func (l *DebugListener) EnterSourceElements(c *parser.SourceElementsContext) {
	l.logger.Info("EnterSourceElements called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayLiteral is called when entering the arrayLiteral production.
func (l *DebugListener) EnterArrayLiteral(c *parser.ArrayLiteralContext) {
	l.logger.Info("EnterArrayLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterElementList is called when entering the elementList production.
func (l *DebugListener) EnterElementList(c *parser.ElementListContext) {
	l.logger.Info("EnterElementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayElement is called when entering the arrayElement production.
func (l *DebugListener) EnterArrayElement(c *parser.ArrayElementContext) {
	l.logger.Info("EnterArrayElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
func (l *DebugListener) EnterPropertyExpressionAssignment(c *parser.PropertyExpressionAssignmentContext) {
	l.logger.Info("EnterPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterComputedPropertyExpressionAssignment is called when entering the ComputedPropertyExpressionAssignment production.
func (l *DebugListener) EnterComputedPropertyExpressionAssignment(c *parser.ComputedPropertyExpressionAssignmentContext) {
	l.logger.Info("EnterComputedPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionProperty is called when entering the FunctionProperty production.
func (l *DebugListener) EnterFunctionProperty(c *parser.FunctionPropertyContext) {
	l.logger.Info("EnterFunctionProperty called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyGetter is called when entering the PropertyGetter production.
func (l *DebugListener) EnterPropertyGetter(c *parser.PropertyGetterContext) {
	l.logger.Info("EnterPropertyGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertySetter is called when entering the PropertySetter production.
func (l *DebugListener) EnterPropertySetter(c *parser.PropertySetterContext) {
	l.logger.Info("EnterPropertySetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyShorthand is called when entering the PropertyShorthand production.
func (l *DebugListener) EnterPropertyShorthand(c *parser.PropertyShorthandContext) {
	l.logger.Info("EnterPropertyShorthand called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyName is called when entering the propertyName production.
func (l *DebugListener) EnterPropertyName(c *parser.PropertyNameContext) {
	l.logger.Info("EnterPropertyName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArguments is called when entering the arguments production.
func (l *DebugListener) EnterArguments(c *parser.ArgumentsContext) {
	l.logger.Info("EnterArguments called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArgument is called when entering the argument production.
func (l *DebugListener) EnterArgument(c *parser.ArgumentContext) {
	l.logger.Info("EnterArgument called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExpressionSequence is called when entering the expressionSequence production.
func (l *DebugListener) EnterExpressionSequence(c *parser.ExpressionSequenceContext) {
	l.logger.Info("EnterExpressionSequence called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringExpression is called when entering the TemplateStringExpression production.
func (l *DebugListener) EnterTemplateStringExpression(c *parser.TemplateStringExpressionContext) {
	l.logger.Info("EnterTemplateStringExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTernaryExpression is called when entering the TernaryExpression production.
func (l *DebugListener) EnterTernaryExpression(c *parser.TernaryExpressionContext) {
	l.logger.Info("EnterTernaryExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
func (l *DebugListener) EnterLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	l.logger.Info("EnterLogicalAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPowerExpression is called when entering the PowerExpression production.
func (l *DebugListener) EnterPowerExpression(c *parser.PowerExpressionContext) {
	l.logger.Info("EnterPowerExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPreIncrementExpression is called when entering the PreIncrementExpression production.
func (l *DebugListener) EnterPreIncrementExpression(c *parser.PreIncrementExpressionContext) {
	l.logger.Info("EnterPreIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
func (l *DebugListener) EnterObjectLiteralExpression(c *parser.ObjectLiteralExpressionContext) {
	l.logger.Info("EnterObjectLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMetaExpression is called when entering the MetaExpression production.
func (l *DebugListener) EnterMetaExpression(c *parser.MetaExpressionContext) {
	l.logger.Info("EnterMetaExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInExpression is called when entering the InExpression production.
func (l *DebugListener) EnterInExpression(c *parser.InExpressionContext) {
	l.logger.Info("EnterInExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
func (l *DebugListener) EnterLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	l.logger.Info("EnterLogicalOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterOptionalChainExpression is called when entering the OptionalChainExpression production.
func (l *DebugListener) EnterOptionalChainExpression(c *parser.OptionalChainExpressionContext) {
	l.logger.Info("EnterOptionalChainExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNotExpression is called when entering the NotExpression production.
func (l *DebugListener) EnterNotExpression(c *parser.NotExpressionContext) {
	l.logger.Info("EnterNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPreDecreaseExpression is called when entering the PreDecreaseExpression production.
func (l *DebugListener) EnterPreDecreaseExpression(c *parser.PreDecreaseExpressionContext) {
	l.logger.Info("EnterPreDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
func (l *DebugListener) EnterArgumentsExpression(c *parser.ArgumentsExpressionContext) {
	l.logger.Info("EnterArgumentsExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAwaitExpression is called when entering the AwaitExpression production.
func (l *DebugListener) EnterAwaitExpression(c *parser.AwaitExpressionContext) {
	l.logger.Info("EnterAwaitExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterThisExpression is called when entering the ThisExpression production.
func (l *DebugListener) EnterThisExpression(c *parser.ThisExpressionContext) {
	l.logger.Info("EnterThisExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionExpression is called when entering the FunctionExpression production.
func (l *DebugListener) EnterFunctionExpression(c *parser.FunctionExpressionContext) {
	l.logger.Info("EnterFunctionExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
func (l *DebugListener) EnterUnaryMinusExpression(c *parser.UnaryMinusExpressionContext) {
	l.logger.Info("EnterUnaryMinusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentExpression is called when entering the AssignmentExpression production.
func (l *DebugListener) EnterAssignmentExpression(c *parser.AssignmentExpressionContext) {
	l.logger.Info("EnterAssignmentExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPostDecreaseExpression is called when entering the PostDecreaseExpression production.
func (l *DebugListener) EnterPostDecreaseExpression(c *parser.PostDecreaseExpressionContext) {
	l.logger.Info("EnterPostDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTypeofExpression is called when entering the TypeofExpression production.
func (l *DebugListener) EnterTypeofExpression(c *parser.TypeofExpressionContext) {
	l.logger.Info("EnterTypeofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInstanceofExpression is called when entering the InstanceofExpression production.
func (l *DebugListener) EnterInstanceofExpression(c *parser.InstanceofExpressionContext) {
	l.logger.Info("EnterInstanceofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
func (l *DebugListener) EnterUnaryPlusExpression(c *parser.UnaryPlusExpressionContext) {
	l.logger.Info("EnterUnaryPlusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDeleteExpression is called when entering the DeleteExpression production.
func (l *DebugListener) EnterDeleteExpression(c *parser.DeleteExpressionContext) {
	l.logger.Info("EnterDeleteExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportExpression is called when entering the ImportExpression production.
func (l *DebugListener) EnterImportExpression(c *parser.ImportExpressionContext) {
	l.logger.Info("EnterImportExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEqualityExpression is called when entering the EqualityExpression production.
func (l *DebugListener) EnterEqualityExpression(c *parser.EqualityExpressionContext) {
	l.logger.Info("EnterEqualityExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitXOrExpression is called when entering the BitXOrExpression production.
func (l *DebugListener) EnterBitXOrExpression(c *parser.BitXOrExpressionContext) {
	l.logger.Info("EnterBitXOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSuperExpression is called when entering the SuperExpression production.
func (l *DebugListener) EnterSuperExpression(c *parser.SuperExpressionContext) {
	l.logger.Info("EnterSuperExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
func (l *DebugListener) EnterMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	l.logger.Info("EnterMultiplicativeExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitShiftExpression is called when entering the BitShiftExpression production.
func (l *DebugListener) EnterBitShiftExpression(c *parser.BitShiftExpressionContext) {
	l.logger.Info("EnterBitShiftExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
func (l *DebugListener) EnterParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {
	l.logger.Info("EnterParenthesizedExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAdditiveExpression is called when entering the AdditiveExpression production.
func (l *DebugListener) EnterAdditiveExpression(c *parser.AdditiveExpressionContext) {
	l.logger.Info("EnterAdditiveExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterRelationalExpression is called when entering the RelationalExpression production.
func (l *DebugListener) EnterRelationalExpression(c *parser.RelationalExpressionContext) {
	l.logger.Info("EnterRelationalExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPostIncrementExpression is called when entering the PostIncrementExpression production.
func (l *DebugListener) EnterPostIncrementExpression(c *parser.PostIncrementExpressionContext) {
	l.logger.Info("EnterPostIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterYieldExpression is called when entering the YieldExpression production.
func (l *DebugListener) EnterYieldExpression(c *parser.YieldExpressionContext) {
	l.logger.Info("EnterYieldExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitNotExpression is called when entering the BitNotExpression production.
func (l *DebugListener) EnterBitNotExpression(c *parser.BitNotExpressionContext) {
	l.logger.Info("EnterBitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNewExpression is called when entering the NewExpression production.
func (l *DebugListener) EnterNewExpression(c *parser.NewExpressionContext) {
	l.logger.Info("EnterNewExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLiteralExpression is called when entering the LiteralExpression production.
func (l *DebugListener) EnterLiteralExpression(c *parser.LiteralExpressionContext) {
	l.logger.Info("EnterLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayLiteralExpression is called when entering the ArrayLiteralExpression production.
func (l *DebugListener) EnterArrayLiteralExpression(c *parser.ArrayLiteralExpressionContext) {
	l.logger.Info("EnterArrayLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMemberDotExpression is called when entering the MemberDotExpression production.
func (l *DebugListener) EnterMemberDotExpression(c *parser.MemberDotExpressionContext) {
	l.logger.Info("EnterMemberDotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassExpression is called when entering the ClassExpression production.
func (l *DebugListener) EnterClassExpression(c *parser.ClassExpressionContext) {
	l.logger.Info("EnterClassExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
func (l *DebugListener) EnterMemberIndexExpression(c *parser.MemberIndexExpressionContext) {
	l.logger.Info("EnterMemberIndexExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifierExpression is called when entering the IdentifierExpression production.
func (l *DebugListener) EnterIdentifierExpression(c *parser.IdentifierExpressionContext) {
	l.logger.Info("EnterIdentifierExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitAndExpression is called when entering the BitAndExpression production.
func (l *DebugListener) EnterBitAndExpression(c *parser.BitAndExpressionContext) {
	l.logger.Info("EnterBitAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitOrExpression is called when entering the BitOrExpression production.
func (l *DebugListener) EnterBitOrExpression(c *parser.BitOrExpressionContext) {
	l.logger.Info("EnterBitOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentOperatorExpression is called when entering the AssignmentOperatorExpression production.
func (l *DebugListener) EnterAssignmentOperatorExpression(c *parser.AssignmentOperatorExpressionContext) {
	l.logger.Info("EnterAssignmentOperatorExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVoidExpression is called when entering the VoidExpression production.
func (l *DebugListener) EnterVoidExpression(c *parser.VoidExpressionContext) {
	l.logger.Info("EnterVoidExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCoalesceExpression is called when entering the CoalesceExpression production.
func (l *DebugListener) EnterCoalesceExpression(c *parser.CoalesceExpressionContext) {
	l.logger.Info("EnterCoalesceExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInitializer is called when entering the initializer production.
func (l *DebugListener) EnterInitializer(c *parser.InitializerContext) {
	l.logger.Info("EnterInitializer called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignable is called when entering the assignable production.
func (l *DebugListener) EnterAssignable(c *parser.AssignableContext) {
	l.logger.Info("EnterAssignable called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterObjectLiteral is called when entering the objectLiteral production.
func (l *DebugListener) EnterObjectLiteral(c *parser.ObjectLiteralContext) {
	l.logger.Info("EnterObjectLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNamedFunction is called when entering the NamedFunction production.
func (l *DebugListener) EnterNamedFunction(c *parser.NamedFunctionContext) {
	l.logger.Info("EnterNamedFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAnonymousFunctionDecl is called when entering the AnonymousFunctionDecl production.
func (l *DebugListener) EnterAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.logger.Info("EnterAnonymousFunctionDecl called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunction is called when entering the ArrowFunction production.
func (l *DebugListener) EnterArrowFunction(c *parser.ArrowFunctionContext) {
	l.logger.Info("EnterArrowFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunctionParameters is called when entering the arrowFunctionParameters production.
func (l *DebugListener) EnterArrowFunctionParameters(c *parser.ArrowFunctionParametersContext) {
	l.logger.Info("EnterArrowFunctionParameters called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunctionBody is called when entering the arrowFunctionBody production.
func (l *DebugListener) EnterArrowFunctionBody(c *parser.ArrowFunctionBodyContext) {
	l.logger.Info("EnterArrowFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentOperator is called when entering the assignmentOperator production.
func (l *DebugListener) EnterAssignmentOperator(c *parser.AssignmentOperatorContext) {
	l.logger.Info("EnterAssignmentOperator called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLiteral is called when entering the literal production.
func (l *DebugListener) EnterLiteral(c *parser.LiteralContext) {
	l.logger.Info("EnterLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringLiteral is called when entering the templateStringLiteral production.
func (l *DebugListener) EnterTemplateStringLiteral(c *parser.TemplateStringLiteralContext) {
	l.logger.Info("EnterTemplateStringLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringAtom is called when entering the templateStringAtom production.
func (l *DebugListener) EnterTemplateStringAtom(c *parser.TemplateStringAtomContext) {
	l.logger.Info("EnterTemplateStringAtom called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNumericLiteral is called when entering the numericLiteral production.
func (l *DebugListener) EnterNumericLiteral(c *parser.NumericLiteralContext) {
	l.logger.Info("EnterNumericLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBigintLiteral is called when entering the bigintLiteral production.
func (l *DebugListener) EnterBigintLiteral(c *parser.BigintLiteralContext) {
	l.logger.Info("EnterBigintLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterGetter is called when entering the getter production.
func (l *DebugListener) EnterGetter(c *parser.GetterContext) {
	l.logger.Info("EnterGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSetter is called when entering the setter production.
func (l *DebugListener) EnterSetter(c *parser.SetterContext) {
	l.logger.Info("EnterSetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifierName is called when entering the identifierName production.
func (l *DebugListener) EnterIdentifierName(c *parser.IdentifierNameContext) {
	l.logger.Info("EnterIdentifierName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifier is called when entering the identifier production.
func (l *DebugListener) EnterIdentifier(c *parser.IdentifierContext) {
	l.logger.Info("EnterIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterReservedWord is called when entering the reservedWord production.
func (l *DebugListener) EnterReservedWord(c *parser.ReservedWordContext) {
	l.logger.Info("EnterReservedWord called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterKeyword is called when entering the keyword production.
func (l *DebugListener) EnterKeyword(c *parser.KeywordContext) {
	l.logger.Info("EnterKeyword called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLet_ is called when entering the let_ production.
func (l *DebugListener) EnterLet_(c *parser.Let_Context) {
	l.logger.Info("EnterLet_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEos is called when entering the eos production.
func (l *DebugListener) EnterEos(c *parser.EosContext) {
	l.logger.Info("EnterEos called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitProgram is called when exiting the program production.
func (l *DebugListener) ExitProgram(c *parser.ProgramContext) {
	l.logger.Info("ExitProgram called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSourceElement is called when exiting the sourceElement production.
func (l *DebugListener) ExitSourceElement(c *parser.SourceElementContext) {
	l.logger.Info("ExitSourceElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitStatement is called when exiting the statement production.
func (l *DebugListener) ExitStatement(c *parser.StatementContext) {
	l.logger.Info("ExitStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBlock is called when exiting the block production.
func (l *DebugListener) ExitBlock(c *parser.BlockContext) {
	l.logger.Info("ExitBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitStatementList is called when exiting the statementList production.
func (l *DebugListener) ExitStatementList(c *parser.StatementListContext) {
	l.logger.Info("ExitStatementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportStatement is called when exiting the importStatement production.
func (l *DebugListener) ExitImportStatement(c *parser.ImportStatementContext) {
	l.logger.Info("ExitImportStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportFromBlock is called when exiting the importFromBlock production.
func (l *DebugListener) ExitImportFromBlock(c *parser.ImportFromBlockContext) {
	l.logger.Info("ExitImportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportModuleItems is called when exiting the importModuleItems production.
func (l *DebugListener) ExitImportModuleItems(c *parser.ImportModuleItemsContext) {
	l.logger.Info("ExitImportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportAliasName is called when exiting the importAliasName production.
func (l *DebugListener) ExitImportAliasName(c *parser.ImportAliasNameContext) {
	l.logger.Info("ExitImportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitModuleExportName is called when exiting the moduleExportName production.
func (l *DebugListener) ExitModuleExportName(c *parser.ModuleExportNameContext) {
	l.logger.Info("ExitModuleExportName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportedBinding is called when exiting the importedBinding production.
func (l *DebugListener) ExitImportedBinding(c *parser.ImportedBindingContext) {
	l.logger.Info("ExitImportedBinding called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportDefault is called when exiting the importDefault production.
func (l *DebugListener) ExitImportDefault(c *parser.ImportDefaultContext) {
	l.logger.Info("ExitImportDefault called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportNamespace is called when exiting the importNamespace production.
func (l *DebugListener) ExitImportNamespace(c *parser.ImportNamespaceContext) {
	l.logger.Info("ExitImportNamespace called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportFrom is called when exiting the importFrom production.
func (l *DebugListener) ExitImportFrom(c *parser.ImportFromContext) {
	l.logger.Info("ExitImportFrom called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAliasName is called when exiting the aliasName production.
func (l *DebugListener) ExitAliasName(c *parser.AliasNameContext) {
	l.logger.Info("ExitAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportDeclaration is called when exiting the ExportDeclaration production.
func (l *DebugListener) ExitExportDeclaration(c *parser.ExportDeclarationContext) {
	l.logger.Info("ExitExportDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportDefaultDeclaration is called when exiting the ExportDefaultDeclaration production.
func (l *DebugListener) ExitExportDefaultDeclaration(c *parser.ExportDefaultDeclarationContext) {
	l.logger.Info("ExitExportDefaultDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportFromBlock is called when exiting the exportFromBlock production.
func (l *DebugListener) ExitExportFromBlock(c *parser.ExportFromBlockContext) {
	l.logger.Info("ExitExportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportModuleItems is called when exiting the exportModuleItems production.
func (l *DebugListener) ExitExportModuleItems(c *parser.ExportModuleItemsContext) {
	l.logger.Info("ExitExportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportAliasName is called when exiting the exportAliasName production.
func (l *DebugListener) ExitExportAliasName(c *parser.ExportAliasNameContext) {
	l.logger.Info("ExitExportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDeclaration is called when exiting the declaration production.
func (l *DebugListener) ExitDeclaration(c *parser.DeclarationContext) {
	l.logger.Info("ExitDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableStatement is called when exiting the variableStatement production.
func (l *DebugListener) ExitVariableStatement(c *parser.VariableStatementContext) {
	l.logger.Info("ExitVariableStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
func (l *DebugListener) ExitVariableDeclarationList(c *parser.VariableDeclarationListContext) {
	l.logger.Info("ExitVariableDeclarationList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableDeclaration is called when exiting the variableDeclaration production.
func (l *DebugListener) ExitVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.logger.Info("ExitVariableDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
func (l *DebugListener) ExitEmptyStatement_(c *parser.EmptyStatement_Context) {
	l.logger.Info("ExitEmptyStatement_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExpressionStatement is called when exiting the expressionStatement production.
func (l *DebugListener) ExitExpressionStatement(c *parser.ExpressionStatementContext) {
	l.logger.Info("ExitExpressionStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIfStatement is called when exiting the ifStatement production.
func (l *DebugListener) ExitIfStatement(c *parser.IfStatementContext) {
	l.logger.Info("ExitIfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDoStatement is called when exiting the DoStatement production.
func (l *DebugListener) ExitDoStatement(c *parser.DoStatementContext) {
	l.logger.Info("ExitDoStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitWhileStatement is called when exiting the WhileStatement production.
func (l *DebugListener) ExitWhileStatement(c *parser.WhileStatementContext) {
	l.logger.Info("ExitWhileStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForStatement is called when exiting the ForStatement production.
func (l *DebugListener) ExitForStatement(c *parser.ForStatementContext) {
	l.logger.Info("ExitForStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForInStatement is called when exiting the ForInStatement production.
func (l *DebugListener) ExitForInStatement(c *parser.ForInStatementContext) {
	l.logger.Info("ExitForInStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForOfStatement is called when exiting the ForOfStatement production.
func (l *DebugListener) ExitForOfStatement(c *parser.ForOfStatementContext) {
	l.logger.Info("ExitForOfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVarModifier is called when exiting the varModifier production.
func (l *DebugListener) ExitVarModifier(c *parser.VarModifierContext) {
	l.logger.Info("ExitVarModifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitContinueStatement is called when exiting the continueStatement production.
func (l *DebugListener) ExitContinueStatement(c *parser.ContinueStatementContext) {
	l.logger.Info("ExitContinueStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBreakStatement is called when exiting the breakStatement production.
func (l *DebugListener) ExitBreakStatement(c *parser.BreakStatementContext) {
	l.logger.Info("ExitBreakStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitReturnStatement is called when exiting the returnStatement production.
func (l *DebugListener) ExitReturnStatement(c *parser.ReturnStatementContext) {
	l.logger.Info("ExitReturnStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitYieldStatement is called when exiting the yieldStatement production.
func (l *DebugListener) ExitYieldStatement(c *parser.YieldStatementContext) {
	l.logger.Info("ExitYieldStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitWithStatement is called when exiting the withStatement production.
func (l *DebugListener) ExitWithStatement(c *parser.WithStatementContext) {
	l.logger.Info("ExitWithStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSwitchStatement is called when exiting the switchStatement production.
func (l *DebugListener) ExitSwitchStatement(c *parser.SwitchStatementContext) {
	l.logger.Info("ExitSwitchStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseBlock is called when exiting the caseBlock production.
func (l *DebugListener) ExitCaseBlock(c *parser.CaseBlockContext) {
	l.logger.Info("ExitCaseBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseClauses is called when exiting the caseClauses production.
func (l *DebugListener) ExitCaseClauses(c *parser.CaseClausesContext) {
	l.logger.Info("ExitCaseClauses called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseClause is called when exiting the caseClause production.
func (l *DebugListener) ExitCaseClause(c *parser.CaseClauseContext) {
	l.logger.Info("ExitCaseClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDefaultClause is called when exiting the defaultClause production.
func (l *DebugListener) ExitDefaultClause(c *parser.DefaultClauseContext) {
	l.logger.Info("ExitDefaultClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLabelledStatement is called when exiting the labelledStatement production.
func (l *DebugListener) ExitLabelledStatement(c *parser.LabelledStatementContext) {
	l.logger.Info("ExitLabelledStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitThrowStatement is called when exiting the throwStatement production.
func (l *DebugListener) ExitThrowStatement(c *parser.ThrowStatementContext) {
	l.logger.Info("ExitThrowStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTryStatement is called when exiting the tryStatement production.
func (l *DebugListener) ExitTryStatement(c *parser.TryStatementContext) {
	l.logger.Info("ExitTryStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCatchProduction is called when exiting the catchProduction production.
func (l *DebugListener) ExitCatchProduction(c *parser.CatchProductionContext) {
	l.logger.Info("ExitCatchProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFinallyProduction is called when exiting the finallyProduction production.
func (l *DebugListener) ExitFinallyProduction(c *parser.FinallyProductionContext) {
	l.logger.Info("ExitFinallyProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDebuggerStatement is called when exiting the debuggerStatement production.
func (l *DebugListener) ExitDebuggerStatement(c *parser.DebuggerStatementContext) {
	l.logger.Info("ExitDebuggerStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
func (l *DebugListener) ExitFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.logger.Info("ExitFunctionDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassDeclaration is called when exiting the classDeclaration production.
func (l *DebugListener) ExitClassDeclaration(c *parser.ClassDeclarationContext) {
	l.logger.Info("ExitClassDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassTail is called when exiting the classTail production.
func (l *DebugListener) ExitClassTail(c *parser.ClassTailContext) {
	l.logger.Info("ExitClassTail called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassElement is called when exiting the classElement production.
func (l *DebugListener) ExitClassElement(c *parser.ClassElementContext) {
	l.logger.Info("ExitClassElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMethodDefinition is called when exiting the methodDefinition production.
func (l *DebugListener) ExitMethodDefinition(c *parser.MethodDefinitionContext) {
	l.logger.Info("ExitMethodDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFieldDefinition is called when exiting the fieldDefinition production.
func (l *DebugListener) ExitFieldDefinition(c *parser.FieldDefinitionContext) {
	l.logger.Info("ExitFieldDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassElementName is called when exiting the classElementName production.
func (l *DebugListener) ExitClassElementName(c *parser.ClassElementNameContext) {
	l.logger.Info("ExitClassElementName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPrivateIdentifier is called when exiting the privateIdentifier production.
func (l *DebugListener) ExitPrivateIdentifier(c *parser.PrivateIdentifierContext) {
	l.logger.Info("ExitPrivateIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFormalParameterList is called when exiting the formalParameterList production.
func (l *DebugListener) ExitFormalParameterList(c *parser.FormalParameterListContext) {
	l.logger.Info("ExitFormalParameterList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFormalParameterArg is called when exiting the formalParameterArg production.
func (l *DebugListener) ExitFormalParameterArg(c *parser.FormalParameterArgContext) {
	l.logger.Info("ExitFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLastFormalParameterArg is called when exiting the lastFormalParameterArg production.
func (l *DebugListener) ExitLastFormalParameterArg(c *parser.LastFormalParameterArgContext) {
	l.logger.Info("ExitLastFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionBody is called when exiting the functionBody production.
func (l *DebugListener) ExitFunctionBody(c *parser.FunctionBodyContext) {
	l.logger.Info("ExitFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSourceElements is called when exiting the sourceElements production.
func (l *DebugListener) ExitSourceElements(c *parser.SourceElementsContext) {
	l.logger.Info("ExitSourceElements called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayLiteral is called when exiting the arrayLiteral production.
func (l *DebugListener) ExitArrayLiteral(c *parser.ArrayLiteralContext) {
	l.logger.Info("ExitArrayLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitElementList is called when exiting the elementList production.
func (l *DebugListener) ExitElementList(c *parser.ElementListContext) {
	l.logger.Info("ExitElementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayElement is called when exiting the arrayElement production.
func (l *DebugListener) ExitArrayElement(c *parser.ArrayElementContext) {
	l.logger.Info("ExitArrayElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
func (l *DebugListener) ExitPropertyExpressionAssignment(c *parser.PropertyExpressionAssignmentContext) {
	l.logger.Info("ExitPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitComputedPropertyExpressionAssignment is called when exiting the ComputedPropertyExpressionAssignment production.
func (l *DebugListener) ExitComputedPropertyExpressionAssignment(c *parser.ComputedPropertyExpressionAssignmentContext) {
	l.logger.Info("ExitComputedPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionProperty is called when exiting the FunctionProperty production.
func (l *DebugListener) ExitFunctionProperty(c *parser.FunctionPropertyContext) {
	l.logger.Info("ExitFunctionProperty called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyGetter is called when exiting the PropertyGetter production.
func (l *DebugListener) ExitPropertyGetter(c *parser.PropertyGetterContext) {
	l.logger.Info("ExitPropertyGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertySetter is called when exiting the PropertySetter production.
func (l *DebugListener) ExitPropertySetter(c *parser.PropertySetterContext) {
	l.logger.Info("ExitPropertySetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyShorthand is called when exiting the PropertyShorthand production.
func (l *DebugListener) ExitPropertyShorthand(c *parser.PropertyShorthandContext) {
	l.logger.Info("ExitPropertyShorthand called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyName is called when exiting the propertyName production.
func (l *DebugListener) ExitPropertyName(c *parser.PropertyNameContext) {
	l.logger.Info("ExitPropertyName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArguments is called when exiting the arguments production.
func (l *DebugListener) ExitArguments(c *parser.ArgumentsContext) {
	l.logger.Info("ExitArguments called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArgument is called when exiting the argument production.
func (l *DebugListener) ExitArgument(c *parser.ArgumentContext) {
	l.logger.Info("ExitArgument called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExpressionSequence is called when exiting the expressionSequence production.
func (l *DebugListener) ExitExpressionSequence(c *parser.ExpressionSequenceContext) {
	l.logger.Info("ExitExpressionSequence called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringExpression is called when exiting the TemplateStringExpression production.
func (l *DebugListener) ExitTemplateStringExpression(c *parser.TemplateStringExpressionContext) {
	l.logger.Info("ExitTemplateStringExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTernaryExpression is called when exiting the TernaryExpression production.
func (l *DebugListener) ExitTernaryExpression(c *parser.TernaryExpressionContext) {
	l.logger.Info("ExitTernaryExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
func (l *DebugListener) ExitLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	l.logger.Info("ExitLogicalAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPowerExpression is called when exiting the PowerExpression production.
func (l *DebugListener) ExitPowerExpression(c *parser.PowerExpressionContext) {
	l.logger.Info("ExitPowerExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPreIncrementExpression is called when exiting the PreIncrementExpression production.
func (l *DebugListener) ExitPreIncrementExpression(c *parser.PreIncrementExpressionContext) {
	l.logger.Info("ExitPreIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
func (l *DebugListener) ExitObjectLiteralExpression(c *parser.ObjectLiteralExpressionContext) {
	l.logger.Info("ExitObjectLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMetaExpression is called when exiting the MetaExpression production.
func (l *DebugListener) ExitMetaExpression(c *parser.MetaExpressionContext) {
	l.logger.Info("ExitMetaExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInExpression is called when exiting the InExpression production.
func (l *DebugListener) ExitInExpression(c *parser.InExpressionContext) {
	l.logger.Info("ExitInExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
func (l *DebugListener) ExitLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	l.logger.Info("ExitLogicalOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitOptionalChainExpression is called when exiting the OptionalChainExpression production.
func (l *DebugListener) ExitOptionalChainExpression(c *parser.OptionalChainExpressionContext) {
	l.logger.Info("ExitOptionalChainExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNotExpression is called when exiting the NotExpression production.
func (l *DebugListener) ExitNotExpression(c *parser.NotExpressionContext) {
	l.logger.Info("ExitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPreDecreaseExpression is called when exiting the PreDecreaseExpression production.
func (l *DebugListener) ExitPreDecreaseExpression(c *parser.PreDecreaseExpressionContext) {
	l.logger.Info("ExitPreDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
func (l *DebugListener) ExitArgumentsExpression(c *parser.ArgumentsExpressionContext) {
	l.logger.Info("ExitArgumentsExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAwaitExpression is called when exiting the AwaitExpression production.
func (l *DebugListener) ExitAwaitExpression(c *parser.AwaitExpressionContext) {
	l.logger.Info("ExitAwaitExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitThisExpression is called when exiting the ThisExpression production.
func (l *DebugListener) ExitThisExpression(c *parser.ThisExpressionContext) {
	l.logger.Info("ExitThisExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionExpression is called when exiting the FunctionExpression production.
func (l *DebugListener) ExitFunctionExpression(c *parser.FunctionExpressionContext) {
	l.logger.Info("ExitFunctionExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
func (l *DebugListener) ExitUnaryMinusExpression(c *parser.UnaryMinusExpressionContext) {
	l.logger.Info("ExitUnaryMinusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
func (l *DebugListener) ExitAssignmentExpression(c *parser.AssignmentExpressionContext) {
	l.logger.Info("ExitAssignmentExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPostDecreaseExpression is called when exiting the PostDecreaseExpression production.
func (l *DebugListener) ExitPostDecreaseExpression(c *parser.PostDecreaseExpressionContext) {
	l.logger.Info("ExitPostDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTypeofExpression is called when exiting the TypeofExpression production.
func (l *DebugListener) ExitTypeofExpression(c *parser.TypeofExpressionContext) {
	l.logger.Info("ExitTypeofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInstanceofExpression is called when exiting the InstanceofExpression production.
func (l *DebugListener) ExitInstanceofExpression(c *parser.InstanceofExpressionContext) {
	l.logger.Info("ExitInstanceofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
func (l *DebugListener) ExitUnaryPlusExpression(c *parser.UnaryPlusExpressionContext) {
	l.logger.Info("ExitUnaryPlusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDeleteExpression is called when exiting the DeleteExpression production.
func (l *DebugListener) ExitDeleteExpression(c *parser.DeleteExpressionContext) {
	l.logger.Info("ExitDeleteExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportExpression is called when exiting the ImportExpression production.
func (l *DebugListener) ExitImportExpression(c *parser.ImportExpressionContext) {
	l.logger.Info("ExitImportExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEqualityExpression is called when exiting the EqualityExpression production.
func (l *DebugListener) ExitEqualityExpression(c *parser.EqualityExpressionContext) {
	l.logger.Info("ExitEqualityExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitXOrExpression is called when exiting the BitXOrExpression production.
func (l *DebugListener) ExitBitXOrExpression(c *parser.BitXOrExpressionContext) {
	l.logger.Info("ExitBitXOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSuperExpression is called when exiting the SuperExpression production.
func (l *DebugListener) ExitSuperExpression(c *parser.SuperExpressionContext) {
	l.logger.Info("ExitSuperExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
func (l *DebugListener) ExitMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	l.logger.Info("ExitMultiplicativeExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitShiftExpression is called when exiting the BitShiftExpression production.
func (l *DebugListener) ExitBitShiftExpression(c *parser.BitShiftExpressionContext) {
	l.logger.Info("ExitBitShiftExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
func (l *DebugListener) ExitParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {
	l.logger.Info("ExitParenthesizedExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
func (l *DebugListener) ExitAdditiveExpression(c *parser.AdditiveExpressionContext) {
	l.logger.Info("ExitAdditiveExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitRelationalExpression is called when exiting the RelationalExpression production.
func (l *DebugListener) ExitRelationalExpression(c *parser.RelationalExpressionContext) {
	l.logger.Info("ExitRelationalExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPostIncrementExpression is called when exiting the PostIncrementExpression production.
func (l *DebugListener) ExitPostIncrementExpression(c *parser.PostIncrementExpressionContext) {
	l.logger.Info("ExitPostIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitYieldExpression is called when exiting the YieldExpression production.
func (l *DebugListener) ExitYieldExpression(c *parser.YieldExpressionContext) {
	l.logger.Info("ExitYieldExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitNotExpression is called when exiting the BitNotExpression production.
func (l *DebugListener) ExitBitNotExpression(c *parser.BitNotExpressionContext) {
	l.logger.Info("ExitBitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNewExpression is called when exiting the NewExpression production.
func (l *DebugListener) ExitNewExpression(c *parser.NewExpressionContext) {
	l.logger.Info("ExitNewExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLiteralExpression is called when exiting the LiteralExpression production.
func (l *DebugListener) ExitLiteralExpression(c *parser.LiteralExpressionContext) {
	l.logger.Info("ExitLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayLiteralExpression is called when exiting the ArrayLiteralExpression production.
func (l *DebugListener) ExitArrayLiteralExpression(c *parser.ArrayLiteralExpressionContext) {
	l.logger.Info("ExitArrayLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMemberDotExpression is called when exiting the MemberDotExpression production.
func (l *DebugListener) ExitMemberDotExpression(c *parser.MemberDotExpressionContext) {
	l.logger.Info("ExitMemberDotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassExpression is called when exiting the ClassExpression production.
func (l *DebugListener) ExitClassExpression(c *parser.ClassExpressionContext) {
	l.logger.Info("ExitClassExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
func (l *DebugListener) ExitMemberIndexExpression(c *parser.MemberIndexExpressionContext) {
	l.logger.Info("ExitMemberIndexExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
func (l *DebugListener) ExitIdentifierExpression(c *parser.IdentifierExpressionContext) {
	l.logger.Info("ExitIdentifierExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitAndExpression is called when exiting the BitAndExpression production.
func (l *DebugListener) ExitBitAndExpression(c *parser.BitAndExpressionContext) {
	l.logger.Info("ExitBitAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitOrExpression is called when exiting the BitOrExpression production.
func (l *DebugListener) ExitBitOrExpression(c *parser.BitOrExpressionContext) {
	l.logger.Info("ExitBitOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentOperatorExpression is called when exiting the AssignmentOperatorExpression production.
func (l *DebugListener) ExitAssignmentOperatorExpression(c *parser.AssignmentOperatorExpressionContext) {
	l.logger.Info("ExitAssignmentOperatorExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVoidExpression is called when exiting the VoidExpression production.
func (l *DebugListener) ExitVoidExpression(c *parser.VoidExpressionContext) {
	l.logger.Info("ExitVoidExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCoalesceExpression is called when exiting the CoalesceExpression production.
func (l *DebugListener) ExitCoalesceExpression(c *parser.CoalesceExpressionContext) {
	l.logger.Info("ExitCoalesceExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInitializer is called when exiting the initializer production.
func (l *DebugListener) ExitInitializer(c *parser.InitializerContext) {
	l.logger.Info("ExitInitializer called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignable is called when exiting the assignable production.
func (l *DebugListener) ExitAssignable(c *parser.AssignableContext) {
	l.logger.Info("ExitAssignable called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitObjectLiteral is called when exiting the objectLiteral production.
func (l *DebugListener) ExitObjectLiteral(c *parser.ObjectLiteralContext) {
	l.logger.Info("ExitObjectLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNamedFunction is called when exiting the NamedFunction production.
func (l *DebugListener) ExitNamedFunction(c *parser.NamedFunctionContext) {
	l.logger.Info("ExitNamedFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAnonymousFunctionDecl is called when exiting the AnonymousFunctionDecl production.
func (l *DebugListener) ExitAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.logger.Info("ExitAnonymousFunctionDecl called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunction is called when exiting the ArrowFunction production.
func (l *DebugListener) ExitArrowFunction(c *parser.ArrowFunctionContext) {
	l.logger.Info("ExitArrowFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunctionParameters is called when exiting the arrowFunctionParameters production.
func (l *DebugListener) ExitArrowFunctionParameters(c *parser.ArrowFunctionParametersContext) {
	l.logger.Info("ExitArrowFunctionParameters called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunctionBody is called when exiting the arrowFunctionBody production.
func (l *DebugListener) ExitArrowFunctionBody(c *parser.ArrowFunctionBodyContext) {
	l.logger.Info("ExitArrowFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentOperator is called when exiting the assignmentOperator production.
func (l *DebugListener) ExitAssignmentOperator(c *parser.AssignmentOperatorContext) {
	l.logger.Info("ExitAssignmentOperator called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLiteral is called when exiting the literal production.
func (l *DebugListener) ExitLiteral(c *parser.LiteralContext) {
	l.logger.Info("ExitLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringLiteral is called when exiting the templateStringLiteral production.
func (l *DebugListener) ExitTemplateStringLiteral(c *parser.TemplateStringLiteralContext) {
	l.logger.Info("ExitTemplateStringLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringAtom is called when exiting the templateStringAtom production.
func (l *DebugListener) ExitTemplateStringAtom(c *parser.TemplateStringAtomContext) {
	l.logger.Info("ExitTemplateStringAtom called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNumericLiteral is called when exiting the numericLiteral production.
func (l *DebugListener) ExitNumericLiteral(c *parser.NumericLiteralContext) {
	l.logger.Info("ExitNumericLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBigintLiteral is called when exiting the bigintLiteral production.
func (l *DebugListener) ExitBigintLiteral(c *parser.BigintLiteralContext) {
	l.logger.Info("ExitBigintLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitGetter is called when exiting the getter production.
func (l *DebugListener) ExitGetter(c *parser.GetterContext) {
	l.logger.Info("ExitGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSetter is called when exiting the setter production.
func (l *DebugListener) ExitSetter(c *parser.SetterContext) {
	l.logger.Info("ExitSetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifierName is called when exiting the identifierName production.
func (l *DebugListener) ExitIdentifierName(c *parser.IdentifierNameContext) {
	l.logger.Info("ExitIdentifierName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifier is called when exiting the identifier production.
func (l *DebugListener) ExitIdentifier(c *parser.IdentifierContext) {
	l.logger.Info("ExitIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitReservedWord is called when exiting the reservedWord production.
func (l *DebugListener) ExitReservedWord(c *parser.ReservedWordContext) {
	l.logger.Info("ExitReservedWord called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitKeyword is called when exiting the keyword production.
func (l *DebugListener) ExitKeyword(c *parser.KeywordContext) {
	l.logger.Info("ExitKeyword called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLet_ is called when exiting the let_ production.
func (l *DebugListener) ExitLet_(c *parser.Let_Context) {
	l.logger.Info("ExitLet_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEos is called when exiting the eos production.
func (l *DebugListener) ExitEos(c *parser.EosContext) {
	l.logger.Info("ExitEos called", slog.Int("RuleIndex", c.RuleIndex))
}
