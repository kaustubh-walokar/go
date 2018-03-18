package main 

import ("fmt"
		"github.com/kaustubh-walokar/go/node")

func main() {
	n := node.NewNode(10)
	l := node.NewNode(20)
	r := node.NewNode(30)
	m := &node.Node{40, nil, nil }

	n.Right = l
	l.Right = r
	r.Right = m

	fmt.Println("n:=", n.Print())
	
}
