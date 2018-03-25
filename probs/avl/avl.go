package main

import ("fmt")

type AvlNode struct {
	Left *AvlNode
	Right *AvlNode
	Height int
	Data int
}

func (n *AvlNode) Print() (string){
	return Print(n)
}

func Print(root *AvlNode) (string) {
	
	if(root == nil) {
		return "."
	} 
		
	l := Print(root.Left)
	r := Print(root.Right)

	return fmt.Sprintf("{%v(%v,%v)%v}",l ,root.Data, root.Height, r) 
}

func newNode(data int) (* AvlNode){
	return &AvlNode{nil, nil, 0, data}
}

func insert(data int, t *AvlNode) (*AvlNode){
	if (t == nil) {
		return newNode(data)
	}
	
	if(data < t.Data ) {
		t.Left = insert(data, t.Left)
	}

	if(data > t.Data ) {
		t.Right = insert(data, t.Right)
	}

	var l,r = 0,0
	if t.Left != nil {
		l= t.Left.Height
	}

	if t.Right != nil {
		r= t.Right.Height
	}

	if(l<r){
		t.Height = r +1 
	} else {
		t.Height = l +1 
	}

	return t;
}

// func leftRotate(t *node.Node, parent *node.Node) (*node.Node){
	
// 	isRight := false;

// 	if(parent.Right == t) {
// 		isRight = true;
// 	}

// 	tmp := t.Right
// 	t.Right = tmp.Left;
// 	tmp.Left = t

// 	if(isRight) {
// 		parent.Right = tmp;
// 	} else {
// 		parent.Left = tmp;
// 	}

// 	return parent;
// }

func main() {
	
	m := insert(1, nil)
	insert(0, m)
	insert(2, m)

	fmt.Println("n:=", m.Print())
	
}


