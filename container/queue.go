package container

type Queue struct {
	Max		int
	Flag 	int	//数组|链表
	Order 	int
	List 	ListInterface
}

func NewQueue(max int,flag int,order int ,debug int)*Queue{
	queue := new(Queue)
	queue.Max = max
	queue.Flag = flag
	if flag == STACK_FLAG_ARRAY{
		queue.List = NewArrayList(order,max,debug)
	}else if flag == STACK_FLAG_LINKED_LIST{
		queue.List = NewLinkedList(order,max,false,1)
	}
	return queue
}

func (queue *Queue)Push(keyword int,data interface{})(int,error){
	return queue.List.InsertNodeByFirst(keyword  ,data  )

}

func  (queue *Queue)Pop(keyword int)(empty bool,searchNode *ListNode){
	return queue.List.FindOneNodeByLocationAndDel(queue.List.Length() - 1)
}