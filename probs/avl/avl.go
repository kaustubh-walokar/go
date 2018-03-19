package main

import ("fmt"
		"github.com/kaustubh-walokar/go/node")


func newNode(data int) (* node.Node){
	return node.NewNode(data)
}

func insert(data int, t *node.Node) (*node.Node){
	if (t == nil) {
		return newNode(data)
	}

	if(data < t.Data.(int) ) {
		if  (t.Left ==nil) {
			t.Left = newNode(data)
			return t.Left
		} else {
			return insert(data, t.Left)
		}

	}

	if(data > t.Data.(int) ) {
		if  (t.Right ==nil) {
			t.Right = newNode(data)
			return t.Right
		} else {
			return insert(data, t.Right)
		}

	}


	return t
}

func main() {
	
	m := insert(1, nil)
	insert(0, m)
	insert(2, m)

	fmt.Println("n:=", m.Print())
	
}


