package parser

import (
	"log/slog"
)

type LoggerListener struct {
	JavaScriptParserListener
	logger *slog.Logger
}

func NewLoggerListener(logger *slog.Logger) *LoggerListener {
	listener := new(LoggerListener)
	listener.JavaScriptParserListener = new(BaseJavaScriptParserListener)
	listener.logger = logger
	return listener
}

// EnterProgram is called when entering the program production.
func (l *LoggerListener) EnterProgram(c *ProgramContext) {
	l.logger.Info("EnterProgram called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSourceElement is called when entering the sourceElement production.
func (l *LoggerListener) EnterSourceElement(c *SourceElementContext) {
	l.logger.Info("EnterSourceElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterStatement is called when entering the statement production.
func (l *LoggerListener) EnterStatement(c *StatementContext) {
	l.logger.Info("EnterStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBlock is called when entering the block production.
func (l *LoggerListener) EnterBlock(c *BlockContext) {
	l.logger.Info("EnterBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterStatementList is called when entering the statementList production.
func (l *LoggerListener) EnterStatementList(c *StatementListContext) {
	l.logger.Info("EnterStatementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportStatement is called when entering the importStatement production.
func (l *LoggerListener) EnterImportStatement(c *ImportStatementContext) {
	l.logger.Info("EnterImportStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportFromBlock is called when entering the importFromBlock production.
func (l *LoggerListener) EnterImportFromBlock(c *ImportFromBlockContext) {
	l.logger.Info("EnterImportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportModuleItems is called when entering the importModuleItems production.
func (l *LoggerListener) EnterImportModuleItems(c *ImportModuleItemsContext) {
	l.logger.Info("EnterImportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportAliasName is called when entering the importAliasName production.
func (l *LoggerListener) EnterImportAliasName(c *ImportAliasNameContext) {
	l.logger.Info("EnterImportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterModuleExportName is called when entering the moduleExportName production.
func (l *LoggerListener) EnterModuleExportName(c *ModuleExportNameContext) {
	l.logger.Info("EnterModuleExportName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportedBinding is called when entering the importedBinding production.
func (l *LoggerListener) EnterImportedBinding(c *ImportedBindingContext) {
	l.logger.Info("EnterImportedBinding called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportDefault is called when entering the importDefault production.
func (l *LoggerListener) EnterImportDefault(c *ImportDefaultContext) {
	l.logger.Info("EnterImportDefault called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportNamespace is called when entering the importNamespace production.
func (l *LoggerListener) EnterImportNamespace(c *ImportNamespaceContext) {
	l.logger.Info("EnterImportNamespace called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportFrom is called when entering the importFrom production.
func (l *LoggerListener) EnterImportFrom(c *ImportFromContext) {
	l.logger.Info("EnterImportFrom called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAliasName is called when entering the aliasName production.
func (l *LoggerListener) EnterAliasName(c *AliasNameContext) {
	l.logger.Info("EnterAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportDeclaration is called when entering the ExportDeclaration production.
func (l *LoggerListener) EnterExportDeclaration(c *ExportDeclarationContext) {
	l.logger.Info("EnterExportDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportDefaultDeclaration is called when entering the ExportDefaultDeclaration production.
func (l *LoggerListener) EnterExportDefaultDeclaration(c *ExportDefaultDeclarationContext) {
	l.logger.Info("EnterExportDefaultDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportFromBlock is called when entering the exportFromBlock production.
func (l *LoggerListener) EnterExportFromBlock(c *ExportFromBlockContext) {
	l.logger.Info("EnterExportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportModuleItems is called when entering the exportModuleItems production.
func (l *LoggerListener) EnterExportModuleItems(c *ExportModuleItemsContext) {
	l.logger.Info("EnterExportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExportAliasName is called when entering the exportAliasName production.
func (l *LoggerListener) EnterExportAliasName(c *ExportAliasNameContext) {
	l.logger.Info("EnterExportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDeclaration is called when entering the declaration production.
func (l *LoggerListener) EnterDeclaration(c *DeclarationContext) {
	l.logger.Info("EnterDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableStatement is called when entering the variableStatement production.
func (l *LoggerListener) EnterVariableStatement(c *VariableStatementContext) {
	l.logger.Info("EnterVariableStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
func (l *LoggerListener) EnterVariableDeclarationList(c *VariableDeclarationListContext) {
	l.logger.Info("EnterVariableDeclarationList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVariableDeclaration is called when entering the variableDeclaration production.
func (l *LoggerListener) EnterVariableDeclaration(c *VariableDeclarationContext) {
	l.logger.Info("EnterVariableDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
func (l *LoggerListener) EnterEmptyStatement_(c *EmptyStatement_Context) {
	l.logger.Info("EnterEmptyStatement_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExpressionStatement is called when entering the expressionStatement production.
func (l *LoggerListener) EnterExpressionStatement(c *ExpressionStatementContext) {
	l.logger.Info("EnterExpressionStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIfStatement is called when entering the ifStatement production.
func (l *LoggerListener) EnterIfStatement(c *IfStatementContext) {
	l.logger.Info("EnterIfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDoStatement is called when entering the DoStatement production.
func (l *LoggerListener) EnterDoStatement(c *DoStatementContext) {
	l.logger.Info("EnterDoStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterWhileStatement is called when entering the WhileStatement production.
func (l *LoggerListener) EnterWhileStatement(c *WhileStatementContext) {
	l.logger.Info("EnterWhileStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForStatement is called when entering the ForStatement production.
func (l *LoggerListener) EnterForStatement(c *ForStatementContext) {
	l.logger.Info("EnterForStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForInStatement is called when entering the ForInStatement production.
func (l *LoggerListener) EnterForInStatement(c *ForInStatementContext) {
	l.logger.Info("EnterForInStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterForOfStatement is called when entering the ForOfStatement production.
func (l *LoggerListener) EnterForOfStatement(c *ForOfStatementContext) {
	l.logger.Info("EnterForOfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVarModifier is called when entering the varModifier production.
func (l *LoggerListener) EnterVarModifier(c *VarModifierContext) {
	l.logger.Info("EnterVarModifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterContinueStatement is called when entering the continueStatement production.
func (l *LoggerListener) EnterContinueStatement(c *ContinueStatementContext) {
	l.logger.Info("EnterContinueStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBreakStatement is called when entering the breakStatement production.
func (l *LoggerListener) EnterBreakStatement(c *BreakStatementContext) {
	l.logger.Info("EnterBreakStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterReturnStatement is called when entering the returnStatement production.
func (l *LoggerListener) EnterReturnStatement(c *ReturnStatementContext) {
	l.logger.Info("EnterReturnStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterYieldStatement is called when entering the yieldStatement production.
func (l *LoggerListener) EnterYieldStatement(c *YieldStatementContext) {
	l.logger.Info("EnterYieldStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterWithStatement is called when entering the withStatement production.
func (l *LoggerListener) EnterWithStatement(c *WithStatementContext) {
	l.logger.Info("EnterWithStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSwitchStatement is called when entering the switchStatement production.
func (l *LoggerListener) EnterSwitchStatement(c *SwitchStatementContext) {
	l.logger.Info("EnterSwitchStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseBlock is called when entering the caseBlock production.
func (l *LoggerListener) EnterCaseBlock(c *CaseBlockContext) {
	l.logger.Info("EnterCaseBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseClauses is called when entering the caseClauses production.
func (l *LoggerListener) EnterCaseClauses(c *CaseClausesContext) {
	l.logger.Info("EnterCaseClauses called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCaseClause is called when entering the caseClause production.
func (l *LoggerListener) EnterCaseClause(c *CaseClauseContext) {
	l.logger.Info("EnterCaseClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDefaultClause is called when entering the defaultClause production.
func (l *LoggerListener) EnterDefaultClause(c *DefaultClauseContext) {
	l.logger.Info("EnterDefaultClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLabelledStatement is called when entering the labelledStatement production.
func (l *LoggerListener) EnterLabelledStatement(c *LabelledStatementContext) {
	l.logger.Info("EnterLabelledStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterThrowStatement is called when entering the throwStatement production.
func (l *LoggerListener) EnterThrowStatement(c *ThrowStatementContext) {
	l.logger.Info("EnterThrowStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTryStatement is called when entering the tryStatement production.
func (l *LoggerListener) EnterTryStatement(c *TryStatementContext) {
	l.logger.Info("EnterTryStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCatchProduction is called when entering the catchProduction production.
func (l *LoggerListener) EnterCatchProduction(c *CatchProductionContext) {
	l.logger.Info("EnterCatchProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFinallyProduction is called when entering the finallyProduction production.
func (l *LoggerListener) EnterFinallyProduction(c *FinallyProductionContext) {
	l.logger.Info("EnterFinallyProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDebuggerStatement is called when entering the debuggerStatement production.
func (l *LoggerListener) EnterDebuggerStatement(c *DebuggerStatementContext) {
	l.logger.Info("EnterDebuggerStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionDeclaration is called when entering the functionDeclaration production.
func (l *LoggerListener) EnterFunctionDeclaration(c *FunctionDeclarationContext) {
	l.logger.Info("EnterFunctionDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassDeclaration is called when entering the classDeclaration production.
func (l *LoggerListener) EnterClassDeclaration(c *ClassDeclarationContext) {
	l.logger.Info("EnterClassDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassTail is called when entering the classTail production.
func (l *LoggerListener) EnterClassTail(c *ClassTailContext) {
	l.logger.Info("EnterClassTail called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassElement is called when entering the classElement production.
func (l *LoggerListener) EnterClassElement(c *ClassElementContext) {
	l.logger.Info("EnterClassElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMethodDefinition is called when entering the methodDefinition production.
func (l *LoggerListener) EnterMethodDefinition(c *MethodDefinitionContext) {
	l.logger.Info("EnterMethodDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFieldDefinition is called when entering the fieldDefinition production.
func (l *LoggerListener) EnterFieldDefinition(c *FieldDefinitionContext) {
	l.logger.Info("EnterFieldDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassElementName is called when entering the classElementName production.
func (l *LoggerListener) EnterClassElementName(c *ClassElementNameContext) {
	l.logger.Info("EnterClassElementName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPrivateIdentifier is called when entering the privateIdentifier production.
func (l *LoggerListener) EnterPrivateIdentifier(c *PrivateIdentifierContext) {
	l.logger.Info("EnterPrivateIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFormalParameterList is called when entering the formalParameterList production.
func (l *LoggerListener) EnterFormalParameterList(c *FormalParameterListContext) {
	l.logger.Info("EnterFormalParameterList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFormalParameterArg is called when entering the formalParameterArg production.
func (l *LoggerListener) EnterFormalParameterArg(c *FormalParameterArgContext) {
	l.logger.Info("EnterFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLastFormalParameterArg is called when entering the lastFormalParameterArg production.
func (l *LoggerListener) EnterLastFormalParameterArg(c *LastFormalParameterArgContext) {
	l.logger.Info("EnterLastFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionBody is called when entering the functionBody production.
func (l *LoggerListener) EnterFunctionBody(c *FunctionBodyContext) {
	l.logger.Info("EnterFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSourceElements is called when entering the sourceElements production.
func (l *LoggerListener) EnterSourceElements(c *SourceElementsContext) {
	l.logger.Info("EnterSourceElements called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayLiteral is called when entering the arrayLiteral production.
func (l *LoggerListener) EnterArrayLiteral(c *ArrayLiteralContext) {
	l.logger.Info("EnterArrayLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterElementList is called when entering the elementList production.
func (l *LoggerListener) EnterElementList(c *ElementListContext) {
	l.logger.Info("EnterElementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayElement is called when entering the arrayElement production.
func (l *LoggerListener) EnterArrayElement(c *ArrayElementContext) {
	l.logger.Info("EnterArrayElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
func (l *LoggerListener) EnterPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext) {
	l.logger.Info("EnterPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterComputedPropertyExpressionAssignment is called when entering the ComputedPropertyExpressionAssignment production.
func (l *LoggerListener) EnterComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext) {
	l.logger.Info("EnterComputedPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionProperty is called when entering the FunctionProperty production.
func (l *LoggerListener) EnterFunctionProperty(c *FunctionPropertyContext) {
	l.logger.Info("EnterFunctionProperty called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyGetter is called when entering the PropertyGetter production.
func (l *LoggerListener) EnterPropertyGetter(c *PropertyGetterContext) {
	l.logger.Info("EnterPropertyGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertySetter is called when entering the PropertySetter production.
func (l *LoggerListener) EnterPropertySetter(c *PropertySetterContext) {
	l.logger.Info("EnterPropertySetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyShorthand is called when entering the PropertyShorthand production.
func (l *LoggerListener) EnterPropertyShorthand(c *PropertyShorthandContext) {
	l.logger.Info("EnterPropertyShorthand called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPropertyName is called when entering the propertyName production.
func (l *LoggerListener) EnterPropertyName(c *PropertyNameContext) {
	l.logger.Info("EnterPropertyName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArguments is called when entering the arguments production.
func (l *LoggerListener) EnterArguments(c *ArgumentsContext) {
	l.logger.Info("EnterArguments called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArgument is called when entering the argument production.
func (l *LoggerListener) EnterArgument(c *ArgumentContext) {
	l.logger.Info("EnterArgument called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterExpressionSequence is called when entering the expressionSequence production.
func (l *LoggerListener) EnterExpressionSequence(c *ExpressionSequenceContext) {
	l.logger.Info("EnterExpressionSequence called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringExpression is called when entering the TemplateStringExpression production.
func (l *LoggerListener) EnterTemplateStringExpression(c *TemplateStringExpressionContext) {
	l.logger.Info("EnterTemplateStringExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTernaryExpression is called when entering the TernaryExpression production.
func (l *LoggerListener) EnterTernaryExpression(c *TernaryExpressionContext) {
	l.logger.Info("EnterTernaryExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
func (l *LoggerListener) EnterLogicalAndExpression(c *LogicalAndExpressionContext) {
	l.logger.Info("EnterLogicalAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPowerExpression is called when entering the PowerExpression production.
func (l *LoggerListener) EnterPowerExpression(c *PowerExpressionContext) {
	l.logger.Info("EnterPowerExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPreIncrementExpression is called when entering the PreIncrementExpression production.
func (l *LoggerListener) EnterPreIncrementExpression(c *PreIncrementExpressionContext) {
	l.logger.Info("EnterPreIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
func (l *LoggerListener) EnterObjectLiteralExpression(c *ObjectLiteralExpressionContext) {
	l.logger.Info("EnterObjectLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMetaExpression is called when entering the MetaExpression production.
func (l *LoggerListener) EnterMetaExpression(c *MetaExpressionContext) {
	l.logger.Info("EnterMetaExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInExpression is called when entering the InExpression production.
func (l *LoggerListener) EnterInExpression(c *InExpressionContext) {
	l.logger.Info("EnterInExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
func (l *LoggerListener) EnterLogicalOrExpression(c *LogicalOrExpressionContext) {
	l.logger.Info("EnterLogicalOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterOptionalChainExpression is called when entering the OptionalChainExpression production.
func (l *LoggerListener) EnterOptionalChainExpression(c *OptionalChainExpressionContext) {
	l.logger.Info("EnterOptionalChainExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNotExpression is called when entering the NotExpression production.
func (l *LoggerListener) EnterNotExpression(c *NotExpressionContext) {
	l.logger.Info("EnterNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPreDecreaseExpression is called when entering the PreDecreaseExpression production.
func (l *LoggerListener) EnterPreDecreaseExpression(c *PreDecreaseExpressionContext) {
	l.logger.Info("EnterPreDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
func (l *LoggerListener) EnterArgumentsExpression(c *ArgumentsExpressionContext) {
	l.logger.Info("EnterArgumentsExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAwaitExpression is called when entering the AwaitExpression production.
func (l *LoggerListener) EnterAwaitExpression(c *AwaitExpressionContext) {
	l.logger.Info("EnterAwaitExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterThisExpression is called when entering the ThisExpression production.
func (l *LoggerListener) EnterThisExpression(c *ThisExpressionContext) {
	l.logger.Info("EnterThisExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterFunctionExpression is called when entering the FunctionExpression production.
func (l *LoggerListener) EnterFunctionExpression(c *FunctionExpressionContext) {
	l.logger.Info("EnterFunctionExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
func (l *LoggerListener) EnterUnaryMinusExpression(c *UnaryMinusExpressionContext) {
	l.logger.Info("EnterUnaryMinusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentExpression is called when entering the AssignmentExpression production.
func (l *LoggerListener) EnterAssignmentExpression(c *AssignmentExpressionContext) {
	l.logger.Info("EnterAssignmentExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPostDecreaseExpression is called when entering the PostDecreaseExpression production.
func (l *LoggerListener) EnterPostDecreaseExpression(c *PostDecreaseExpressionContext) {
	l.logger.Info("EnterPostDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTypeofExpression is called when entering the TypeofExpression production.
func (l *LoggerListener) EnterTypeofExpression(c *TypeofExpressionContext) {
	l.logger.Info("EnterTypeofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInstanceofExpression is called when entering the InstanceofExpression production.
func (l *LoggerListener) EnterInstanceofExpression(c *InstanceofExpressionContext) {
	l.logger.Info("EnterInstanceofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
func (l *LoggerListener) EnterUnaryPlusExpression(c *UnaryPlusExpressionContext) {
	l.logger.Info("EnterUnaryPlusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterDeleteExpression is called when entering the DeleteExpression production.
func (l *LoggerListener) EnterDeleteExpression(c *DeleteExpressionContext) {
	l.logger.Info("EnterDeleteExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterImportExpression is called when entering the ImportExpression production.
func (l *LoggerListener) EnterImportExpression(c *ImportExpressionContext) {
	l.logger.Info("EnterImportExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEqualityExpression is called when entering the EqualityExpression production.
func (l *LoggerListener) EnterEqualityExpression(c *EqualityExpressionContext) {
	l.logger.Info("EnterEqualityExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitXOrExpression is called when entering the BitXOrExpression production.
func (l *LoggerListener) EnterBitXOrExpression(c *BitXOrExpressionContext) {
	l.logger.Info("EnterBitXOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSuperExpression is called when entering the SuperExpression production.
func (l *LoggerListener) EnterSuperExpression(c *SuperExpressionContext) {
	l.logger.Info("EnterSuperExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
func (l *LoggerListener) EnterMultiplicativeExpression(c *MultiplicativeExpressionContext) {
	l.logger.Info("EnterMultiplicativeExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitShiftExpression is called when entering the BitShiftExpression production.
func (l *LoggerListener) EnterBitShiftExpression(c *BitShiftExpressionContext) {
	l.logger.Info("EnterBitShiftExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
func (l *LoggerListener) EnterParenthesizedExpression(c *ParenthesizedExpressionContext) {
	l.logger.Info("EnterParenthesizedExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAdditiveExpression is called when entering the AdditiveExpression production.
func (l *LoggerListener) EnterAdditiveExpression(c *AdditiveExpressionContext) {
	l.logger.Info("EnterAdditiveExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterRelationalExpression is called when entering the RelationalExpression production.
func (l *LoggerListener) EnterRelationalExpression(c *RelationalExpressionContext) {
	l.logger.Info("EnterRelationalExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterPostIncrementExpression is called when entering the PostIncrementExpression production.
func (l *LoggerListener) EnterPostIncrementExpression(c *PostIncrementExpressionContext) {
	l.logger.Info("EnterPostIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterYieldExpression is called when entering the YieldExpression production.
func (l *LoggerListener) EnterYieldExpression(c *YieldExpressionContext) {
	l.logger.Info("EnterYieldExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitNotExpression is called when entering the BitNotExpression production.
func (l *LoggerListener) EnterBitNotExpression(c *BitNotExpressionContext) {
	l.logger.Info("EnterBitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNewExpression is called when entering the NewExpression production.
func (l *LoggerListener) EnterNewExpression(c *NewExpressionContext) {
	l.logger.Info("EnterNewExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLiteralExpression is called when entering the LiteralExpression production.
func (l *LoggerListener) EnterLiteralExpression(c *LiteralExpressionContext) {
	l.logger.Info("EnterLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrayLiteralExpression is called when entering the ArrayLiteralExpression production.
func (l *LoggerListener) EnterArrayLiteralExpression(c *ArrayLiteralExpressionContext) {
	l.logger.Info("EnterArrayLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMemberDotExpression is called when entering the MemberDotExpression production.
func (l *LoggerListener) EnterMemberDotExpression(c *MemberDotExpressionContext) {
	l.logger.Info("EnterMemberDotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterClassExpression is called when entering the ClassExpression production.
func (l *LoggerListener) EnterClassExpression(c *ClassExpressionContext) {
	l.logger.Info("EnterClassExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
func (l *LoggerListener) EnterMemberIndexExpression(c *MemberIndexExpressionContext) {
	l.logger.Info("EnterMemberIndexExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifierExpression is called when entering the IdentifierExpression production.
func (l *LoggerListener) EnterIdentifierExpression(c *IdentifierExpressionContext) {
	l.logger.Info("EnterIdentifierExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitAndExpression is called when entering the BitAndExpression production.
func (l *LoggerListener) EnterBitAndExpression(c *BitAndExpressionContext) {
	l.logger.Info("EnterBitAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBitOrExpression is called when entering the BitOrExpression production.
func (l *LoggerListener) EnterBitOrExpression(c *BitOrExpressionContext) {
	l.logger.Info("EnterBitOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentOperatorExpression is called when entering the AssignmentOperatorExpression production.
func (l *LoggerListener) EnterAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext) {
	l.logger.Info("EnterAssignmentOperatorExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterVoidExpression is called when entering the VoidExpression production.
func (l *LoggerListener) EnterVoidExpression(c *VoidExpressionContext) {
	l.logger.Info("EnterVoidExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterCoalesceExpression is called when entering the CoalesceExpression production.
func (l *LoggerListener) EnterCoalesceExpression(c *CoalesceExpressionContext) {
	l.logger.Info("EnterCoalesceExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterInitializer is called when entering the initializer production.
func (l *LoggerListener) EnterInitializer(c *InitializerContext) {
	l.logger.Info("EnterInitializer called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignable is called when entering the assignable production.
func (l *LoggerListener) EnterAssignable(c *AssignableContext) {
	l.logger.Info("EnterAssignable called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterObjectLiteral is called when entering the objectLiteral production.
func (l *LoggerListener) EnterObjectLiteral(c *ObjectLiteralContext) {
	l.logger.Info("EnterObjectLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNamedFunction is called when entering the NamedFunction production.
func (l *LoggerListener) EnterNamedFunction(c *NamedFunctionContext) {
	l.logger.Info("EnterNamedFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAnonymousFunctionDecl is called when entering the AnonymousFunctionDecl production.
func (l *LoggerListener) EnterAnonymousFunctionDecl(c *AnonymousFunctionDeclContext) {
	l.logger.Info("EnterAnonymousFunctionDecl called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunction is called when entering the ArrowFunction production.
func (l *LoggerListener) EnterArrowFunction(c *ArrowFunctionContext) {
	l.logger.Info("EnterArrowFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunctionParameters is called when entering the arrowFunctionParameters production.
func (l *LoggerListener) EnterArrowFunctionParameters(c *ArrowFunctionParametersContext) {
	l.logger.Info("EnterArrowFunctionParameters called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterArrowFunctionBody is called when entering the arrowFunctionBody production.
func (l *LoggerListener) EnterArrowFunctionBody(c *ArrowFunctionBodyContext) {
	l.logger.Info("EnterArrowFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterAssignmentOperator is called when entering the assignmentOperator production.
func (l *LoggerListener) EnterAssignmentOperator(c *AssignmentOperatorContext) {
	l.logger.Info("EnterAssignmentOperator called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLiteral is called when entering the literal production.
func (l *LoggerListener) EnterLiteral(c *LiteralContext) {
	l.logger.Info("EnterLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringLiteral is called when entering the templateStringLiteral production.
func (l *LoggerListener) EnterTemplateStringLiteral(c *TemplateStringLiteralContext) {
	l.logger.Info("EnterTemplateStringLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterTemplateStringAtom is called when entering the templateStringAtom production.
func (l *LoggerListener) EnterTemplateStringAtom(c *TemplateStringAtomContext) {
	l.logger.Info("EnterTemplateStringAtom called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterNumericLiteral is called when entering the numericLiteral production.
func (l *LoggerListener) EnterNumericLiteral(c *NumericLiteralContext) {
	l.logger.Info("EnterNumericLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterBigintLiteral is called when entering the bigintLiteral production.
func (l *LoggerListener) EnterBigintLiteral(c *BigintLiteralContext) {
	l.logger.Info("EnterBigintLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterGetter is called when entering the getter production.
func (l *LoggerListener) EnterGetter(c *GetterContext) {
	l.logger.Info("EnterGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterSetter is called when entering the setter production.
func (l *LoggerListener) EnterSetter(c *SetterContext) {
	l.logger.Info("EnterSetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifierName is called when entering the identifierName production.
func (l *LoggerListener) EnterIdentifierName(c *IdentifierNameContext) {
	l.logger.Info("EnterIdentifierName called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterIdentifier is called when entering the identifier production.
func (l *LoggerListener) EnterIdentifier(c *IdentifierContext) {
	l.logger.Info("EnterIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterReservedWord is called when entering the reservedWord production.
func (l *LoggerListener) EnterReservedWord(c *ReservedWordContext) {
	l.logger.Info("EnterReservedWord called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterKeyword is called when entering the keyword production.
func (l *LoggerListener) EnterKeyword(c *KeywordContext) {
	l.logger.Info("EnterKeyword called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterLet_ is called when entering the let_ production.
func (l *LoggerListener) EnterLet_(c *Let_Context) {
	l.logger.Info("EnterLet_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// EnterEos is called when entering the eos production.
func (l *LoggerListener) EnterEos(c *EosContext) {
	l.logger.Info("EnterEos called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitProgram is called when exiting the program production.
func (l *LoggerListener) ExitProgram(c *ProgramContext) {
	l.logger.Info("ExitProgram called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSourceElement is called when exiting the sourceElement production.
func (l *LoggerListener) ExitSourceElement(c *SourceElementContext) {
	l.logger.Info("ExitSourceElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitStatement is called when exiting the statement production.
func (l *LoggerListener) ExitStatement(c *StatementContext) {
	l.logger.Info("ExitStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBlock is called when exiting the block production.
func (l *LoggerListener) ExitBlock(c *BlockContext) {
	l.logger.Info("ExitBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitStatementList is called when exiting the statementList production.
func (l *LoggerListener) ExitStatementList(c *StatementListContext) {
	l.logger.Info("ExitStatementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportStatement is called when exiting the importStatement production.
func (l *LoggerListener) ExitImportStatement(c *ImportStatementContext) {
	l.logger.Info("ExitImportStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportFromBlock is called when exiting the importFromBlock production.
func (l *LoggerListener) ExitImportFromBlock(c *ImportFromBlockContext) {
	l.logger.Info("ExitImportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportModuleItems is called when exiting the importModuleItems production.
func (l *LoggerListener) ExitImportModuleItems(c *ImportModuleItemsContext) {
	l.logger.Info("ExitImportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportAliasName is called when exiting the importAliasName production.
func (l *LoggerListener) ExitImportAliasName(c *ImportAliasNameContext) {
	l.logger.Info("ExitImportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitModuleExportName is called when exiting the moduleExportName production.
func (l *LoggerListener) ExitModuleExportName(c *ModuleExportNameContext) {
	l.logger.Info("ExitModuleExportName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportedBinding is called when exiting the importedBinding production.
func (l *LoggerListener) ExitImportedBinding(c *ImportedBindingContext) {
	l.logger.Info("ExitImportedBinding called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportDefault is called when exiting the importDefault production.
func (l *LoggerListener) ExitImportDefault(c *ImportDefaultContext) {
	l.logger.Info("ExitImportDefault called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportNamespace is called when exiting the importNamespace production.
func (l *LoggerListener) ExitImportNamespace(c *ImportNamespaceContext) {
	l.logger.Info("ExitImportNamespace called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportFrom is called when exiting the importFrom production.
func (l *LoggerListener) ExitImportFrom(c *ImportFromContext) {
	l.logger.Info("ExitImportFrom called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAliasName is called when exiting the aliasName production.
func (l *LoggerListener) ExitAliasName(c *AliasNameContext) {
	l.logger.Info("ExitAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportDeclaration is called when exiting the ExportDeclaration production.
func (l *LoggerListener) ExitExportDeclaration(c *ExportDeclarationContext) {
	l.logger.Info("ExitExportDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportDefaultDeclaration is called when exiting the ExportDefaultDeclaration production.
func (l *LoggerListener) ExitExportDefaultDeclaration(c *ExportDefaultDeclarationContext) {
	l.logger.Info("ExitExportDefaultDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportFromBlock is called when exiting the exportFromBlock production.
func (l *LoggerListener) ExitExportFromBlock(c *ExportFromBlockContext) {
	l.logger.Info("ExitExportFromBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportModuleItems is called when exiting the exportModuleItems production.
func (l *LoggerListener) ExitExportModuleItems(c *ExportModuleItemsContext) {
	l.logger.Info("ExitExportModuleItems called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExportAliasName is called when exiting the exportAliasName production.
func (l *LoggerListener) ExitExportAliasName(c *ExportAliasNameContext) {
	l.logger.Info("ExitExportAliasName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDeclaration is called when exiting the declaration production.
func (l *LoggerListener) ExitDeclaration(c *DeclarationContext) {
	l.logger.Info("ExitDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableStatement is called when exiting the variableStatement production.
func (l *LoggerListener) ExitVariableStatement(c *VariableStatementContext) {
	l.logger.Info("ExitVariableStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
func (l *LoggerListener) ExitVariableDeclarationList(c *VariableDeclarationListContext) {
	l.logger.Info("ExitVariableDeclarationList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVariableDeclaration is called when exiting the variableDeclaration production.
func (l *LoggerListener) ExitVariableDeclaration(c *VariableDeclarationContext) {
	l.logger.Info("ExitVariableDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
func (l *LoggerListener) ExitEmptyStatement_(c *EmptyStatement_Context) {
	l.logger.Info("ExitEmptyStatement_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExpressionStatement is called when exiting the expressionStatement production.
func (l *LoggerListener) ExitExpressionStatement(c *ExpressionStatementContext) {
	l.logger.Info("ExitExpressionStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIfStatement is called when exiting the ifStatement production.
func (l *LoggerListener) ExitIfStatement(c *IfStatementContext) {
	l.logger.Info("ExitIfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDoStatement is called when exiting the DoStatement production.
func (l *LoggerListener) ExitDoStatement(c *DoStatementContext) {
	l.logger.Info("ExitDoStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitWhileStatement is called when exiting the WhileStatement production.
func (l *LoggerListener) ExitWhileStatement(c *WhileStatementContext) {
	l.logger.Info("ExitWhileStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForStatement is called when exiting the ForStatement production.
func (l *LoggerListener) ExitForStatement(c *ForStatementContext) {
	l.logger.Info("ExitForStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForInStatement is called when exiting the ForInStatement production.
func (l *LoggerListener) ExitForInStatement(c *ForInStatementContext) {
	l.logger.Info("ExitForInStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitForOfStatement is called when exiting the ForOfStatement production.
func (l *LoggerListener) ExitForOfStatement(c *ForOfStatementContext) {
	l.logger.Info("ExitForOfStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVarModifier is called when exiting the varModifier production.
func (l *LoggerListener) ExitVarModifier(c *VarModifierContext) {
	l.logger.Info("ExitVarModifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitContinueStatement is called when exiting the continueStatement production.
func (l *LoggerListener) ExitContinueStatement(c *ContinueStatementContext) {
	l.logger.Info("ExitContinueStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBreakStatement is called when exiting the breakStatement production.
func (l *LoggerListener) ExitBreakStatement(c *BreakStatementContext) {
	l.logger.Info("ExitBreakStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitReturnStatement is called when exiting the returnStatement production.
func (l *LoggerListener) ExitReturnStatement(c *ReturnStatementContext) {
	l.logger.Info("ExitReturnStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitYieldStatement is called when exiting the yieldStatement production.
func (l *LoggerListener) ExitYieldStatement(c *YieldStatementContext) {
	l.logger.Info("ExitYieldStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitWithStatement is called when exiting the withStatement production.
func (l *LoggerListener) ExitWithStatement(c *WithStatementContext) {
	l.logger.Info("ExitWithStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSwitchStatement is called when exiting the switchStatement production.
func (l *LoggerListener) ExitSwitchStatement(c *SwitchStatementContext) {
	l.logger.Info("ExitSwitchStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseBlock is called when exiting the caseBlock production.
func (l *LoggerListener) ExitCaseBlock(c *CaseBlockContext) {
	l.logger.Info("ExitCaseBlock called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseClauses is called when exiting the caseClauses production.
func (l *LoggerListener) ExitCaseClauses(c *CaseClausesContext) {
	l.logger.Info("ExitCaseClauses called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCaseClause is called when exiting the caseClause production.
func (l *LoggerListener) ExitCaseClause(c *CaseClauseContext) {
	l.logger.Info("ExitCaseClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDefaultClause is called when exiting the defaultClause production.
func (l *LoggerListener) ExitDefaultClause(c *DefaultClauseContext) {
	l.logger.Info("ExitDefaultClause called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLabelledStatement is called when exiting the labelledStatement production.
func (l *LoggerListener) ExitLabelledStatement(c *LabelledStatementContext) {
	l.logger.Info("ExitLabelledStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitThrowStatement is called when exiting the throwStatement production.
func (l *LoggerListener) ExitThrowStatement(c *ThrowStatementContext) {
	l.logger.Info("ExitThrowStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTryStatement is called when exiting the tryStatement production.
func (l *LoggerListener) ExitTryStatement(c *TryStatementContext) {
	l.logger.Info("ExitTryStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCatchProduction is called when exiting the catchProduction production.
func (l *LoggerListener) ExitCatchProduction(c *CatchProductionContext) {
	l.logger.Info("ExitCatchProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFinallyProduction is called when exiting the finallyProduction production.
func (l *LoggerListener) ExitFinallyProduction(c *FinallyProductionContext) {
	l.logger.Info("ExitFinallyProduction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDebuggerStatement is called when exiting the debuggerStatement production.
func (l *LoggerListener) ExitDebuggerStatement(c *DebuggerStatementContext) {
	l.logger.Info("ExitDebuggerStatement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
func (l *LoggerListener) ExitFunctionDeclaration(c *FunctionDeclarationContext) {
	l.logger.Info("ExitFunctionDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassDeclaration is called when exiting the classDeclaration production.
func (l *LoggerListener) ExitClassDeclaration(c *ClassDeclarationContext) {
	l.logger.Info("ExitClassDeclaration called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassTail is called when exiting the classTail production.
func (l *LoggerListener) ExitClassTail(c *ClassTailContext) {
	l.logger.Info("ExitClassTail called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassElement is called when exiting the classElement production.
func (l *LoggerListener) ExitClassElement(c *ClassElementContext) {
	l.logger.Info("ExitClassElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMethodDefinition is called when exiting the methodDefinition production.
func (l *LoggerListener) ExitMethodDefinition(c *MethodDefinitionContext) {
	l.logger.Info("ExitMethodDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFieldDefinition is called when exiting the fieldDefinition production.
func (l *LoggerListener) ExitFieldDefinition(c *FieldDefinitionContext) {
	l.logger.Info("ExitFieldDefinition called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassElementName is called when exiting the classElementName production.
func (l *LoggerListener) ExitClassElementName(c *ClassElementNameContext) {
	l.logger.Info("ExitClassElementName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPrivateIdentifier is called when exiting the privateIdentifier production.
func (l *LoggerListener) ExitPrivateIdentifier(c *PrivateIdentifierContext) {
	l.logger.Info("ExitPrivateIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFormalParameterList is called when exiting the formalParameterList production.
func (l *LoggerListener) ExitFormalParameterList(c *FormalParameterListContext) {
	l.logger.Info("ExitFormalParameterList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFormalParameterArg is called when exiting the formalParameterArg production.
func (l *LoggerListener) ExitFormalParameterArg(c *FormalParameterArgContext) {
	l.logger.Info("ExitFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLastFormalParameterArg is called when exiting the lastFormalParameterArg production.
func (l *LoggerListener) ExitLastFormalParameterArg(c *LastFormalParameterArgContext) {
	l.logger.Info("ExitLastFormalParameterArg called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionBody is called when exiting the functionBody production.
func (l *LoggerListener) ExitFunctionBody(c *FunctionBodyContext) {
	l.logger.Info("ExitFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSourceElements is called when exiting the sourceElements production.
func (l *LoggerListener) ExitSourceElements(c *SourceElementsContext) {
	l.logger.Info("ExitSourceElements called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayLiteral is called when exiting the arrayLiteral production.
func (l *LoggerListener) ExitArrayLiteral(c *ArrayLiteralContext) {
	l.logger.Info("ExitArrayLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitElementList is called when exiting the elementList production.
func (l *LoggerListener) ExitElementList(c *ElementListContext) {
	l.logger.Info("ExitElementList called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayElement is called when exiting the arrayElement production.
func (l *LoggerListener) ExitArrayElement(c *ArrayElementContext) {
	l.logger.Info("ExitArrayElement called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
func (l *LoggerListener) ExitPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext) {
	l.logger.Info("ExitPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitComputedPropertyExpressionAssignment is called when exiting the ComputedPropertyExpressionAssignment production.
func (l *LoggerListener) ExitComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext) {
	l.logger.Info("ExitComputedPropertyExpressionAssignment called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionProperty is called when exiting the FunctionProperty production.
func (l *LoggerListener) ExitFunctionProperty(c *FunctionPropertyContext) {
	l.logger.Info("ExitFunctionProperty called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyGetter is called when exiting the PropertyGetter production.
func (l *LoggerListener) ExitPropertyGetter(c *PropertyGetterContext) {
	l.logger.Info("ExitPropertyGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertySetter is called when exiting the PropertySetter production.
func (l *LoggerListener) ExitPropertySetter(c *PropertySetterContext) {
	l.logger.Info("ExitPropertySetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyShorthand is called when exiting the PropertyShorthand production.
func (l *LoggerListener) ExitPropertyShorthand(c *PropertyShorthandContext) {
	l.logger.Info("ExitPropertyShorthand called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPropertyName is called when exiting the propertyName production.
func (l *LoggerListener) ExitPropertyName(c *PropertyNameContext) {
	l.logger.Info("ExitPropertyName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArguments is called when exiting the arguments production.
func (l *LoggerListener) ExitArguments(c *ArgumentsContext) {
	l.logger.Info("ExitArguments called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArgument is called when exiting the argument production.
func (l *LoggerListener) ExitArgument(c *ArgumentContext) {
	l.logger.Info("ExitArgument called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitExpressionSequence is called when exiting the expressionSequence production.
func (l *LoggerListener) ExitExpressionSequence(c *ExpressionSequenceContext) {
	l.logger.Info("ExitExpressionSequence called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringExpression is called when exiting the TemplateStringExpression production.
func (l *LoggerListener) ExitTemplateStringExpression(c *TemplateStringExpressionContext) {
	l.logger.Info("ExitTemplateStringExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTernaryExpression is called when exiting the TernaryExpression production.
func (l *LoggerListener) ExitTernaryExpression(c *TernaryExpressionContext) {
	l.logger.Info("ExitTernaryExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
func (l *LoggerListener) ExitLogicalAndExpression(c *LogicalAndExpressionContext) {
	l.logger.Info("ExitLogicalAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPowerExpression is called when exiting the PowerExpression production.
func (l *LoggerListener) ExitPowerExpression(c *PowerExpressionContext) {
	l.logger.Info("ExitPowerExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPreIncrementExpression is called when exiting the PreIncrementExpression production.
func (l *LoggerListener) ExitPreIncrementExpression(c *PreIncrementExpressionContext) {
	l.logger.Info("ExitPreIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
func (l *LoggerListener) ExitObjectLiteralExpression(c *ObjectLiteralExpressionContext) {
	l.logger.Info("ExitObjectLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMetaExpression is called when exiting the MetaExpression production.
func (l *LoggerListener) ExitMetaExpression(c *MetaExpressionContext) {
	l.logger.Info("ExitMetaExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInExpression is called when exiting the InExpression production.
func (l *LoggerListener) ExitInExpression(c *InExpressionContext) {
	l.logger.Info("ExitInExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
func (l *LoggerListener) ExitLogicalOrExpression(c *LogicalOrExpressionContext) {
	l.logger.Info("ExitLogicalOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitOptionalChainExpression is called when exiting the OptionalChainExpression production.
func (l *LoggerListener) ExitOptionalChainExpression(c *OptionalChainExpressionContext) {
	l.logger.Info("ExitOptionalChainExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNotExpression is called when exiting the NotExpression production.
func (l *LoggerListener) ExitNotExpression(c *NotExpressionContext) {
	l.logger.Info("ExitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPreDecreaseExpression is called when exiting the PreDecreaseExpression production.
func (l *LoggerListener) ExitPreDecreaseExpression(c *PreDecreaseExpressionContext) {
	l.logger.Info("ExitPreDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
func (l *LoggerListener) ExitArgumentsExpression(c *ArgumentsExpressionContext) {
	l.logger.Info("ExitArgumentsExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAwaitExpression is called when exiting the AwaitExpression production.
func (l *LoggerListener) ExitAwaitExpression(c *AwaitExpressionContext) {
	l.logger.Info("ExitAwaitExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitThisExpression is called when exiting the ThisExpression production.
func (l *LoggerListener) ExitThisExpression(c *ThisExpressionContext) {
	l.logger.Info("ExitThisExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitFunctionExpression is called when exiting the FunctionExpression production.
func (l *LoggerListener) ExitFunctionExpression(c *FunctionExpressionContext) {
	l.logger.Info("ExitFunctionExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
func (l *LoggerListener) ExitUnaryMinusExpression(c *UnaryMinusExpressionContext) {
	l.logger.Info("ExitUnaryMinusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
func (l *LoggerListener) ExitAssignmentExpression(c *AssignmentExpressionContext) {
	l.logger.Info("ExitAssignmentExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPostDecreaseExpression is called when exiting the PostDecreaseExpression production.
func (l *LoggerListener) ExitPostDecreaseExpression(c *PostDecreaseExpressionContext) {
	l.logger.Info("ExitPostDecreaseExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTypeofExpression is called when exiting the TypeofExpression production.
func (l *LoggerListener) ExitTypeofExpression(c *TypeofExpressionContext) {
	l.logger.Info("ExitTypeofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInstanceofExpression is called when exiting the InstanceofExpression production.
func (l *LoggerListener) ExitInstanceofExpression(c *InstanceofExpressionContext) {
	l.logger.Info("ExitInstanceofExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
func (l *LoggerListener) ExitUnaryPlusExpression(c *UnaryPlusExpressionContext) {
	l.logger.Info("ExitUnaryPlusExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitDeleteExpression is called when exiting the DeleteExpression production.
func (l *LoggerListener) ExitDeleteExpression(c *DeleteExpressionContext) {
	l.logger.Info("ExitDeleteExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitImportExpression is called when exiting the ImportExpression production.
func (l *LoggerListener) ExitImportExpression(c *ImportExpressionContext) {
	l.logger.Info("ExitImportExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEqualityExpression is called when exiting the EqualityExpression production.
func (l *LoggerListener) ExitEqualityExpression(c *EqualityExpressionContext) {
	l.logger.Info("ExitEqualityExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitXOrExpression is called when exiting the BitXOrExpression production.
func (l *LoggerListener) ExitBitXOrExpression(c *BitXOrExpressionContext) {
	l.logger.Info("ExitBitXOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSuperExpression is called when exiting the SuperExpression production.
func (l *LoggerListener) ExitSuperExpression(c *SuperExpressionContext) {
	l.logger.Info("ExitSuperExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
func (l *LoggerListener) ExitMultiplicativeExpression(c *MultiplicativeExpressionContext) {
	l.logger.Info("ExitMultiplicativeExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitShiftExpression is called when exiting the BitShiftExpression production.
func (l *LoggerListener) ExitBitShiftExpression(c *BitShiftExpressionContext) {
	l.logger.Info("ExitBitShiftExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
func (l *LoggerListener) ExitParenthesizedExpression(c *ParenthesizedExpressionContext) {
	l.logger.Info("ExitParenthesizedExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
func (l *LoggerListener) ExitAdditiveExpression(c *AdditiveExpressionContext) {
	l.logger.Info("ExitAdditiveExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitRelationalExpression is called when exiting the RelationalExpression production.
func (l *LoggerListener) ExitRelationalExpression(c *RelationalExpressionContext) {
	l.logger.Info("ExitRelationalExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitPostIncrementExpression is called when exiting the PostIncrementExpression production.
func (l *LoggerListener) ExitPostIncrementExpression(c *PostIncrementExpressionContext) {
	l.logger.Info("ExitPostIncrementExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitYieldExpression is called when exiting the YieldExpression production.
func (l *LoggerListener) ExitYieldExpression(c *YieldExpressionContext) {
	l.logger.Info("ExitYieldExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitNotExpression is called when exiting the BitNotExpression production.
func (l *LoggerListener) ExitBitNotExpression(c *BitNotExpressionContext) {
	l.logger.Info("ExitBitNotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNewExpression is called when exiting the NewExpression production.
func (l *LoggerListener) ExitNewExpression(c *NewExpressionContext) {
	l.logger.Info("ExitNewExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLiteralExpression is called when exiting the LiteralExpression production.
func (l *LoggerListener) ExitLiteralExpression(c *LiteralExpressionContext) {
	l.logger.Info("ExitLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrayLiteralExpression is called when exiting the ArrayLiteralExpression production.
func (l *LoggerListener) ExitArrayLiteralExpression(c *ArrayLiteralExpressionContext) {
	l.logger.Info("ExitArrayLiteralExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMemberDotExpression is called when exiting the MemberDotExpression production.
func (l *LoggerListener) ExitMemberDotExpression(c *MemberDotExpressionContext) {
	l.logger.Info("ExitMemberDotExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitClassExpression is called when exiting the ClassExpression production.
func (l *LoggerListener) ExitClassExpression(c *ClassExpressionContext) {
	l.logger.Info("ExitClassExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
func (l *LoggerListener) ExitMemberIndexExpression(c *MemberIndexExpressionContext) {
	l.logger.Info("ExitMemberIndexExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
func (l *LoggerListener) ExitIdentifierExpression(c *IdentifierExpressionContext) {
	l.logger.Info("ExitIdentifierExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitAndExpression is called when exiting the BitAndExpression production.
func (l *LoggerListener) ExitBitAndExpression(c *BitAndExpressionContext) {
	l.logger.Info("ExitBitAndExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBitOrExpression is called when exiting the BitOrExpression production.
func (l *LoggerListener) ExitBitOrExpression(c *BitOrExpressionContext) {
	l.logger.Info("ExitBitOrExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentOperatorExpression is called when exiting the AssignmentOperatorExpression production.
func (l *LoggerListener) ExitAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext) {
	l.logger.Info("ExitAssignmentOperatorExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitVoidExpression is called when exiting the VoidExpression production.
func (l *LoggerListener) ExitVoidExpression(c *VoidExpressionContext) {
	l.logger.Info("ExitVoidExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitCoalesceExpression is called when exiting the CoalesceExpression production.
func (l *LoggerListener) ExitCoalesceExpression(c *CoalesceExpressionContext) {
	l.logger.Info("ExitCoalesceExpression called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitInitializer is called when exiting the initializer production.
func (l *LoggerListener) ExitInitializer(c *InitializerContext) {
	l.logger.Info("ExitInitializer called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignable is called when exiting the assignable production.
func (l *LoggerListener) ExitAssignable(c *AssignableContext) {
	l.logger.Info("ExitAssignable called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitObjectLiteral is called when exiting the objectLiteral production.
func (l *LoggerListener) ExitObjectLiteral(c *ObjectLiteralContext) {
	l.logger.Info("ExitObjectLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNamedFunction is called when exiting the NamedFunction production.
func (l *LoggerListener) ExitNamedFunction(c *NamedFunctionContext) {
	l.logger.Info("ExitNamedFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAnonymousFunctionDecl is called when exiting the AnonymousFunctionDecl production.
func (l *LoggerListener) ExitAnonymousFunctionDecl(c *AnonymousFunctionDeclContext) {
	l.logger.Info("ExitAnonymousFunctionDecl called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunction is called when exiting the ArrowFunction production.
func (l *LoggerListener) ExitArrowFunction(c *ArrowFunctionContext) {
	l.logger.Info("ExitArrowFunction called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunctionParameters is called when exiting the arrowFunctionParameters production.
func (l *LoggerListener) ExitArrowFunctionParameters(c *ArrowFunctionParametersContext) {
	l.logger.Info("ExitArrowFunctionParameters called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitArrowFunctionBody is called when exiting the arrowFunctionBody production.
func (l *LoggerListener) ExitArrowFunctionBody(c *ArrowFunctionBodyContext) {
	l.logger.Info("ExitArrowFunctionBody called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitAssignmentOperator is called when exiting the assignmentOperator production.
func (l *LoggerListener) ExitAssignmentOperator(c *AssignmentOperatorContext) {
	l.logger.Info("ExitAssignmentOperator called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLiteral is called when exiting the literal production.
func (l *LoggerListener) ExitLiteral(c *LiteralContext) {
	l.logger.Info("ExitLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringLiteral is called when exiting the templateStringLiteral production.
func (l *LoggerListener) ExitTemplateStringLiteral(c *TemplateStringLiteralContext) {
	l.logger.Info("ExitTemplateStringLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitTemplateStringAtom is called when exiting the templateStringAtom production.
func (l *LoggerListener) ExitTemplateStringAtom(c *TemplateStringAtomContext) {
	l.logger.Info("ExitTemplateStringAtom called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitNumericLiteral is called when exiting the numericLiteral production.
func (l *LoggerListener) ExitNumericLiteral(c *NumericLiteralContext) {
	l.logger.Info("ExitNumericLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitBigintLiteral is called when exiting the bigintLiteral production.
func (l *LoggerListener) ExitBigintLiteral(c *BigintLiteralContext) {
	l.logger.Info("ExitBigintLiteral called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitGetter is called when exiting the getter production.
func (l *LoggerListener) ExitGetter(c *GetterContext) {
	l.logger.Info("ExitGetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitSetter is called when exiting the setter production.
func (l *LoggerListener) ExitSetter(c *SetterContext) {
	l.logger.Info("ExitSetter called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifierName is called when exiting the identifierName production.
func (l *LoggerListener) ExitIdentifierName(c *IdentifierNameContext) {
	l.logger.Info("ExitIdentifierName called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitIdentifier is called when exiting the identifier production.
func (l *LoggerListener) ExitIdentifier(c *IdentifierContext) {
	l.logger.Info("ExitIdentifier called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitReservedWord is called when exiting the reservedWord production.
func (l *LoggerListener) ExitReservedWord(c *ReservedWordContext) {
	l.logger.Info("ExitReservedWord called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitKeyword is called when exiting the keyword production.
func (l *LoggerListener) ExitKeyword(c *KeywordContext) {
	l.logger.Info("ExitKeyword called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitLet_ is called when exiting the let_ production.
func (l *LoggerListener) ExitLet_(c *Let_Context) {
	l.logger.Info("ExitLet_ called", slog.Int("RuleIndex", c.RuleIndex))
}

// ExitEos is called when exiting the eos production.
func (l *LoggerListener) ExitEos(c *EosContext) {
	l.logger.Info("ExitEos called", slog.Int("RuleIndex", c.RuleIndex))
}
