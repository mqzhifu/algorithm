package container

import (
	"fmt"
	"os"
)

const (
	STACK_FLAG_ARRAY = 1
	STACK_FLAG_LINKED_LIST = 1
)

type Stack struct {
	Max		int
	Flag 	int	//数组|链表
	List 	ListInterface
}

func NewStack(max int,flag int ,debug int)*Stack{
	stack := new(Stack)
	stack.Max = max
	stack.Flag = flag
	if flag == STACK_FLAG_ARRAY{
		stack.List = NewArrayList(ORDER_NONE,max,debug)
	}else if flag == STACK_FLAG_LINKED_LIST{
		stack.List = NewLinkedList(ORDER_NONE,max,false,debug)
	}else{
		fmt.Println("flag err.")
		os.Exit(11)
	}
	return stack
}

func (stack *Stack)Push(keyword int,data interface{})(int,error){
	return stack.List.InsertNodeByFirst(keyword  ,data  )
}

func  (stack *Stack)Pop( )(empty bool,searchNode *ListNode){
	return stack.List.FindOneNodeByLocationAndDel(0)
}

//判断链表是否为：空节点
func (stack *Stack)IsEmpty()bool{
	return stack.List.IsEmpty()
}