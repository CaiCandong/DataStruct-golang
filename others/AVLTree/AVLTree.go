package AVLTree

import (
	"fmt"
)

//平衡二叉树
type TreeNode struct {
	Val    int //值
	height int //当前高度
	//factor      int       //平衡因子
	left, right *TreeNode //左右指针
}

//树结构 保存根节点和节点数量
type AVLTree struct {
	root *TreeNode //指向根节点的指针
	//size int       //总的节点个数
}

func NewAVLTree(val int) *AVLTree {
	return &AVLTree{
		root: &TreeNode{Val: val},
	}
}
func (t *AVLTree) Insert(val int) {
	t.root = insert(t.root, val)
	//t.size++
}
func (t *AVLTree) Delete(val int) {
	t.root = delete(t.root, val)
	t.root = adjust(t.root)
}

/* LL
//                A
//              /  \
//             B   	AR    ----->          B
//            / \                       /   \
//			 C   BR					   C     A
//	  		/ \						  / \   / \
//		   CL  CR                   CL  CR BR  AR
*/
func LL(A *TreeNode) *TreeNode {
	//1.完成结构转变
	B := A.left
	A.left = B.right
	B.right = A
	//2.更新A,B的高度
	A.height = max(getHeight(A.left), getHeight(A.right)) + 1
	//B.height = max(getHeight(B.left), getHeight(B.right)) + 1
	//C.height = max(getHeight(C.left), getHeight(C.right)) + 1
	//3.返回新的根节点
	return B
}

/*
//   	  A
//      /  \
//     AL   B                        B
//         /  \   ------->         /   \
//        BL   C                  A     C
//            / \               /  \   / \
//          CL   CR            AL  BL CL  CR
//
*/
func RR(A *TreeNode) *TreeNode {
	//调整位置
	B := A.right
	A.right = B.left
	B.left = A
	//更新高度
	//B.height = max(getHeight(B.left), getHeight(B.right)) + 1
	A.height = max(getHeight(A.left), getHeight(A.right)) + 1
	//C.height = max(getHeight(C.left), getHeight(C.right)) + 1
	return B
}

/*
//   	  A
//      /  \
//     B    AR                     C
//    /  \      ------->         /   \
//   BL   C                     B     A
//       /  \                 /  \   / \
//      CL  CR               BL  CL CR AR
//
*/
func LR(A *TreeNode) *TreeNode {
	B, C := A.left, A.left.right
	B.right, A.left = C.left, C.right
	C.left, C.right = B, A
	A.height = max(getHeight(A.left), getHeight(A.right)) + 1
	B.height = max(getHeight(B.left), getHeight(B.right)) + 1
	//C.height = max(getHeight(C.left), getHeight(C.right)) + 1
	return C
}

/*
//   	  A
//      /  \
//     AL   B                      C
//         / \  ------->         /   \
//        C   BR                A     B
//       /  \                 /  \   / \
//      CL  CR               AL  CL CR BR
//
*/
func RL(A *TreeNode) *TreeNode {
	//1.调整结构
	B, C := A.right, A.right.left
	B.left, A.right = C.right, C.left
	C.left, C.right = A, B
	//2.更新高度
	A.height = max(getHeight(A.left), getHeight(A.right)) + 1
	B.height = max(getHeight(B.left), getHeight(B.right)) + 1
	//C.height = max(getHeight(C.left), getHeight(C.right)) + 1
	return C
}
func adjust(root *TreeNode) *TreeNode {
	rl, rr := getHeight(root.left), getHeight(root.right)
	//1.左子树比右子树高2层
	//if getHeight(root.left)-getHeight(root.right) == 2 {
	if rl-rr == 2 {
		if getHeight(root.left.left)-getHeight(root.left.right) > 0 {
			//在左子树的左节点上进行插入导致不平衡
			//LL
			return LL(root)
		} else {
			//在左子树的右节点上进行插入导致不平衡
			//LR
			return LR(root)
		}
	} else if rl-rr == -2 {
		//2.左子树比右子树低两层
		if getHeight(root.right.left)-getHeight(root.right.right) > 0 {
			//在右子树的左节点上进行插入导致不平衡
			//RL
			return RL(root)
		} else {
			//在右子树的右节点上进行插入导致不平衡
			//RR
			return RR(root)
		}
	}
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	return root
}
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val, height: 1}
	}
	if val < root.Val {
		root.left = insert(root.left, val)
		root = adjust(root)
	} else if val > root.Val {
		root.right = insert(root.right, val)
		root = adjust(root)
	}
	//root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	return root
}
func delete(root *TreeNode, val int) *TreeNode {
	//递归边界结束条件
	if root == nil {
		return nil
	}
	//1.val<root.Val
	if val < root.Val {
		root.left = delete(root.left, val)
		return adjust(root)
		//2.val>root.Val
	} else if val > root.Val {
		root.right = delete(root.right, val)
		return adjust(root)
		//3.val==root.Val
	} else {
		switch {
		//(1).root为叶子节点
		case root.left == nil && root.right == nil:
			return nil
		//(2).root左子树为空
		case root.left == nil:
			return root.right
		//(3).root右子树为空
		case root.right == nil:
			return root.left
		//(4).root左右子树都非空
		default:
			//左子树的最大节点(root的直接前驱)肯定为没有右子树 temp
			//temp
			//然后交换root和temp 不影响root.left的二叉搜索树结构
			//递归删除temp
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}
			root.Val, temp.Val = temp.Val, root.Val
			root.left = delete(root.left, val)
			return adjust(root)
		}
	}
}
func getHeight(root *TreeNode) int {
	if root != nil {
		return root.height
	}
	return 0
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//树的中序遍历(迭代)
func (t AVLTree) InOrder() {
	inOrder(t.root)
}
func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root)
	inOrder(root.right)
}

//func main() {
//	tree := NewAVLTree(6)
//	data := []int{3, 8, 2, 5, 7, 9, 1, 4, 11, 10, 12}
//	for i := 0; i < len(data); i++ {
//		tree.Insert(data[i])
//	}
//	tree.Delete(6)
//	tree.InOrder()
//}
