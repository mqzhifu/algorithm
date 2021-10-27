package container

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

const (
	DIRECTION_LEFT = 1
	DIRECTION_RIGHT = 2

	NODE_KEYWORD_PLACEHOLDER = 3


	TREE_NODE_MAX = 100
	KEYWORD_NIL = 999//有些需求，某些空节点必须得有，但是keyword得给一个int 占位符
)

type TreeNode struct{
	Parent *TreeNode
	Left *TreeNode
	Right *TreeNode
	Keyword int
	Data interface{}
}

type BinaryTree struct {
	RootNode *TreeNode
	Balance bool
	NodeMax int
	Len 	int
	Debug 	int
}

func NewBinaryTree(nodeMax int ,balance bool,debug int)*BinaryTree{
	binaryTree :=  new (BinaryTree)

	if nodeMax > TREE_NODE_MAX{
		nodeMax = TREE_NODE_MAX
	}

	binaryTree.NodeMax = nodeMax
	binaryTree.Balance = balance
	binaryTree.Debug = debug

	return binaryTree
}


func(binaryTree *BinaryTree) IsEmpty()bool{
	if  binaryTree.GetLength() <= 0 {
		return true
	}
	return false
}

//输出信息，用于debug
func (binaryTree *BinaryTree) Print(a ...interface{}) (n int, err error) {
	if binaryTree.Debug > 0{
		return fmt.Println(a)
	}
	return
}
//创建一个error,统一管理
func  (binaryTree *BinaryTree)makeError(msg string)error{
	binaryTree.Print("[errors] " + msg)
	return errors.New(msg)
}

func (binaryTree *BinaryTree)GetLength()int{
	return binaryTree.Len
}

func (binaryTree *BinaryTree)GetRootNode()*TreeNode{
	return binaryTree.RootNode
}
//创建一个新的节点
func (binaryTree *BinaryTree)NewOneNode(keyword int ,data interface{})*TreeNode{
	treeNode := TreeNode{
		Keyword: keyword,
		Data: data,
		Left: nil,
		Right: nil,
		Parent: nil,
	}
	return &treeNode
}
//插入一个新节点
func (binaryTree *BinaryTree)InsertOneNode(keyword int ,data interface{})error{
	binaryTree.Print("InsertOneNode  keyword:",keyword)

	if binaryTree.NodeMax > TREE_NODE_MAX{
		msg := "NodeMax > "+strconv.Itoa(TREE_NODE_MAX)
		return binaryTree.makeError(msg)
	}

	newNode := binaryTree.NewOneNode(keyword,data)
	if binaryTree.IsEmpty(){
		binaryTree.RootNode = newNode
		binaryTree.Len = 1
		return nil
	}

	node := binaryTree.GetRootNode()
	searchNode,direction,err := binaryTree.InsertOneNodeRecursionCompare(node,newNode,0,nil)
	if err != nil{
		return err
	}
	//if searchNode == nil{
	//	binaryTree.Print("InsertOneNodeRecursionCompare return: ",searchNode, " dir:" ,direction)
	//}else{
	//	binaryTree.Print("InsertOneNodeRecursionCompare return: ",searchNode.Keyword," dir:",direction)
	//}

	if searchNode == nil { //只有根节点,因为根节点的父节点为空
		searchNode = node
	}
	if direction == DIRECTION_LEFT{
		if searchNode.Left != nil{
			searchNode.Left.Parent = newNode
		}
		searchNode.Left = newNode

	}else{
		if searchNode.Right != nil{
			searchNode.Right.Parent = newNode
		}
		searchNode.Right = newNode
	}
	newNode.Parent = searchNode

	binaryTree.Len++
	return nil
}
//插入节点时，递归查找新元素应该插入到哪个元素的左右
//node:当前比对节点
//insertNode:要插入的新节点
//direction:算是个上下文保留值，记录 当前节点是从上个节点的 左/右 方向过来的，因为最后一次循环，肯定节点是nil
//parentNode:算是个上下文保留值，记录 当前节点的父节点，因为最后一次循环，肯定节点是nil
func (binaryTree *BinaryTree) InsertOneNodeRecursionCompare(node *TreeNode,insertNode *TreeNode ,direction int,parentNode  *TreeNode)(downNode *TreeNode , nodeDirection int ,err error){
	if node == nil{
		return parentNode,direction,nil
	}
	if insertNode.Keyword < node.Keyword{
		return binaryTree.InsertOneNodeRecursionCompare(node.Left,insertNode,DIRECTION_LEFT,node)
	}else if insertNode.Keyword > node.Keyword {
		return binaryTree.InsertOneNodeRecursionCompare(node.Right,insertNode,DIRECTION_RIGHT,node)
	}else{
		msg := "NodeKeyword: not allow repeat ."
		return downNode,direction,binaryTree.makeError(msg)
	}
}
//删除一个节点，根据keyword
func (binaryTree *BinaryTree)DelOneByKeyword(keyword int)(empty bool){
	if binaryTree.IsEmpty(){
		return true
	}
	searchNode,empty := binaryTree.FindOneByKeyword(keyword)
	if empty {
		return true
	}

	if searchNode.Left == nil{
		searchNode.Parent.Left = nil
	}else{
		searchNode.Left.Parent = searchNode.Parent
	}

	if searchNode.Right == nil{
		searchNode.Parent.Right = nil
	}else{
		searchNode.Right.Parent = searchNode.Parent
	}

	binaryTree.Len--

	return false
}
//递归查找一个节点
func (binaryTree *BinaryTree)FindOneByKeywordRecursionCompare(keyword int ,node *TreeNode)(downNode *TreeNode,empty bool){
	if node.Keyword == keyword{
		return node,false
	}

	if node == nil{
		return node,true
	}

	if keyword > node.Keyword {
		return binaryTree.FindOneByKeywordRecursionCompare(keyword,node.Right)
	}else if keyword < node.Keyword {
		return binaryTree.FindOneByKeywordRecursionCompare(keyword,node.Left)
	}else{//这里是出现了 相等的情况，按说insert时做了限制，但为防万一，还是加条输出吧
		binaryTree.makeError("FindOneByKeywordRecursionCompare in case :else...")
	}
	return downNode,true
}

