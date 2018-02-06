package node

import "fmt"

type node struct {
	Data interface{}
	Left *node
	Right *node
}

func NewNode(d interface {}) (* node){
	n := new(node)
	n.Data = d
	return n
}