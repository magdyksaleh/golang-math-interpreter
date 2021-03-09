package main

type NodeType int

const (
   unknown NodeType = iota
   NUMBER
   ADD
   SUB
   MUL
   DIV
   PLUS
   MINUS
   sentinel
)

type Node struct {
   NodeType NodeType
   Nodes []*Node
   Value float32
}

