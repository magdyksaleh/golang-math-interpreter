package main

import (
   "fmt"
	"strconv"
)

type NodeType int

const (
	unknown_NODE NodeType = iota
	NUMBER_NODE
	ADD_NODE
	SUB_NODE
	MUL_NODE
	DIV_NODE
	PLUS_NODE
	MINUS_NODE
	sentinel_NODE
)

type Node struct {
	NodeType NodeType
	Nodes    []*Node
	Value    float64
}

type Parser struct {
	CurrentToken *Token
	Tokens       []Token
	Pos          int
}

func NewParser(tokens []Token) *Parser {
	if len(tokens) > 0 {
		return &Parser{CurrentToken: &tokens[0], Pos: 0, Tokens: tokens}
	}
	return &Parser{Pos: 0, Tokens: tokens}
}

func AdvanceParser(parser *Parser) {
	parser.Pos++
	if len(parser.Tokens) == parser.Pos {
		parser.Pos = -1
		parser.CurrentToken = nil
		return
	}
	parser.CurrentToken = &parser.Tokens[parser.Pos]
}

func ParseFactor(parser *Parser) (FactorNode *Node) {
	switch {
	case parser.CurrentToken.TokenType == LPAREN_TOKEN:
		AdvanceParser(parser)
		FactorNode = ParseExpression(parser)
		if parser.CurrentToken.TokenType != RPAREN_TOKEN {
			panic("Syntax Error: Unmatched parens")
		}
		AdvanceParser(parser)
	case parser.CurrentToken.TokenType == NUMBER_TOKEN:
      fValue, _ := strconv.ParseFloat(parser.CurrentToken.Value, 32)
		FactorNode = &Node{NodeType: NUMBER_NODE, Value: fValue}
		AdvanceParser(parser)
	case parser.CurrentToken.TokenType == PLUS_TOKEN:
		AdvanceParser(parser)
		FactorNode = &Node{NodeType: PLUS_NODE, Nodes: []*Node{ParseFactor(parser)}}
	case parser.CurrentToken.TokenType == MINUS_TOKEN:
		AdvanceParser(parser)
		FactorNode = &Node{NodeType: MINUS_NODE, Nodes: []*Node{ParseFactor(parser)}}
	}
   return
}

func ParseTerm(parser *Parser) *Node {
	TermNode := ParseFactor(parser)
	for parser.CurrentToken != nil &&
		(parser.CurrentToken.TokenType == MULTIPLY_TOKEN ||
			parser.CurrentToken.TokenType == DIVIDE_TOKEN) {

		if parser.CurrentToken.TokenType == MULTIPLY_TOKEN {
			AdvanceParser(parser)
			TermNode = &Node{NodeType: MUL_NODE, Nodes: []*Node{TermNode, ParseFactor(parser)}}
		} else if parser.CurrentToken.TokenType == DIVIDE_TOKEN {
			AdvanceParser(parser)
			TermNode = &Node{NodeType: DIV_NODE, Nodes: []*Node{TermNode, ParseFactor(parser)}}
		}
	}
	return TermNode
}

func ParseExpression(parser *Parser) *Node {
	ExpressionNode := ParseTerm(parser)
	for parser.CurrentToken != nil &&
		(parser.CurrentToken.TokenType == PLUS_TOKEN ||
			parser.CurrentToken.TokenType == MINUS_TOKEN) {

		if parser.CurrentToken.TokenType == PLUS_TOKEN {
			AdvanceParser(parser)
			ExpressionNode = &Node{NodeType: ADD_NODE, Nodes: []*Node{ExpressionNode, ParseTerm(parser)}}
		} else if parser.CurrentToken.TokenType == MINUS_TOKEN {
			AdvanceParser(parser)
			ExpressionNode = &Node{NodeType: SUB_NODE, Nodes: []*Node{ExpressionNode, ParseTerm(parser)}}
		}
	}
	return ExpressionNode
}

func StringTree(node *Node) (exprStr string) {
   if len(node.Nodes) == 2 {
      lstring := StringTree(node.Nodes[0])
      var OpsString string
      switch node.NodeType {
      case ADD_NODE:
         OpsString = "+"
      case SUB_NODE:
         OpsString = "-"
      case MUL_NODE:
         OpsString = "*"
      case DIV_NODE:
         OpsString = "/"
      }
      rstring := StringTree(node.Nodes[1])
      exprStr = "(" + lstring + OpsString +  rstring + ")"
   } else if len(node.Nodes) == 1 {
      rstring := StringTree(node.Nodes[0])
      var OpsString string
      switch node.NodeType {
      case PLUS_NODE:
         OpsString = "+"
      case MINUS_NODE:
         OpsString = "-"
      }
      exprStr = "(" + OpsString + rstring + ")"
   } else {
      exprStr = fmt.Sprintf("%.2f", node.Value)
   }
   return
}

func Parse(parser *Parser) *Node {
	if parser.CurrentToken == nil {
		return nil
	}

	RootNode := ParseExpression(parser)

	if parser.CurrentToken != nil {
		panic("Syntax Error")
	}
   fmt.Println(StringTree(RootNode))
	return RootNode
}
