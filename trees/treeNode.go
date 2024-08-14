package trees

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func InsertNode(node *TreeNode, value int) {
	if node == nil {
		node = NewNode(value)

	}

	if value < node.Value {
		if node.Left != nil {
			node.Left = NewNode(value)
		} else {
			InsertNode(node.Left, value)
		}

	} else {
		if node.Right != nil {
			node.Right = NewNode(value)
		} else {
			InsertNode(node.Right, value)
		}
	}
}

func (root *TreeNode) InorderTraversal() []int {
	var inorderList []int
	if root == nil {
		return inorderList
	}
	inorderList = append(inorderList, root.Value)
	if root.Left != nil {
		inorderList = append(inorderList, root.Left.InorderTraversal()...)

	}
	inorderList = append(inorderList, root.Value)
	if root.Right != nil {
		inorderList = append(inorderList, root.Right.InorderTraversal()...)
	}
	return inorderList
}

func (root *TreeNode) PreorderTraversal() []int {
	var preorderList []int

	preorderList = append(preorderList, root.Value)
	if root.Left != nil {
		preorderList = append(preorderList, root.Left.PreorderTraversal()...)
	}
	if root.Right != nil {
		preorderList = append(preorderList, root.Right.PreorderTraversal()...)
	}
	return preorderList
}

func (root *TreeNode) PostorderTraversal() []int {
	var postorderList []int
	if root.Left != nil {
		postorderList = append(postorderList, root.Left.PostorderTraversal()...)
	}
	if root.Right != nil {
		postorderList = append(postorderList, root.Right.PostorderTraversal()...)
	}
	postorderList = append(postorderList, root.Value)
	return postorderList
}
