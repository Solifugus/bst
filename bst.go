package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

type Deepest struct {
	keys  []int
	depth int
}

func main() {
	inputs := os.Args[1:]
	//inputs := [6]string{"12", "11", "90", "82", "7", "9"}
	input, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic("Inputs should be integers (not \"" + inputs[0] + "\")")
	}
	root := &Node{input, nil, nil}
	depth := 0
	deepest := Deepest{[]int{input}, 0}
	for i := 1; i < len(inputs); i += 1 {
		newKey, err := strconv.Atoi(inputs[i])
		if err != nil {
			panic("Inputs should be integers (not \"" + inputs[i] + "\")")
		}
		depth = insert(0, root, newKey)
		if depth > deepest.depth {
			deepest = Deepest{[]int{newKey}, depth}
		} else if depth == deepest.depth {
			deepest.keys = append(deepest.keys, newKey)
		}
	}
	fmt.Print("Deepest: ")
	fmt.Print(deepest.keys)
	fmt.Print(", Depth: ")
	fmt.Println(deepest.depth)
	//fmt.Println("And, btw, Hello World!")
}

func insert(depth int, node *Node, num int) int {
	if num == node.key {
		return depth
	}
	if num < node.key {
		if node.left == nil {
			node.left = &Node{key: num, left: nil, right: nil}
			return depth + 1
		} else {
			return insert(depth+1, node.left, num)
		}
	}
	if node.right == nil {
		node.right = &Node{key: num, left: nil, right: nil}
		return depth + 1
	} else {
		return insert(depth+1, node.right, num)
	}
}
