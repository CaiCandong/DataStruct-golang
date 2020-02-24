package SqList

import (
	"fmt"
	"testing"
)

func TestNewSqList(t *testing.T) {
	list:=NewSqList()//init测试
	fmt.Println(list)
}
func TestSqList_Destroy(t *testing.T) {
	list:=NewSqList()//init测试
	list.Destroy()//destroy测试
	fmt.Println(list)
	list=nil
	list.Destroy()//destroy nil测试
	fmt.Println(list)
}

