package container

import (
	"errors"
	"fmt"
	"reflect"
)

type Huffman struct {
	RootNode *HuffmanNode			//根节点
	Length 	int						//当前树的节点总数
	Content string					//要压缩的内容
	Debug int
	DictCode map[string]string		//每个字符最后对应的-字典压缩码
	DictCntTimes map[string]int		//每个字符的出现次数
	DictLocationSort []string
	Stack *Stack
}

type HuffmanNode struct{
	Left *HuffmanNode	//左节点
	Right *HuffmanNode	//右节点
	Keyword 	int
	Char 		string	//当前字符
	CharTimes 	int		//当前字符出现的次数
}

func NewHuffman ( debug int)*Huffman{
	huffman :=  new (Huffman)
	huffman.Debug = debug
	huffman.DictCode = make(map[string]string)
	huffman.DictCntTimes = make(map[string]int)
	return huffman
}
//设置需要压缩的字符
func(huffman *Huffman) SetContent(str string){
	huffman.Content = str
}

//输出信息，用于debug
func (huffman *Huffman) Print(a ...interface{}) (n int, err error) {
	if huffman.Debug > 0{
		return fmt.Println(a)
	}
	return
}

func  (huffman *Huffman)makeError(msg string)error{
	huffman.Print("[errors] " + msg)
	return errors.New(msg)
}
//开始统计每个字符出现的次数
func (huffman *Huffman)ContentCntTime(){
	dictCnt := make(map[string]int)
	for i:=0;i<len(huffman.Content );i++{
		oneChar := string( huffman.Content[i])
		//huffman.Print(huffman.Content[i])
		//aaa := huffman.Content[i]
		_ ,ok := dictCnt[oneChar]
		if ok {
			dictCnt[oneChar] ++
		}else{
			dictCnt[oneChar] = 1
		}
	}
	huffman.DictCntTimes = dictCnt
	huffman.Print("char dict map , Cnt times : ",dictCnt)
}

func (huffman *Huffman)ContentSort(){
	//开始，给字符排序，根据：出现次数
	dictCntCharTimesArr := []int{}//将每个字符的出现次数，放入到一个新的数组中
	for _,v:=range huffman.DictCntTimes{
		dictCntCharTimesArr = append(dictCntCharTimesArr,v)
	}
	//循环比对大小，最终是一个有序数组
	for i:=0;i<len(dictCntCharTimesArr) - 1;i++{
		for j:=i;j<len(dictCntCharTimesArr) ;j++{
			if dictCntCharTimesArr[j] < dictCntCharTimesArr[i]{
				tmp := dictCntCharTimesArr[j]
				dictCntCharTimesArr[j] = dictCntCharTimesArr[i]
				dictCntCharTimesArr[i] = tmp
			}
		}
	}
	huffman.Print("dictCntCharTimesArr :",dictCntCharTimesArr)
	//上面只是计算了一个数组中的数字的大小，还得将数字与字符关联起来
	dictLocationSort := []string{}
	exceptKeys := make(map[string]int)
	for j:=0;j<len(dictCntCharTimesArr) ;j++{
		for k,v := range huffman.DictCntTimes{
			_ ,ok := exceptKeys[k]
			if ok {
				continue
			}
			if v == dictCntCharTimesArr[j]{
				dictLocationSort = append(dictLocationSort,k)
				exceptKeys[k] = 1
				break
			}
		}
	}
	huffman.Print("dictLocationSort:",dictLocationSort)
	huffman.DictLocationSort = dictLocationSort
}

