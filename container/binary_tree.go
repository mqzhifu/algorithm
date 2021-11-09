package container

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

const (
	DIRECTION_LEFT = 1
	DIRECTION_RIGHT = 2

	NODE_KEYWORD_PLACEHOLDER = 3


	TREE_NODE_MAX = 100
	KEYWORD_NIL = 999//有些需求，某些空节点必须得有，但是keyword得给一个int 占位符

	FLAG_NORMAL = 1
	FLAG_BALANCE = 2
	FLAG_RED_BLACK = 3
)

type TreeNode struct{
	Parent 		*TreeNode
	Left 		*TreeNode
	Right 		*TreeNode
	Keyword 	int
	Data 		interface{}
	DeepDesc 	int //这个元素有个问题：1. 每次发生变化，连带着都得跟着改=》影响性能。先加上，方便测试。降序
	DeepAsc		int //同上，是升序
	LeftDeep 	int	//左节点深度
	RightDeep 	int	//右节点深度
}

type BinaryTree struct {
	RootNode *TreeNode
	Flag int
	NodeMax int
	Len 	int
	Debug 	int
}

func NewBinaryTree(nodeMax int ,flag int,debug int)*BinaryTree{
	binaryTree :=  new (BinaryTree)

	if nodeMax > TREE_NODE_MAX{
		nodeMax = TREE_NODE_MAX
	}

	binaryTree.NodeMax = nodeMax
	binaryTree.Flag = flag
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
func (binaryTree *BinaryTree)InsertOneNode(keyword int ,data interface{})(compare int ,err error){
	binaryTree.Print("InsertOneNode  keyword:",keyword)

	if binaryTree.NodeMax > TREE_NODE_MAX{
		msg := "NodeMax > "+strconv.Itoa(TREE_NODE_MAX)
		return compare,binaryTree.makeError(msg)
	}

	newNode := binaryTree.NewOneNode(keyword,data)
	if binaryTree.IsEmpty(){
		//newNode.Deep = 1
		binaryTree.RootNode = newNode
		binaryTree.Len = 1
		return compare,nil
	}

	node := binaryTree.GetRootNode()
	searchNode,direction,compare,err := binaryTree.InsertOneNodeRecursionCompare(node,newNode,0,nil,0)
	if err != nil{
		return compare,err
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

	binaryTree.Print("compare times:",compare , " downNode:",searchNode.Keyword)

	binaryTree.Len++
	if binaryTree.Flag == FLAG_BALANCE{
		binaryTree.CheckInsertBalance(newNode)
	}

	return compare,nil
}
//插入节点时，递归查找新元素应该插入到哪个元素的左右
//node:当前比对节点
//insertNode:要插入的新节点
//direction:算是个上下文保留值，记录 当前节点是从上个节点的 左/右 方向过来的，因为最后一次循环，肯定节点是nil
//parentNode:算是个上下文保留值，记录 当前节点的父节点，因为最后一次循环，肯定节点是nil
//compare:比较次数
func (binaryTree *BinaryTree) InsertOneNodeRecursionCompare(node *TreeNode,insertNode *TreeNode ,direction int,parentNode  *TreeNode,compare int)(downNode *TreeNode , nodeDirection int ,compareTimes int,err error){
	if node == nil{
		return parentNode,direction,compare,nil
	}
	compare++
	if insertNode.Keyword < node.Keyword{
		return binaryTree.InsertOneNodeRecursionCompare(node.Left,insertNode,DIRECTION_LEFT,node,compare)
	}else if insertNode.Keyword > node.Keyword {
		return binaryTree.InsertOneNodeRecursionCompare(node.Right,insertNode,DIRECTION_RIGHT,node,compare)
	}else{
		msg := "NodeKeyword: not allow repeat ."
		return downNode,direction,compare,binaryTree.makeError(msg)
	}
}
//删除一个节点，根据keyword
func (binaryTree *BinaryTree)DelOneByKeyword(keyword int)(empty bool){
	binaryTree.Print("DelOneByKeyword keyword:",keyword)
	if binaryTree.IsEmpty(){
		return true
	}
	searchNode,empty ,_ := binaryTree.FindOneByKeyword(keyword)
	if empty {
		return true
	}
	//是否为  页子节点，即：无左右节点
	isLeaf := false
	if searchNode.Left == nil && searchNode.Right == nil{
		isLeaf = true
	}
	//是否为：根节点
	isRooNode := false
	if searchNode == binaryTree.GetRootNode(){
		isRooNode = true
	}
	//各种情况的不同处理
	if isRooNode && isLeaf {//即是根节点也是页子节点，即：该树只有一个节点了...
		binaryTree.RootNode = nil
	}else if searchNode.Left != nil && searchNode.Right != nil {//左节点也有值，右节点也有值
		//查找右节点的：最左节点
		rightTheMinNode := binaryTree.GetTheLeftNodeByNode(searchNode.Right)
		binaryTree.Print("GetTheLeftNodeByNode:",rightTheMinNode.Keyword , " p :",searchNode.Right.Parent.Keyword)
		//先把这个节点与原父节点的 关联 切断
		//rightTheMinNodeFromParentDir,_ := binaryTree.GetNodeFromParentDir(rightTheMinNode)
		//if rightTheMinNodeFromParentDir == "left"{
		//	rightTheMinNode.Parent.Left = nil
		//}else{
		//	rightTheMinNode.Parent.Right = nil
		//}
		if rightTheMinNode == searchNode.Right{
			rightTheMinNode.Parent = searchNode.Parent
			searchNode.Left.Parent = rightTheMinNode
			rightTheMinNode.Left = searchNode.Left

		}else{
			searchNode.Left.Parent = rightTheMinNode//将左节点的父节点 更新成 最小节点
			rightTheMinNode.Left = searchNode.Left//将最小节点的左节点 更新成 删除节点的左节点
			//rightTheMinNode.Parent = searchNode.Parent//把删除节点的父节点与 最小节点  连起来

			if rightTheMinNode.Right != nil{
				rightTheMinNode.Parent.Left = rightTheMinNode.Right
				rightTheMinNode.Right.Parent = rightTheMinNode.Parent
			}else{
				rightTheMinNode.Parent.Left = nil
			}

			binaryTree.Print("searchNode.Right:",searchNode.Right.Keyword)
			rightTheMinNode.Right = searchNode.Right
			rightTheMinNode.Right.Parent = rightTheMinNode
		}
		//更新 最小节点的 右节点
		//if searchNode.Right == nil || searchNode.Right.Right == nil{
		//	rightTheMinNode.Right = nil
		//}else{
		//	rightTheMinNode.Right = searchNode.Right.Right
		//	rightTheMinNode.Right.Parent = rightTheMinNode
		//}
		//binaryTree.Print( searchNode.Right.Right.Keyword)
		if searchNode == binaryTree.GetRootNode() {
			binaryTree.RootNode = rightTheMinNode
		}else{
			//把删除节点的父节点与 最小节点  连起来
			fromParentDir,_ := binaryTree.GetNodeFromParentDir(searchNode)
			binaryTree.Print("fromParentDir:",fromParentDir)
			if fromParentDir == "left"{
				searchNode.Parent.Left = rightTheMinNode
			}else{
				searchNode.Parent.Right = rightTheMinNode
			}
		}

		//binaryTree.Print("rightTheMinNode.Parent.Keyword:",rightTheMinNode.Parent.Keyword , " rightTheMinNode.left.Keyword:",rightTheMinNode.Left.Keyword, " rightTheMinNode.Right.Keyword:",rightTheMinNode.Right.Keyword, " rightTheMinNode.Parent.Left.Keyword:",rightTheMinNode.Parent.Left.Keyword)
	}else if isLeaf{//是页子节点，但不是根节点
		fromParentDir,_ := binaryTree.GetNodeFromParentDir(searchNode)
		if fromParentDir == "left"{
			searchNode.Parent.Left = nil
		}else{
			searchNode.Parent.Right = nil
		}
	}else{
		//左节点存在，右节点为空
		if searchNode.Left != nil && searchNode.Left == nil{
			if !isRooNode{
				fromParentDir,_ := binaryTree.GetNodeFromParentDir(searchNode)
				//更新：父节点的子节点
				if fromParentDir == "left"{
					searchNode.Parent.Left = searchNode.Left
				}else{
					searchNode.Parent.Right = searchNode.Left
				}
			}

			searchNode.Left.Parent = searchNode.Parent//把左节点挂到删除节点的父节点
		}else  if searchNode.Left != nil && searchNode.Left == nil{//左节点为空，右节点存在
			if !isRooNode{
				fromParentDir,_ := binaryTree.GetNodeFromParentDir(searchNode)
				//更新：父节点的子节点
				if fromParentDir == "left"{
					searchNode.Parent.Left = searchNode.Right
				}else{
					searchNode.Parent.Right = searchNode.Right
				}
			}
			searchNode.Right.Parent = searchNode.Parent//把右节点挂到删除节点的父节点
		}else{
			binaryTree.makeError("del err...")
		}
	}

	binaryTree.Len--

	return false
}
func (binaryTree *BinaryTree)GetNodeFromParentDir(node *TreeNode)(string,error){
	if node.Parent == nil{
		return "",binaryTree.makeError("no parent")
	}

	parent := node.Parent
	if parent.Left == node {
		return "left",nil
	}else if parent.Right == node{
		return "right",nil
	}else{
		return "",binaryTree.makeError("GetNodeFromParentDir no search keyword:"+ strconv.Itoa( node.Keyword))
	}
}
//递归查找一个节点
func (binaryTree *BinaryTree)FindOneByKeywordRecursionCompare(keyword int ,node *TreeNode,compare int)(downNode *TreeNode,empty bool,compareTimes int){
	if node.Keyword == keyword{
		return node,false,compareTimes
	}

	if node == nil{
		return node,true,compareTimes
	}
	compareTimes++
	if keyword > node.Keyword {
		return binaryTree.FindOneByKeywordRecursionCompare(keyword,node.Right,compareTimes)
	}else if keyword < node.Keyword {
		return binaryTree.FindOneByKeywordRecursionCompare(keyword,node.Left,compareTimes)
	}else{//这里是出现了 相等的情况，按说insert时做了限制，但为防万一，还是加条输出吧
		binaryTree.makeError("FindOneByKeywordRecursionCompare in case :else...")
	}
	return downNode,true,compareTimes
}

func (binaryTree *BinaryTree)FindOneByKeyword(keyword int )(downNode *TreeNode,empty bool,compare int){
	binaryTree.Print("FindOneByKeyword:",keyword)
	if binaryTree.IsEmpty(){
		return downNode,true,compare
	}
	node := binaryTree.GetRootNode()
	searchNode,empty,compare := binaryTree.FindOneByKeywordRecursionCompare(keyword,node,0)
	binaryTree.Print("FindOneByKeyword compare times:",compare)
	return searchNode,empty,compare
}
//层级遍历-广度优先
//nodeNilFill:空节点的元素，是否需要填充
func (binaryTree *BinaryTree)EachDeepByBreadthFirst(nodeNilFill bool)(empty bool,finalNode map[int][]*TreeNode){
	//保存最终结果
	nodeContainer := make(map[int][]*ListNode)

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
			isEmpty , queueNode := list.Pop(0)
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
				if treeNode.Left != nil{
					list.Push(treeNode.Left.Keyword,treeNode.Left)
				}
				if treeNode.Right != nil{
					list.Push(treeNode.Right.Keyword,treeNode.Right)
				}
			}
		}
	}
	finalNode = make(map[int][]*TreeNode)
	for k,nodeListRowArr:= range nodeContainer{
		var finalNodeListArr []*TreeNode
		for _,nodeList :=range nodeListRowArr{
			treeNode ,ok := nodeList.Data.(*TreeNode)
			if !ok {
				treeNode = nil
			}
			finalNodeListArr = append(finalNodeListArr,treeNode)
		}
		finalNode[k] = finalNodeListArr
	}

	for i,v:=range finalNode{
		for _,v2:=range v{
			//fmt.Print(k,v,k2)
			if v2 == nil{
				//binaryTree.Print()
			}else{
				parent := 0
				if v2.Parent != nil{
					parent =v2.Parent.Keyword
				}
				fmt.Println(i,v2.Keyword , " parent:",parent)
			}

		}

	}

	return false,finalNode
}
//层级遍历-深度优先
func (binaryTree *BinaryTree)EachDeepByDeepFirst(node  *TreeNode)int{
	if node == nil{
		return 0
	}
	left := binaryTree.EachDeepByDeepFirst(node.Left)
	right := binaryTree.EachDeepByDeepFirst(node.Right)
	node.LeftDeep = left
	node.RightDeep = right

	result := 0
	if left > right{
		result = left + 1
	}else{
		result = right + 1
	}

	node.DeepDesc = result

	return result
}
//获取当前树的深度/高度
func (binaryTree *BinaryTree)GetDeep(flag int)int{
	if binaryTree.IsEmpty(){
		return 0
	}
	var nodeList map[int][]*TreeNode
	if flag == 1{
		_,nodeList = binaryTree.EachDeepByBreadthFirst(false)
		return len(nodeList)
	}else{
		deep := binaryTree.EachDeepByDeepFirst(binaryTree.GetRootNode())
		return deep
	}
}
//递归根据方向：遍历树
func (binaryTree *BinaryTree)EachByOrder( node *TreeNode , list *[]*TreeNode , order int){
	if binaryTree.IsEmpty(){
		return
	}

	if node == nil{
		return
	}
	if order == 1{
		binaryTree.EachByOrder(node.Left,list,order)
		InsertIntArray(list,node)
		//binaryTree.Print(node.Keyword)
		binaryTree.EachByOrder(node.Right,list,order)
	}else if order == 2{
		InsertIntArray(list,node)
		//binaryTree.Print(node.Keyword)
		binaryTree.EachByOrder(node.Left,list,order)
		binaryTree.EachByOrder(node.Right,list,order)
	}else if order == 3{
		binaryTree.EachByOrder(node.Left,list,order)
		binaryTree.EachByOrder(node.Right,list,order)
		InsertIntArray(list,node)
		//binaryTree.Print(node.Keyword)
	}else{
		binaryTree.makeError("EachByOrder case ")
	}
}
//先序遍历
func (binaryTree *BinaryTree)EachByFirst(  )[]*TreeNode{
	list :=GetNewIntArrayAndFillEmpty(binaryTree.GetLength())
	node := binaryTree.GetRootNode()
	binaryTree.EachByOrder(node,&list,1)

	return list
}
//中序遍历
func (binaryTree *BinaryTree)EachByMiddle(  )[]*TreeNode{
	list :=GetNewIntArrayAndFillEmpty(binaryTree.GetLength())
	node := binaryTree.GetRootNode()
	binaryTree.EachByOrder(node,&list,2)

	return list
}
//后序遍历
func (binaryTree *BinaryTree)EachByAfter(  )[]*TreeNode{
	list :=GetNewIntArrayAndFillEmpty(binaryTree.GetLength())
	node := binaryTree.GetRootNode()
	binaryTree.EachByOrder(node,&list,3)

	return list
}
//获取树的最左/最右极值节点
func (binaryTree *BinaryTree)GetLimitNodeRecursionCompare(node *TreeNode,dir int )*TreeNode{
	if binaryTree.IsEmpty(){
		return nil
	}
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
//获取某个节点的~最左侧的节点（最小值）
func (binaryTree *BinaryTree)GetTheLeftNodeByNode(node *TreeNode )*TreeNode{
	return binaryTree.GetLimitNodeRecursionCompare(node,DIRECTION_LEFT)
}
//获取某个节点的~最右侧的节点（最大值）
func (binaryTree *BinaryTree)GetTheRightNodeByNode( node *TreeNode)*TreeNode{
	return binaryTree.GetLimitNodeRecursionCompare(binaryTree.RootNode,DIRECTION_RIGHT)
}
//检查一棵树是否是平衡的
func (binaryTree *BinaryTree)CheckInsertBalance(node *TreeNode){
	binaryTree.Print("CheckInsertBalance newNode keyword:",node.Keyword)
	//先做深度遍历层级，并给每个节点打上层级标识
	binaryTree.GetDeep(2)


	searchNode := node
	searchTimes := 0
	searchRs := false
	distance := 0
	for{
		distance = searchNode.LeftDeep - searchNode.RightDeep
		binaryTree.Print("distance:",distance)
		if math.Abs(float64(distance)) > 1{
			searchRs = true
			//minSubRootNode := searchNode
			break
		}
		if searchNode == binaryTree.GetRootNode() {
			break
		}

		searchNode = searchNode.Parent
		searchTimes++
	}

	if searchRs{
		binaryTree.Print(" start balance , distance:",distance , " searchNode keyword :",searchNode.Keyword)
		if distance > 1 && node.Keyword < searchNode.Left.Keyword{
			//当前node右转一次
			binaryTree.RightRotate(searchNode,"left")
		}else if distance > 1 && node.Keyword > searchNode.Left.Keyword{
			binaryTree.LeftRotate(searchNode.Left,"left")
			binaryTree.RightRotate(searchNode,"right")
		}else if distance < -1 && node.Keyword > searchNode.Right.Keyword{
			binaryTree.LeftRotate(searchNode,"right")
		}else if distance < -1 && node.Keyword < searchNode.Right.Keyword{
			binaryTree.RightRotate(searchNode.Right,"right")
			binaryTree.LeftRotate(searchNode,"left")
		}
		//if distance > 0 {//证明左侧比右侧要高，主要调度左侧
		//	binaryTree.Print("case in left")
		//	if searchNode.Right == nil {
		//		binaryTree.Print("case in nil")
		//		if searchNode.Left != nil && searchNode.Left.Left != nil{
		//			//当前node右转一次
		//			binaryTree.RightRotate(searchNode,"left")
		//		}else{
		//			////当前node右转一次
		//			//binaryTree.RightRotate(searchNode)
		//			////再右旋转一次
		//			//binaryTree.LeftRotate(searchNode)
		//		}
		//	}else {
		//		binaryTree.Print("case in not nil")
		//		leftDistance := searchNode.Left.LeftDeep - searchNode.Left.RightDeep
		//		if leftDistance > 0 {
		//			//当前node右转一次
		//			binaryTree.RightRotate(searchNode,"left")
		//		}else{
		//
		//		}
		//	}
		//}else{//证明右侧比比侧要高，主要调度右侧
		//	binaryTree.Print("case in right")
		//	if searchNode.Left == nil {
		//		binaryTree.Print("case in nil")
		//		if searchNode.Right != nil && searchNode.Right.Right != nil{
		//			//当前node右转一次
		//			binaryTree.LeftRotate(searchNode,"right")
		//		}else{
		//			binaryTree.ShowPrintTree()
		//			//当前node右转一次
		//			binaryTree.RightRotate(searchNode.Right,"right")
		//			binaryTree.ShowPrintTree()
		//			//os.Exit(44)
		//			//再右旋转一次
		//			binaryTree.LeftRotate(searchNode,"left")
		//
		//		}
		//	}else {
		//		binaryTree.Print("case in not nil")
		//		leftDistance := searchNode.Right.LeftDeep - searchNode.Right.RightDeep
		//		if leftDistance < 0 {
		//			//当前node右转一次
		//			binaryTree.LeftRotate(searchNode,"right")
		//		}else{
		//
		//		}
		//	}
		//}
	}else{
		binaryTree.Print(" has balance ...no need ")
	}
	binaryTree.Print("balance finish.")

}
//打印树形状
func (binaryTree *BinaryTree)ShowPrintTree()[][]string{
	nodeMiddleSpaceNum := 5//最后一行，两个节点间距离（空格数）
	numberPlaceHoder := 3 //输出一个数字，占多少位
	//先获取，树的每一层的节点列表
	_,deepList := binaryTree.EachDeepByBreadthFirst(true)
	//deepList中的行号是从1开始的，不是0
	levelCnt := len(deepList) - 1//最后一行：用<空,nil>做占位符，所以不用输出了

	recordLocation := make(map[int]map[int]int)
	//计算：每个元素的位置（元素之间的空格数），从下到上，最后一行算完后，之上的每一行由下面一行计算而来，所以是倒序的
	for i:=levelCnt ;i >= 1;i--{
		//记录：处理一行的一组的节点的位置（处理到第几个节点了）
		nodeLocation := 0
		//初始化 map
		recordLocation[i] = make(map[int]int)
		for j:=0;j<len(deepList[i]);j++ {
			if i == levelCnt{//最后一行不需要添加首空格了，同时也是最先计算的
				if j == 0 {
					recordLocation[i][0] = 0
				}
				//每两个子节点有一个父节点，两个子节点之间的距离是固定的
				if j % 2 == 0 {//偶数代表：一对子节点的：第一个节点的位置
					//当前处理过的节点位置 + 数字占位符数 + 一对（两个子节点）间的距离
					nodeLocation = nodeLocation + numberPlaceHoder + nodeMiddleSpaceNum
				}else{
					nodeLocation = nodeLocation + 3
				}
				recordLocation[i][j+1] = nodeLocation
				//binaryTree.Print(recordLocation)
			}else{//偶数代表：一对子节点的：第二个节点的位置
				arrMaxIndex := j * 2 + 1
				arrMinIndex := arrMaxIndex - 1
				dis := (recordLocation[i+1][arrMaxIndex] - recordLocation[i+1][arrMinIndex]) / 2
				spaceNum := recordLocation[i+1][arrMaxIndex] - dis
				recordLocation[i][j] = spaceNum
			}
		}
		//if i ==  levelCnt - 1{
		//	binaryTree.Print(recordLocation[i])
		//	recordLocation[i][0] = 0
		//}
		//fmt.Println("")
	}

	//binaryTree.Print(recordLocation[2])

	maxIndex := int(math.Pow(2 ,float64(levelCnt  - 1) ))
	//binaryTree.Print("maxIndex:",maxIndex)
	matrixRowMax := recordLocation[levelCnt][ maxIndex]

	//fmt.Println("matrixRowMax:",matrixRowMax , " maxIndex:", maxIndex)
	//创建一个二维 矩阵，并填充0
	matrix := make([][]string,levelCnt)
	for i:=0;i<levelCnt;i++{
		matrix[i] = make([]string,matrixRowMax)
	}
	//binaryTree.Print("aaaa",matrixRowMax)
	for i:=0;i<levelCnt ;i++{
		fmt.Print(i," ")
		jumpSpackPrintNum := 0
		for j:=0;j < matrixRowMax;j++ {
			if jumpSpackPrintNum > 0 {
				jumpSpackPrintNum--
				continue
			}
			//binaryTree.Print(i," ",j)
			f := false
			location := 0
			for k,v :=range recordLocation[i+1]{
				if v == j{
					f = true
					location = k
					//fmt.Print(location)
					break
				}
			}
			if f {
				jumpSpackPrintNum = numberPlaceHoder - 1
				if deepList[i+1][location] == nil{
					fmt.Print( "nil" )
				}else{
					value :=  deepList[i+1][location].Keyword
					if value == KEYWORD_NIL{
						fmt.Print( "nil" )
						matrix[i][j] = "nil"
					}else{
						fmt.Print( binaryTree.FormatNumberAlign(value) )
						matrix[i][j] = strconv.Itoa(value)
					}
				}

			}else{
				fmt.Print("-")
			}
		}
		//binaryTree.Print(i,matrixRowMax,levelCnt)
		//最后一个组，会少一个值
		//if i == levelCnt - 1 {
		//	ll := len( deepList[i+1])
		//	if ll >= 1{
		//		if deepList[i+1][ll - 1] != nil{
		//			fmt.Print(deepList[i+1][ll - 1].Keyword)
		//		}else{
		//			fmt.Print("nil")
		//		}
		//	}
		//}
		fmt.Println( )
	}
	return matrix
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

//传递值时，得用到 引用传参，能保存每次计算后的结果
//又因为是递归，并且是数组，各种地址参数重复改变一个指针的值
//所以这里，每次对数据的操作都是指针操作，得特殊处理下
func InsertIntArray( list *[]*TreeNode,treeNode *TreeNode ){
	for i:=0;i<len((*list));i++{
		if (*list)[i] == nil{
			(*list)[i] = treeNode
			break
		}
	}
}
//动态创建一个数组，并把所有元素值：初始化为 KEYWORD_NIL
func GetNewIntArrayAndFillEmpty(len int)[]*TreeNode{
	list :=  []*TreeNode{}
	for i:=0;i<len;i++{
		list = append(list,nil)
	}
	return list
}

//完整二叉树
//满二叉树

//左旋转
func (binaryTree *BinaryTree)LeftRotate(topNode *TreeNode,fromParentDir string){
	binaryTree.Print("start LeftRotate topNode keyword:",topNode.Keyword, topNode,topNode.Left)


	if topNode.Parent == nil {
		binaryTree.Print("im root.")
		binaryTree.RootNode = topNode.Right
	}else{
		if fromParentDir == "left"{
			topNode.Parent.Left = topNode.Right
		}else{
			topNode.Parent.Right = topNode.Right
		}

	}

	oldLeftRightNode := topNode.Right.Left
	topNode.Right.Left = topNode
	topNode.Right.Parent = topNode.Parent
	topNode.Parent = topNode.Right
	if oldLeftRightNode == nil{
		topNode.Right = nil
	}else{
		topNode.Right = oldLeftRightNode
		oldLeftRightNode.Parent = topNode
	}
}
//右旋转
func (binaryTree *BinaryTree)RightRotate(topNode *TreeNode,fromParentDir string){

	binaryTree.Print("start RightRotate topNode keyword:",topNode.Keyword, topNode,topNode.Left)

	if topNode.Parent == nil {
		binaryTree.Print("im root.")
		binaryTree.RootNode = topNode.Left
	}else{
		if fromParentDir == "left"{
			topNode.Parent.Left = topNode.Left
		}else{
			topNode.Parent.Right = topNode.Left
		}
	}

	oldLeftRightNode := topNode.Left.Right
	topNode.Left.Right = topNode
	topNode.Left.Parent = topNode.Parent
	topNode.Parent = topNode.Left
	if oldLeftRightNode == nil{
		topNode.Left = nil
	}else{
		topNode.Left = oldLeftRightNode
		oldLeftRightNode.Parent = topNode
	}
	//binaryTree.Print(topNode.Parent.Keyword)
}

//动态生成N个空格的字符串
//func (binaryTree *BinaryTree)GetSpaceStr(length int)string{
//	str := ""
//	for i:=0;i<length;i++{
//		str += "-"
//	}
//	return str
//}