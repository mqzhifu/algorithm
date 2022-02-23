package container

import (
	"errors"
	"fmt"
	"math"
)

const (
	HEAP_SORT_ASC = 1
	HEAP_SORT_DESC = 2
)

type HeapNode struct{
	//Parent *HeapNode
	//Left *HeapNode
	//Right *HeapNode
	Keyword int
}

type Heap struct {
	//RootNode *HeapNode
	//Length int
	Debug int
	Sort int
	Pool []*HeapNode
}

func NewHeap(debug int,sort int)*Heap{
	heap := new(Heap)
	heap.Debug = debug
	heap.Sort = sort
	return heap
}

func (heap *Heap)IsEmpty()bool{
	if  heap.GetLength() <= 0 {
		return true
	}
	return false
}

//输出信息，用于debug
func (heap *Heap) Print(a ...interface{}) (n int, err error) {
	if heap.Debug > 0{
		return fmt.Println(a)
	}
	return
}
//创建一个error,统一管理
func (heap *Heap)makeError(msg string)error{
	heap.Print("[errors] " + msg)
	return errors.New(msg)
}

func (heap *Heap)GetLength()int{
	return len(heap.Pool)
}

func(heap *Heap)GetFirstNode()*HeapNode{
	return heap.Pool[0]
}

func(heap *Heap)GetLastNode()*HeapNode{
	return heap.Pool[heap.GetLength() - 1]
}

func(heap *Heap) InsertNode(keyword int){
	heap.Print("InsertNode : ",keyword)
	newHeapNode := heap.NewHeapNode(keyword)
	if heap.IsEmpty(){
		//heap.Length = 1
		//heap.RootNode = newHeapNode
		heap.Pool = []*HeapNode{newHeapNode}
		return
	}
	lastLocation := heap.GetLength()
	heap.Print("lastLocation :",lastLocation)
	heap.Pool = append(heap.Pool,newHeapNode)
	for{
		parentIndex := (lastLocation - 1) / 2
		heap.Print("parentIndex : ",parentIndex)
		if heap.Sort == HEAP_SORT_ASC{
			if heap.Pool[lastLocation].Keyword < heap.Pool[parentIndex].Keyword{
				heap.SwapHeapNode(lastLocation,parentIndex)
			}
		}else{
			if heap.Pool[lastLocation].Keyword > heap.Pool[parentIndex].Keyword{
				heap.SwapHeapNode(lastLocation,parentIndex)
			}
		}
		if parentIndex == 0{
			break
		}
		lastLocation = parentIndex
	}
}

func(heap *Heap)SwapHeapNode(index1 int ,index2 int){
	tmp := heap.Pool[index1]
	heap.Pool[index1] = heap.Pool[index2]
	heap.Pool[index2] = tmp
}

func(heap *Heap)NewHeapNode(keyword int)*HeapNode{
	newHeapNode := HeapNode{Keyword: keyword}
	return &newHeapNode
}

func(heap *Heap)GetDeep()int{
	if heap.IsEmpty(){
		return 0
	}
	times := 0
	base := 2
	length := heap.GetLength()
	//heap.Print("GetDeep times:",times , " base:",base , " length:",length)
	// 1 2 4 8 16 32 89
	for{
		length = length / base
		if length <= 0{
			times++
			break
		}

		times++
	}
	//heap.Print("times:",times)
	return times
}


func(heap *Heap)ForEach(){
	if heap.GetLength() <=0 {
		return
	}
	deep :=heap.GetDeep()
	var everyLevelStart int
	var everyLevelEnd   int
	for i:=1;i<=deep;i++{
		fmt.Print(i,"-")
		power := i - 1
		if i == 1{
			everyLevelStart = 0
			everyLevelEnd = 0
		}else{
			everyLevelStart = int( math.Pow(2,float64(power ))) - 1
			everyLevelEnd = everyLevelStart * 2
			if everyLevelEnd >= heap.GetLength() - 1{
				everyLevelEnd = heap.GetLength() - 1
			}
		}
		//fmt.Println("everyLevelStart:",everyLevelStart ," everyLevelEnd:",everyLevelEnd)
		for j:=everyLevelStart;j<= everyLevelEnd;j++{
			fmt.Print(heap.Pool[j].Keyword, " ")
		}

		fmt.Println()

	}

	//newHeapNode := HeapNode{Keyword: keyword}
	//return &newHeapNode
}

