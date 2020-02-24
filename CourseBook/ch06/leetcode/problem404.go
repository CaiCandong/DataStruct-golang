package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	return helper(root,false)
}
func helper(root *TreeNode,isLeft bool)int{
	if root==nil{
		return 0
	}
	if root.Left==nil&&root.Right==nil&&isLeft{//左子叶
		return root.Val
	}
	return helper(root.Left,true)+helper(root.Right,false)
}
