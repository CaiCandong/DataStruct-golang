package BinarySortTree

import (
	"fmt"
)

// Binary Sort Tree 二叉排序树
type TreeNode struct {
	Val         int
	left, right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
	size int
}

func (t BinarySearchTree) String() string {
	return t.InOrder()
}
func (t BinarySearchTree) InOrder() string {
	// 非递归 使用栈
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	curr := t.root
	for len(stack) != 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.right
	}
	fmt.Println(res)
	return fmt.Sprintf("%v", res)
}

// 查找算法
// 查找成功返回 指向该节点的指针,true
// 查找失败返回 指向插入位置的父亲节点，false
func Search(root *TreeNode, parent *TreeNode, key int) (p *TreeNode, ok bool) {
	if root == nil {
		return parent, false
	}
	switch {
	case root.Val == key:
		return root, true
	case key < root.Val:
		return Search(root.left, root, key)
	default:
		return Search(root.right, root, key)
	}
}

// 插入算法
func (t *BinarySearchTree) Insert(elem int) bool {
	if t.root == nil {
		t.root = &TreeNode{Val: elem}
	}
	parent, root := t.root, t.root
	for root != nil {
		if root.Val < elem {
			parent = root
			root = root.right
		} else if root.Val > elem {
			parent = root
			root = root.left
		} else {
			return false
		}
	}
	if parent.Val > elem {
		parent.left = &TreeNode{Val: elem}
	} else {
		parent.right = &TreeNode{Val: elem}
	}
	return true
}
func (t *BinarySearchTree) Inserts(nums []int) {
	for _, item := range nums {
		t.Insert(item)
	}
}

// 删除算法
// 删除root 节点
// 将root.right子树 连接到root.left子树的最大元素后面
func (t *BinarySearchTree) Delete(elem int) {
	if t.root == nil {
		return
	}
	if t.root.Val == elem { //删除根节点元素
		p := t.root.left
		for p.right != nil {
			p = p.right
		}
		p.right = t.root.right
		t.root = t.root.left
	}
	//1.寻找删除元素位置,执行完毕后root指向要删除的元素
	root, parent := t.root, t.root
	for root != nil && root.Val != elem {
		if root.Val > elem {
			parent = root
			root = root.left
		} else if root.Val < elem {
			parent = root
			root = root.right
		}
	}
	//2.执行删除任务
	if root == nil {
		return
	}

	//root在parent的左子树上
	if parent.Val > elem {
		//1.root的左子树为空
		//3.root的左右子树都为空
		if root.left == nil {
			parent.left = root.right
		}
		//2.root的右子树为空,左子树不为空
		if root.right == nil {
			parent.left = root.left
		}
		//4.root的左右子树都不为空
		if root.right != nil && root.left != nil {
			//将右子树接到左子树最大的元素后面
			p := root.left
			for p.right != nil {
				p = p.right
			}
			parent.left = root.left
			p.right = root.right
		}
	} else if parent.Val < elem { //右树上
		//1.root的左子树为空
		//3.root的左右子树都为空
		if root.left == nil {
			parent.right = root.right
		}
		//4.root的右子树为空
		if root.right == nil {
			parent.right = root.left
		}
		if root.right != nil && root.left != nil {
			p := root.left
			for p.right != nil {
				p = p.right
			}
			parent.right = root.left
			p.right = root.right
		}
	}
}

func main() {
	tree := BinarySearchTree{}
	tree.Inserts([]int{45, 24, 53, 45, 12, 24, 90})
	fmt.Println(tree)
}
