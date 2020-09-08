package parser

import (
	"fmt"
)

type Printer struct {
}

func NewPrinter() Printer {
	return Printer{}
}

func (w Printer) EnterNode(n Node) bool {
	switch s := n.(type) {
	case *AssignmentStatement:
		s.Variable.Walk(w)
		fmt.Print(" = ")
		s.Expression.Walk(w)
		fmt.Println()
		return false

	case *ScopeVar:
		fmt.Print(s.Name)

	case *BinaryNode:
		s.op1.Walk(w)
		fmt.Print(s.operation)
		s.op1.Walk(w)
	case *ValueNode:
		switch s := s.value.(type) {
		case *NumberValue:
			fmt.Print(s.value)
		case *StringValue:
			fmt.Print(s.value)
		}
	}

	return true
}

func (w Printer) LeaveNode(n Node) {}
