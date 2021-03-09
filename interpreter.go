package main

func InterpretTree(root *Node) float64 {
	switch root.NodeType {
	case NUMBER_NODE:
		return root.Value
	case ADD_NODE:
		if len(root.Nodes) != 2 {
			panic("Parse Error! Got add node with not 2 children")
		}
		return InterpretTree(root.Nodes[0]) + InterpretTree(root.Nodes[1])
	case SUB_NODE:
		if len(root.Nodes) != 2 {
			panic("Parse Error! Got sub node with not 2 children")
		}
		return InterpretTree(root.Nodes[0]) - InterpretTree(root.Nodes[1])
	case MUL_NODE:
		if len(root.Nodes) != 2 {
			panic("Parse Error! Got mul node with not 2 children")
		}
		return InterpretTree(root.Nodes[0]) * InterpretTree(root.Nodes[1])
	case DIV_NODE:
		if len(root.Nodes) != 2 {
			panic("Parse Error! Got div node with not 2 children")
		}
		return InterpretTree(root.Nodes[0]) / InterpretTree(root.Nodes[1])
	case PLUS_NODE:
		if len(root.Nodes) != 1 {
			panic("Parse Error! Got plus node with not 1 children")
		}
		return InterpretTree(root.Nodes[0])
	case MINUS_NODE:
		if len(root.Nodes) != 1 {
			panic("Parse Error! Got minus node with not 1 children")
		}
		return -1 * InterpretTree(root.Nodes[0])
	}
	return 0.0
}
