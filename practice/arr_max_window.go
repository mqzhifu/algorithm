package practice

import (
	"algorithm/container"
	"errors"
	"fmt"
)

type ArrMaxWindow struct {
	Debug int
	ArrPool []int
	Window	int
	Queue *container.Queue
}

func NewArrMaxWindow(windows int)*ArrMaxWindow{
	arrMaxWindow := new(ArrMaxWindow)
	arrMaxWindow.Debug = 1
	arrMaxWindow.Window = windows
	arrMaxWindow.ArrPool = []int{0}
	arrMaxWindow.Queue = container.NewQueue(100,container.STACK_FLAG_LINKED_LIST,container.ORDER_NONE,0)
	return arrMaxWindow
}

//输出信息，用于debug
func (arrMaxWindow *ArrMaxWindow) Print(a ...interface{}) (n int, err error) {
	if arrMaxWindow.Debug > 0{
		return fmt.Println(a)
	}
	return
}

func  (arrMaxWindow *ArrMaxWindow)makeError(msg string)error{
	arrMaxWindow.Print("[errors] " + msg)
	return errors.New(msg)
}

func (arrMaxWindow *ArrMaxWindow)AddNumber(number int){
	arrMaxWindow.ArrPool = append(arrMaxWindow.ArrPool,number)
}

func  (arrMaxWindow *ArrMaxWindow)PrintPushInfo(index int,from string){
	arrMaxWindow.Print("push index:",index , " value:",arrMaxWindow.ArrPool[index], " from",from)
}

func  (arrMaxWindow *ArrMaxWindow)PrintPopInfo(index int,from string){
	arrMaxWindow.Print("pop index:",index , " value:",arrMaxWindow.ArrPool[index], " from",from)
}

func  (arrMaxWindow *ArrMaxWindow)PrintMax(i,index int,from string){
	if i >= arrMaxWindow.Window  {
		arrMaxWindow.Print("max index:",index , " value:",arrMaxWindow.ArrPool[index], " from",from)
	}
}

func  (arrMaxWindow *ArrMaxWindow)GetMaxWindow(){
	if !arrMaxWindow.Queue.IsEmpty(){
		arrMaxWindow.makeError("queue not empty~")
		return
	}

	if arrMaxWindow.Window <= 0{
		arrMaxWindow.makeError("Window <= 0")
		return
	}

	if len(arrMaxWindow.ArrPool) <= 0{
		arrMaxWindow.makeError("ArrPool <= 0")
		return
	}
	arrMaxWindow.Print("arrMaxWindow len: ",len(arrMaxWindow.ArrPool))
	for i:=1;i<len(arrMaxWindow.ArrPool);i++{
		empty,node := arrMaxWindow.Queue.GetOneByEnd()
		if empty{
			arrMaxWindow.PrintMax(i,i,"empty1")
			arrMaxWindow.PrintPushInfo(i,"empty1")
			arrMaxWindow.Queue.Push(i,"")
			continue
		}

		if i - node.Keyword >= arrMaxWindow.Window{
			arrMaxWindow.Queue.Pop()
			empty,node = arrMaxWindow.Queue.Pop()
			if empty{
				arrMaxWindow.PrintMax(i,i,"empty2")
				arrMaxWindow.PrintPushInfo(i,"empty2")
				continue
			}
		}

		if arrMaxWindow.ArrPool[i] < arrMaxWindow.ArrPool[node.Keyword]{
			arrMaxWindow.PrintMax(i,node.Keyword,"<")
			arrMaxWindow.PrintPushInfo(i,"<")
			arrMaxWindow.Queue.Push(i,"")
			continue
		}
		arrMaxWindow.PrintPopInfo(node.Keyword,"")
		arrMaxWindow.Queue.PopByFirst()
		arrMaxWindow.PrintPushInfo(i,"last")
		arrMaxWindow.Queue.PushByEnd(i,"")


		arrMaxWindow.PrintMax(i,i,"last")

		//if i >= arrMaxWindow.Window - 1 {
		//	_,maxNode := arrMaxWindow.Queue.GetOneByEnd()
		//	arrMaxWindow.Print(arrMaxWindow.ArrPool[maxNode.Keyword])
		//}
	}

}

func  TestArrMaxWindow(){
	arrMaxWindow := NewArrMaxWindow(3)

	arrMaxWindow.AddNumber( 4 )
	arrMaxWindow.AddNumber( 3 )
	arrMaxWindow.AddNumber( 5 )
	arrMaxWindow.AddNumber( 4 )
	arrMaxWindow.AddNumber( 3 )
	arrMaxWindow.AddNumber( 3 )
	arrMaxWindow.AddNumber( 6 )
	arrMaxWindow.AddNumber( 7 )

	arrMaxWindow.GetMaxWindow()

	fmt.Println(arrMaxWindow)

	// 1 4 3 7 2 9 8
}
