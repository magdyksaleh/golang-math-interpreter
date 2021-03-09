package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculator")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		lexer := NewLexer(text)
		Tokens := ParseText(lexer)
		parser := NewParser(Tokens)
		RootNode := Parse(parser)
		val := InterpretTree(RootNode)
		fmt.Println(val)
	}
}
