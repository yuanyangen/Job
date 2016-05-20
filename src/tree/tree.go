package tree

type TreeNode struct {
	value string
	left  *TreeNode
	right *TreeNode
}

var RootNode = &TreeNode{}

//初始化一颗二叉树
func init() {
	RootNode.left = &TreeNode{}
	RootNode.left.left = &TreeNode{}
	RootNode.left.right = &TreeNode{}
	RootNode.right = &TreeNode{}
	RootNode.left.right.left = &TreeNode{}

}

//获取一棵树的深度
func GetDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftNum := GetDepth(node.left)
	rightNum := GetDepth(node.right)
	if leftNum > rightNum {
		return leftNum + 1
	} else {
		return rightNum + 1
	}
}

//获取一棵树的节点个数
func GetNodeNum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return GetNodeNum(node.left) + GetNodeNum(node.right) + 1
}
