// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging
parser grammar FigScriptParser;

options {
    tokenVocab = FigScriptLexer;
    superClass = FigScriptParserBase;
}

program
    : HashBangLine? sourceElements? EOF
    ;

sourceElement
    : statement
    ;

statement
    : block
    | variableStatement
    | importStatement
    | exportStatement
    | emptyStatement_
    | classDeclaration
    | functionDeclaration
    | expressionStatement
    | ifStatement
    | iterationStatement
    | continueStatement
    | breakStatement
    | returnStatement
    | yieldStatement
    | withStatement
    | labelledStatement
    | switchStatement
    | throwStatement
    // | tryStatement
    | debuggerStatement
    ;

block
    : '{' statementList? '}'
    ;

statementList
    : statement+
    ;

importStatement
    : Import importFromBlock
    ;

importFromBlock
    : importDefault? (importNamespace | importModuleItems) importFrom eos
    | StringLiteral eos
    ;

importModuleItems
    : '{' (importAliasName ',')* (importAliasName ','?)? '}'
    ;

importAliasName
    : moduleExportName (As importedBinding)?
    ;

moduleExportName
    : identifierName
    | StringLiteral
    ;

// yield and await are permitted as BindingIdentifier in the grammar
importedBinding
    : Identifier
    | Yield
    | Await
    ;

importDefault
    : aliasName ','
    ;

importNamespace
    : ('*' | identifierName) (As identifierName)?
    ;

importFrom
    : From StringLiteral
    ;

aliasName
    : identifierName (As identifierName)?
    ;

exportStatement
    : Export Default? (exportFromBlock | declaration) eos # ExportDeclaration
    | Export Default singleExpression eos                 # ExportDefaultDeclaration
    ;

exportFromBlock
    : importNamespace importFrom eos
    | exportModuleItems importFrom? eos
    ;

exportModuleItems
    : '{' (exportAliasName ',')* (exportAliasName ','?)? '}'
    ;

exportAliasName
    : moduleExportName (As moduleExportName)?
    ;

declaration
    : variableStatement
    | classDeclaration
    | functionDeclaration
    ;

variableStatement
    : variableDeclarationList eos
    ;

variableDeclarationList
    : varModifier variableDeclaration (',' variableDeclaration)*
    ;

variableDeclaration
    : assignable ('=' singleExpression)? // ECMAScript 6: Array & Object Matching
    ;

emptyStatement_
    : SemiColon
    ;

expressionStatement
    : {p.notOpenBraceAndNotFunction()}? expressionSequence eos
    ;

ifStatement
    : If '(' expressionSequence ')' statement (Else statement)?
    ;

iterationStatement
    : Do statement While '(' expressionSequence ')' eos                                                                     # DoStatement
    | While '(' expressionSequence ')' statement                                                                            # WhileStatement
    | For '(' (expressionSequence | variableDeclarationList)? ';' expressionSequence? ';' expressionSequence? ')' statement # ForStatement
    | For '(' (singleExpression | variableDeclarationList) In expressionSequence ')' statement                              # ForInStatement
    | For Await? '(' (singleExpression | variableDeclarationList) Of expressionSequence ')' statement                       # ForOfStatement
    ;

varModifier // let, const - ECMAScript 6
    : Var
    | let_
    | Const
    ;

continueStatement
    : Continue ({p.notLineTerminator()}? identifier)? eos
    ;

breakStatement
    : Break ({p.notLineTerminator()}? identifier)? eos
    ;

returnStatement
    : Return ({p.notLineTerminator()}? expressionSequence)? eos
    ;

yieldStatement
    : (Yield | YieldStar) ({p.notLineTerminator()}? expressionSequence)? eos
    ;

withStatement
    : With '(' expressionSequence ')' statement
    ;

switchStatement
    : Switch '(' expressionSequence ')' caseBlock
    ;

caseBlock
    : '{' caseClauses? (defaultClause caseClauses?)? '}'
    ;

caseClauses
    : caseClause+
    ;

caseClause
    : Case expressionSequence ':' statementList?
    ;

defaultClause
    : Default ':' statementList?
    ;

labelledStatement
    : identifier ':' statement
    ;

throwStatement
    : Throw {p.notLineTerminator()}? expressionSequence eos
    ;

tryStatement
    : Try block (catchProduction finallyProduction? | finallyProduction)
    ;

catchProduction
    : Catch ('(' assignable? ')')? block
    ;

finallyProduction
    : Finally block
    ;

debuggerStatement
    : Debugger eos
    ;

functionDeclaration
    : Async? Function_ '*'? identifier '(' formalParameterList? ')' functionBody
    ;

