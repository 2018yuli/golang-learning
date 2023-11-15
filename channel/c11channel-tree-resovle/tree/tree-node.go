package tree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (e *TreeNode) SetValue(i int) {
	e.Value = i
}

func (node *TreeNode) TraverseFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *TreeNode) TraverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	// 使用协程，遍历树
	go func() {
		node.TraverseFunc(func(tn *TreeNode) {
			out <- tn
		})
		// 记得完成之后关闭协程
		close(out)
	}()
	return out
}

func CrateNode(i int) *TreeNode {
	return &TreeNode{i, nil, nil}
}
