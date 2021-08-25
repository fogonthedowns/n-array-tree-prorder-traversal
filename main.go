package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	// a stack to store node, an output to store result
	stack := []*Node{root}
	output := []int{}

	// while the stack is not empty
	for len(stack) > 0 {
		root = stack[0]                   // get the first root
		output = append(output, root.Val) // store a value
		stack = stack[1:]                 // update stack, shift array right

		if len(root.Children) > 0 {
			stack = append(root.Children, stack...) // prepend Children to stack
		}
	}
	return output
}
