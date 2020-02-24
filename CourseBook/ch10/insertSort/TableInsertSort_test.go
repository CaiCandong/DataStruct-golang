package insertSort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableInsertSort(t *testing.T) {
	data := getData()
	ass := assert.New(t)
	for _, q := range data {
		tableInsertSortAdapter(q.question)
		ass.Equal(q.question, q.ans)
	}
}
