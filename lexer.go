package main

import (
	"strings"
)

type TokenType int

const (
	WHITESPACE string = " \n\t"
	DIGITS     string = "0123456789"
)

const (
	unknown_TOKEN TokenType = iota
	NUMBER_TOKEN
	PLUS_TOKEN
	MINUS_TOKEN
	MULTIPLY_TOKEN
	DIVIDE_TOKEN
	LPAREN_TOKEN
	RPAREN_TOKEN
	sentinel_TOKEN
)

type Token struct {
	TokenType TokenType
	Value     string
}

type Lexer struct {
	Pos     int
	RawText []rune
	Tokens  []Token
}

func IsDigitChar(c rune) bool {
	return strings.ContainsRune(DIGITS, c) || strings.ContainsRune(".", c)
}

func NewLexer(text string) *Lexer {
	return &Lexer{Pos: 0, RawText: ([]rune)(text)}
}

func AdvanceLexer(lexer *Lexer) {
	lexer.Pos++
	if len(lexer.RawText) == lexer.Pos {
		lexer.Pos = -1
	}
}

func GenerateNumberToken(lexer *Lexer) Token {
	dotCount := 0
	NumberString := string(lexer.RawText[lexer.Pos])
	AdvanceLexer(lexer)

	// Look for end of the string
	for lexer.Pos >= 0 && IsDigitChar(lexer.RawText[lexer.Pos]) {
		if strings.ContainsRune(".", lexer.RawText[lexer.Pos]) {
			dotCount++
			if dotCount > 1 {
				break
			}
		}
		NumberString += string(lexer.RawText[lexer.Pos])
		AdvanceLexer(lexer)
	}
	if strings.HasPrefix(NumberString, ".") {
		NumberString = "0" + NumberString
	}
	if strings.HasSuffix(NumberString, ".") {
		NumberString += "0"
	}
	return Token{NUMBER_TOKEN, NumberString}
}

func GenerateSymbolToken(lexer *Lexer) (token Token) {
	switch {
	case strings.ContainsRune("+", lexer.RawText[lexer.Pos]):
		token = Token{TokenType: PLUS_TOKEN}
	case strings.ContainsRune("-", lexer.RawText[lexer.Pos]):
		token = Token{TokenType: MINUS_TOKEN}
	case strings.ContainsRune("*", lexer.RawText[lexer.Pos]):
		token = Token{TokenType: MULTIPLY_TOKEN}
	case strings.ContainsRune("/", lexer.RawText[lexer.Pos]):
		token = Token{TokenType: DIVIDE_TOKEN}
	case strings.ContainsRune("(", lexer.RawText[lexer.Pos]):
		token = Token{TokenType: LPAREN_TOKEN}
	case strings.ContainsRune(")", lexer.RawText[lexer.Pos]):
		token = Token{TokenType: RPAREN_TOKEN}
	}
	return
}

func ParseText(lexer *Lexer) []Token {
	for lexer.Pos >= 0 {
		if IsDigitChar(lexer.RawText[lexer.Pos]) {
			lexer.Tokens = append(lexer.Tokens, GenerateNumberToken(lexer))
		} else {
			if !strings.ContainsRune(WHITESPACE, lexer.RawText[lexer.Pos]) {
				lexer.Tokens = append(lexer.Tokens, GenerateSymbolToken(lexer))
			}
			AdvanceLexer(lexer)
		}
	}
	return lexer.Tokens
}
