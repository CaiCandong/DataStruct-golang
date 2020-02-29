package sameStructTree

/*
 树的同构问题
 1.树怎么表示?什么数据结构存储树？
 2.怎么寻找到树的根节点?
 3.怎么判别树是否同构
*/
const Null = -1

type rawData struct { //存储原始数据
	Val   int32 //存储的值:字符
	left  int   //左节点在数组中的位置
	right int   //右节点在数组中的位置
}

func sameStructTree(leftTree, rightTree []rawData) bool {
	//leftTree, rightTree := getDefaultDataDiff()
	leftRoot, rightRoot := getRoot(leftTree), getRoot(rightTree)
	//fmt.Println(leftRoot, rightRoot)
	return isSameTree(leftRoot, rightRoot, leftTree, rightTree)
	//return false
}
func isSameTree(leftRoot, rightRoot int, leftTree, rightTree []rawData) bool {
	if leftRoot == Null && rightRoot == Null {
		return true
	}
	if leftRoot == Null || rightRoot == Null {
		return false
	}
	//fmt.Printf("当前比较:%v-%v,%c,%c\n", leftRoot, rightRoot, leftTree[leftRoot].Val, rightTree[rightRoot].Val)
	if leftTree[leftRoot].Val != rightTree[rightRoot].Val {
		return false
	}
	ll, lr := leftTree[leftRoot].left, leftTree[leftRoot].right
	rl, rr := rightTree[rightRoot].left, rightTree[rightRoot].right
	flag1 := isSameTree(ll, rl, leftTree, rightTree) && isSameTree(lr, rr, leftTree, rightTree)
	flag2 := isSameTree(ll, rr, leftTree, rightTree) && isSameTree(lr, rl, leftTree, rightTree)
	return flag1 || flag2
}
func getRoot(tree []rawData) int {
	length := len(tree)
	check := make([]bool, length)
	for i := 0; i < length; i++ {
		if tree[i].left != Null {
			check[tree[i].left] = true
		}
		if tree[i].right != Null {
			check[tree[i].right] = true
		}
	}
	for i := 0; i < length; i++ {
		if check[i] == false {
			return i
		}
	}
	return -1
}