classDeclaration
    : Class identifier classTail
    ;

classTail
    : (Extends singleExpression)? '{' classElement* '}'
    ;

classElement
    : (Static | {p.n("static")}? identifier)? methodDefinition
    | (Static | {p.n("static")}? identifier)? fieldDefinition
    | (Static | {p.n("static")}? identifier) block
    | emptyStatement_
    ;

methodDefinition
    : (Async {p.notLineTerminator()}?)? '*'? classElementName '(' formalParameterList? ')' functionBody
    | '*'? getter '(' ')' functionBody
    | '*'? setter '(' formalParameterList? ')' functionBody
    ;

fieldDefinition
    : classElementName initializer?
    ;

classElementName
    : propertyName
    | privateIdentifier
    ;

privateIdentifier
    : '#' identifierName
    ;

formalParameterList
    : formalParameterArg (',' formalParameterArg)* (',' lastFormalParameterArg)?
    | lastFormalParameterArg
    ;

formalParameterArg
    : assignable ('=' singleExpression)? // ECMAScript 6: Initialization
    ;

lastFormalParameterArg // ECMAScript 6: Rest Parameter
    : Ellipsis singleExpression
    ;

functionBody
    : '{' sourceElements? '}'
    ;

sourceElements
    : sourceElement+
    ;

arrayLiteral
    : ('[' elementList ']')
    ;

elementList
    : ','* arrayElement? (','+ arrayElement)* ','*
    ;

arrayElement
    : Ellipsis? singleExpression
    ;

propertyAssignment
    : propertyName ':' singleExpression                                  # PropertyExpressionAssignment
    | '[' singleExpression ']' ':' singleExpression                      # ComputedPropertyExpressionAssignment
    | Async? '*'? propertyName '(' formalParameterList? ')' functionBody # FunctionProperty
    | getter '(' ')' functionBody                                        # PropertyGetter
    | setter '(' formalParameterArg ')' functionBody                     # PropertySetter
    | Ellipsis? singleExpression                                         # PropertyShorthand
    ;

propertyName
    : identifierName
    | StringLiteral
    | numericLiteral
    | '[' singleExpression ']'
    ;

arguments
    : '(' (argument (',' argument)* ','?)? ')'
    ;

argument
    : Ellipsis? (singleExpression | identifier)
    ;

expressionSequence
    : singleExpression (',' singleExpression)*
    ;

singleExpression
    : anonymousFunction                                                    # FunctionExpression
    | Class identifier? classTail                                          # ClassExpression
    | singleExpression '.' Swizzle                                         # SwizzleExpression
    | singleExpression '?.' singleExpression                               # OptionalChainExpression
    | singleExpression '?.'? '[' expressionSequence ']'                    # MemberIndexExpression
    | singleExpression '?'? '.' '#'? identifierName                        # MemberDotExpression
    | New identifier arguments                                             # NewExpression
    | New singleExpression arguments                                       # NewExpression
    | New singleExpression                                                 # NewExpression
    | singleExpression arguments? trailingLambda                           # TrailingLambdaExpression
    | singleExpression arguments                                           # ArgumentsExpression
    | New '.' identifier                                                   # MetaExpression
    | singleExpression {p.notLineTerminator()}? '++'                       # PostIncrementExpression
    | singleExpression {p.notLineTerminator()}? '--'                       # PostDecreaseExpression
    | Delete singleExpression                                              # DeleteExpression
    | Void singleExpression                                                # VoidExpression
    | Typeof singleExpression                                              # TypeofExpression
    | '++' singleExpression                                                # PreIncrementExpression
    | '--' singleExpression                                                # PreDecreaseExpression
    | '+' singleExpression                                                 # UnaryPlusExpression
    | '-' singleExpression                                                 # UnaryMinusExpression
    | '~' singleExpression                                                 # BitNotExpression
    | '!' singleExpression                                                 # NotExpression
    | Await singleExpression                                               # AwaitExpression
    | <assoc = right> singleExpression '**' singleExpression               # PowerExpression
    | singleExpression ('*' | '/' | '%') singleExpression                  # MultiplicativeExpression
    | singleExpression ('+' | '-') singleExpression                        # AdditiveExpression
    | singleExpression '??' singleExpression                               # CoalesceExpression
    | singleExpression ('<<' | '>>' | '>>>') singleExpression              # BitShiftExpression
    | singleExpression ('<' | '>' | '<=' | '>=') singleExpression          # RelationalExpression
    | singleExpression Instanceof singleExpression                         # InstanceofExpression
    | singleExpression In singleExpression                                 # InExpression
    | singleExpression ('==' | '!=' | '===' | '!==') singleExpression      # EqualityExpression
    | singleExpression '&' singleExpression                                # BitAndExpression
    | singleExpression '^' singleExpression                                # BitXOrExpression
    | singleExpression '|' singleExpression                                # BitOrExpression
    | singleExpression '&&' singleExpression                               # LogicalAndExpression
    | singleExpression '||' singleExpression                               # LogicalOrExpression
    | singleExpression '?' singleExpression ':' singleExpression           # TernaryExpression
    | <assoc = right> singleExpression '=' singleExpression                # AssignmentExpression
    | <assoc = right> singleExpression assignmentOperator singleExpression # AssignmentOperatorExpression
    | Import '(' singleExpression ')'                                      # ImportExpression
    | singleExpression templateStringLiteral                               # TemplateStringExpression // ECMAScript 6
    | yieldStatement                                                       # YieldExpression          // ECMAScript 6
    | This                                                                 # ThisExpression
    | identifier                                                           # IdentifierExpression
    | Super                                                                # SuperExpression
    | literal                                                              # LiteralExpression
    | arrayLiteral                                                         # ArrayLiteralExpression
    | objectLiteral                                                        # ObjectLiteralExpression
    | vectorLiteral                                                        # VectorLiteralExpression
    | '(' expressionSequence ')'                                           # ParenthesizedExpression
    ;

