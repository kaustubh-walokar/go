package node

import ("fmt")

type Node struct {
	Data interface{}
	Left *Node
	Right *Node
}

func NewNode(d interface {}) (* Node){
	n := new(Node)
	n.Data = d
	return n
}

func (n *Node) Print() (string){
	return Print(n)
}

func Print(root *Node) (string) {
	
	if(root == nil) {
		return "."
	} 
		
	l := Print(root.Left)
	r := Print(root.Right)

	return fmt.Sprintf("{%v %v %v}",l ,root.Data, r) 
}
