package DisjointSet

type Graph struct { //存储图的数据结构
	point int      //点的个数
	edge  [][2]int //存储边
}

type node struct {
	level int   //记录层级,使树平衡
	root  *node //指向代表
}

func CycleDetect(graph Graph) bool { //环检测
	data := make([]node, graph.point)
	for _, edge := range graph.edge {
		left, right := &data[edge[0]], &data[edge[1]]
		for left.root != nil {
			left = left.root
		}
		for right.root != nil {
			right = right.root
		}
		if left == right {
			//fmt.Println("检测到环")
			return true
		}
		if left.level > right.level {
			right.root = left
		} else if left.level < right.level {
			left.root = right
		} else {
			left.root = right
			left.level++
		}
	}
	//fmt.Println("未检测到环")
	return false
}