func (binaryTree *BinaryTree)FindOneByKeyword(keyword int )(downNode *TreeNode,empty bool){
	if binaryTree.IsEmpty(){
		return downNode,true
	}
	node := binaryTree.GetRootNode()
	searchNode,empty := binaryTree.FindOneByKeywordRecursionCompare(keyword,node)
	return searchNode,empty
}
//层级遍历-广度优先
//nodeNilFill:空节点的元素，是否需要填充
func (binaryTree *BinaryTree)EachDeepByBreadthFirst(nodeNilFill bool)(empty bool,finalNode map[int][]int){
	//保存最终结果
	finalNode = make(map[int][]int)

	if binaryTree.IsEmpty(){
		return true,finalNode
	}
	//binaryTree.Print("tree len:",binaryTree.Len)
	//创建一个 无序 队列(数组类型)
	list := NewQueue(binaryTree.NodeMax,STACK_FLAG_ARRAY,ORDER_NONE,0)
	firstNode := binaryTree.GetRootNode()
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
			empty , data := list.Pop(0)
			if empty{
				break
			}
			finalNodeArr,ok := finalNode[level]
			if ok{
				finalNode[level] = append(finalNodeArr,data.Keyword)
			}else{
				finalNode[level] = []int{data.Keyword}
			}
			//保存本次弹出的节点
			nodeList = append(nodeList,data)
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
			for k,node:=range nodeList{
				binaryTree.Print("level:",level, " k:", k , " ", node,node.Keyword)
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
				dateValueOf :=  reflect.ValueOf(node.Data)
				if dateValueOf.IsNil(){//空节点，直接丢弃
					continue
				}
				//断言
				treeNode,ok := node.Data.(*TreeNode)
				if !ok{
					binaryTree.Print("assertions failed.")
					continue
				}
				//leftKeyword := KEYWORD_NIL
				//if treeNode.Left != nil{
				//	leftKeyword = treeNode.Left.Keyword
				//}
				//rightKeyword := KEYWORD_NIL
				//if treeNode.Right != nil{
				//	rightKeyword = treeNode.Right.Keyword
				//}
				//list.Push(leftKeyword,treeNode.Left)
				//list.Push(rightKeyword,treeNode.Right)
				if treeNode.Left != nil{
					list.Push(treeNode.Left.Keyword,treeNode.Left)
				}
				if treeNode.Right != nil{
					list.Push(treeNode.Right.Keyword,treeNode.Right)
				}
			}
		}
	}

	return false,finalNode
}
//层级遍历-深度优先
func (binaryTree *BinaryTree)EachDeepByDeepFirst(node  *TreeNode){

}
//获取当前树的深度/高度
func (binaryTree *BinaryTree)GetDeep(flag int)int{
	if binaryTree.IsEmpty(){
		return 0
	}
	var nodeList map[int][]int
	if flag == 1{
		_,nodeList = binaryTree.EachDeepByBreadthFirst(false)
	}else{

	}

	return len(nodeList)
}
//先序遍历
//func (binaryTree *BinaryTree)EachByOrder( node *TreeNode , list map[int]*TreeNode , order int){
func (binaryTree *BinaryTree)EachByOrder( node *TreeNode , list *[]int , order int){
	if binaryTree.IsEmpty(){
		return
	}

	if node == nil{
		return
	}
	if order == 1{
		binaryTree.EachByOrder(node.Left,list,order)
		InsertIntArray(list,node.Keyword)
		//binaryTree.Print(node.Keyword)
		binaryTree.EachByOrder(node.Right,list,order)
	}else if order == 2{
		InsertIntArray(list,node.Keyword)
		//binaryTree.Print(node.Keyword)
		binaryTree.EachByOrder(node.Left,list,order)
		binaryTree.EachByOrder(node.Right,list,order)
	}else if order == 3{
		binaryTree.EachByOrder(node.Left,list,order)
		binaryTree.EachByOrder(node.Right,list,order)
		InsertIntArray(list,node.Keyword)
		//binaryTree.Print(node.Keyword)
	}else{
		binaryTree.makeError("EachByOrder case ")
	}
}
func InsertIntArray( list *[]int,keyword int){
	for i:=0;i<len((*list));i++{
		if (*list)[i] == KEYWORD_NIL{
			(*list)[i] = keyword
			break
		}
	}
}
func GetNewIntArrayAndFillEmpty(len int)[]int{
	list :=  []int{}
	for i:=0;i<len;i++{
		list = append(list,KEYWORD_NIL)
	}
	return list
}

