package practice

import (
	"fmt"
)

/*
	汉诺塔，比较经典的一个游戏
	分析：
	先看下移动步数
	1个元素->1
	2个元素->2
	3个元素->7
	....
	总移动步数就是：2^N - 1

	再来看下移动过程：
	1. 把当前节点上面的所有节点，统一按照顺序，借用缓存，移到B
	2. 将当前节点直接移到C
	3. 再把已经移到B的所有节点，再统一移到C


*/

type Hanoi struct {
	number int
}

func NewHanoi(n int)*Hanoi{
	hanoi := new(Hanoi)
	hanoi.number = n
	return hanoi
}

func (hanoi *Hanoi)Start(){
	hanoi.Move(hanoi.number,"A","C","B")
}

func (hanoi *Hanoi)Move(n int,from string ,to string, cache string ){
	//hanoi.Print(n,from,to)
	//if n != hanoi.number {
	//	hanoi.Print(n,from,to)
	//}

	if n == 1{
		hanoi.Print(1,from,to)
		return
	}
	hanoi.Move( n - 1,from,cache,to)
	//hanoi.Print(n,from,to)
	//hanoi.Move( n,    "A","C")
	//hanoi.Move(  1,"A","C")
	hanoi.Print(n,from ,to)
	hanoi.Move( n - 1,    cache,to,from)
}

func TestHanoi(){
	hanoi := NewHanoi(3)
	hanoi.Start()
}

func (hanoi *Hanoi)Print(n int,from string,to string){
	fmt.Println("n" , ":",n , "  from:" ,from," to:",to)
}