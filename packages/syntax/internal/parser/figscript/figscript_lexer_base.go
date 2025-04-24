package figscript

import "github.com/antlr4-go/antlr/v4"

// FigScriptLexerBase state
type FigScriptLexerBase struct {
	*antlr.BaseLexer

	scopeStrictModes []bool
	stackLength      int
	stackIx          int

	lastToken          antlr.Token
	useStrictDefault   bool
	useStrictCurrent   bool
	currentDepth       int
	templateDepthStack []int
}

func (l *FigScriptLexerBase) IsStartOfFile() bool {
	return l.lastToken == nil
}

func (l *FigScriptLexerBase) pushStrictModeScope(v bool) {
	if l.stackIx == l.stackLength {
		l.scopeStrictModes = append(l.scopeStrictModes, v)
		l.stackLength++
	} else {
		l.scopeStrictModes[l.stackIx] = v
	}
	l.stackIx++
}

func (l *FigScriptLexerBase) popStrictModeScope() bool {
	l.stackIx--
	v := l.scopeStrictModes[l.stackIx]
	l.scopeStrictModes[l.stackIx] = false
	return v
}

// IsStrictMode is self explanatory.
func (l *FigScriptLexerBase) IsStrictMode() bool {
	return l.useStrictCurrent
}

// NextToken from the character stream.
func (l *FigScriptLexerBase) NextToken() antlr.Token {
	next := l.BaseLexer.NextToken() // Get next token
	if next.GetChannel() == antlr.TokenDefaultChannel {
		// Keep track of the last token on default channel
		l.lastToken = next
	}
	return next
}

// ProcessOpenBrace is called when a { is encountered during
// lexing, we push a new scope everytime.
func (l *FigScriptLexerBase) ProcessOpenBrace() {
	l.currentDepth++
	l.useStrictCurrent = l.useStrictDefault
	if l.stackIx > 0 && l.scopeStrictModes[l.stackIx-1] {
		l.useStrictCurrent = true
	}
	l.pushStrictModeScope(l.useStrictCurrent)
}

// ProcessCloseBrace is called when a } is encountered during
// lexing, we pop a scope unless we're inside global scope.
func (l *FigScriptLexerBase) ProcessCloseBrace() {
	l.useStrictCurrent = l.useStrictDefault
	if l.stackIx > 0 {
		l.useStrictCurrent = l.popStrictModeScope()
	}
	l.currentDepth--
}

// ProcessStringLiteral is called when lexing a string literal.
func (l *FigScriptLexerBase) ProcessStringLiteral() {
	if l.lastToken == nil || l.lastToken.GetTokenType() == FigScriptLexerOpenBrace {
		if l.GetText() == `"use strict"` || l.GetText() == "'use strict'" {
			if l.stackIx > 0 {
				l.popStrictModeScope()
			}
			l.useStrictCurrent = true
			l.pushStrictModeScope(l.useStrictCurrent)
		}
	}
}

func (l *FigScriptLexerBase) ProcessTemplateOpenBrace() {
	l.currentDepth++
	l.templateDepthStack = append(l.templateDepthStack, l.currentDepth)
}

func (l *FigScriptLexerBase) ProcessTemplateCloseBrace() {
	l.templateDepthStack = l.templateDepthStack[:len(l.templateDepthStack)-1]
	l.currentDepth--
}

// IsRegexPossible returns true if the lexer can match a
// regex literal.
func (l *FigScriptLexerBase) IsRegexPossible() bool {
	if l.lastToken == nil {
		return true
	}
	switch l.lastToken.GetTokenType() {
	case FigScriptLexerIdentifier, FigScriptLexerNullLiteral,
		FigScriptLexerBooleanLiteral, FigScriptLexerThis,
		FigScriptLexerCloseBracket, FigScriptLexerCloseParen,
		FigScriptLexerOctalIntegerLiteral, FigScriptLexerDecimalLiteral,
		FigScriptLexerHexIntegerLiteral, FigScriptLexerStringLiteral,
		FigScriptLexerPlusPlus, FigScriptLexerMinusMinus:
		return false
	default:
		return true
	}
}

func (l *FigScriptLexerBase) IsInTemplateString() bool {
	return len(l.templateDepthStack) > 0 && l.templateDepthStack[len(l.templateDepthStack)-1] == l.currentDepth
}

func (l *FigScriptLexerBase) Reset() {
	l.scopeStrictModes = nil
	l.stackLength = 0
	l.stackIx = 0
	l.lastToken = nil
	l.useStrictDefault = false
	l.useStrictCurrent = false
	l.currentDepth = 0
	l.templateDepthStack = make([]int, 0)
	l.BaseLexer.Reset()
}
