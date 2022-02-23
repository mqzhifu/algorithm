package practice

import (
	"algorithm/container"
	"sync"
)

type StackQueue struct {
	StackPush *container.Stack
	StackPop *container.Stack
	Debug int
	MoveLock 	sync.Mutex
}


func NewStackQueue(sort int)*StackQueue{
	stackQueue := new(StackQueue)
	stackQueue.Debug 	 = 1

	stackQueue.StackPush 	 = container.NewStack(100,container.STACK_FLAG_LINKED_LIST,1,container.ORDER_NONE)
	stackQueue.StackPop = container.NewStack(100,container.STACK_FLAG_LINKED_LIST,1,container.ORDER_NONE)
	return stackQueue
}

func (stackQueue *StackQueue ) Pop()(empty bool,searchNode *container.ListNode){
	stackQueue.MoveLock.Lock()
	defer stackQueue.MoveLock.Unlock()

	if !stackQueue.StackPop.IsEmpty(){
		return stackQueue.StackPop.Pop()
	}else{
		if stackQueue.StackPush.IsEmpty(){
			return true,searchNode
		}else{
			stackQueue.Move()
			return stackQueue.StackPop.Pop()
		}
	}

}

func (stackQueue *StackQueue ) Push(keyword int ){
	stackQueue.MoveLock.Lock()
	defer stackQueue.MoveLock.Unlock()

	stackQueue.StackPush.Push(keyword,"")
}

func (stackQueue *StackQueue )Move(){
	for{
		if stackQueue.StackPush.IsEmpty(){
			break
		}
		_, popNode :=  stackQueue.StackPush.Pop()
		stackQueue.StackPop.Push(popNode.Keyword,popNode.Data)
	}
}
