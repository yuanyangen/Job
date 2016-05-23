package tree

//实现了二叉查找树的基本逻辑和功能

//但是二叉树的旋转怎么办呢？

type TreeNode struct {
	key   int
	left  *TreeNode
	right *TreeNode
}

var RootNode = &TreeNode{}

//初始化一颗二叉树
func init() {
	RootNode = &TreeNode{}
	RootNode.key = 10
	addNode(1, nil)
	addNode(2, nil)
	addNode(3, nil)
	addNode(20, nil)

}

//添加一个节点， 如果Node为nil 表示为空节点
func addNode(val int, Node * TreeNode) {
	if Node == nil {
		Node = &RootNode
	}
	//val已经存在
	if Node.key == val {
		panic("value already exists")
	} else if Node.key > val {
		if Node.left != nil {
			addNode(val, Node.left)
		} else {
			Node.left = &TreeNode{}
			Node.left.key = val
		}

	} else {
		if Node.right != nil {
			addNode(val,Node.right)
		} else {
			Node.right = &TreeNode{}
			Node.right.key = val
		}
	}
}

func Query(key int ) (int, error) {
	return QueryNode(key, RootNode)
}

 func QueryNode(key, Node *TreeNode) {
	 if key  == Node.key {
		 return Node.key, nil
	 } else if key < Node.key {
		 return QueryNode(key, Node.left)
	 } else {
		 return QueryNode(key, Node.right)
	 }
 }


