package parser

type BlockContext struct {
	Vars map[string]struct{}
}

func NewBlockContext() *BlockContext {
	return &BlockContext{Vars: make(map[string]struct{})}
}
