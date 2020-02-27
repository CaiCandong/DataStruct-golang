package HeapSort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	data := []int{49, 38, 65, 97, 76, 13, 27, 49}
	rs := MaxHeap{MAX, 97, 76, 65, 49, 49, 13, 27, 38}
	ass := assert.New(t)
	ass.Equal(rs, *Init(data))
}
func TestMaxHeap_Add(t *testing.T) {
	h := MaxHeap{MAX, 97, 76, 65, 49, 49, 13, 27, 38}
	h.Add(100)
	fmt.Print(h)
}
func TestHeapSort(t *testing.T) {
	data := []int{49, 38, 65, 97, 76, 13, 27, 49}
	heap := Init(data)
	ass := assert.New(t)
	ass.Equal(97, heap.Delete())
}
