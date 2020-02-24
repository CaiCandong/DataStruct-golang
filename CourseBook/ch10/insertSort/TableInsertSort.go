package insertSort

type node struct {
	key  int
	next int
}

func tableInsertSortAdapter(nums []int) { //适配器，使用课本的表插入排序算法
	//特判
	if len(nums) <= 1 {
		return
	}
	//初始化
	tables := make([]node, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		tables[i+1].key = nums[i]
	}
	tables[0].key = 1 << 31
	tables[0].next = 1
	tables[1].next = 0
	// 调用表排序算法
	tableInsertSort(tables)
	// 调用arrange算法
	arrange(tables)

	for i := 0; i < len(nums); i++ {
		nums[i] = tables[i+1].key
	}
}
func tableInsertSort(tables []node) {
	for i := 2; i < len(tables); i++ {
		pre, p := tables[0].next, 0
		for tables[pre].key < tables[i].key {
			pre = tables[pre].next
			p = tables[p].next
		}
		tables[p].next = i
		tables[i].next = pre
	}
}

func arrange(tables []node) { // 按顺序排序tables
	p := tables[0].next //p指示第一个记录的当前位置
	for i := 1; i < len(tables); i++ {
		for p < i { //该循环的作用是避免p指向错误的位置
			p = tables[p].next
		}
		q := tables[p].next //暂存第i+1个元素的位置
		if p != i {
			tables[p], tables[i] = tables[i], tables[p]
			tables[i].next = p //指向原本元素当前所在的位置
		}
		p = q
	}
}

/*
	指针实现
*/
//type node struct {
//	key  int
//	next *node
//}
//func tableInsertSort(nums []int) {
//	if len(nums) <= 1 {
//		return
//	}
//	tables := make([]node, len(nums))
//	head := node{0, &tables[0]} //头节点
//	tables[0].next = &head
//	for i := 0; i < len(nums); i++ {
//		tables[i].key = nums[i]
//	}
//
//	for i := 1; i < len(nums); i++ {
//		pre, p := head.next, &head
//		for pre != &head && pre.key < nums[i] {
//			pre = pre.next
//			p = p.next
//		}
//		tables[i].next = pre
//		p.next = &tables[i]
//	}
//	i := 0
//	for p := head.next; p != &head; p = p.next {
//		nums[i] = p.key
//		i++
//	}
//}
