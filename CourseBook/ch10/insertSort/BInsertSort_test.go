package insertSort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBInsertSort(t *testing.T) {
	qs := getData()
	ass := assert.New(t)
	for _, q := range qs {
		BInsertSort(q.question)
		ass.Equal(q.ans, q.question)
	}
}
