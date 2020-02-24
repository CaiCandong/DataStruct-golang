package insertSort

// 提供排序算法的测试数据
type data struct {
	question []int
	ans      []int
}

func getData() []data {
	return []data{
		{[]int{5, 2, 3, 1},
			[]int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0},
			[]int{0, 0, 1, 1, 2, 5}},
		{ []int{49, 38, 65, 97, 76, 13, 27, 49}, //课本数据
			[]int{13, 27, 38, 49, 49, 65, 76, 97}},
		{[]int{49, 38, 65, 97, 76, 13, 27, 49,55,4}, //课本数据
			[]int{4,13, 27, 38, 49, 49, 55,65, 76, 97}},
		//{question: []int{}, ans: []int{}},
	}
}
