package insertSort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBPathsInsertSort(t *testing.T) {
	qs := getData()
	ass := assert.New(t)
	for _, q := range qs {
		BPathsInsertSort(q.question)
		ass.Equal(q.ans, q.question)
	}
}
