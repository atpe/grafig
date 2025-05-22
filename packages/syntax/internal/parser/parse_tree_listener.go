package parser

import (
	"fmt"
	"log/slog"

	"github.com/antlr4-go/antlr/v4"
)

type ParseTreeListener struct {
	JavaScriptParserListener
	logger *slog.Logger
	indent string
}

func NewParseTreeListener(logger *slog.Logger) *ParseTreeListener {
	listener := new(ParseTreeListener)
	listener.JavaScriptParserListener = new(BaseJavaScriptParserListener)
	listener.logger = logger
	return listener
}

func (l *ParseTreeListener) enter() {
	l.indent += "  "
}

func (l *ParseTreeListener) exit() {
	l.indent = l.indent[:len(l.indent)-2]
}

// EnterProgram is called when entering the program production.
func (l *ParseTreeListener) EnterProgram(c *ProgramContext) {
	l.enter()
	fmt.Println(l.indent + "Program")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterSourceElement is called when entering the sourceElement production.
func (l *ParseTreeListener) EnterSourceElement(c *SourceElementContext) {
	l.enter()
	fmt.Println(l.indent + "SourceElement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterStatement is called when entering the statement production.
func (l *ParseTreeListener) EnterStatement(c *StatementContext) {
	l.enter()
	fmt.Println(l.indent + "Statement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterBlock is called when entering the block production.
func (l *ParseTreeListener) EnterBlock(c *BlockContext) {
	l.enter()
	fmt.Println(l.indent + "Block")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterStatementList is called when entering the statementList production.
func (l *ParseTreeListener) EnterStatementList(c *StatementListContext) {
	l.enter()
	fmt.Println(l.indent + "StatementList")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportStatement is called when entering the importStatement production.
func (l *ParseTreeListener) EnterImportStatement(c *ImportStatementContext) {
	l.enter()
	fmt.Println(l.indent + "ImportStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportFromBlock is called when entering the importFromBlock production.
func (l *ParseTreeListener) EnterImportFromBlock(c *ImportFromBlockContext) {
	l.enter()
	fmt.Println(l.indent + "ImportFromBlock")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportModuleItems is called when entering the importModuleItems production.
func (l *ParseTreeListener) EnterImportModuleItems(c *ImportModuleItemsContext) {
	l.enter()
	fmt.Println(l.indent + "ImportModuleItems")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportAliasName is called when entering the importAliasName production.
func (l *ParseTreeListener) EnterImportAliasName(c *ImportAliasNameContext) {
	l.enter()
	fmt.Println(l.indent + "ImportAliasName")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterModuleExportName is called when entering the moduleExportName production.
func (l *ParseTreeListener) EnterModuleExportName(c *ModuleExportNameContext) {
	l.enter()
	fmt.Println(l.indent + "ModuleExportName")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportedBinding is called when entering the importedBinding production.
func (l *ParseTreeListener) EnterImportedBinding(c *ImportedBindingContext) {
	l.enter()
	fmt.Println(l.indent + "ImportedBinding")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportDefault is called when entering the importDefault production.
func (l *ParseTreeListener) EnterImportDefault(c *ImportDefaultContext) {
	l.enter()
	fmt.Println(l.indent + "ImportDefault")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportNamespace is called when entering the importNamespace production.
func (l *ParseTreeListener) EnterImportNamespace(c *ImportNamespaceContext) {
	l.enter()
	fmt.Println(l.indent + "ImportNamespace")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportFrom is called when entering the importFrom production.
func (l *ParseTreeListener) EnterImportFrom(c *ImportFromContext) {
	l.enter()
	fmt.Println(l.indent + "ImportFrom")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterAliasName is called when entering the aliasName production.
func (l *ParseTreeListener) EnterAliasName(c *AliasNameContext) {
	l.enter()
	fmt.Println(l.indent + "AliasName")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterExportDeclaration is called when entering the ExportDeclaration production.
func (l *ParseTreeListener) EnterExportDeclaration(c *ExportDeclarationContext) {
	l.enter()
	fmt.Println(l.indent + "ExportDeclaration")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterExportDefaultDeclaration is called when entering the ExportDefaultDeclaration production.
func (l *ParseTreeListener) EnterExportDefaultDeclaration(c *ExportDefaultDeclarationContext) {
	l.enter()
	fmt.Println(l.indent + "ExportDefaultDeclaration")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterExportFromBlock is called when entering the exportFromBlock production.
func (l *ParseTreeListener) EnterExportFromBlock(c *ExportFromBlockContext) {
	l.enter()
	fmt.Println(l.indent + "ExportFromBlock")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterExportModuleItems is called when entering the exportModuleItems production.
func (l *ParseTreeListener) EnterExportModuleItems(c *ExportModuleItemsContext) {
	l.enter()
	fmt.Println(l.indent + "ExportModuleItems")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterExportAliasName is called when entering the exportAliasName production.
func (l *ParseTreeListener) EnterExportAliasName(c *ExportAliasNameContext) {
	l.enter()
	fmt.Println(l.indent + "ExportAliasName")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterDeclaration is called when entering the declaration production.
func (l *ParseTreeListener) EnterDeclaration(c *DeclarationContext) {
	l.enter()
	fmt.Println(l.indent + "Declaration")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterVariableStatement is called when entering the variableStatement production.
func (l *ParseTreeListener) EnterVariableStatement(c *VariableStatementContext) {
	l.enter()
	fmt.Println(l.indent + "VariableStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
func (l *ParseTreeListener) EnterVariableDeclarationList(c *VariableDeclarationListContext) {
	l.enter()
	fmt.Println(l.indent + "VariableDeclarationList")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterVariableDeclaration is called when entering the variableDeclaration production.
func (l *ParseTreeListener) EnterVariableDeclaration(c *VariableDeclarationContext) {
	l.enter()
	fmt.Println(l.indent + "VariableDeclaration")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
func (l *ParseTreeListener) EnterEmptyStatement_(c *EmptyStatement_Context) {
	l.enter()
	fmt.Println(l.indent + "EmptyStatement_")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterExpressionStatement is called when entering the expressionStatement production.
func (l *ParseTreeListener) EnterExpressionStatement(c *ExpressionStatementContext) {
	l.enter()
	fmt.Println(l.indent + "ExpressionStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterIfStatement is called when entering the ifStatement production.
func (l *ParseTreeListener) EnterIfStatement(c *IfStatementContext) {
	l.enter()
	fmt.Println(l.indent + "IfStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterDoStatement is called when entering the DoStatement production.
func (l *ParseTreeListener) EnterDoStatement(c *DoStatementContext) {
	l.enter()
	fmt.Println(l.indent + "DoStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterWhileStatement is called when entering the WhileStatement production.
func (l *ParseTreeListener) EnterWhileStatement(c *WhileStatementContext) {
	l.enter()
	fmt.Println(l.indent + "WhileStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterForStatement is called when entering the ForStatement production.
func (l *ParseTreeListener) EnterForStatement(c *ForStatementContext) {
	l.enter()
	fmt.Println(l.indent + "ForStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterForInStatement is called when entering the ForInStatement production.
func (l *ParseTreeListener) EnterForInStatement(c *ForInStatementContext) {
	l.enter()
	fmt.Println(l.indent + "ForInStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterForOfStatement is called when entering the ForOfStatement production.
func (l *ParseTreeListener) EnterForOfStatement(c *ForOfStatementContext) {
	l.enter()
	fmt.Println(l.indent + "ForOfStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterVarModifier is called when entering the varModifier production.
func (l *ParseTreeListener) EnterVarModifier(c *VarModifierContext) {
	l.enter()
	fmt.Println(l.indent + "VarModifier")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterContinueStatement is called when entering the continueStatement production.
func (l *ParseTreeListener) EnterContinueStatement(c *ContinueStatementContext) {
	l.enter()
	fmt.Println(l.indent + "ContinueStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterBreakStatement is called when entering the breakStatement production.
func (l *ParseTreeListener) EnterBreakStatement(c *BreakStatementContext) {
	l.enter()
	fmt.Println(l.indent + "BreakStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterReturnStatement is called when entering the returnStatement production.
func (l *ParseTreeListener) EnterReturnStatement(c *ReturnStatementContext) {
	l.enter()
	fmt.Println(l.indent + "ReturnStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterYieldStatement is called when entering the yieldStatement production.
func (l *ParseTreeListener) EnterYieldStatement(c *YieldStatementContext) {
	l.enter()
	fmt.Println(l.indent + "YieldStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterWithStatement is called when entering the withStatement production.
func (l *ParseTreeListener) EnterWithStatement(c *WithStatementContext) {
	l.enter()
	fmt.Println(l.indent + "WithStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterSwitchStatement is called when entering the switchStatement production.
func (l *ParseTreeListener) EnterSwitchStatement(c *SwitchStatementContext) {
	l.enter()
	fmt.Println(l.indent + "SwitchStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterCaseBlock is called when entering the caseBlock production.
func (l *ParseTreeListener) EnterCaseBlock(c *CaseBlockContext) {
	l.enter()
	fmt.Println(l.indent + "CaseBlock")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterCaseClauses is called when entering the caseClauses production.
func (l *ParseTreeListener) EnterCaseClauses(c *CaseClausesContext) {
	l.enter()
	fmt.Println(l.indent + "CaseClauses")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterCaseClause is called when entering the caseClause production.
func (l *ParseTreeListener) EnterCaseClause(c *CaseClauseContext) {
	l.enter()
	fmt.Println(l.indent + "CaseClause")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterDefaultClause is called when entering the defaultClause production.
func (l *ParseTreeListener) EnterDefaultClause(c *DefaultClauseContext) {
	l.enter()
	fmt.Println(l.indent + "DefaultClause")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterLabelledStatement is called when entering the labelledStatement production.
func (l *ParseTreeListener) EnterLabelledStatement(c *LabelledStatementContext) {
	l.enter()
	fmt.Println(l.indent + "LabelledStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterThrowStatement is called when entering the throwStatement production.
func (l *ParseTreeListener) EnterThrowStatement(c *ThrowStatementContext) {
	l.enter()
	fmt.Println(l.indent + "ThrowStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterTryStatement is called when entering the tryStatement production.
func (l *ParseTreeListener) EnterTryStatement(c *TryStatementContext) {
	l.enter()
	fmt.Println(l.indent + "TryStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterCatchProduction is called when entering the catchProduction production.
func (l *ParseTreeListener) EnterCatchProduction(c *CatchProductionContext) {
	l.enter()
	fmt.Println(l.indent + "CatchProduction")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterFinallyProduction is called when entering the finallyProduction production.
func (l *ParseTreeListener) EnterFinallyProduction(c *FinallyProductionContext) {
	l.enter()
	fmt.Println(l.indent + "FinallyProduction")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterDebuggerStatement is called when entering the debuggerStatement production.
func (l *ParseTreeListener) EnterDebuggerStatement(c *DebuggerStatementContext) {
	l.enter()
	fmt.Println(l.indent + "DebuggerStatement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterFunctionDeclaration is called when entering the functionDeclaration production.
func (l *ParseTreeListener) EnterFunctionDeclaration(c *FunctionDeclarationContext) {
	l.enter()
	fmt.Println(l.indent + "FunctionDeclaration")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterClassDeclaration is called when entering the classDeclaration production.
func (l *ParseTreeListener) EnterClassDeclaration(c *ClassDeclarationContext) {
	l.enter()
	fmt.Println(l.indent + "ClassDeclaration")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterClassTail is called when entering the classTail production.
func (l *ParseTreeListener) EnterClassTail(c *ClassTailContext) {
	l.enter()
	fmt.Println(l.indent + "ClassTail")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterClassElement is called when entering the classElement production.
func (l *ParseTreeListener) EnterClassElement(c *ClassElementContext) {
	l.enter()
	fmt.Println(l.indent + "ClassElement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterMethodDefinition is called when entering the methodDefinition production.
func (l *ParseTreeListener) EnterMethodDefinition(c *MethodDefinitionContext) {
	l.enter()
	fmt.Println(l.indent + "MethodDefinition")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterFieldDefinition is called when entering the fieldDefinition production.
func (l *ParseTreeListener) EnterFieldDefinition(c *FieldDefinitionContext) {
	l.enter()
	fmt.Println(l.indent + "FieldDefinition")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterClassElementName is called when entering the classElementName production.
func (l *ParseTreeListener) EnterClassElementName(c *ClassElementNameContext) {
	l.enter()
	fmt.Println(l.indent + "ClassElementName")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPrivateIdentifier is called when entering the privateIdentifier production.
func (l *ParseTreeListener) EnterPrivateIdentifier(c *PrivateIdentifierContext) {
	l.enter()
	fmt.Println(l.indent + "PrivateIdentifier")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterFormalParameterList is called when entering the formalParameterList production.
func (l *ParseTreeListener) EnterFormalParameterList(c *FormalParameterListContext) {
	l.enter()
	fmt.Println(l.indent + "FormalParameterList")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterFormalParameterArg is called when entering the formalParameterArg production.
func (l *ParseTreeListener) EnterFormalParameterArg(c *FormalParameterArgContext) {
	l.enter()
	fmt.Println(l.indent + "FormalParameterArg")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterLastFormalParameterArg is called when entering the lastFormalParameterArg production.
func (l *ParseTreeListener) EnterLastFormalParameterArg(c *LastFormalParameterArgContext) {
	l.enter()
	fmt.Println(l.indent + "LastFormalParameterArg")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterFunctionBody is called when entering the functionBody production.
func (l *ParseTreeListener) EnterFunctionBody(c *FunctionBodyContext) {
	l.enter()
	fmt.Println(l.indent + "FunctionBody")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterSourceElements is called when entering the sourceElements production.
func (l *ParseTreeListener) EnterSourceElements(c *SourceElementsContext) {
	l.enter()
	fmt.Println(l.indent + "SourceElements")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArrayLiteral is called when entering the arrayLiteral production.
func (l *ParseTreeListener) EnterArrayLiteral(c *ArrayLiteralContext) {
	l.enter()
	fmt.Println(l.indent + "ArrayLiteral")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterElementList is called when entering the elementList production.
func (l *ParseTreeListener) EnterElementList(c *ElementListContext) {
	l.enter()
	fmt.Println(l.indent + "ElementList")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArrayElement is called when entering the arrayElement production.
func (l *ParseTreeListener) EnterArrayElement(c *ArrayElementContext) {
	l.enter()
	fmt.Println(l.indent + "ArrayElement")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
func (l *ParseTreeListener) EnterPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext) {
	l.enter()
	fmt.Println(l.indent + "PropertyExpressionAssignment")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterComputedPropertyExpressionAssignment is called when entering the ComputedPropertyExpressionAssignment production.
func (l *ParseTreeListener) EnterComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext) {
	l.enter()
	fmt.Println(l.indent + "ComputedPropertyExpressionAssignment")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterFunctionProperty is called when entering the FunctionProperty production.
func (l *ParseTreeListener) EnterFunctionProperty(c *FunctionPropertyContext) {
	l.enter()
	fmt.Println(l.indent + "FunctionProperty")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPropertyGetter is called when entering the PropertyGetter production.
func (l *ParseTreeListener) EnterPropertyGetter(c *PropertyGetterContext) {
	l.enter()
	fmt.Println(l.indent + "PropertyGetter")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPropertySetter is called when entering the PropertySetter production.
func (l *ParseTreeListener) EnterPropertySetter(c *PropertySetterContext) {
	l.enter()
	fmt.Println(l.indent + "PropertySetter")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPropertyShorthand is called when entering the PropertyShorthand production.
func (l *ParseTreeListener) EnterPropertyShorthand(c *PropertyShorthandContext) {
	l.enter()
	fmt.Println(l.indent + "PropertyShorthand")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPropertyName is called when entering the propertyName production.
func (l *ParseTreeListener) EnterPropertyName(c *PropertyNameContext) {
	l.enter()
	fmt.Println(l.indent + "PropertyName")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArguments is called when entering the arguments production.
func (l *ParseTreeListener) EnterArguments(c *ArgumentsContext) {
	l.enter()
	fmt.Println(l.indent + "Arguments")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArgument is called when entering the argument production.
func (l *ParseTreeListener) EnterArgument(c *ArgumentContext) {
	l.enter()
	fmt.Println(l.indent + "Argument")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterExpressionSequence is called when entering the expressionSequence production.
func (l *ParseTreeListener) EnterExpressionSequence(c *ExpressionSequenceContext) {
	l.enter()
	fmt.Println(l.indent + "ExpressionSequence")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterTemplateStringExpression is called when entering the TemplateStringExpression production.
func (l *ParseTreeListener) EnterTemplateStringExpression(c *TemplateStringExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "TemplateStringExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterTernaryExpression is called when entering the TernaryExpression production.
func (l *ParseTreeListener) EnterTernaryExpression(c *TernaryExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "TernaryExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
func (l *ParseTreeListener) EnterLogicalAndExpression(c *LogicalAndExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "LogicalAndExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPowerExpression is called when entering the PowerExpression production.
func (l *ParseTreeListener) EnterPowerExpression(c *PowerExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "PowerExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPreIncrementExpression is called when entering the PreIncrementExpression production.
func (l *ParseTreeListener) EnterPreIncrementExpression(c *PreIncrementExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "PreIncrementExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
func (l *ParseTreeListener) EnterObjectLiteralExpression(c *ObjectLiteralExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "ObjectLiteralExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterMetaExpression is called when entering the MetaExpression production.
func (l *ParseTreeListener) EnterMetaExpression(c *MetaExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "MetaExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterInExpression is called when entering the InExpression production.
func (l *ParseTreeListener) EnterInExpression(c *InExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "InExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
func (l *ParseTreeListener) EnterLogicalOrExpression(c *LogicalOrExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "LogicalOrExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterOptionalChainExpression is called when entering the OptionalChainExpression production.
func (l *ParseTreeListener) EnterOptionalChainExpression(c *OptionalChainExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "OptionalChainExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterNotExpression is called when entering the NotExpression production.
func (l *ParseTreeListener) EnterNotExpression(c *NotExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "NotExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPreDecreaseExpression is called when entering the PreDecreaseExpression production.
func (l *ParseTreeListener) EnterPreDecreaseExpression(c *PreDecreaseExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "PreDecreaseExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
func (l *ParseTreeListener) EnterArgumentsExpression(c *ArgumentsExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "ArgumentsExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterAwaitExpression is called when entering the AwaitExpression production.
func (l *ParseTreeListener) EnterAwaitExpression(c *AwaitExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "AwaitExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterThisExpression is called when entering the ThisExpression production.
func (l *ParseTreeListener) EnterThisExpression(c *ThisExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "ThisExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterFunctionExpression is called when entering the FunctionExpression production.
func (l *ParseTreeListener) EnterFunctionExpression(c *FunctionExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "FunctionExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
func (l *ParseTreeListener) EnterUnaryMinusExpression(c *UnaryMinusExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "UnaryMinusExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterAssignmentExpression is called when entering the AssignmentExpression production.
func (l *ParseTreeListener) EnterAssignmentExpression(c *AssignmentExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "AssignmentExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPostDecreaseExpression is called when entering the PostDecreaseExpression production.
func (l *ParseTreeListener) EnterPostDecreaseExpression(c *PostDecreaseExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "PostDecreaseExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterTypeofExpression is called when entering the TypeofExpression production.
func (l *ParseTreeListener) EnterTypeofExpression(c *TypeofExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "TypeofExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterInstanceofExpression is called when entering the InstanceofExpression production.
func (l *ParseTreeListener) EnterInstanceofExpression(c *InstanceofExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "InstanceofExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
func (l *ParseTreeListener) EnterUnaryPlusExpression(c *UnaryPlusExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "UnaryPlusExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterDeleteExpression is called when entering the DeleteExpression production.
func (l *ParseTreeListener) EnterDeleteExpression(c *DeleteExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "DeleteExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterImportExpression is called when entering the ImportExpression production.
func (l *ParseTreeListener) EnterImportExpression(c *ImportExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "ImportExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterEqualityExpression is called when entering the EqualityExpression production.
func (l *ParseTreeListener) EnterEqualityExpression(c *EqualityExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "EqualityExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterBitXOrExpression is called when entering the BitXOrExpression production.
func (l *ParseTreeListener) EnterBitXOrExpression(c *BitXOrExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "BitXOrExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterSuperExpression is called when entering the SuperExpression production.
func (l *ParseTreeListener) EnterSuperExpression(c *SuperExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "SuperExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
func (l *ParseTreeListener) EnterMultiplicativeExpression(c *MultiplicativeExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "MultiplicativeExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterBitShiftExpression is called when entering the BitShiftExpression production.
func (l *ParseTreeListener) EnterBitShiftExpression(c *BitShiftExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "BitShiftExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
func (l *ParseTreeListener) EnterParenthesizedExpression(c *ParenthesizedExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "ParenthesizedExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterAdditiveExpression is called when entering the AdditiveExpression production.
func (l *ParseTreeListener) EnterAdditiveExpression(c *AdditiveExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "AdditiveExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterRelationalExpression is called when entering the RelationalExpression production.
func (l *ParseTreeListener) EnterRelationalExpression(c *RelationalExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "RelationalExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterPostIncrementExpression is called when entering the PostIncrementExpression production.
func (l *ParseTreeListener) EnterPostIncrementExpression(c *PostIncrementExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "PostIncrementExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterYieldExpression is called when entering the YieldExpression production.
func (l *ParseTreeListener) EnterYieldExpression(c *YieldExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "YieldExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterBitNotExpression is called when entering the BitNotExpression production.
func (l *ParseTreeListener) EnterBitNotExpression(c *BitNotExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "BitNotExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterNewExpression is called when entering the NewExpression production.
func (l *ParseTreeListener) EnterNewExpression(c *NewExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "NewExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterLiteralExpression is called when entering the LiteralExpression production.
func (l *ParseTreeListener) EnterLiteralExpression(c *LiteralExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "LiteralExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArrayLiteralExpression is called when entering the ArrayLiteralExpression production.
func (l *ParseTreeListener) EnterArrayLiteralExpression(c *ArrayLiteralExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "ArrayLiteralExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterMemberDotExpression is called when entering the MemberDotExpression production.
func (l *ParseTreeListener) EnterMemberDotExpression(c *MemberDotExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "MemberDotExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterClassExpression is called when entering the ClassExpression production.
func (l *ParseTreeListener) EnterClassExpression(c *ClassExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "ClassExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
func (l *ParseTreeListener) EnterMemberIndexExpression(c *MemberIndexExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "MemberIndexExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterIdentifierExpression is called when entering the IdentifierExpression production.
func (l *ParseTreeListener) EnterIdentifierExpression(c *IdentifierExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "IdentifierExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterBitAndExpression is called when entering the BitAndExpression production.
func (l *ParseTreeListener) EnterBitAndExpression(c *BitAndExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "BitAndExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterBitOrExpression is called when entering the BitOrExpression production.
func (l *ParseTreeListener) EnterBitOrExpression(c *BitOrExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "BitOrExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterAssignmentOperatorExpression is called when entering the AssignmentOperatorExpression production.
func (l *ParseTreeListener) EnterAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "AssignmentOperatorExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterVoidExpression is called when entering the VoidExpression production.
func (l *ParseTreeListener) EnterVoidExpression(c *VoidExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "VoidExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterCoalesceExpression is called when entering the CoalesceExpression production.
func (l *ParseTreeListener) EnterCoalesceExpression(c *CoalesceExpressionContext) {
	l.enter()
	fmt.Println(l.indent + "CoalesceExpression")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterInitializer is called when entering the initializer production.
func (l *ParseTreeListener) EnterInitializer(c *InitializerContext) {
	l.enter()
	fmt.Println(l.indent + "Initializer")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterAssignable is called when entering the assignable production.
func (l *ParseTreeListener) EnterAssignable(c *AssignableContext) {
	l.enter()
	fmt.Println(l.indent + "Assignable")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterObjectLiteral is called when entering the objectLiteral production.
func (l *ParseTreeListener) EnterObjectLiteral(c *ObjectLiteralContext) {
	l.enter()
	fmt.Println(l.indent + "ObjectLiteral")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterNamedFunction is called when entering the NamedFunction production.
func (l *ParseTreeListener) EnterNamedFunction(c *NamedFunctionContext) {
	l.enter()
	fmt.Println(l.indent + "NamedFunction")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterAnonymousFunctionDecl is called when entering the AnonymousFunctionDecl production.
func (l *ParseTreeListener) EnterAnonymousFunctionDecl(c *AnonymousFunctionDeclContext) {
	l.enter()
	fmt.Println(l.indent + "AnonymousFunctionDecl")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArrowFunction is called when entering the ArrowFunction production.
func (l *ParseTreeListener) EnterArrowFunction(c *ArrowFunctionContext) {
	l.enter()
	fmt.Println(l.indent + "ArrowFunction")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArrowFunctionParameters is called when entering the arrowFunctionParameters production.
func (l *ParseTreeListener) EnterArrowFunctionParameters(c *ArrowFunctionParametersContext) {
	l.enter()
	fmt.Println(l.indent + "ArrowFunctionParameters")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterArrowFunctionBody is called when entering the arrowFunctionBody production.
func (l *ParseTreeListener) EnterArrowFunctionBody(c *ArrowFunctionBodyContext) {
	l.enter()
	fmt.Println(l.indent + "ArrowFunctionBody")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterAssignmentOperator is called when entering the assignmentOperator production.
func (l *ParseTreeListener) EnterAssignmentOperator(c *AssignmentOperatorContext) {
	l.enter()
	fmt.Println(l.indent + "AssignmentOperator")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterLiteral is called when entering the literal production.
func (l *ParseTreeListener) EnterLiteral(c *LiteralContext) {
	l.enter()
	fmt.Println(l.indent + "Literal")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterTemplateStringLiteral is called when entering the templateStringLiteral production.
func (l *ParseTreeListener) EnterTemplateStringLiteral(c *TemplateStringLiteralContext) {
	l.enter()
	fmt.Println(l.indent + "TemplateStringLiteral")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterTemplateStringAtom is called when entering the templateStringAtom production.
func (l *ParseTreeListener) EnterTemplateStringAtom(c *TemplateStringAtomContext) {
	l.enter()
	fmt.Println(l.indent + "TemplateStringAtom")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterNumericLiteral is called when entering the numericLiteral production.
func (l *ParseTreeListener) EnterNumericLiteral(c *NumericLiteralContext) {
	l.enter()
	fmt.Println(l.indent + "NumericLiteral")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterBigintLiteral is called when entering the bigintLiteral production.
func (l *ParseTreeListener) EnterBigintLiteral(c *BigintLiteralContext) {
	l.enter()
	fmt.Println(l.indent + "BigintLiteral")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterGetter is called when entering the getter production.
func (l *ParseTreeListener) EnterGetter(c *GetterContext) {
	l.enter()
	fmt.Println(l.indent + "Getter")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterSetter is called when entering the setter production.
func (l *ParseTreeListener) EnterSetter(c *SetterContext) {
	l.enter()
	fmt.Println(l.indent + "Setter")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterIdentifierName is called when entering the identifierName production.
func (l *ParseTreeListener) EnterIdentifierName(c *IdentifierNameContext) {
	l.enter()
	fmt.Println(l.indent + "IdentifierName")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterIdentifier is called when entering the identifier production.
func (l *ParseTreeListener) EnterIdentifier(c *IdentifierContext) {
	l.enter()
	fmt.Println(l.indent + "Identifier")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterReservedWord is called when entering the reservedWord production.
func (l *ParseTreeListener) EnterReservedWord(c *ReservedWordContext) {
	l.enter()
	fmt.Println(l.indent + "ReservedWord")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterKeyword is called when entering the keyword production.
func (l *ParseTreeListener) EnterKeyword(c *KeywordContext) {
	l.enter()
	fmt.Println(l.indent + "Keyword")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterLet_ is called when entering the let_ production.
func (l *ParseTreeListener) EnterLet_(c *Let_Context) {
	l.enter()
	fmt.Println(l.indent + "Let_")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// EnterEos is called when entering the eos production.
func (l *ParseTreeListener) EnterEos(c *EosContext) {
	l.enter()
	fmt.Println(l.indent + "Eos")
	for _, child := range c.GetChildren() {
		if c, ok := child.(antlr.TerminalNode); ok {
			fmt.Println(l.indent + c.GetText())
		}
	}

}

// ExitProgram is called when exiting the program production.
func (l *ParseTreeListener) ExitProgram(c *ProgramContext) {
	l.exit()
}

// ExitSourceElement is called when exiting the sourceElement production.
func (l *ParseTreeListener) ExitSourceElement(c *SourceElementContext) {
	l.exit()
}

// ExitStatement is called when exiting the statement production.
func (l *ParseTreeListener) ExitStatement(c *StatementContext) {
	l.exit()
}

// ExitBlock is called when exiting the block production.
func (l *ParseTreeListener) ExitBlock(c *BlockContext) {
	l.exit()
}

// ExitStatementList is called when exiting the statementList production.
func (l *ParseTreeListener) ExitStatementList(c *StatementListContext) {
	l.exit()
}

// ExitImportStatement is called when exiting the importStatement production.
func (l *ParseTreeListener) ExitImportStatement(c *ImportStatementContext) {
	l.exit()
}

// ExitImportFromBlock is called when exiting the importFromBlock production.
func (l *ParseTreeListener) ExitImportFromBlock(c *ImportFromBlockContext) {
	l.exit()
}

// ExitImportModuleItems is called when exiting the importModuleItems production.
func (l *ParseTreeListener) ExitImportModuleItems(c *ImportModuleItemsContext) {
	l.exit()
}

// ExitImportAliasName is called when exiting the importAliasName production.
func (l *ParseTreeListener) ExitImportAliasName(c *ImportAliasNameContext) {
	l.exit()
}

// ExitModuleExportName is called when exiting the moduleExportName production.
func (l *ParseTreeListener) ExitModuleExportName(c *ModuleExportNameContext) {
	l.exit()
}

// ExitImportedBinding is called when exiting the importedBinding production.
func (l *ParseTreeListener) ExitImportedBinding(c *ImportedBindingContext) {
	l.exit()
}

// ExitImportDefault is called when exiting the importDefault production.
func (l *ParseTreeListener) ExitImportDefault(c *ImportDefaultContext) {
	l.exit()
}

// ExitImportNamespace is called when exiting the importNamespace production.
func (l *ParseTreeListener) ExitImportNamespace(c *ImportNamespaceContext) {
	l.exit()
}

// ExitImportFrom is called when exiting the importFrom production.
func (l *ParseTreeListener) ExitImportFrom(c *ImportFromContext) {
	l.exit()
}

// ExitAliasName is called when exiting the aliasName production.
func (l *ParseTreeListener) ExitAliasName(c *AliasNameContext) {
	l.exit()
}

// ExitExportDeclaration is called when exiting the ExportDeclaration production.
func (l *ParseTreeListener) ExitExportDeclaration(c *ExportDeclarationContext) {
	l.exit()
}

// ExitExportDefaultDeclaration is called when exiting the ExportDefaultDeclaration production.
func (l *ParseTreeListener) ExitExportDefaultDeclaration(c *ExportDefaultDeclarationContext) {
	l.exit()
}

// ExitExportFromBlock is called when exiting the exportFromBlock production.
func (l *ParseTreeListener) ExitExportFromBlock(c *ExportFromBlockContext) {
	l.exit()
}

// ExitExportModuleItems is called when exiting the exportModuleItems production.
func (l *ParseTreeListener) ExitExportModuleItems(c *ExportModuleItemsContext) {
	l.exit()
}

// ExitExportAliasName is called when exiting the exportAliasName production.
func (l *ParseTreeListener) ExitExportAliasName(c *ExportAliasNameContext) {
	l.exit()
}

// ExitDeclaration is called when exiting the declaration production.
func (l *ParseTreeListener) ExitDeclaration(c *DeclarationContext) {
	l.exit()
}

// ExitVariableStatement is called when exiting the variableStatement production.
func (l *ParseTreeListener) ExitVariableStatement(c *VariableStatementContext) {
	l.exit()
}

// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
func (l *ParseTreeListener) ExitVariableDeclarationList(c *VariableDeclarationListContext) {
	l.exit()
}

// ExitVariableDeclaration is called when exiting the variableDeclaration production.
func (l *ParseTreeListener) ExitVariableDeclaration(c *VariableDeclarationContext) {
	l.exit()
}

// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
func (l *ParseTreeListener) ExitEmptyStatement_(c *EmptyStatement_Context) {
	l.exit()
}

// ExitExpressionStatement is called when exiting the expressionStatement production.
func (l *ParseTreeListener) ExitExpressionStatement(c *ExpressionStatementContext) {
	l.exit()
}

// ExitIfStatement is called when exiting the ifStatement production.
func (l *ParseTreeListener) ExitIfStatement(c *IfStatementContext) {
	l.exit()
}

// ExitDoStatement is called when exiting the DoStatement production.
func (l *ParseTreeListener) ExitDoStatement(c *DoStatementContext) {
	l.exit()
}

// ExitWhileStatement is called when exiting the WhileStatement production.
func (l *ParseTreeListener) ExitWhileStatement(c *WhileStatementContext) {
	l.exit()
}

// ExitForStatement is called when exiting the ForStatement production.
func (l *ParseTreeListener) ExitForStatement(c *ForStatementContext) {
	l.exit()
}

// ExitForInStatement is called when exiting the ForInStatement production.
func (l *ParseTreeListener) ExitForInStatement(c *ForInStatementContext) {
	l.exit()
}

// ExitForOfStatement is called when exiting the ForOfStatement production.
func (l *ParseTreeListener) ExitForOfStatement(c *ForOfStatementContext) {
	l.exit()
}

// ExitVarModifier is called when exiting the varModifier production.
func (l *ParseTreeListener) ExitVarModifier(c *VarModifierContext) {
	l.exit()
}

// ExitContinueStatement is called when exiting the continueStatement production.
func (l *ParseTreeListener) ExitContinueStatement(c *ContinueStatementContext) {
	l.exit()
}

// ExitBreakStatement is called when exiting the breakStatement production.
func (l *ParseTreeListener) ExitBreakStatement(c *BreakStatementContext) {
	l.exit()
}

// ExitReturnStatement is called when exiting the returnStatement production.
func (l *ParseTreeListener) ExitReturnStatement(c *ReturnStatementContext) {
	l.exit()
}

// ExitYieldStatement is called when exiting the yieldStatement production.
func (l *ParseTreeListener) ExitYieldStatement(c *YieldStatementContext) {
	l.exit()
}

// ExitWithStatement is called when exiting the withStatement production.
func (l *ParseTreeListener) ExitWithStatement(c *WithStatementContext) {
	l.exit()
}

// ExitSwitchStatement is called when exiting the switchStatement production.
func (l *ParseTreeListener) ExitSwitchStatement(c *SwitchStatementContext) {
	l.exit()
}

// ExitCaseBlock is called when exiting the caseBlock production.
func (l *ParseTreeListener) ExitCaseBlock(c *CaseBlockContext) {
	l.exit()
}

// ExitCaseClauses is called when exiting the caseClauses production.
func (l *ParseTreeListener) ExitCaseClauses(c *CaseClausesContext) {
	l.exit()
}

// ExitCaseClause is called when exiting the caseClause production.
func (l *ParseTreeListener) ExitCaseClause(c *CaseClauseContext) {
	l.exit()
}

// ExitDefaultClause is called when exiting the defaultClause production.
func (l *ParseTreeListener) ExitDefaultClause(c *DefaultClauseContext) {
	l.exit()
}

// ExitLabelledStatement is called when exiting the labelledStatement production.
func (l *ParseTreeListener) ExitLabelledStatement(c *LabelledStatementContext) {
	l.exit()
}

// ExitThrowStatement is called when exiting the throwStatement production.
func (l *ParseTreeListener) ExitThrowStatement(c *ThrowStatementContext) {
	l.exit()
}

// ExitTryStatement is called when exiting the tryStatement production.
func (l *ParseTreeListener) ExitTryStatement(c *TryStatementContext) {
	l.exit()
}

// ExitCatchProduction is called when exiting the catchProduction production.
func (l *ParseTreeListener) ExitCatchProduction(c *CatchProductionContext) {
	l.exit()
}

// ExitFinallyProduction is called when exiting the finallyProduction production.
func (l *ParseTreeListener) ExitFinallyProduction(c *FinallyProductionContext) {
	l.exit()
}

// ExitDebuggerStatement is called when exiting the debuggerStatement production.
func (l *ParseTreeListener) ExitDebuggerStatement(c *DebuggerStatementContext) {
	l.exit()
}

// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
func (l *ParseTreeListener) ExitFunctionDeclaration(c *FunctionDeclarationContext) {
	l.exit()
}

// ExitClassDeclaration is called when exiting the classDeclaration production.
func (l *ParseTreeListener) ExitClassDeclaration(c *ClassDeclarationContext) {
	l.exit()
}

// ExitClassTail is called when exiting the classTail production.
func (l *ParseTreeListener) ExitClassTail(c *ClassTailContext) {
	l.exit()
}

// ExitClassElement is called when exiting the classElement production.
func (l *ParseTreeListener) ExitClassElement(c *ClassElementContext) {
	l.exit()
}

// ExitMethodDefinition is called when exiting the methodDefinition production.
func (l *ParseTreeListener) ExitMethodDefinition(c *MethodDefinitionContext) {
	l.exit()
}

// ExitFieldDefinition is called when exiting the fieldDefinition production.
func (l *ParseTreeListener) ExitFieldDefinition(c *FieldDefinitionContext) {
	l.exit()
}

// ExitClassElementName is called when exiting the classElementName production.
func (l *ParseTreeListener) ExitClassElementName(c *ClassElementNameContext) {
	l.exit()
}

// ExitPrivateIdentifier is called when exiting the privateIdentifier production.
func (l *ParseTreeListener) ExitPrivateIdentifier(c *PrivateIdentifierContext) {
	l.exit()
}

// ExitFormalParameterList is called when exiting the formalParameterList production.
func (l *ParseTreeListener) ExitFormalParameterList(c *FormalParameterListContext) {
	l.exit()
}

// ExitFormalParameterArg is called when exiting the formalParameterArg production.
func (l *ParseTreeListener) ExitFormalParameterArg(c *FormalParameterArgContext) {
	l.exit()
}

// ExitLastFormalParameterArg is called when exiting the lastFormalParameterArg production.
func (l *ParseTreeListener) ExitLastFormalParameterArg(c *LastFormalParameterArgContext) {
	l.exit()
}

// ExitFunctionBody is called when exiting the functionBody production.
func (l *ParseTreeListener) ExitFunctionBody(c *FunctionBodyContext) {
	l.exit()
}

// ExitSourceElements is called when exiting the sourceElements production.
func (l *ParseTreeListener) ExitSourceElements(c *SourceElementsContext) {
	l.exit()
}

// ExitArrayLiteral is called when exiting the arrayLiteral production.
func (l *ParseTreeListener) ExitArrayLiteral(c *ArrayLiteralContext) {
	l.exit()
}

// ExitElementList is called when exiting the elementList production.
func (l *ParseTreeListener) ExitElementList(c *ElementListContext) {
	l.exit()
}

// ExitArrayElement is called when exiting the arrayElement production.
func (l *ParseTreeListener) ExitArrayElement(c *ArrayElementContext) {
	l.exit()
}

// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
func (l *ParseTreeListener) ExitPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext) {
	l.exit()
}

// ExitComputedPropertyExpressionAssignment is called when exiting the ComputedPropertyExpressionAssignment production.
func (l *ParseTreeListener) ExitComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext) {
	l.exit()
}

// ExitFunctionProperty is called when exiting the FunctionProperty production.
func (l *ParseTreeListener) ExitFunctionProperty(c *FunctionPropertyContext) {
	l.exit()
}

// ExitPropertyGetter is called when exiting the PropertyGetter production.
func (l *ParseTreeListener) ExitPropertyGetter(c *PropertyGetterContext) {
	l.exit()
}

// ExitPropertySetter is called when exiting the PropertySetter production.
func (l *ParseTreeListener) ExitPropertySetter(c *PropertySetterContext) {
	l.exit()
}

// ExitPropertyShorthand is called when exiting the PropertyShorthand production.
func (l *ParseTreeListener) ExitPropertyShorthand(c *PropertyShorthandContext) {
	l.exit()
}

// ExitPropertyName is called when exiting the propertyName production.
func (l *ParseTreeListener) ExitPropertyName(c *PropertyNameContext) {
	l.exit()
}

// ExitArguments is called when exiting the arguments production.
func (l *ParseTreeListener) ExitArguments(c *ArgumentsContext) {
	l.exit()
}

// ExitArgument is called when exiting the argument production.
func (l *ParseTreeListener) ExitArgument(c *ArgumentContext) {
	l.exit()
}

// ExitExpressionSequence is called when exiting the expressionSequence production.
func (l *ParseTreeListener) ExitExpressionSequence(c *ExpressionSequenceContext) {
	l.exit()
}

// ExitTemplateStringExpression is called when exiting the TemplateStringExpression production.
func (l *ParseTreeListener) ExitTemplateStringExpression(c *TemplateStringExpressionContext) {
	l.exit()
}

// ExitTernaryExpression is called when exiting the TernaryExpression production.
func (l *ParseTreeListener) ExitTernaryExpression(c *TernaryExpressionContext) {
	l.exit()
}

// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
func (l *ParseTreeListener) ExitLogicalAndExpression(c *LogicalAndExpressionContext) {
	l.exit()
}

// ExitPowerExpression is called when exiting the PowerExpression production.
func (l *ParseTreeListener) ExitPowerExpression(c *PowerExpressionContext) {
	l.exit()
}

// ExitPreIncrementExpression is called when exiting the PreIncrementExpression production.
func (l *ParseTreeListener) ExitPreIncrementExpression(c *PreIncrementExpressionContext) {
	l.exit()
}

// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
func (l *ParseTreeListener) ExitObjectLiteralExpression(c *ObjectLiteralExpressionContext) {
	l.exit()
}

// ExitMetaExpression is called when exiting the MetaExpression production.
func (l *ParseTreeListener) ExitMetaExpression(c *MetaExpressionContext) {
	l.exit()
}

// ExitInExpression is called when exiting the InExpression production.
func (l *ParseTreeListener) ExitInExpression(c *InExpressionContext) {
	l.exit()
}

// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
func (l *ParseTreeListener) ExitLogicalOrExpression(c *LogicalOrExpressionContext) {
	l.exit()
}

// ExitOptionalChainExpression is called when exiting the OptionalChainExpression production.
func (l *ParseTreeListener) ExitOptionalChainExpression(c *OptionalChainExpressionContext) {
	l.exit()
}

// ExitNotExpression is called when exiting the NotExpression production.
func (l *ParseTreeListener) ExitNotExpression(c *NotExpressionContext) {
	l.exit()
}

// ExitPreDecreaseExpression is called when exiting the PreDecreaseExpression production.
func (l *ParseTreeListener) ExitPreDecreaseExpression(c *PreDecreaseExpressionContext) {
	l.exit()
}

// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
func (l *ParseTreeListener) ExitArgumentsExpression(c *ArgumentsExpressionContext) {
	l.exit()
}

// ExitAwaitExpression is called when exiting the AwaitExpression production.
func (l *ParseTreeListener) ExitAwaitExpression(c *AwaitExpressionContext) {
	l.exit()
}

// ExitThisExpression is called when exiting the ThisExpression production.
func (l *ParseTreeListener) ExitThisExpression(c *ThisExpressionContext) {
	l.exit()
}

// ExitFunctionExpression is called when exiting the FunctionExpression production.
func (l *ParseTreeListener) ExitFunctionExpression(c *FunctionExpressionContext) {
	l.exit()
}

// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
func (l *ParseTreeListener) ExitUnaryMinusExpression(c *UnaryMinusExpressionContext) {
	l.exit()
}

// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
func (l *ParseTreeListener) ExitAssignmentExpression(c *AssignmentExpressionContext) {
	l.exit()
}

// ExitPostDecreaseExpression is called when exiting the PostDecreaseExpression production.
func (l *ParseTreeListener) ExitPostDecreaseExpression(c *PostDecreaseExpressionContext) {
	l.exit()
}

// ExitTypeofExpression is called when exiting the TypeofExpression production.
func (l *ParseTreeListener) ExitTypeofExpression(c *TypeofExpressionContext) {
	l.exit()
}

// ExitInstanceofExpression is called when exiting the InstanceofExpression production.
func (l *ParseTreeListener) ExitInstanceofExpression(c *InstanceofExpressionContext) {
	l.exit()
}

// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
func (l *ParseTreeListener) ExitUnaryPlusExpression(c *UnaryPlusExpressionContext) {
	l.exit()
}

// ExitDeleteExpression is called when exiting the DeleteExpression production.
func (l *ParseTreeListener) ExitDeleteExpression(c *DeleteExpressionContext) {
	l.exit()
}

// ExitImportExpression is called when exiting the ImportExpression production.
func (l *ParseTreeListener) ExitImportExpression(c *ImportExpressionContext) {
	l.exit()
}

// ExitEqualityExpression is called when exiting the EqualityExpression production.
func (l *ParseTreeListener) ExitEqualityExpression(c *EqualityExpressionContext) {
	l.exit()
}

// ExitBitXOrExpression is called when exiting the BitXOrExpression production.
func (l *ParseTreeListener) ExitBitXOrExpression(c *BitXOrExpressionContext) {
	l.exit()
}

// ExitSuperExpression is called when exiting the SuperExpression production.
func (l *ParseTreeListener) ExitSuperExpression(c *SuperExpressionContext) {
	l.exit()
}

// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
func (l *ParseTreeListener) ExitMultiplicativeExpression(c *MultiplicativeExpressionContext) {
	l.exit()
}

// ExitBitShiftExpression is called when exiting the BitShiftExpression production.
func (l *ParseTreeListener) ExitBitShiftExpression(c *BitShiftExpressionContext) {
	l.exit()
}

// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
func (l *ParseTreeListener) ExitParenthesizedExpression(c *ParenthesizedExpressionContext) {
	l.exit()
}

// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
func (l *ParseTreeListener) ExitAdditiveExpression(c *AdditiveExpressionContext) {
	l.exit()
}

// ExitRelationalExpression is called when exiting the RelationalExpression production.
func (l *ParseTreeListener) ExitRelationalExpression(c *RelationalExpressionContext) {
	l.exit()
}

// ExitPostIncrementExpression is called when exiting the PostIncrementExpression production.
func (l *ParseTreeListener) ExitPostIncrementExpression(c *PostIncrementExpressionContext) {
	l.exit()
}

// ExitYieldExpression is called when exiting the YieldExpression production.
func (l *ParseTreeListener) ExitYieldExpression(c *YieldExpressionContext) {
	l.exit()
}

// ExitBitNotExpression is called when exiting the BitNotExpression production.
func (l *ParseTreeListener) ExitBitNotExpression(c *BitNotExpressionContext) {
	l.exit()
}

// ExitNewExpression is called when exiting the NewExpression production.
func (l *ParseTreeListener) ExitNewExpression(c *NewExpressionContext) {
	l.exit()
}

// ExitLiteralExpression is called when exiting the LiteralExpression production.
func (l *ParseTreeListener) ExitLiteralExpression(c *LiteralExpressionContext) {
	l.exit()
}

// ExitArrayLiteralExpression is called when exiting the ArrayLiteralExpression production.
func (l *ParseTreeListener) ExitArrayLiteralExpression(c *ArrayLiteralExpressionContext) {
	l.exit()
}

// ExitMemberDotExpression is called when exiting the MemberDotExpression production.
func (l *ParseTreeListener) ExitMemberDotExpression(c *MemberDotExpressionContext) {
	l.exit()
}

// ExitClassExpression is called when exiting the ClassExpression production.
func (l *ParseTreeListener) ExitClassExpression(c *ClassExpressionContext) {
	l.exit()
}

// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
func (l *ParseTreeListener) ExitMemberIndexExpression(c *MemberIndexExpressionContext) {
	l.exit()
}

// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
func (l *ParseTreeListener) ExitIdentifierExpression(c *IdentifierExpressionContext) {
	l.exit()
}

// ExitBitAndExpression is called when exiting the BitAndExpression production.
func (l *ParseTreeListener) ExitBitAndExpression(c *BitAndExpressionContext) {
	l.exit()
}

// ExitBitOrExpression is called when exiting the BitOrExpression production.
func (l *ParseTreeListener) ExitBitOrExpression(c *BitOrExpressionContext) {
	l.exit()
}

// ExitAssignmentOperatorExpression is called when exiting the AssignmentOperatorExpression production.
func (l *ParseTreeListener) ExitAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext) {
	l.exit()
}

// ExitVoidExpression is called when exiting the VoidExpression production.
func (l *ParseTreeListener) ExitVoidExpression(c *VoidExpressionContext) {
	l.exit()
}

// ExitCoalesceExpression is called when exiting the CoalesceExpression production.
func (l *ParseTreeListener) ExitCoalesceExpression(c *CoalesceExpressionContext) {
	l.exit()
}

// ExitInitializer is called when exiting the initializer production.
func (l *ParseTreeListener) ExitInitializer(c *InitializerContext) {
	l.exit()
}

// ExitAssignable is called when exiting the assignable production.
func (l *ParseTreeListener) ExitAssignable(c *AssignableContext) {
	l.exit()
}

// ExitObjectLiteral is called when exiting the objectLiteral production.
func (l *ParseTreeListener) ExitObjectLiteral(c *ObjectLiteralContext) {
	l.exit()
}

// ExitNamedFunction is called when exiting the NamedFunction production.
func (l *ParseTreeListener) ExitNamedFunction(c *NamedFunctionContext) {
	l.exit()
}

// ExitAnonymousFunctionDecl is called when exiting the AnonymousFunctionDecl production.
func (l *ParseTreeListener) ExitAnonymousFunctionDecl(c *AnonymousFunctionDeclContext) {
	l.exit()
}

// ExitArrowFunction is called when exiting the ArrowFunction production.
func (l *ParseTreeListener) ExitArrowFunction(c *ArrowFunctionContext) {
	l.exit()
}

// ExitArrowFunctionParameters is called when exiting the arrowFunctionParameters production.
func (l *ParseTreeListener) ExitArrowFunctionParameters(c *ArrowFunctionParametersContext) {
	l.exit()
}

// ExitArrowFunctionBody is called when exiting the arrowFunctionBody production.
func (l *ParseTreeListener) ExitArrowFunctionBody(c *ArrowFunctionBodyContext) {
	l.exit()
}

// ExitAssignmentOperator is called when exiting the assignmentOperator production.
func (l *ParseTreeListener) ExitAssignmentOperator(c *AssignmentOperatorContext) {
	l.exit()
}

// ExitLiteral is called when exiting the literal production.
func (l *ParseTreeListener) ExitLiteral(c *LiteralContext) {
	l.exit()
}

// ExitTemplateStringLiteral is called when exiting the templateStringLiteral production.
func (l *ParseTreeListener) ExitTemplateStringLiteral(c *TemplateStringLiteralContext) {
	l.exit()
}

// ExitTemplateStringAtom is called when exiting the templateStringAtom production.
func (l *ParseTreeListener) ExitTemplateStringAtom(c *TemplateStringAtomContext) {
	l.exit()
}

// ExitNumericLiteral is called when exiting the numericLiteral production.
func (l *ParseTreeListener) ExitNumericLiteral(c *NumericLiteralContext) {
	l.exit()
}

// ExitBigintLiteral is called when exiting the bigintLiteral production.
func (l *ParseTreeListener) ExitBigintLiteral(c *BigintLiteralContext) {
	l.exit()
}

// ExitGetter is called when exiting the getter production.
func (l *ParseTreeListener) ExitGetter(c *GetterContext) {
	l.exit()
}

// ExitSetter is called when exiting the setter production.
func (l *ParseTreeListener) ExitSetter(c *SetterContext) {
	l.exit()
}

// ExitIdentifierName is called when exiting the identifierName production.
func (l *ParseTreeListener) ExitIdentifierName(c *IdentifierNameContext) {
	l.exit()
}

// ExitIdentifier is called when exiting the identifier production.
func (l *ParseTreeListener) ExitIdentifier(c *IdentifierContext) {
	l.exit()
}

// ExitReservedWord is called when exiting the reservedWord production.
func (l *ParseTreeListener) ExitReservedWord(c *ReservedWordContext) {
	l.exit()
}

// ExitKeyword is called when exiting the keyword production.
func (l *ParseTreeListener) ExitKeyword(c *KeywordContext) {
	l.exit()
}

// ExitLet_ is called when exiting the let_ production.
func (l *ParseTreeListener) ExitLet_(c *Let_Context) {
	l.exit()
}

// ExitEos is called when exiting the eos production.
func (l *ParseTreeListener) ExitEos(c *EosContext) {
	l.exit()
}
