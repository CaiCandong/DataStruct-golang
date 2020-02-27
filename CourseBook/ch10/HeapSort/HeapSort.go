package HeapSort

//测试样例:堆排序:int 大根堆
// 1.如何由一个无序序列建成一个堆
// 2.如何输出堆顶元素之后，调整剩余元素成为一个新的堆
// 完全二叉树
type MaxHeap []int

const MAX = 1<<31 - 1

func Init(data []int) *MaxHeap {
	h := &MaxHeap{MAX}
	*h = append(*h, data...)
	for i := len(*h) / 2; i >= 1; i-- {
		h.AdjustDown(i)
	}
	//2.对每个非叶子节点执行下沉操作 选最大

	return h
}
func (h *MaxHeap) Add(elem int) {
	*h = append(*h, elem)
	h.AdjustUp()
}
func (h *MaxHeap) Delete() (res int) {
	res = (*h)[1]
	(*h)[1] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	if !h.Empty() {
		h.AdjustDown(1)
	}
	return
}
func (h *MaxHeap) AdjustUp() { //上浮，添加元素时使用
	i := len(*h) - 1
	temp := (*h)[i] //每次上浮最后一个元素
	for temp > (*h)[i/2] {
		(*h)[i] = (*h)[i/2]
		i /= 2
	}
	(*h)[i] = temp
}
func (h *MaxHeap) AdjustDown(i int) { //下沉，删除元素和初始化时使用
	//对以i为根的堆进行调整
	temp := (*h)[i]
	for i*2 < len(*h) { //有儿子
		index := i * 2
		if index+1 < len(*h) && (*h)[index] < (*h)[index+1] {
			index++
		}
		if temp > (*h)[index] {
			break
		}
		(*h)[i] = (*h)[index]
		i = index
	}
	(*h)[i] = temp
}
func (h *MaxHeap) Empty() bool {
	return len(*h) == 1
}
