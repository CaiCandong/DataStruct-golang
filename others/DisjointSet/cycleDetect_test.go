package DisjointSet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getDefaultData() *Graph {
	return &Graph{
		point: 6,
		edge: [][2]int{
			{0, 1},
			{1, 2},
			{1, 3},
			{2, 5},
			{3, 4},
			{2, 4},
		},
	}
}
func TestCycleDetect(t *testing.T) {
	data:=getDefaultData()
	ass:=assert.New(t)
	ass.Equal(true,CycleDetect(*data))
}