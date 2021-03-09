package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TokenType int

const (
	WHITESPACE string = " \n\t"
	DIGITS     string = "0123456789"
)

const (
	unknown TokenType = iota
	NUMBER
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
	sentinel
)

type Token struct {
	TokenType TokenType
	Value     string
}

type Lexer struct {
	Pos          int
	RawText      []rune
	Tokens       []Token
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
	return Token{NUMBER, NumberString}
}

func GenerateSymbolToken(lexer *Lexer) (token Token) {
   switch {
      case strings.ContainsRune("+", lexer.RawText[lexer.Pos]):
         token = Token{TokenType: PLUS}
      case strings.ContainsRune("-", lexer.RawText[lexer.Pos]):
         token = Token{TokenType: MINUS}
      case strings.ContainsRune("*", lexer.RawText[lexer.Pos]):
         token = Token{TokenType: MULTIPLY}
      case strings.ContainsRune("/", lexer.RawText[lexer.Pos]):
         token = Token{TokenType: DIVIDE}
      case strings.ContainsRune("(", lexer.RawText[lexer.Pos]):
         token = Token{TokenType: LPAREN}
      case strings.ContainsRune(")", lexer.RawText[lexer.Pos]):
         token = Token{TokenType: RPAREN}
		}
      return
}

func ParseText(lexer *Lexer) {
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
   fmt.Println(lexer)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculator")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		lexer := NewLexer(text)
		ParseText(lexer)
		break
	}

}
