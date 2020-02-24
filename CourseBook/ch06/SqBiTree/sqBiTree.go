package SqBiTree

var (
	MaxTreeSize int = 1024 //数组大小
	NULL        int = 0    //代表未使用
)

type sqBiTree struct {
	elem []int //存储树结构的切片
}

// left:i*2+1 right:i*2+2

func New() *sqBiTree { //预分配MaxTreeSize个空间
	return &sqBiTree{elem: make([]int, MaxTreeSize)}
}

func (root sqBiTree) Empty() bool {
	return root.elem[0] == NULL
}
