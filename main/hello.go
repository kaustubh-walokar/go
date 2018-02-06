package main 

import ("fmt"
		"github.com/kaustubh-walokar/go/node")

func main() {
	n := node.NewNode(10)
	fmt.Println("hello", n)
	n.Data = 20
	fmt.Println("hello", n)
}