initializer
    : '=' singleExpression
    ;

assignable
    : identifier
    | keyword
    | arrayLiteral
    | objectLiteral
    | vectorLiteral
    ;

vectorLiteral
    : '(' componentList ')'
    ;

componentList
    : component ',' component (',' component)?
    ;

component
    : (singleExpression | identifier)
    ;

trailingLambda
    : '{' lambdaStatement+ '}'
    ;

lambdaStatement
    : singleExpression arguments
    ;

objectLiteral
    : '{' (propertyAssignment (',' propertyAssignment)* ','?)? '}'
    ;

anonymousFunction
    : functionDeclaration                                             # NamedFunction
    | Async? Function_ '*'? '(' formalParameterList? ')' functionBody # AnonymousFunctionDecl
    | Async? arrowFunctionParameters '=>' arrowFunctionBody           # ArrowFunction
    ;

arrowFunctionParameters
    : propertyName
    | '(' formalParameterList? ')'
    ;

arrowFunctionBody
    : singleExpression
    | functionBody
    ;

assignmentOperator
    : '*='
    | '/='
    | '%='
    | '+='
    | '-='
    | '<<='
    | '>>='
    | '>>>='
    | '&='
    | '^='
    | '|='
    | '**='
    | '??='
    ;

literal
    : NullLiteral
    | BooleanLiteral
    | StringLiteral
    | templateStringLiteral
    | RegularExpressionLiteral
    | numericLiteral
    | bigintLiteral
    ;

templateStringLiteral
    : BackTick templateStringAtom* BackTick
    ;

templateStringAtom
    : TemplateStringAtom
    | TemplateStringStartExpression singleExpression TemplateCloseBrace
    ;

numericLiteral
    : DecimalLiteral
    | HexIntegerLiteral
    | OctalIntegerLiteral
    | OctalIntegerLiteral2
    | BinaryIntegerLiteral
    ;

bigintLiteral
    : BigDecimalIntegerLiteral
    | BigHexIntegerLiteral
    | BigOctalIntegerLiteral
    | BigBinaryIntegerLiteral
    ;

getter
    : {p.n("get")}? identifier classElementName
    ;

setter
    : {p.n("set")}? identifier classElementName
    ;

identifierName
    : identifier
    | reservedWord
    ;

identifier
    : Identifier
    | NonStrictLet
    | Async
    | As
    | From
    | Yield
    | Of
    ;

reservedWord
    : keyword
    | NullLiteral
    | BooleanLiteral
    ;

keyword
    : Break
    | Do
    | Instanceof
    | Typeof
    | Case
    | Else
    | New
    | Var
    | Catch
    | Finally
    | Return
    | Void
    | Continue
    | For
    | Switch
    | While
    | Debugger
    | Function_
    | This
    | With
    | Default
    | If
    | Throw
    | Delete
    | In
    | Try
    | Class
    | Enum
    | Extends
    | Super
    | Const
    | Export
    | Import
    | Implements
    | let_
    | Private
    | Public
    | Interface
    | Package
    | Protected
    | Static
    | Yield
    | YieldStar
    | Async
    | Await
    | From
    | As
    | Of
    ;

let_
    : NonStrictLet
    | StrictLet
    ;

eos
    : SemiColon
    | EOF
    | {p.lineTerminatorAhead()}?
    | {p.closeBrace()}?
    ;