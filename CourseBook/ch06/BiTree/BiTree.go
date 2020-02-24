package BiTree

var NULL = -1 << 63

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Empty() bool { //二叉树是否为空
	return t == nil
}

func New(Val int) *TreeNode {
	return &TreeNode{Val: Val}
}

// 创建二叉树方法
//1.层次遍历构造树
func Ints2TreeNode(ints []int) *TreeNode {
	n:= len(ints)
	if n<1{
		return nil
	}
	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root}
