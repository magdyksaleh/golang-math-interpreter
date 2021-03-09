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
	   ParseText(lexer)
		break
	}

}
