package main

import (
	"fmt"
	"learning/channel/c11channel-tree-resovle/tree"
)

func traverseWithFunc() {
	root := tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{Value: 5, Left: nil, Right: nil}
	root.Left.Right = tree.CrateNode(2)
	root.Right.Left = new(tree.TreeNode)
	root.Right.Left.SetValue(4)
	root.TraverseFunc(func(tn *tree.TreeNode) {
		fmt.Printf("%d\t", tn.Value)
	})
}

func traverseWithChannel() {
	root := tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{Value: 5, Left: nil, Right: nil}
	root.Left.Right = tree.CrateNode(2)
	root.Right.Left = new(tree.TreeNode)
	root.Right.Left.SetValue(4)

	maxValue := 0
	c := root.TraverseWithChannel()
	for node := range c {
		if node.Value > maxValue {
			maxValue = node.Value
		}
	}
	fmt.Printf("max value = %d", maxValue)
}

func main() {
	traverseWithFunc()
	fmt.Println()
	traverseWithChannel()
}