func (huffman *Huffman)BuildTree(){
	//将排序好的字符压到堆栈里
	for i:=len(huffman.DictLocationSort) - 1;i>=0;i--{
		huffman.Print("list push times:",huffman.DictCntTimes[huffman.DictLocationSort[i]] , " char:",huffman.DictLocationSort[i])
		NewHuffmanNode := huffman.NewHuffmanNode(huffman.DictLocationSort[i],huffman.DictCntTimes[huffman.DictLocationSort[i]])
		//频率 / 字符
		huffman.Stack.Push(huffman.DictCntTimes[huffman.DictLocationSort[i]],NewHuffmanNode)
	}
	for {
		//每次弹出两个
		empty1 ,popOne := huffman.Stack.Pop()
		if empty1{
			huffman.makeError(" list.pop is empty 1~")
			break
		}

		empty2 ,popTwo:= huffman.Stack.Pop()
		if empty2{
			huffman.makeError(" list.pop is empty 2~")
			break
		}
		//一次：弹出两个栈节点，然后，新建一个树节点
		NewHuffmanNode1 := popOne.Data.(*HuffmanNode)
		NewHuffmanNode2 := popTwo.Data.(*HuffmanNode)
		huffman.Print("stack pop ,1 :",NewHuffmanNode1.Keyword,"-",NewHuffmanNode1.Char, " , 2:",NewHuffmanNode2.Keyword,"-",NewHuffmanNode2.Char)
		//将两个节点的：出现次数 相累加，再创建一个新的节点(树节点)
		NewHuffmanNode3 := huffman.NewHuffmanNode("",popOne.Keyword +popTwo.Keyword )
		//将两个节点挂在新的节点下面
		NewHuffmanNode3.Left = NewHuffmanNode1
		NewHuffmanNode3.Right = NewHuffmanNode2
		//如果栈里不再有元素，证明已经是最后两个节点了，再创建一个空的树节点，将两个值连起来，即完成了整个树的构建了
		if huffman.Stack.IsEmpty() {
			huffman.RootNode = NewHuffmanNode3
			break
		}
		//现在:弹出的2个栈节点、创建了新的1个节点(树节点)
		//新的树节点还要做为栈节点，重新压回到栈里，但得有规则：要顺序的

		popNodesRecord := []*ListNode{}//暂存每次弹出的栈节点
		for{
			//先弹出一个元素
			empty ,popNode := huffman.Stack.Pop()
			if empty{//如果为空，证明新的树节点是栈中最大的结点，栈中的节点没有比新节点更大的了
				huffman.Print(" list.pop is empty 3~")
				//先把新建的树节点压到栈中
				huffman.PushOneBack( NewHuffmanNode3)
				huffman.PushListBack( popNodesRecord)
				break
			}
			//huffman.Print("popNode:",popNode.Keyword,popNode.Data)
			popNodesRecord =  append(popNodesRecord,popNode)
			//对比两个元素的  频率  大小，符合条件即是要压入的位置
			huffman.Print(popNode.Keyword ,"（pop） >= ", NewHuffmanNode3.CharTimes , " (new)")
			//已找到合适位置，停止弹出，压入新节点，将旧节点再压回去
			if popNode.Keyword >= NewHuffmanNode3.CharTimes{
				huffman.Print("push back:",popNode.Keyword," - ",popNode.Data)
				huffman.Stack.Push(popNode.Keyword,popNode.Data)
				huffman.PushOneBack( NewHuffmanNode3)
				if len(popNodesRecord) > 1{
					//删除最后一个元素
					popNodesRecord = append(popNodesRecord[:len(popNodesRecord)-1])
					huffman.PushListBack(popNodesRecord)
				}
				break
			}
		}
	}
	huffman.Print("huffmanTreeRoot:",huffman.RootNode.CharTimes)
}

func (huffman *Huffman)Init(){
	if len(huffman.Content ) == 0 {
		huffman.makeError("content is empty , please first set content....")
		return
	}
	//创建一个堆栈容器
	huffman.Stack = NewStack(len(huffman.Content ),STACK_FLAG_LINKED_LIST,0,ORDER_NONE)
	huffman.ContentCntTime()
	huffman.ContentSort()
	huffman.BuildTree()
	huffman.BuildDictCode()

}
func (huffman *Huffman)ShowTree(){

	_,nowTreeListNode := huffman.EachDeepByBreadthFirst(false)
	for i:=1; i<=len(nowTreeListNode);i++{
		fmt.Print(i,"-")
		for j:=0;j<len(nowTreeListNode[i]);j++{
			fmt.Print(nowTreeListNode[i][j].Char," ",nowTreeListNode[i][j].Keyword)
		}
		fmt.Println(" ")
	}
}