func(heap *Heap)PopOneNode()(backNode HeapNode){
	if heap.IsEmpty(){
		return backNode
	}

	if heap.GetLength() == 1{
		node := heap.GetFirstNode()
		heap.Pool = []*HeapNode{}
		return *node
	}

	first := *heap.GetFirstNode()
	heap.Pool[0] = heap.GetLastNode()
	//删除最后一个元素
	heap.Pool = append(heap.Pool[:len(heap.Pool)-1])

	location := 0
	for {
		node := heap.Pool[location]
		leftIndex := location * 2 + 1
		rightIndex := location * 2 + 2

		//heap.Print("location:",location , " leftIndex:",leftIndex , " rightIndex:",rightIndex)

		if leftIndex > heap.GetLength() - 1 {
			//heap.Print("break in 1")
		}else  if leftIndex == heap.GetLength() - 1{
			//heap.Print("break in 2")
			if heap.Sort == HEAP_SORT_DESC{
				if node.Keyword < heap.Pool[leftIndex].Keyword{
					heap.SwapHeapNode(location,leftIndex)
				}
			}else{
				if node.Keyword > heap.Pool[leftIndex].Keyword{
					heap.SwapHeapNode(location,leftIndex)
				}
			}

			break
		}else{
			//heap.Print("location:",location , heap.Pool[location].Keyword,  " leftIndex:",heap.Pool[leftIndex].Keyword , " rightIndex:",heap.Pool[rightIndex].Keyword)
			min := node
			minLocation := location
			//一定有右节点
			if heap.Sort == HEAP_SORT_DESC{
				if min.Keyword < heap.Pool[leftIndex].Keyword{
					min = heap.Pool[leftIndex]
					minLocation = leftIndex
				}

				if min.Keyword < heap.Pool[rightIndex].Keyword{
					min = heap.Pool[rightIndex]
					minLocation = rightIndex
				}
			}else{
				if min.Keyword > heap.Pool[leftIndex].Keyword{
					min = heap.Pool[leftIndex]
					minLocation = leftIndex
				}

				if min.Keyword > heap.Pool[rightIndex].Keyword{
					min = heap.Pool[rightIndex]
					minLocation = rightIndex
				}
			}

			if minLocation == location{
				break
			}
			heap.SwapHeapNode(location,minLocation)
			location = minLocation
		}

		//if leftIndex < heap.GetLength() - 1 || rightIndex > heap.GetLength() - 1{
		//
		//}
		//if node.Keyword <= heap.Pool[leftIndex].Keyword && node.Keyword <= heap.Pool[rightIndex].Keyword{
		//	break
		//}

		if leftIndex > heap.GetLength() - 1 || rightIndex > heap.GetLength() - 1{
			//heap.Print("break in 3")
			break
		}
	}
	//heap.Print("-----")
	//heap.ForEach()

	return first
}



func TestHeap(){
	heap := NewHeap(1,HEAP_SORT_DESC)
	//heap := NewHeap(1,HEAP_SORT_ASC)
	forEnd := 12
	for i:=0;i<forEnd;i++{
		num :=GetRandIntNum(100)
		heap.InsertNode(num)
	}
	for i:=0;i<len(heap.Pool);i++{
		fmt.Println(heap.Pool[i].Keyword)
	}


	fmt.Println(" --------- ")

	//heap.GetDeep()
	heap.ForEach()
	fmt.Println(" --------- ")

	for i:=0;i<forEnd;i++{
		popNode := heap.PopOneNode()
		fmt.Println("popNode:",popNode.Keyword)
	}

	//fmt.Println(" ")
	//for i:=0;i<len(heap.Pool);i++{
	//	fmt.Println(heap.Pool[i].Keyword)
	//}


}