package main

import (
	"fmt"
	"encoding/json"
)

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

func Walk(t * Tree, ch chan int){
	if t.Left != nil { Walk(t.Left, ch)} 
	fmt.Println(t.Value)
	if t.Left != nil { Walk(t.Right, ch)}
}

func Same(t1, t2 * Tree) bool{
	return false
}

func main() {
	c := make(chan int, 3)
	
	tree := Tree{
		Left: &Tree{
			Left: &Tree{Value: 0},
			Value: 2,
			Right: &Tree{Value: 4}}, 
		Value: 5, 
		Right: &Tree{
			Left: &Tree{Value: 6},
			Value: 7,
			Right: &Tree{Value: 10}}}

	out, err := json.Marshal(tree)
    if err != nil {
        panic (err)
	}

	fmt.Println(string(out))
	
	Walk(&tree, c)	
}