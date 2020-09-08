package parser

import (
	"fmt"
)

type BlockWalker struct {
	ctx *BlockContext
}

func NewBlockWalker() BlockWalker {
	return BlockWalker{ctx: NewBlockContext()}
}

func (w BlockWalker) EnterNode(n Node) bool {
	switch s := n.(type) {
	case *AssignmentStatement:
		fmt.Println("sss")

	case *ScopeVar:
		_, ok := w.ctx.Vars[s.Name]
		if !ok {
			w.ctx.Vars[s.Name] = struct{}{}
		}
		fmt.Println(s.Name)

	case *BinaryNode:
		switch s.operation {
		case '+':

		}

	}

	return true
}

func (w BlockWalker) LeaveNode(n Node) {}
