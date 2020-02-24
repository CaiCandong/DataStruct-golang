package insertSort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 直接插入排序
func TestStraightInsertSort(t *testing.T) {
	qs := getData()
	ass := assert.New(t)
	for _, q := range qs {
		StraightInsertSort(q.question)
		ass.Equal(q.ans, q.question)
	}
}
