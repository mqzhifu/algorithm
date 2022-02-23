package practice

import (
	"algorithm/container"
	"fmt"
)

const (
	LIMIT_STACK_MAX = 1
	LIMIT_STACK_MIN = 2
)

type LimitStack struct {
	Stack *container.Stack
	Sort int
	SortStack *container.Stack
	Debug int
}

func NewLimitStack(sort int)*LimitStack{
	limitStack := new(LimitStack)
	limitStack.Sort 	 = sort
	limitStack.Debug 	 = 1

	limitStack.Stack 	 = container.NewStack(100,container.STACK_FLAG_LINKED_LIST,1,container.ORDER_NONE)
	limitStack.SortStack = container.NewStack(100,container.STACK_FLAG_LINKED_LIST,1,container.ORDER_NONE)

	return limitStack
}

func(limitStack *LimitStack) IsEmpty()bool{
	return limitStack.Stack.IsEmpty()
}

//输出信息，用于debug
func (limitStack *LimitStack) Print(a ...interface{}) (n int, err error) {
	if limitStack.Debug > 0{
		return fmt.Println(a)
	}
	return
}

func (limitStack *LimitStack) Push(keyword int){
	data := "push_aaa"
	if limitStack.IsEmpty(){
		limitStack.Stack.Push(keyword,data)
		limitStack.SortStack.Push(keyword,data)
		return
	}

	limitStack.Stack.Push(keyword,data)
	_,sortStackNode := limitStack.SortStack.Pop()
	if limitStack.Sort == LIMIT_STACK_MAX{
		if keyword >= sortStackNode.Keyword{
			limitStack.SortStack.Push(sortStackNode.Keyword,sortStackNode.Data)
			limitStack.SortStack.Push(keyword,data)
		}else{
			limitStack.SortStack.Push(sortStackNode.Keyword,sortStackNode.Data)
		}
	}else{
		if keyword <= sortStackNode.Keyword{
			limitStack.SortStack.Push(sortStackNode.Keyword,sortStackNode.Data)
			limitStack.SortStack.Push(keyword,data)
		}else{
			limitStack.SortStack.Push(sortStackNode.Keyword,sortStackNode.Data)
		}
	}
}
func (limitStack *LimitStack) Pop()(node *container.ListNode){
	empty,stackNode := limitStack.Stack.Pop()
	if empty{
		return stackNode
	}

	empty,sortStackNode := limitStack.SortStack.Pop()
	fmt.Println(empty)
	if stackNode.Keyword == sortStackNode.Keyword{
		return stackNode
	}else{
		limitStack.SortStack.Push(sortStackNode.Keyword,sortStackNode.Data)
	}
	return node
}

func  TestLimitStack(){
	limitStack :=  NewLimitStack(LIMIT_STACK_MIN)
	limitStack.Push(10)
	limitStack.Push(1)
	limitStack.Push(4)
	limitStack.Push(5)
	limitStack.Push(3)

	TestLimitStackShowNodeList(limitStack)

	limitStack.Pop()
	limitStack.Pop()
	limitStack.Pop()
	limitStack.Pop()

	TestLimitStackShowNodeList(limitStack)
}

func TestLimitStackShowNodeList(limitStack *LimitStack){
	_,list := limitStack.Stack.GetAllByFirst()
	for _,v:=range list{
		fmt.Print(v.Keyword," ")
	}
	fmt.Println( "" )
	_,listSort := limitStack.SortStack.GetAllByFirst()
	for _,v:=range listSort{
		fmt.Print(v.Keyword," ")
	}
	fmt.Println( "" )
}