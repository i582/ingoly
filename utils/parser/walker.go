package parser

import (
	"fmt"
)

type BlockWalker struct {
	Vars map[string]struct{}
}

func NewBlockWalker() BlockWalker {
	return BlockWalker{Vars: make(map[string]struct{})}
}

func (w BlockWalker) EnterNode(n Node) bool {
	switch s := n.(type) {
	case *AssignmentStatement:
		fmt.Println("sss")

	case *ScopeVar:
		_, ok := w.Vars[s.Name]
		if !ok {
			w.Vars[s.Name] = struct{}{}
		}
		fmt.Println(s.Name)
	}

	return true
}

func (w BlockWalker) LeaveNode(n Node) {}
