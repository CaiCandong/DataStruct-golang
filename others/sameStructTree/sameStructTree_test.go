package sameStructTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSameStructTree(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(true, sameStructTree(getDefaultDataSame()))
	ass.Equal(false, sameStructTree(getDefaultDataDiff()))
}
func getDefaultDataDiff() (tree1, tree2 []rawData) {
	//默认测试数据
	return []rawData{
			{'A', 1, 2},       //0
			{'B', 3, 4},       //1
			{'C', 6, Null},    //2
			{'D', Null, Null}, //3
			{'E', 5, Null},    //4
			{'F', Null, Null}, //5
			{'G', 7, Null},    //6
			{'H', Null, Null}, //7
		},
		[]rawData{
			{'G', Null, 1},    //0
			{'H', Null, Null}, //1
			{'C', 7, 5},       //2
			{'A', 4, 2},       //3
			{'B', 0, Null},    //4
			{'E', Null, Null}, //5
			{'D', 7, Null},    //6
			{'F', Null, Null}, //7
		}
}

func getDefaultDataSame() (tree1, tree2 []rawData) {
	//默认测试数据
	return []rawData{
			{'A', 1, 2},       //0
			{'B', 3, 4},       //1
			{'C', 6, Null},    //2
			{'D', Null, Null}, //3
			{'E', 5, Null},    //4
			{'F', Null, Null}, //5
			{'G', 7, Null},    //6
			{'H', Null, Null}, //7
		},
		[]rawData{
			{'G', Null, 1},    //0
			{'H', Null, Null}, //1
			{'C', 0, Null},    //2
			{'A', 2, 4},       //3
			{'B', 5, 6},       //4
			{'E', 7, Null},    //5
			{'D', Null, Null}, //6
			{'F', Null, Null}, //7
		}
}
