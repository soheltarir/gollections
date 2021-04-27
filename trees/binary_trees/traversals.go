package binary_trees

/**
Recursive Auxiliary functions
*/

// inorderTraversalAuxiliary is a recursive function to traverse the tree InOrder depth first
func inorderTraversalAuxiliary(node *Node, result []interface{}) []interface{} {
	if node.Left != nil {
		result = inorderTraversalAuxiliary(node.Left, result)
	}
	result = append(result, node.Value.Key())
	if node.Right != nil {
		result = inorderTraversalAuxiliary(node.Right, result)
	}
	return result
}

// preorderTraversalAuxiliary is a recursive function to traverse the tree PreOrder depth first
func preorderTraversalAuxiliary(node *Node, result []interface{}) []interface{} {
	result = append(result, node.Value.Key())
	if node.Left != nil {
		result = preorderTraversalAuxiliary(node.Left, result)
	}
	if node.Right != nil {
		result = preorderTraversalAuxiliary(node.Right, result)
	}
	return result
}

// postorderTraversalAuxiliary is a recursive function to traverse the tree PostOrder depth first
func postorderTraversalAuxiliary(node *Node, result []interface{}) []interface{} {
	if node.Left != nil {
		result = postorderTraversalAuxiliary(node.Left, result)
	}
	if node.Right != nil {
		result = postorderTraversalAuxiliary(node.Right, result)
	}
	result = append(result, node.Value.Key())
	return result
}

/*****************************************************************************************************/

func (t *Tree) InOrderTraversal() []interface{} {
	var result []interface{}
	return inorderTraversalAuxiliary(t.Root, result)
}

func (t *Tree) PreOrderTraversal() []interface{} {
	var result []interface{}
	return preorderTraversalAuxiliary(t.Root, result)
}

func (t *Tree) PostOrderTraversal() []interface{} {
	var result []interface{}
	return postorderTraversalAuxiliary(t.Root, result)
}