//先序遍历
func (binaryTree *BinaryTree)EachByFirst(  )[]int{
	list :=GetNewIntArrayAndFillEmpty(binaryTree.GetLength())
	node := binaryTree.GetRootNode()
	binaryTree.EachByOrder(node,&list,1)

	return list
}
//中序遍历
func (binaryTree *BinaryTree)EachByMiddle(  )[]int{
	list :=GetNewIntArrayAndFillEmpty(binaryTree.GetLength())
	node := binaryTree.GetRootNode()
	binaryTree.EachByOrder(node,&list,2)

	return list
}
//后序遍历
func (binaryTree *BinaryTree)EachByAfter(  )[]int{
	list :=GetNewIntArrayAndFillEmpty(binaryTree.GetLength())
	node := binaryTree.GetRootNode()
	binaryTree.EachByOrder(node,&list,3)

	return list
}
//获取树的最左/最右极值节点
func (binaryTree *BinaryTree)GetLimitNodeRecursionCompare(node *TreeNode,dir int )*TreeNode{
	if dir == DIRECTION_RIGHT{
		if node.Right == nil{
			return node
		}
		return binaryTree.GetLimitNodeRecursionCompare(node.Right,dir)
	}else{
		if node.Left == nil{
			return node
		}
		return binaryTree.GetLimitNodeRecursionCompare(node.Left,dir)
	}
}
//获取最左侧的节点（最小值）
func (binaryTree *BinaryTree)GetTheLeftNode( )*TreeNode{
	return binaryTree.GetLimitNodeRecursionCompare(binaryTree.RootNode,DIRECTION_LEFT)
}
//获取最右侧的节点（最大值）
func (binaryTree *BinaryTree)GetTheRightNode( )*TreeNode{
	return binaryTree.GetLimitNodeRecursionCompare(binaryTree.RootNode,DIRECTION_RIGHT)
}
//打印树形状
func (binaryTree *BinaryTree)ShowPrintTree(){
	nodeMiddleSpaceNum := 5//最后一行，两个节点间距离（空格数）
	_,deepList := binaryTree.EachDeepByBreadthFirst(true)
	levelCnt := len(deepList)
	//最后一列肯定全是nil，所以不用输出了
	recordLocation := make(map[int]map[int]int)
	for i:=levelCnt - 1;i >= 1;i--{
	//for i:=1;i< levelCnt;i++{
		fmt.Print(strconv.Itoa(i) + " ")
		//prefSpace := ""
		nodeLocation := 0
		oneRowLocationMap := make(map[int]int)
		recordLocation[i] = oneRowLocationMap
		for j:=0;j<len(deepList[i]);j++ {
			number := deepList[i][j]
			if number == 999{
				number = 0
			}
			numberStr := binaryTree.FormatNumberAlign(number)
			if i == len(deepList) - 1{//最后一行不需要添加首空格了
				if j % 2 == 0 {
					nodeLocation = nodeLocation + 3 + nodeMiddleSpaceNum
					fmt.Print(numberStr , binaryTree.GetSpaceStr(nodeMiddleSpaceNum))
				}else{
					nodeLocation = nodeLocation + 3
					fmt.Print(numberStr)
				}
				recordLocation[i][j+1] = nodeLocation
			}else{
				//tmp := NODE_KEYWORD_PLACEHOLDER + nodeMiddleSpaceNum + NODE_KEYWORD_PLACEHOLDER
				//firstPreSpace := (levelCnt - 2 - i ) * tmp
				//prefSpaceNum := 0
				//if j == 0 {
				//	prefSpaceNum =  firstPreSpace  + NODE_KEYWORD_PLACEHOLDER + 2
				//}else{
				//	prefSpaceNum =     NODE_KEYWORD_PLACEHOLDER + 2 + 3 +   firstPreSpace
				//}
				//
				//prefSpace := binaryTree.GetSpaceStr(prefSpaceNum)
				//fmt.Print(prefSpace,numberStr)
				max := j * 2 + 1
				//fmt.Println(max)
				min := max - 1
				spaceNum := (recordLocation[i+1][max] - recordLocation[i+1][min]) / 2
				if j > 0 {
					spaceNum = spaceNum  + 3
				}
				//fmt.Println(recordLocation[i+1][max] )
				//os.Exit(11)
				fmt.Print(binaryTree.GetSpaceStr(spaceNum),numberStr)
			}

		}
		if i == len(deepList) - 1{
			recordLocation[i][0] = 0
			fmt.Println(recordLocation[i])
		}
		//fmt.Println(recordLocation[i])
		fmt.Println("")
	}
}
func (binaryTree *BinaryTree)GetSpaceStr(length int)string{
	str := ""
	for i:=0;i<length;i++{
		str += "-"
	}
	return str
}
//将数字 格式化等长
func(binaryTree *BinaryTree)FormatNumberAlign(number int)string{
	numberStr := strconv.Itoa(number)
	max := NODE_KEYWORD_PLACEHOLDER
	if len(numberStr) > max{
		binaryTree.makeError("len(numberStr) > 1")
		return ""
	}

	str := numberStr
	for i:=0;i<max - len(numberStr) ;i++{
		str += " "
	}

	return str
}

//完整二叉树
//满二叉树
//左旋转
func (binaryTree *BinaryTree)LeftRotate(){

}
//右旋转
func (binaryTree *BinaryTree)RightRotate(){

}