func  (huffman *Huffman)Compress()string{
	str := ""
	for i:=0;i<len(huffman.Content);i++ {
		char := string(huffman.Content[i])
		code := huffman.DictCode[char]
		//huffman.Print(code)
		str += code
	}
	return str
}

func  (huffman *Huffman)Unpack(str string){
	node := huffman.RootNode
	for i:=0;i<len(str);i++ {
		char := string(str[i])
		if char == "0"{
			node  = node.Left
		}else{
			node = node.Right
		}

		if node.Left == nil && node.Right == nil{
			huffman.Print(node.Char)
			node = huffman.RootNode
		}
	}
}
//构建 每个字符的 编码
func  (huffman *Huffman)BuildDictCode(){
	huffman.EachAfter(huffman.RootNode,"-1","")
}
//后序遍历
func  (huffman *Huffman)EachAfter(node *HuffmanNode,fromNodeDir string,history string){
	if node == nil{
		//huffman.Print("fromNode:",fromNode.CharTimes,"|",fromNode.Char,fromNode.Left,fromNode.Right)
		return
	}
	if fromNodeDir == "left"{
		history += "0"
	}else if fromNodeDir == "right"{
		history += "1"
	}else{
		//history += "-"
	}

	if node.Left == nil && node.Right == nil{
		huffman.Print("path:",history, " char:",node.Char)
		huffman.DictCode[node.Char] = history
	}

	huffman.EachAfter(node.Left,"left",history)
	huffman.EachAfter(node.Right,"right",history)
	//huffman.Print(node.CharTimes,"-",node.Char)
}

func  (huffman *Huffman)PushOneBack( node *HuffmanNode){
	huffman.Print("PushOneBack :",node.Keyword)
	huffman.Stack.Push(node.Keyword,node)
}

func  (huffman *Huffman)PushListBack( listNodeRecord []*ListNode){
	if len(listNodeRecord) > 0{
		for i:=len(listNodeRecord) - 1;i>=0;i--{
			huffman.Print("PushListBack :",listNodeRecord[i].Keyword,listNodeRecord[i].Data)
			huffman.Stack.Push(listNodeRecord[i].Keyword,listNodeRecord[i].Data)
		}
	}else{
		huffman.Print("no need push back.")
	}
}


func  (huffman *Huffman)NewHuffmanNode(char string,charTimes int)*HuffmanNode{
	huffman.Print("NewHuffmanNode ,char :",char," charTimes:",charTimes)
	huffman.Length++
	return &HuffmanNode{Char: char,CharTimes: charTimes,Keyword: charTimes}
}

func TestHuffman(){
	huffman := NewHuffman(1)
	huffman.Content = "abcdefasdfadsfasdlkfjasd;lfkjasd;lfkjasd;lkflj"
	huffman.Init()

	str := huffman.Compress()
	huffman.Unpack(str)
}


