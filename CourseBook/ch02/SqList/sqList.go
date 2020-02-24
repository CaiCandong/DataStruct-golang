package SqList

import "fmt"

const ListInitSize int = 5
const ListIncrement int = 10

type ElemType int
type sqList struct {
	elem   []ElemType
	length int
	size   int
}
type Compare interface {
	compare(i, j int) bool
}

func NewSqList() *sqList {
	res := &sqList{}
	res.init()
	return res
}
func (L *sqList) init() {
	L.elem = make([]ElemType, ListInitSize)
	L.length = 0
	L.size = ListInitSize
}
func (L *sqList) Destroy() {
	if L == nil {
		return
	}
	L.elem = nil //直接将其置为nil,相应内存等待自行释放
	L.length = 0
	L.size = 0
}
func (L *sqList) Clear() {
	if L == nil {
		return
	}
	L.length = 0
}
func (L sqList) IsEmpty() bool {
	if L.length == 0 {
		return true
	}
	return false
}
func (L sqList) Length() int {
	return L.length
}
func (L sqList) GetElem(i int) ElemType {
	return L.elem[i]
}

// 在第i个位置插入新的元素e,插入元素的索引为i
func (L *sqList) Insert(i int, e ElemType) bool {
	if i < 0 || i > L.length {
		return false
	}
	if L.length == L.size && !L.realloc() { //容量已满并且重新分配失败,利用短路与特性
		return false
	}
	for j := L.length; j > i; j-- {
		L.elem[j] = L.elem[j-1]
	}
	L.elem[i] = e
	L.length++
	return true
}

func (L *sqList) realloc() bool {
	L.size += ListIncrement
	newBase := make([]ElemType, L.size)
	for i := 0; i < L.length; i++ {
		newBase[i] = L.elem[i]
	}
	L.elem = newBase
	return true
}
func (L *sqList) Delete(i int) (ElemType, bool) {
	var res ElemType
	if i < 0 || i >= L.length {
		return res, false
	}
	res = L.elem[i]
	for ; i < L.length-1; i++ {
		L.elem[i] = L.elem[i+1]
	}
	L.length--
	return res, true
}
func (L *sqList) Locate(e ElemType, f func(e1, e2 ElemType) bool) (int, bool) {
	for i, elem := range L.elem {
		if f(e, elem) {
			return i, true
		}
	}
	return -1, false
}

// 返回元素e的前一个元素的索引
func (L *sqList) PriorElem(e ElemType) (int, bool) {
	for i, elem := range L.elem {
		if e == elem && i != 0 {
			return i - 1, true
		}
	}
	return -1, false
}
func (L *sqList) NextElem(e ElemType) (int, bool) {
	for i, elem := range L.elem {
		if e == elem && i != L.length-1 {
			return i + 1, true
		}
	}
	return -1, false
}
func (L sqList)String()string{
	return fmt.Sprintf("{elem:%v,length:%v,size:%v}",L.elem[0:L.length],
		L.length,L.size)
}
