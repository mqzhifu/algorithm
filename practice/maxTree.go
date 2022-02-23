package practice

import (
	"algorithm/container"
	"fmt"
)

type MaxTree struct {
	Debug int
	Arr []int
	LeftMax map[int]int
	RightMax map[int]int
	Stack *container.Stack
}

func NewMaxTree()*MaxTree{
	maxTree := new(MaxTree)
	maxTree.LeftMax = make(map[int]int)
	maxTree.RightMax = make(map[int]int)
	maxTree.Stack= container.NewStack(100,container.STACK_FLAG_LINKED_LIST,1,container.ORDER_NONE)
	return maxTree
}

func (maxTree *MaxTree) Print(a ...interface{}) (n int, err error) {
	if maxTree.Debug > 0{
		return fmt.Println(a)
	}
	return
}

func (maxTree *MaxTree)SetArr(arr []int){
	maxTree.Arr = arr
}

func (maxTree *MaxTree)Start(){
	if len(maxTree.Arr) <=0 {
		return
	}

	for i:=0;i<len(maxTree.Arr);i++ {
		for  {
			empty ,popOnde := maxTree.Stack.Pop()
			if empty{
				maxTree.Stack.Push(i,"")
				break
			}

			if maxTree.Arr[i] < popOnde.Keyword{
				maxTree.Stack.Push(popOnde.Keyword,"")
				maxTree.Stack.Push(i,"")
				break
			}
		}

	}
}

func TestMaxTree(){
	maxTree := NewMaxTree()
	arr := []int{4,8,5,2,4,9,0,3}
	maxTree.SetArr(arr)
	maxTree.Start()
}