//层级遍历-广度优先
//nodeNilFill:空节点的元素，是否需要填充
func (huffman *Huffman)EachDeepByBreadthFirst(nodeNilFill bool)(empty bool,finalNode map[int][]*HuffmanNode){
	//保存最终结果
	nodeContainer := make(map[int][]*ListNode)

	//if binaryTree.IsEmpty(){
	//	return true,finalNode
	//}
	//binaryTree.Print("tree len:",binaryTree.Len)
	//创建一个 无序 队列(数组类型)
	list := NewQueue(huffman.Length,STACK_FLAG_ARRAY,ORDER_NONE,0)
	firstNode := huffman.RootNode
	//先把首节点压入队列中
	list.Push(firstNode.Keyword,firstNode)
	//当前层级
	level := 0
	for  {
		level ++
		//一次进队列/出队列，遍历出来的一层的所有数据列表
		var nodeList []*ListNode
		//每次弹出一个节点，保存，后面再把该节点的左右节点继续压到队列中
		for{
			isEmpty , queueNode := list.Pop()
			if isEmpty{
				break
			}
			if queueNode != nil{
				treeNode,ok  := queueNode.Data.(*TreeNode)
				if ok {
					treeNode.DeepAsc = level
				}
			}
			nodeContainerOne,ok := nodeContainer[level]
			if ok{
				nodeContainer[level] = append(nodeContainerOne,queueNode)
			}else{
				nodeContainer[level] = []*ListNode{queueNode }
			}
			//保存本次弹出的节点
			nodeList = append(nodeList,queueNode)
		}
		if !nodeNilFill{
			if len(nodeList) <= 0 {
				break
			}
		}else{
			isEmpty := true
			for _,v:= range nodeList{
				if v.Keyword != KEYWORD_NIL{
					isEmpty = false
					break
				}
			}
			if isEmpty{
				break
			}
		}
		if nodeNilFill{
			for _,node:=range nodeList{
				//binaryTree.Print("level:",level, " k:", k , " ", node,node.Keyword)
				leftKeyword := KEYWORD_NIL
				rightKeyword := KEYWORD_NIL

				if node == nil {
					list.Push(leftKeyword,nil)
					list.Push(rightKeyword,nil)
					continue
				}

				if node.Keyword == KEYWORD_NIL{
					list.Push(leftKeyword,nil)
					list.Push(leftKeyword,nil)
					continue
				}

				treeNode := node.Data.(*TreeNode)
				if treeNode.Left != nil{
					list.Push(treeNode.Left.Keyword,treeNode.Left)
				}else{
					list.Push(leftKeyword,nil)
				}

				if treeNode.Right != nil{
					list.Push( treeNode.Right.Keyword,treeNode.Right)
				}else{
					list.Push(rightKeyword,nil)
				}
			}
		}else{
			//开始将上面弹出的节点的：左右子节点再重新压回到队列中
			for _,node:=range nodeList{
				//for k,node:=range nodeList{
				//binaryTree.Print("level:",level, " k:", k , " ", node,node.Keyword)
				if node == nil {//空节点，直接丢弃
					continue
				}
				//这里是对interface 做 nil 判断
				dateValueOf := reflect.ValueOf(node.Data)
				if dateValueOf.IsNil(){//空节点，直接丢弃
					continue
				}
				//断言
				treeNode,ok := node.Data.(*HuffmanNode)
				if !ok{
					huffman.Print("assertions failed.")
					continue
				}
				if treeNode.Left != nil{
					list.Push(treeNode.Left.Keyword,treeNode.Left)
				}
				if treeNode.Right != nil{
					list.Push(treeNode.Right.Keyword,treeNode.Right)
				}
			}
		}
	}
	finalNode = make(map[int][]*HuffmanNode)
	for k,nodeListRowArr:= range nodeContainer{
		var finalNodeListArr []*HuffmanNode
		for _,nodeList :=range nodeListRowArr{
			treeNode ,ok := nodeList.Data.(*HuffmanNode)
			if !ok {
				treeNode = nil
			}
			finalNodeListArr = append(finalNodeListArr,treeNode)
		}
		finalNode[k] = finalNodeListArr
	}

	//for i,v:=range finalNode{
	//	for _,v2:=range v{
	//		//fmt.Print(k,v,k2)
	//		if v2 == nil{
	//			//binaryTree.Print()
	//		}else{
	//			parent := 0
	//			if v2.Parent != nil{
	//				parent =v2.Parent.Keyword
	//			}
	//			fmt.Println(i,v2.Keyword , " parent:",parent)
	//		}
	//
	//	}
	//
	//}

	return false,finalNode
}
