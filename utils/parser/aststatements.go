package parser

import "fmt"

type Statement interface {
	Execute() (Value, error)
	ToString() string
	getNodesList() []Node
	Walk(v Visitor)
}

//////////////////

type ScopeVar struct {
	Name string
}

func NewScopeVar(name string) *ScopeVar {
	return &ScopeVar{Name: name}
}

func (as *ScopeVar) Execute() (Value, error) {
	return nil, nil
}

func (as *ScopeVar) ToString() string {
	return as.Name
}

func (as *ScopeVar) getNodesList() []Node {
	return nil
}

func (as *ScopeVar) Walk(v Visitor) {
	v.EnterNode(as)

	v.LeaveNode(as)
}

/* Assignment Statement */

type AssignmentStatement struct {
	Variable   *ScopeVar
	Expression Node
}

func (as *AssignmentStatement) New(variable string, node Node) *AssignmentStatement {
	return &AssignmentStatement{Variable: NewScopeVar(variable), Expression: node}
}

func (as *AssignmentStatement) Execute() (Value, error) {
	// result, err := as.Expression.Execute()
	// if err == nil {
	// 	VarTable[as.Variable] = result
	// }
	return nil, nil
}

func (as *AssignmentStatement) ToString() string {
	return " = "
}

func (as *AssignmentStatement) getNodesList() []Node {
	return []Node{as.Expression}
}

func (as *AssignmentStatement) Walk(v Visitor) {
	v.EnterNode(as)

	as.Variable.Walk(v)
	as.Expression.Walk(v)

	v.LeaveNode(as)
}

//////////////////

/* Print Statement */

type PrintStatement struct {
	node Node
}

func (ps *PrintStatement) Execute() (Value, error) {
	result, _ := ps.node.Execute()
	fmt.Println(result.AsString())
	return nil, nil
}

func (ps *PrintStatement) ToString() string {
	return " print "
}

func (ps *PrintStatement) getNodesList() []Node {
	return []Node{ps.node}
}

func (ps *PrintStatement) Walk(v Visitor) {
	v.EnterNode(ps)

	ps.node.Walk(v)

	v.LeaveNode(ps)
}
