package listeners

import (
	"figsyntax/internal/parser"
	"log/slog"
)

type DebugListener struct {
	parser.JavaScriptParserListener
	logger *slog.Logger
}

func NewDebugListener(logger *slog.Logger) *DebugListener {
	listener := new(DebugListener)
	listener.JavaScriptParserListener = new(parser.BaseJavaScriptParserListener)
	listener.logger = logger
	return listener
}

// EnterProgram is called when entering the program production.
func (l *DebugListener) EnterProgram(c *parser.ProgramContext) {
	l.logger.Debug("EnterProgram called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSourceElement is called when entering the sourceElement production.
func (l *DebugListener) EnterSourceElement(c *parser.SourceElementContext) {
	l.logger.Debug("EnterSourceElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterStatement is called when entering the statement production.
func (l *DebugListener) EnterStatement(c *parser.StatementContext) {
	l.logger.Debug("EnterStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBlock is called when entering the block production.
func (l *DebugListener) EnterBlock(c *parser.BlockContext) {
	l.logger.Debug("EnterBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterStatementList is called when entering the statementList production.
func (l *DebugListener) EnterStatementList(c *parser.StatementListContext) {
	l.logger.Debug("EnterStatementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportStatement is called when entering the importStatement production.
func (l *DebugListener) EnterImportStatement(c *parser.ImportStatementContext) {
	l.logger.Debug("EnterImportStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportFromBlock is called when entering the importFromBlock production.
func (l *DebugListener) EnterImportFromBlock(c *parser.ImportFromBlockContext) {
	l.logger.Debug("EnterImportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportModuleItems is called when entering the importModuleItems production.
func (l *DebugListener) EnterImportModuleItems(c *parser.ImportModuleItemsContext) {
	l.logger.Debug("EnterImportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportAliasName is called when entering the importAliasName production.
func (l *DebugListener) EnterImportAliasName(c *parser.ImportAliasNameContext) {
	l.logger.Debug("EnterImportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterModuleExportName is called when entering the moduleExportName production.
func (l *DebugListener) EnterModuleExportName(c *parser.ModuleExportNameContext) {
	l.logger.Debug("EnterModuleExportName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportedBinding is called when entering the importedBinding production.
func (l *DebugListener) EnterImportedBinding(c *parser.ImportedBindingContext) {
	l.logger.Debug("EnterImportedBinding called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportDefault is called when entering the importDefault production.
func (l *DebugListener) EnterImportDefault(c *parser.ImportDefaultContext) {
	l.logger.Debug("EnterImportDefault called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportNamespace is called when entering the importNamespace production.
func (l *DebugListener) EnterImportNamespace(c *parser.ImportNamespaceContext) {
	l.logger.Debug("EnterImportNamespace called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportFrom is called when entering the importFrom production.
func (l *DebugListener) EnterImportFrom(c *parser.ImportFromContext) {
	l.logger.Debug("EnterImportFrom called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAliasName is called when entering the aliasName production.
func (l *DebugListener) EnterAliasName(c *parser.AliasNameContext) {
	l.logger.Debug("EnterAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportDeclaration is called when entering the ExportDeclaration production.
func (l *DebugListener) EnterExportDeclaration(c *parser.ExportDeclarationContext) {
	l.logger.Debug("EnterExportDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportDefaultDeclaration is called when entering the ExportDefaultDeclaration production.
func (l *DebugListener) EnterExportDefaultDeclaration(c *parser.ExportDefaultDeclarationContext) {
	l.logger.Debug("EnterExportDefaultDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportFromBlock is called when entering the exportFromBlock production.
func (l *DebugListener) EnterExportFromBlock(c *parser.ExportFromBlockContext) {
	l.logger.Debug("EnterExportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportModuleItems is called when entering the exportModuleItems production.
func (l *DebugListener) EnterExportModuleItems(c *parser.ExportModuleItemsContext) {
	l.logger.Debug("EnterExportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportAliasName is called when entering the exportAliasName production.
func (l *DebugListener) EnterExportAliasName(c *parser.ExportAliasNameContext) {
	l.logger.Debug("EnterExportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDeclaration is called when entering the declaration production.
func (l *DebugListener) EnterDeclaration(c *parser.DeclarationContext) {
	l.logger.Debug("EnterDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableStatement is called when entering the variableStatement production.
func (l *DebugListener) EnterVariableStatement(c *parser.VariableStatementContext) {
	l.logger.Debug("EnterVariableStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
func (l *DebugListener) EnterVariableDeclarationList(c *parser.VariableDeclarationListContext) {
	l.logger.Debug("EnterVariableDeclarationList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableDeclaration is called when entering the variableDeclaration production.
func (l *DebugListener) EnterVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.logger.Debug("EnterVariableDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
func (l *DebugListener) EnterEmptyStatement_(c *parser.EmptyStatement_Context) {
	l.logger.Debug("EnterEmptyStatement_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExpressionStatement is called when entering the expressionStatement production.
func (l *DebugListener) EnterExpressionStatement(c *parser.ExpressionStatementContext) {
	l.logger.Debug("EnterExpressionStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIfStatement is called when entering the ifStatement production.
func (l *DebugListener) EnterIfStatement(c *parser.IfStatementContext) {
	l.logger.Debug("EnterIfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDoStatement is called when entering the DoStatement production.
func (l *DebugListener) EnterDoStatement(c *parser.DoStatementContext) {
	l.logger.Debug("EnterDoStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterWhileStatement is called when entering the WhileStatement production.
func (l *DebugListener) EnterWhileStatement(c *parser.WhileStatementContext) {
	l.logger.Debug("EnterWhileStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForStatement is called when entering the ForStatement production.
func (l *DebugListener) EnterForStatement(c *parser.ForStatementContext) {
	l.logger.Debug("EnterForStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForInStatement is called when entering the ForInStatement production.
func (l *DebugListener) EnterForInStatement(c *parser.ForInStatementContext) {
	l.logger.Debug("EnterForInStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForOfStatement is called when entering the ForOfStatement production.
func (l *DebugListener) EnterForOfStatement(c *parser.ForOfStatementContext) {
	l.logger.Debug("EnterForOfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVarModifier is called when entering the varModifier production.
func (l *DebugListener) EnterVarModifier(c *parser.VarModifierContext) {
	l.logger.Debug("EnterVarModifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterContinueStatement is called when entering the continueStatement production.
func (l *DebugListener) EnterContinueStatement(c *parser.ContinueStatementContext) {
	l.logger.Debug("EnterContinueStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBreakStatement is called when entering the breakStatement production.
func (l *DebugListener) EnterBreakStatement(c *parser.BreakStatementContext) {
	l.logger.Debug("EnterBreakStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterReturnStatement is called when entering the returnStatement production.
func (l *DebugListener) EnterReturnStatement(c *parser.ReturnStatementContext) {
	l.logger.Debug("EnterReturnStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterYieldStatement is called when entering the yieldStatement production.
func (l *DebugListener) EnterYieldStatement(c *parser.YieldStatementContext) {
	l.logger.Debug("EnterYieldStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterWithStatement is called when entering the withStatement production.
func (l *DebugListener) EnterWithStatement(c *parser.WithStatementContext) {
	l.logger.Debug("EnterWithStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSwitchStatement is called when entering the switchStatement production.
func (l *DebugListener) EnterSwitchStatement(c *parser.SwitchStatementContext) {
	l.logger.Debug("EnterSwitchStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseBlock is called when entering the caseBlock production.
func (l *DebugListener) EnterCaseBlock(c *parser.CaseBlockContext) {
	l.logger.Debug("EnterCaseBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseClauses is called when entering the caseClauses production.
func (l *DebugListener) EnterCaseClauses(c *parser.CaseClausesContext) {
	l.logger.Debug("EnterCaseClauses called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseClause is called when entering the caseClause production.
func (l *DebugListener) EnterCaseClause(c *parser.CaseClauseContext) {
	l.logger.Debug("EnterCaseClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDefaultClause is called when entering the defaultClause production.
func (l *DebugListener) EnterDefaultClause(c *parser.DefaultClauseContext) {
	l.logger.Debug("EnterDefaultClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLabelledStatement is called when entering the labelledStatement production.
func (l *DebugListener) EnterLabelledStatement(c *parser.LabelledStatementContext) {
	l.logger.Debug("EnterLabelledStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterThrowStatement is called when entering the throwStatement production.
func (l *DebugListener) EnterThrowStatement(c *parser.ThrowStatementContext) {
	l.logger.Debug("EnterThrowStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTryStatement is called when entering the tryStatement production.
func (l *DebugListener) EnterTryStatement(c *parser.TryStatementContext) {
	l.logger.Debug("EnterTryStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCatchProduction is called when entering the catchProduction production.
func (l *DebugListener) EnterCatchProduction(c *parser.CatchProductionContext) {
	l.logger.Debug("EnterCatchProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFinallyProduction is called when entering the finallyProduction production.
func (l *DebugListener) EnterFinallyProduction(c *parser.FinallyProductionContext) {
	l.logger.Debug("EnterFinallyProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDebuggerStatement is called when entering the debuggerStatement production.
func (l *DebugListener) EnterDebuggerStatement(c *parser.DebuggerStatementContext) {
	l.logger.Debug("EnterDebuggerStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionDeclaration is called when entering the functionDeclaration production.
func (l *DebugListener) EnterFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.logger.Debug("EnterFunctionDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassDeclaration is called when entering the classDeclaration production.
func (l *DebugListener) EnterClassDeclaration(c *parser.ClassDeclarationContext) {
	l.logger.Debug("EnterClassDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassTail is called when entering the classTail production.
func (l *DebugListener) EnterClassTail(c *parser.ClassTailContext) {
	l.logger.Debug("EnterClassTail called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassElement is called when entering the classElement production.
func (l *DebugListener) EnterClassElement(c *parser.ClassElementContext) {
	l.logger.Debug("EnterClassElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMethodDefinition is called when entering the methodDefinition production.
func (l *DebugListener) EnterMethodDefinition(c *parser.MethodDefinitionContext) {
	l.logger.Debug("EnterMethodDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFieldDefinition is called when entering the fieldDefinition production.
func (l *DebugListener) EnterFieldDefinition(c *parser.FieldDefinitionContext) {
	l.logger.Debug("EnterFieldDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassElementName is called when entering the classElementName production.
func (l *DebugListener) EnterClassElementName(c *parser.ClassElementNameContext) {
	l.logger.Debug("EnterClassElementName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPrivateIdentifier is called when entering the privateIdentifier production.
func (l *DebugListener) EnterPrivateIdentifier(c *parser.PrivateIdentifierContext) {
	l.logger.Debug("EnterPrivateIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFormalParameterList is called when entering the formalParameterList production.
func (l *DebugListener) EnterFormalParameterList(c *parser.FormalParameterListContext) {
	l.logger.Debug("EnterFormalParameterList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFormalParameterArg is called when entering the formalParameterArg production.
func (l *DebugListener) EnterFormalParameterArg(c *parser.FormalParameterArgContext) {
	l.logger.Debug("EnterFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLastFormalParameterArg is called when entering the lastFormalParameterArg production.
func (l *DebugListener) EnterLastFormalParameterArg(c *parser.LastFormalParameterArgContext) {
	l.logger.Debug("EnterLastFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionBody is called when entering the functionBody production.
func (l *DebugListener) EnterFunctionBody(c *parser.FunctionBodyContext) {
	l.logger.Debug("EnterFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSourceElements is called when entering the sourceElements production.
func (l *DebugListener) EnterSourceElements(c *parser.SourceElementsContext) {
	l.logger.Debug("EnterSourceElements called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayLiteral is called when entering the arrayLiteral production.
func (l *DebugListener) EnterArrayLiteral(c *parser.ArrayLiteralContext) {
	l.logger.Debug("EnterArrayLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterElementList is called when entering the elementList production.
func (l *DebugListener) EnterElementList(c *parser.ElementListContext) {
	l.logger.Debug("EnterElementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayElement is called when entering the arrayElement production.
func (l *DebugListener) EnterArrayElement(c *parser.ArrayElementContext) {
	l.logger.Debug("EnterArrayElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
func (l *DebugListener) EnterPropertyExpressionAssignment(c *parser.PropertyExpressionAssignmentContext) {
	l.logger.Debug("EnterPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterComputedPropertyExpressionAssignment is called when entering the ComputedPropertyExpressionAssignment production.
func (l *DebugListener) EnterComputedPropertyExpressionAssignment(c *parser.ComputedPropertyExpressionAssignmentContext) {
	l.logger.Debug("EnterComputedPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionProperty is called when entering the FunctionProperty production.
func (l *DebugListener) EnterFunctionProperty(c *parser.FunctionPropertyContext) {
	l.logger.Debug("EnterFunctionProperty called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyGetter is called when entering the PropertyGetter production.
func (l *DebugListener) EnterPropertyGetter(c *parser.PropertyGetterContext) {
	l.logger.Debug("EnterPropertyGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertySetter is called when entering the PropertySetter production.
func (l *DebugListener) EnterPropertySetter(c *parser.PropertySetterContext) {
	l.logger.Debug("EnterPropertySetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyShorthand is called when entering the PropertyShorthand production.
func (l *DebugListener) EnterPropertyShorthand(c *parser.PropertyShorthandContext) {
	l.logger.Debug("EnterPropertyShorthand called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyName is called when entering the propertyName production.
func (l *DebugListener) EnterPropertyName(c *parser.PropertyNameContext) {
	l.logger.Debug("EnterPropertyName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArguments is called when entering the arguments production.
func (l *DebugListener) EnterArguments(c *parser.ArgumentsContext) {
	l.logger.Debug("EnterArguments called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArgument is called when entering the argument production.
func (l *DebugListener) EnterArgument(c *parser.ArgumentContext) {
	l.logger.Debug("EnterArgument called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExpressionSequence is called when entering the expressionSequence production.
func (l *DebugListener) EnterExpressionSequence(c *parser.ExpressionSequenceContext) {
	l.logger.Debug("EnterExpressionSequence called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringExpression is called when entering the TemplateStringExpression production.
func (l *DebugListener) EnterTemplateStringExpression(c *parser.TemplateStringExpressionContext) {
	l.logger.Debug("EnterTemplateStringExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTernaryExpression is called when entering the TernaryExpression production.
func (l *DebugListener) EnterTernaryExpression(c *parser.TernaryExpressionContext) {
	l.logger.Debug("EnterTernaryExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
func (l *DebugListener) EnterLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	l.logger.Debug("EnterLogicalAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPowerExpression is called when entering the PowerExpression production.
func (l *DebugListener) EnterPowerExpression(c *parser.PowerExpressionContext) {
	l.logger.Debug("EnterPowerExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPreIncrementExpression is called when entering the PreIncrementExpression production.
func (l *DebugListener) EnterPreIncrementExpression(c *parser.PreIncrementExpressionContext) {
	l.logger.Debug("EnterPreIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
func (l *DebugListener) EnterObjectLiteralExpression(c *parser.ObjectLiteralExpressionContext) {
	l.logger.Debug("EnterObjectLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMetaExpression is called when entering the MetaExpression production.
func (l *DebugListener) EnterMetaExpression(c *parser.MetaExpressionContext) {
	l.logger.Debug("EnterMetaExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInExpression is called when entering the InExpression production.
func (l *DebugListener) EnterInExpression(c *parser.InExpressionContext) {
	l.logger.Debug("EnterInExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
func (l *DebugListener) EnterLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	l.logger.Debug("EnterLogicalOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterOptionalChainExpression is called when entering the OptionalChainExpression production.
func (l *DebugListener) EnterOptionalChainExpression(c *parser.OptionalChainExpressionContext) {
	l.logger.Debug("EnterOptionalChainExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNotExpression is called when entering the NotExpression production.
func (l *DebugListener) EnterNotExpression(c *parser.NotExpressionContext) {
	l.logger.Debug("EnterNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPreDecreaseExpression is called when entering the PreDecreaseExpression production.
func (l *DebugListener) EnterPreDecreaseExpression(c *parser.PreDecreaseExpressionContext) {
	l.logger.Debug("EnterPreDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
func (l *DebugListener) EnterArgumentsExpression(c *parser.ArgumentsExpressionContext) {
	l.logger.Debug("EnterArgumentsExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAwaitExpression is called when entering the AwaitExpression production.
func (l *DebugListener) EnterAwaitExpression(c *parser.AwaitExpressionContext) {
	l.logger.Debug("EnterAwaitExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterThisExpression is called when entering the ThisExpression production.
func (l *DebugListener) EnterThisExpression(c *parser.ThisExpressionContext) {
	l.logger.Debug("EnterThisExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionExpression is called when entering the FunctionExpression production.
func (l *DebugListener) EnterFunctionExpression(c *parser.FunctionExpressionContext) {
	l.logger.Debug("EnterFunctionExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
func (l *DebugListener) EnterUnaryMinusExpression(c *parser.UnaryMinusExpressionContext) {
	l.logger.Debug("EnterUnaryMinusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentExpression is called when entering the AssignmentExpression production.
func (l *DebugListener) EnterAssignmentExpression(c *parser.AssignmentExpressionContext) {
	l.logger.Debug("EnterAssignmentExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPostDecreaseExpression is called when entering the PostDecreaseExpression production.
func (l *DebugListener) EnterPostDecreaseExpression(c *parser.PostDecreaseExpressionContext) {
	l.logger.Debug("EnterPostDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTypeofExpression is called when entering the TypeofExpression production.
func (l *DebugListener) EnterTypeofExpression(c *parser.TypeofExpressionContext) {
	l.logger.Debug("EnterTypeofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInstanceofExpression is called when entering the InstanceofExpression production.
func (l *DebugListener) EnterInstanceofExpression(c *parser.InstanceofExpressionContext) {
	l.logger.Debug("EnterInstanceofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
func (l *DebugListener) EnterUnaryPlusExpression(c *parser.UnaryPlusExpressionContext) {
	l.logger.Debug("EnterUnaryPlusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDeleteExpression is called when entering the DeleteExpression production.
func (l *DebugListener) EnterDeleteExpression(c *parser.DeleteExpressionContext) {
	l.logger.Debug("EnterDeleteExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportExpression is called when entering the ImportExpression production.
func (l *DebugListener) EnterImportExpression(c *parser.ImportExpressionContext) {
	l.logger.Debug("EnterImportExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEqualityExpression is called when entering the EqualityExpression production.
func (l *DebugListener) EnterEqualityExpression(c *parser.EqualityExpressionContext) {
	l.logger.Debug("EnterEqualityExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitXOrExpression is called when entering the BitXOrExpression production.
func (l *DebugListener) EnterBitXOrExpression(c *parser.BitXOrExpressionContext) {
	l.logger.Debug("EnterBitXOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSuperExpression is called when entering the SuperExpression production.
func (l *DebugListener) EnterSuperExpression(c *parser.SuperExpressionContext) {
	l.logger.Debug("EnterSuperExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
func (l *DebugListener) EnterMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	l.logger.Debug("EnterMultiplicativeExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitShiftExpression is called when entering the BitShiftExpression production.
func (l *DebugListener) EnterBitShiftExpression(c *parser.BitShiftExpressionContext) {
	l.logger.Debug("EnterBitShiftExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
func (l *DebugListener) EnterParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {
	l.logger.Debug("EnterParenthesizedExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAdditiveExpression is called when entering the AdditiveExpression production.
func (l *DebugListener) EnterAdditiveExpression(c *parser.AdditiveExpressionContext) {
	l.logger.Debug("EnterAdditiveExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterRelationalExpression is called when entering the RelationalExpression production.
func (l *DebugListener) EnterRelationalExpression(c *parser.RelationalExpressionContext) {
	l.logger.Debug("EnterRelationalExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPostIncrementExpression is called when entering the PostIncrementExpression production.
func (l *DebugListener) EnterPostIncrementExpression(c *parser.PostIncrementExpressionContext) {
	l.logger.Debug("EnterPostIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterYieldExpression is called when entering the YieldExpression production.
func (l *DebugListener) EnterYieldExpression(c *parser.YieldExpressionContext) {
	l.logger.Debug("EnterYieldExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitNotExpression is called when entering the BitNotExpression production.
func (l *DebugListener) EnterBitNotExpression(c *parser.BitNotExpressionContext) {
	l.logger.Debug("EnterBitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNewExpression is called when entering the NewExpression production.
func (l *DebugListener) EnterNewExpression(c *parser.NewExpressionContext) {
	l.logger.Debug("EnterNewExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLiteralExpression is called when entering the LiteralExpression production.
func (l *DebugListener) EnterLiteralExpression(c *parser.LiteralExpressionContext) {
	l.logger.Debug("EnterLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayLiteralExpression is called when entering the ArrayLiteralExpression production.
func (l *DebugListener) EnterArrayLiteralExpression(c *parser.ArrayLiteralExpressionContext) {
	l.logger.Debug("EnterArrayLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMemberDotExpression is called when entering the MemberDotExpression production.
func (l *DebugListener) EnterMemberDotExpression(c *parser.MemberDotExpressionContext) {
	l.logger.Debug("EnterMemberDotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassExpression is called when entering the ClassExpression production.
func (l *DebugListener) EnterClassExpression(c *parser.ClassExpressionContext) {
	l.logger.Debug("EnterClassExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
func (l *DebugListener) EnterMemberIndexExpression(c *parser.MemberIndexExpressionContext) {
	l.logger.Debug("EnterMemberIndexExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifierExpression is called when entering the IdentifierExpression production.
func (l *DebugListener) EnterIdentifierExpression(c *parser.IdentifierExpressionContext) {
	l.logger.Debug("EnterIdentifierExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitAndExpression is called when entering the BitAndExpression production.
func (l *DebugListener) EnterBitAndExpression(c *parser.BitAndExpressionContext) {
	l.logger.Debug("EnterBitAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitOrExpression is called when entering the BitOrExpression production.
func (l *DebugListener) EnterBitOrExpression(c *parser.BitOrExpressionContext) {
	l.logger.Debug("EnterBitOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentOperatorExpression is called when entering the AssignmentOperatorExpression production.
func (l *DebugListener) EnterAssignmentOperatorExpression(c *parser.AssignmentOperatorExpressionContext) {
	l.logger.Debug("EnterAssignmentOperatorExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVoidExpression is called when entering the VoidExpression production.
func (l *DebugListener) EnterVoidExpression(c *parser.VoidExpressionContext) {
	l.logger.Debug("EnterVoidExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCoalesceExpression is called when entering the CoalesceExpression production.
func (l *DebugListener) EnterCoalesceExpression(c *parser.CoalesceExpressionContext) {
	l.logger.Debug("EnterCoalesceExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInitializer is called when entering the initializer production.
func (l *DebugListener) EnterInitializer(c *parser.InitializerContext) {
	l.logger.Debug("EnterInitializer called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignable is called when entering the assignable production.
func (l *DebugListener) EnterAssignable(c *parser.AssignableContext) {
	l.logger.Debug("EnterAssignable called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterObjectLiteral is called when entering the objectLiteral production.
func (l *DebugListener) EnterObjectLiteral(c *parser.ObjectLiteralContext) {
	l.logger.Debug("EnterObjectLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNamedFunction is called when entering the NamedFunction production.
func (l *DebugListener) EnterNamedFunction(c *parser.NamedFunctionContext) {
	l.logger.Debug("EnterNamedFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAnonymousFunctionDecl is called when entering the AnonymousFunctionDecl production.
func (l *DebugListener) EnterAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.logger.Debug("EnterAnonymousFunctionDecl called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunction is called when entering the ArrowFunction production.
func (l *DebugListener) EnterArrowFunction(c *parser.ArrowFunctionContext) {
	l.logger.Debug("EnterArrowFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunctionParameters is called when entering the arrowFunctionParameters production.
func (l *DebugListener) EnterArrowFunctionParameters(c *parser.ArrowFunctionParametersContext) {
	l.logger.Debug("EnterArrowFunctionParameters called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunctionBody is called when entering the arrowFunctionBody production.
func (l *DebugListener) EnterArrowFunctionBody(c *parser.ArrowFunctionBodyContext) {
	l.logger.Debug("EnterArrowFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentOperator is called when entering the assignmentOperator production.
func (l *DebugListener) EnterAssignmentOperator(c *parser.AssignmentOperatorContext) {
	l.logger.Debug("EnterAssignmentOperator called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLiteral is called when entering the literal production.
func (l *DebugListener) EnterLiteral(c *parser.LiteralContext) {
	l.logger.Debug("EnterLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringLiteral is called when entering the templateStringLiteral production.
func (l *DebugListener) EnterTemplateStringLiteral(c *parser.TemplateStringLiteralContext) {
	l.logger.Debug("EnterTemplateStringLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringAtom is called when entering the templateStringAtom production.
func (l *DebugListener) EnterTemplateStringAtom(c *parser.TemplateStringAtomContext) {
	l.logger.Debug("EnterTemplateStringAtom called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNumericLiteral is called when entering the numericLiteral production.
func (l *DebugListener) EnterNumericLiteral(c *parser.NumericLiteralContext) {
	l.logger.Debug("EnterNumericLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBigintLiteral is called when entering the bigintLiteral production.
func (l *DebugListener) EnterBigintLiteral(c *parser.BigintLiteralContext) {
	l.logger.Debug("EnterBigintLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterGetter is called when entering the getter production.
func (l *DebugListener) EnterGetter(c *parser.GetterContext) {
	l.logger.Debug("EnterGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSetter is called when entering the setter production.
func (l *DebugListener) EnterSetter(c *parser.SetterContext) {
	l.logger.Debug("EnterSetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifierName is called when entering the identifierName production.
func (l *DebugListener) EnterIdentifierName(c *parser.IdentifierNameContext) {
	l.logger.Debug("EnterIdentifierName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifier is called when entering the identifier production.
func (l *DebugListener) EnterIdentifier(c *parser.IdentifierContext) {
	l.logger.Debug("EnterIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterReservedWord is called when entering the reservedWord production.
func (l *DebugListener) EnterReservedWord(c *parser.ReservedWordContext) {
	l.logger.Debug("EnterReservedWord called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterKeyword is called when entering the keyword production.
func (l *DebugListener) EnterKeyword(c *parser.KeywordContext) {
	l.logger.Debug("EnterKeyword called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLet_ is called when entering the let_ production.
func (l *DebugListener) EnterLet_(c *parser.Let_Context) {
	l.logger.Debug("EnterLet_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEos is called when entering the eos production.
func (l *DebugListener) EnterEos(c *parser.EosContext) {
	l.logger.Debug("EnterEos called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitProgram is called when exiting the program production.
func (l *DebugListener) ExitProgram(c *parser.ProgramContext) {
	l.logger.Debug("ExitProgram called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSourceElement is called when exiting the sourceElement production.
func (l *DebugListener) ExitSourceElement(c *parser.SourceElementContext) {
	l.logger.Debug("ExitSourceElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitStatement is called when exiting the statement production.
func (l *DebugListener) ExitStatement(c *parser.StatementContext) {
	l.logger.Debug("ExitStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBlock is called when exiting the block production.
func (l *DebugListener) ExitBlock(c *parser.BlockContext) {
	l.logger.Debug("ExitBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitStatementList is called when exiting the statementList production.
func (l *DebugListener) ExitStatementList(c *parser.StatementListContext) {
	l.logger.Debug("ExitStatementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportStatement is called when exiting the importStatement production.
func (l *DebugListener) ExitImportStatement(c *parser.ImportStatementContext) {
	l.logger.Debug("ExitImportStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportFromBlock is called when exiting the importFromBlock production.
func (l *DebugListener) ExitImportFromBlock(c *parser.ImportFromBlockContext) {
	l.logger.Debug("ExitImportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportModuleItems is called when exiting the importModuleItems production.
func (l *DebugListener) ExitImportModuleItems(c *parser.ImportModuleItemsContext) {
	l.logger.Debug("ExitImportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportAliasName is called when exiting the importAliasName production.
func (l *DebugListener) ExitImportAliasName(c *parser.ImportAliasNameContext) {
	l.logger.Debug("ExitImportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitModuleExportName is called when exiting the moduleExportName production.
func (l *DebugListener) ExitModuleExportName(c *parser.ModuleExportNameContext) {
	l.logger.Debug("ExitModuleExportName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportedBinding is called when exiting the importedBinding production.
func (l *DebugListener) ExitImportedBinding(c *parser.ImportedBindingContext) {
	l.logger.Debug("ExitImportedBinding called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportDefault is called when exiting the importDefault production.
func (l *DebugListener) ExitImportDefault(c *parser.ImportDefaultContext) {
	l.logger.Debug("ExitImportDefault called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportNamespace is called when exiting the importNamespace production.
func (l *DebugListener) ExitImportNamespace(c *parser.ImportNamespaceContext) {
	l.logger.Debug("ExitImportNamespace called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportFrom is called when exiting the importFrom production.
func (l *DebugListener) ExitImportFrom(c *parser.ImportFromContext) {
	l.logger.Debug("ExitImportFrom called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAliasName is called when exiting the aliasName production.
func (l *DebugListener) ExitAliasName(c *parser.AliasNameContext) {
	l.logger.Debug("ExitAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportDeclaration is called when exiting the ExportDeclaration production.
func (l *DebugListener) ExitExportDeclaration(c *parser.ExportDeclarationContext) {
	l.logger.Debug("ExitExportDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportDefaultDeclaration is called when exiting the ExportDefaultDeclaration production.
func (l *DebugListener) ExitExportDefaultDeclaration(c *parser.ExportDefaultDeclarationContext) {
	l.logger.Debug("ExitExportDefaultDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportFromBlock is called when exiting the exportFromBlock production.
func (l *DebugListener) ExitExportFromBlock(c *parser.ExportFromBlockContext) {
	l.logger.Debug("ExitExportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportModuleItems is called when exiting the exportModuleItems production.
func (l *DebugListener) ExitExportModuleItems(c *parser.ExportModuleItemsContext) {
	l.logger.Debug("ExitExportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportAliasName is called when exiting the exportAliasName production.
func (l *DebugListener) ExitExportAliasName(c *parser.ExportAliasNameContext) {
	l.logger.Debug("ExitExportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDeclaration is called when exiting the declaration production.
func (l *DebugListener) ExitDeclaration(c *parser.DeclarationContext) {
	l.logger.Debug("ExitDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableStatement is called when exiting the variableStatement production.
func (l *DebugListener) ExitVariableStatement(c *parser.VariableStatementContext) {
	l.logger.Debug("ExitVariableStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
func (l *DebugListener) ExitVariableDeclarationList(c *parser.VariableDeclarationListContext) {
	l.logger.Debug("ExitVariableDeclarationList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableDeclaration is called when exiting the variableDeclaration production.
func (l *DebugListener) ExitVariableDeclaration(c *parser.VariableDeclarationContext) {
	l.logger.Debug("ExitVariableDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
func (l *DebugListener) ExitEmptyStatement_(c *parser.EmptyStatement_Context) {
	l.logger.Debug("ExitEmptyStatement_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExpressionStatement is called when exiting the expressionStatement production.
func (l *DebugListener) ExitExpressionStatement(c *parser.ExpressionStatementContext) {
	l.logger.Debug("ExitExpressionStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIfStatement is called when exiting the ifStatement production.
func (l *DebugListener) ExitIfStatement(c *parser.IfStatementContext) {
	l.logger.Debug("ExitIfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDoStatement is called when exiting the DoStatement production.
func (l *DebugListener) ExitDoStatement(c *parser.DoStatementContext) {
	l.logger.Debug("ExitDoStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitWhileStatement is called when exiting the WhileStatement production.
func (l *DebugListener) ExitWhileStatement(c *parser.WhileStatementContext) {
	l.logger.Debug("ExitWhileStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForStatement is called when exiting the ForStatement production.
func (l *DebugListener) ExitForStatement(c *parser.ForStatementContext) {
	l.logger.Debug("ExitForStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForInStatement is called when exiting the ForInStatement production.
func (l *DebugListener) ExitForInStatement(c *parser.ForInStatementContext) {
	l.logger.Debug("ExitForInStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForOfStatement is called when exiting the ForOfStatement production.
func (l *DebugListener) ExitForOfStatement(c *parser.ForOfStatementContext) {
	l.logger.Debug("ExitForOfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVarModifier is called when exiting the varModifier production.
func (l *DebugListener) ExitVarModifier(c *parser.VarModifierContext) {
	l.logger.Debug("ExitVarModifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitContinueStatement is called when exiting the continueStatement production.
func (l *DebugListener) ExitContinueStatement(c *parser.ContinueStatementContext) {
	l.logger.Debug("ExitContinueStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBreakStatement is called when exiting the breakStatement production.
func (l *DebugListener) ExitBreakStatement(c *parser.BreakStatementContext) {
	l.logger.Debug("ExitBreakStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitReturnStatement is called when exiting the returnStatement production.
func (l *DebugListener) ExitReturnStatement(c *parser.ReturnStatementContext) {
	l.logger.Debug("ExitReturnStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitYieldStatement is called when exiting the yieldStatement production.
func (l *DebugListener) ExitYieldStatement(c *parser.YieldStatementContext) {
	l.logger.Debug("ExitYieldStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitWithStatement is called when exiting the withStatement production.
func (l *DebugListener) ExitWithStatement(c *parser.WithStatementContext) {
	l.logger.Debug("ExitWithStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSwitchStatement is called when exiting the switchStatement production.
func (l *DebugListener) ExitSwitchStatement(c *parser.SwitchStatementContext) {
	l.logger.Debug("ExitSwitchStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseBlock is called when exiting the caseBlock production.
func (l *DebugListener) ExitCaseBlock(c *parser.CaseBlockContext) {
	l.logger.Debug("ExitCaseBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseClauses is called when exiting the caseClauses production.
func (l *DebugListener) ExitCaseClauses(c *parser.CaseClausesContext) {
	l.logger.Debug("ExitCaseClauses called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseClause is called when exiting the caseClause production.
func (l *DebugListener) ExitCaseClause(c *parser.CaseClauseContext) {
	l.logger.Debug("ExitCaseClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDefaultClause is called when exiting the defaultClause production.
func (l *DebugListener) ExitDefaultClause(c *parser.DefaultClauseContext) {
	l.logger.Debug("ExitDefaultClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLabelledStatement is called when exiting the labelledStatement production.
func (l *DebugListener) ExitLabelledStatement(c *parser.LabelledStatementContext) {
	l.logger.Debug("ExitLabelledStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitThrowStatement is called when exiting the throwStatement production.
func (l *DebugListener) ExitThrowStatement(c *parser.ThrowStatementContext) {
	l.logger.Debug("ExitThrowStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTryStatement is called when exiting the tryStatement production.
func (l *DebugListener) ExitTryStatement(c *parser.TryStatementContext) {
	l.logger.Debug("ExitTryStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCatchProduction is called when exiting the catchProduction production.
func (l *DebugListener) ExitCatchProduction(c *parser.CatchProductionContext) {
	l.logger.Debug("ExitCatchProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFinallyProduction is called when exiting the finallyProduction production.
func (l *DebugListener) ExitFinallyProduction(c *parser.FinallyProductionContext) {
	l.logger.Debug("ExitFinallyProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDebuggerStatement is called when exiting the debuggerStatement production.
func (l *DebugListener) ExitDebuggerStatement(c *parser.DebuggerStatementContext) {
	l.logger.Debug("ExitDebuggerStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
func (l *DebugListener) ExitFunctionDeclaration(c *parser.FunctionDeclarationContext) {
	l.logger.Debug("ExitFunctionDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassDeclaration is called when exiting the classDeclaration production.
func (l *DebugListener) ExitClassDeclaration(c *parser.ClassDeclarationContext) {
	l.logger.Debug("ExitClassDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassTail is called when exiting the classTail production.
func (l *DebugListener) ExitClassTail(c *parser.ClassTailContext) {
	l.logger.Debug("ExitClassTail called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassElement is called when exiting the classElement production.
func (l *DebugListener) ExitClassElement(c *parser.ClassElementContext) {
	l.logger.Debug("ExitClassElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMethodDefinition is called when exiting the methodDefinition production.
func (l *DebugListener) ExitMethodDefinition(c *parser.MethodDefinitionContext) {
	l.logger.Debug("ExitMethodDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFieldDefinition is called when exiting the fieldDefinition production.
func (l *DebugListener) ExitFieldDefinition(c *parser.FieldDefinitionContext) {
	l.logger.Debug("ExitFieldDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassElementName is called when exiting the classElementName production.
func (l *DebugListener) ExitClassElementName(c *parser.ClassElementNameContext) {
	l.logger.Debug("ExitClassElementName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPrivateIdentifier is called when exiting the privateIdentifier production.
func (l *DebugListener) ExitPrivateIdentifier(c *parser.PrivateIdentifierContext) {
	l.logger.Debug("ExitPrivateIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFormalParameterList is called when exiting the formalParameterList production.
func (l *DebugListener) ExitFormalParameterList(c *parser.FormalParameterListContext) {
	l.logger.Debug("ExitFormalParameterList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFormalParameterArg is called when exiting the formalParameterArg production.
func (l *DebugListener) ExitFormalParameterArg(c *parser.FormalParameterArgContext) {
	l.logger.Debug("ExitFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLastFormalParameterArg is called when exiting the lastFormalParameterArg production.
func (l *DebugListener) ExitLastFormalParameterArg(c *parser.LastFormalParameterArgContext) {
	l.logger.Debug("ExitLastFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionBody is called when exiting the functionBody production.
func (l *DebugListener) ExitFunctionBody(c *parser.FunctionBodyContext) {
	l.logger.Debug("ExitFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSourceElements is called when exiting the sourceElements production.
func (l *DebugListener) ExitSourceElements(c *parser.SourceElementsContext) {
	l.logger.Debug("ExitSourceElements called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayLiteral is called when exiting the arrayLiteral production.
func (l *DebugListener) ExitArrayLiteral(c *parser.ArrayLiteralContext) {
	l.logger.Debug("ExitArrayLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitElementList is called when exiting the elementList production.
func (l *DebugListener) ExitElementList(c *parser.ElementListContext) {
	l.logger.Debug("ExitElementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayElement is called when exiting the arrayElement production.
func (l *DebugListener) ExitArrayElement(c *parser.ArrayElementContext) {
	l.logger.Debug("ExitArrayElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
func (l *DebugListener) ExitPropertyExpressionAssignment(c *parser.PropertyExpressionAssignmentContext) {
	l.logger.Debug("ExitPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitComputedPropertyExpressionAssignment is called when exiting the ComputedPropertyExpressionAssignment production.
func (l *DebugListener) ExitComputedPropertyExpressionAssignment(c *parser.ComputedPropertyExpressionAssignmentContext) {
	l.logger.Debug("ExitComputedPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionProperty is called when exiting the FunctionProperty production.
func (l *DebugListener) ExitFunctionProperty(c *parser.FunctionPropertyContext) {
	l.logger.Debug("ExitFunctionProperty called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyGetter is called when exiting the PropertyGetter production.
func (l *DebugListener) ExitPropertyGetter(c *parser.PropertyGetterContext) {
	l.logger.Debug("ExitPropertyGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertySetter is called when exiting the PropertySetter production.
func (l *DebugListener) ExitPropertySetter(c *parser.PropertySetterContext) {
	l.logger.Debug("ExitPropertySetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyShorthand is called when exiting the PropertyShorthand production.
func (l *DebugListener) ExitPropertyShorthand(c *parser.PropertyShorthandContext) {
	l.logger.Debug("ExitPropertyShorthand called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyName is called when exiting the propertyName production.
func (l *DebugListener) ExitPropertyName(c *parser.PropertyNameContext) {
	l.logger.Debug("ExitPropertyName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArguments is called when exiting the arguments production.
func (l *DebugListener) ExitArguments(c *parser.ArgumentsContext) {
	l.logger.Debug("ExitArguments called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArgument is called when exiting the argument production.
func (l *DebugListener) ExitArgument(c *parser.ArgumentContext) {
	l.logger.Debug("ExitArgument called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExpressionSequence is called when exiting the expressionSequence production.
func (l *DebugListener) ExitExpressionSequence(c *parser.ExpressionSequenceContext) {
	l.logger.Debug("ExitExpressionSequence called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringExpression is called when exiting the TemplateStringExpression production.
func (l *DebugListener) ExitTemplateStringExpression(c *parser.TemplateStringExpressionContext) {
	l.logger.Debug("ExitTemplateStringExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTernaryExpression is called when exiting the TernaryExpression production.
func (l *DebugListener) ExitTernaryExpression(c *parser.TernaryExpressionContext) {
	l.logger.Debug("ExitTernaryExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
func (l *DebugListener) ExitLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	l.logger.Debug("ExitLogicalAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPowerExpression is called when exiting the PowerExpression production.
func (l *DebugListener) ExitPowerExpression(c *parser.PowerExpressionContext) {
	l.logger.Debug("ExitPowerExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPreIncrementExpression is called when exiting the PreIncrementExpression production.
func (l *DebugListener) ExitPreIncrementExpression(c *parser.PreIncrementExpressionContext) {
	l.logger.Debug("ExitPreIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
func (l *DebugListener) ExitObjectLiteralExpression(c *parser.ObjectLiteralExpressionContext) {
	l.logger.Debug("ExitObjectLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMetaExpression is called when exiting the MetaExpression production.
func (l *DebugListener) ExitMetaExpression(c *parser.MetaExpressionContext) {
	l.logger.Debug("ExitMetaExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInExpression is called when exiting the InExpression production.
func (l *DebugListener) ExitInExpression(c *parser.InExpressionContext) {
	l.logger.Debug("ExitInExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
func (l *DebugListener) ExitLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	l.logger.Debug("ExitLogicalOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitOptionalChainExpression is called when exiting the OptionalChainExpression production.
func (l *DebugListener) ExitOptionalChainExpression(c *parser.OptionalChainExpressionContext) {
	l.logger.Debug("ExitOptionalChainExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNotExpression is called when exiting the NotExpression production.
func (l *DebugListener) ExitNotExpression(c *parser.NotExpressionContext) {
	l.logger.Debug("ExitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPreDecreaseExpression is called when exiting the PreDecreaseExpression production.
func (l *DebugListener) ExitPreDecreaseExpression(c *parser.PreDecreaseExpressionContext) {
	l.logger.Debug("ExitPreDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
func (l *DebugListener) ExitArgumentsExpression(c *parser.ArgumentsExpressionContext) {
	l.logger.Debug("ExitArgumentsExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAwaitExpression is called when exiting the AwaitExpression production.
func (l *DebugListener) ExitAwaitExpression(c *parser.AwaitExpressionContext) {
	l.logger.Debug("ExitAwaitExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitThisExpression is called when exiting the ThisExpression production.
func (l *DebugListener) ExitThisExpression(c *parser.ThisExpressionContext) {
	l.logger.Debug("ExitThisExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionExpression is called when exiting the FunctionExpression production.
func (l *DebugListener) ExitFunctionExpression(c *parser.FunctionExpressionContext) {
	l.logger.Debug("ExitFunctionExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
func (l *DebugListener) ExitUnaryMinusExpression(c *parser.UnaryMinusExpressionContext) {
	l.logger.Debug("ExitUnaryMinusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
func (l *DebugListener) ExitAssignmentExpression(c *parser.AssignmentExpressionContext) {
	l.logger.Debug("ExitAssignmentExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPostDecreaseExpression is called when exiting the PostDecreaseExpression production.
func (l *DebugListener) ExitPostDecreaseExpression(c *parser.PostDecreaseExpressionContext) {
	l.logger.Debug("ExitPostDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTypeofExpression is called when exiting the TypeofExpression production.
func (l *DebugListener) ExitTypeofExpression(c *parser.TypeofExpressionContext) {
	l.logger.Debug("ExitTypeofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInstanceofExpression is called when exiting the InstanceofExpression production.
func (l *DebugListener) ExitInstanceofExpression(c *parser.InstanceofExpressionContext) {
	l.logger.Debug("ExitInstanceofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
func (l *DebugListener) ExitUnaryPlusExpression(c *parser.UnaryPlusExpressionContext) {
	l.logger.Debug("ExitUnaryPlusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDeleteExpression is called when exiting the DeleteExpression production.
func (l *DebugListener) ExitDeleteExpression(c *parser.DeleteExpressionContext) {
	l.logger.Debug("ExitDeleteExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportExpression is called when exiting the ImportExpression production.
func (l *DebugListener) ExitImportExpression(c *parser.ImportExpressionContext) {
	l.logger.Debug("ExitImportExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEqualityExpression is called when exiting the EqualityExpression production.
func (l *DebugListener) ExitEqualityExpression(c *parser.EqualityExpressionContext) {
	l.logger.Debug("ExitEqualityExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitXOrExpression is called when exiting the BitXOrExpression production.
func (l *DebugListener) ExitBitXOrExpression(c *parser.BitXOrExpressionContext) {
	l.logger.Debug("ExitBitXOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSuperExpression is called when exiting the SuperExpression production.
func (l *DebugListener) ExitSuperExpression(c *parser.SuperExpressionContext) {
	l.logger.Debug("ExitSuperExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
func (l *DebugListener) ExitMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	l.logger.Debug("ExitMultiplicativeExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitShiftExpression is called when exiting the BitShiftExpression production.
func (l *DebugListener) ExitBitShiftExpression(c *parser.BitShiftExpressionContext) {
	l.logger.Debug("ExitBitShiftExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
func (l *DebugListener) ExitParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {
	l.logger.Debug("ExitParenthesizedExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
func (l *DebugListener) ExitAdditiveExpression(c *parser.AdditiveExpressionContext) {
	l.logger.Debug("ExitAdditiveExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitRelationalExpression is called when exiting the RelationalExpression production.
func (l *DebugListener) ExitRelationalExpression(c *parser.RelationalExpressionContext) {
	l.logger.Debug("ExitRelationalExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPostIncrementExpression is called when exiting the PostIncrementExpression production.
func (l *DebugListener) ExitPostIncrementExpression(c *parser.PostIncrementExpressionContext) {
	l.logger.Debug("ExitPostIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitYieldExpression is called when exiting the YieldExpression production.
func (l *DebugListener) ExitYieldExpression(c *parser.YieldExpressionContext) {
	l.logger.Debug("ExitYieldExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitNotExpression is called when exiting the BitNotExpression production.
func (l *DebugListener) ExitBitNotExpression(c *parser.BitNotExpressionContext) {
	l.logger.Debug("ExitBitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNewExpression is called when exiting the NewExpression production.
func (l *DebugListener) ExitNewExpression(c *parser.NewExpressionContext) {
	l.logger.Debug("ExitNewExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLiteralExpression is called when exiting the LiteralExpression production.
func (l *DebugListener) ExitLiteralExpression(c *parser.LiteralExpressionContext) {
	l.logger.Debug("ExitLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayLiteralExpression is called when exiting the ArrayLiteralExpression production.
func (l *DebugListener) ExitArrayLiteralExpression(c *parser.ArrayLiteralExpressionContext) {
	l.logger.Debug("ExitArrayLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMemberDotExpression is called when exiting the MemberDotExpression production.
func (l *DebugListener) ExitMemberDotExpression(c *parser.MemberDotExpressionContext) {
	l.logger.Debug("ExitMemberDotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassExpression is called when exiting the ClassExpression production.
func (l *DebugListener) ExitClassExpression(c *parser.ClassExpressionContext) {
	l.logger.Debug("ExitClassExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
func (l *DebugListener) ExitMemberIndexExpression(c *parser.MemberIndexExpressionContext) {
	l.logger.Debug("ExitMemberIndexExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
func (l *DebugListener) ExitIdentifierExpression(c *parser.IdentifierExpressionContext) {
	l.logger.Debug("ExitIdentifierExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitAndExpression is called when exiting the BitAndExpression production.
func (l *DebugListener) ExitBitAndExpression(c *parser.BitAndExpressionContext) {
	l.logger.Debug("ExitBitAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitOrExpression is called when exiting the BitOrExpression production.
func (l *DebugListener) ExitBitOrExpression(c *parser.BitOrExpressionContext) {
	l.logger.Debug("ExitBitOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentOperatorExpression is called when exiting the AssignmentOperatorExpression production.
func (l *DebugListener) ExitAssignmentOperatorExpression(c *parser.AssignmentOperatorExpressionContext) {
	l.logger.Debug("ExitAssignmentOperatorExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVoidExpression is called when exiting the VoidExpression production.
func (l *DebugListener) ExitVoidExpression(c *parser.VoidExpressionContext) {
	l.logger.Debug("ExitVoidExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCoalesceExpression is called when exiting the CoalesceExpression production.
func (l *DebugListener) ExitCoalesceExpression(c *parser.CoalesceExpressionContext) {
	l.logger.Debug("ExitCoalesceExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInitializer is called when exiting the initializer production.
func (l *DebugListener) ExitInitializer(c *parser.InitializerContext) {
	l.logger.Debug("ExitInitializer called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignable is called when exiting the assignable production.
func (l *DebugListener) ExitAssignable(c *parser.AssignableContext) {
	l.logger.Debug("ExitAssignable called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitObjectLiteral is called when exiting the objectLiteral production.
func (l *DebugListener) ExitObjectLiteral(c *parser.ObjectLiteralContext) {
	l.logger.Debug("ExitObjectLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNamedFunction is called when exiting the NamedFunction production.
func (l *DebugListener) ExitNamedFunction(c *parser.NamedFunctionContext) {
	l.logger.Debug("ExitNamedFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAnonymousFunctionDecl is called when exiting the AnonymousFunctionDecl production.
func (l *DebugListener) ExitAnonymousFunctionDecl(c *parser.AnonymousFunctionDeclContext) {
	l.logger.Debug("ExitAnonymousFunctionDecl called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunction is called when exiting the ArrowFunction production.
func (l *DebugListener) ExitArrowFunction(c *parser.ArrowFunctionContext) {
	l.logger.Debug("ExitArrowFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunctionParameters is called when exiting the arrowFunctionParameters production.
func (l *DebugListener) ExitArrowFunctionParameters(c *parser.ArrowFunctionParametersContext) {
	l.logger.Debug("ExitArrowFunctionParameters called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunctionBody is called when exiting the arrowFunctionBody production.
func (l *DebugListener) ExitArrowFunctionBody(c *parser.ArrowFunctionBodyContext) {
	l.logger.Debug("ExitArrowFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentOperator is called when exiting the assignmentOperator production.
func (l *DebugListener) ExitAssignmentOperator(c *parser.AssignmentOperatorContext) {
	l.logger.Debug("ExitAssignmentOperator called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLiteral is called when exiting the literal production.
func (l *DebugListener) ExitLiteral(c *parser.LiteralContext) {
	l.logger.Debug("ExitLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringLiteral is called when exiting the templateStringLiteral production.
func (l *DebugListener) ExitTemplateStringLiteral(c *parser.TemplateStringLiteralContext) {
	l.logger.Debug("ExitTemplateStringLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringAtom is called when exiting the templateStringAtom production.
func (l *DebugListener) ExitTemplateStringAtom(c *parser.TemplateStringAtomContext) {
	l.logger.Debug("ExitTemplateStringAtom called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNumericLiteral is called when exiting the numericLiteral production.
func (l *DebugListener) ExitNumericLiteral(c *parser.NumericLiteralContext) {
	l.logger.Debug("ExitNumericLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBigintLiteral is called when exiting the bigintLiteral production.
func (l *DebugListener) ExitBigintLiteral(c *parser.BigintLiteralContext) {
	l.logger.Debug("ExitBigintLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitGetter is called when exiting the getter production.
func (l *DebugListener) ExitGetter(c *parser.GetterContext) {
	l.logger.Debug("ExitGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSetter is called when exiting the setter production.
func (l *DebugListener) ExitSetter(c *parser.SetterContext) {
	l.logger.Debug("ExitSetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifierName is called when exiting the identifierName production.
func (l *DebugListener) ExitIdentifierName(c *parser.IdentifierNameContext) {
	l.logger.Debug("ExitIdentifierName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifier is called when exiting the identifier production.
func (l *DebugListener) ExitIdentifier(c *parser.IdentifierContext) {
	l.logger.Debug("ExitIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitReservedWord is called when exiting the reservedWord production.
func (l *DebugListener) ExitReservedWord(c *parser.ReservedWordContext) {
	l.logger.Debug("ExitReservedWord called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitKeyword is called when exiting the keyword production.
func (l *DebugListener) ExitKeyword(c *parser.KeywordContext) {
	l.logger.Debug("ExitKeyword called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLet_ is called when exiting the let_ production.
func (l *DebugListener) ExitLet_(c *parser.Let_Context) {
	l.logger.Debug("ExitLet_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEos is called when exiting the eos production.
func (l *DebugListener) ExitEos(c *parser.EosContext) {
	l.logger.Debug("ExitEos called", slog.Int("RuleIndex", c.RuleIndex))
}
