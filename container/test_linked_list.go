package container

import (
	"fmt"
	"os"
)

func TestLinkedList(){
	//TestLinkedListNewError()
	//TestLinkedListByOrderNone()
	TestLinkedListByOrder()
	//listSearchCondition := container.ListSearchCondition{
	//	Empty: 1,
	//	Offset: 8,
	//	Limit: 100,
	//}
	//TestLinkedListGetAllPrint(linkedList,listSearchCondition)

}
//测试 创建链表错误的case
func TestLinkedListNewError(){
	nodeMax := 99999
	debug := 1
	loop := true
	//order := container.ORDER_ASC
	//order := container.ORDER_DESC
	order := 999
	NewLinkedList(order,nodeMax,loop,debug)
}

func TestLinkedListByOrderNone(){
	prefix := "测试无序 "
	nodeMax := 100
	debug := 1
	loop := true
	order := ORDER_NONE
	linkedList := NewLinkedList(order,nodeMax,loop,debug)
	//先获取一下列表：看看是否空
	empty,nodeList  := linkedList.GetAllByFirst(ListSearchCondition{})
	if !empty {
		fmt.Println(prefix + " first GetAllByFirst not empty: ",empty,nodeList)
	}
	//开始测试节点插入
	fmt.Println(prefix + "start insert node...")
	//一共要插入多少个节点，最好是3的整数倍(头、尾、任意位置、指定关键字)，方便测试
	insertForEnd := 40
	insertBat := 4
	insertBatEvery :=  insertForEnd / insertBat
	//先测试头部插入
	for i:=0 ; i < insertBatEvery;i++{
		keyword := GetRandIntNum(insertForEnd)
		data := keyword
		linkedList.InsertNodeByFirst(keyword ,data)
	}
	TestLinkedListCheckLength(prefix,insertBatEvery,linkedList)
	//输出一下本次插入后的列表节点情况
	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})
	//尾部插入
	for i:=0 ; i < insertBatEvery ;i++{
		keyword := GetRandIntNum(insertForEnd)
		data := keyword
		linkedList.InsertNodeByEnd(keyword ,data)
	}
	TestLinkedListCheckLength(prefix,insertBatEvery * 2,linkedList)
	//输出一下本次插入后的列表节点情况
	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})
	//在某个位置插入
	for i:=0 ; i < insertBatEvery ;i++{
		keyword := GetRandIntNum(insertForEnd)
		data := keyword
		location := GetRandIntNum(insertBatEvery * 2 - 1)
		linkedList.InsertNodeByLocation(location,keyword,data)
	}
	TestLinkedListCheckLength(prefix,insertBatEvery * 3,linkedList)
	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})
	//指定关键字
	insertNodeByKewyordSuccessNum := 0
	for i:=0 ; i < insertBatEvery ;i++{
		keyword := GetRandIntNum(insertForEnd)
		data := keyword
		offerKeyword := GetRandIntNum(insertForEnd - 1)
		_,err := linkedList.InsertNodeByKeyword(offerKeyword,keyword,data)
		if err == nil{
			insertNodeByKewyordSuccessNum++
		}
	}
	fmt.Print("insertNodeByKewyordSuccessNum:",insertNodeByKewyordSuccessNum)
	nowLinkedListLen := insertBatEvery * 3 + insertNodeByKewyordSuccessNum
	TestLinkedListCheckLength(prefix,nowLinkedListLen,linkedList)
	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})

	//开始测试 节点删除
	fmt.Println(prefix +"start delete--============================")
	delNodeCnt := 0//统计一共删除了几个节点
	//尾部删除一个
	linkedList.DelEndNode()
	delNodeCnt++
	//头部删除一个
	linkedList.DelFirstNode()
	delNodeCnt++
	//固定位置删除N个
	delNodeCnt++;delNodeCnt++
	hasDeletedNode,_ := linkedList.DelOneNodeByLocation(DIRECTION_FIRST,4,2)
	fmt.Println("hasDeletedNode len:",len(hasDeletedNode))
	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})
	//获取 keyword 为重复的列表
	repeatKeywordList,_ := linkedList.NodeRepeatTotal()
	//随便从重复列表里拿出一个keyword ，用于下面删除操作
	var repeatKeywordOneNodeKeyword int
	for k,_ :=range repeatKeywordList{
		repeatKeywordOneNodeKeyword = k
	}
	fmt.Println("repeatKeywordList: len:",len(repeatKeywordList),repeatKeywordList)
	//根据关键字删除N个,这里测试删除一个keyword有重复的情况
	linkedList.DelNodeByKeyword(repeatKeywordOneNodeKeyword)
	delNodeCnt = delNodeCnt + repeatKeywordList[repeatKeywordOneNodeKeyword]
	nowLinkedListLen = nowLinkedListLen - delNodeCnt

	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})
	TestLinkedListCheckLength(prefix,nowLinkedListLen,linkedList)

	//开始查找测试
	fmt.Println("start search: now nowLinkedListLen : ",nowLinkedListLen , " linkedList.len:" ,linkedList.Len)

	//查找第一个位置的元素
	empty,node := linkedList.FindOneNodeByLocation(0)
	fmt.Println(prefix +"FindOneNodeByLocation first:",node.Keyword)
	//查找第最后一个位置的元素
	empty,node = linkedList.FindOneNodeByLocation(linkedList.Len -  1 )
	fmt.Println(prefix +"FindOneNodeByLocation end:",node.Keyword)

	_ ,list := linkedList.GetAllByFirst(ListSearchCondition{})
	for k,v :=range list{
		fmt.Println(k," ",v.Keyword)
	}

	var insertKeywords []int
	_ ,list = linkedList.GetAllByEnd(ListSearchCondition{})
	for k,v :=range list{
		insertKeywords = append(insertKeywords,v.Keyword)
		fmt.Println(k," ",v.Keyword)
	}

	randNum := GetRandIntNum(len(insertKeywords) - 1)
	searchKeyword := insertKeywords[randNum]
	empty,node = linkedList.FindOneNodeByKeyword(searchKeyword)
	fmt.Println(prefix +"test FindOneNodeByKeyword : ",empty,node.Keyword)

	//查找中位节点
	empty , node = linkedList.GetMiddleNode()
	fmt.Println(prefix ,"GetMiddleNode:", node.Keyword)

}

func TestLinkedListCheckLength(prefix string,length int,linkedList *LinkedList)bool{
	fmt.Println(prefix +"len:",linkedList.Len)
	if linkedList.Len != length {
		fmt.Println(prefix +"len err , hope:",length," real_leh:",linkedList.Len)
		os.Exit(111)
		return false
	}else{
		fmt.Println(prefix +"len ok ~")
		return true
	}
}

func TestLinkedListByOrder(){
	prefix := "测试有序 "
	nodeMax := 100
	debug := 1
	loop := true
	order := ORDER_DESC
	linkedList := NewLinkedList(order,nodeMax,loop,debug)
	//先获取一下列表：看看是否空
	empty,nodeList  := linkedList.GetAllByFirst(ListSearchCondition{})
	fmt.Println(prefix + " first GetAllByFirst : ",empty,nodeList)
	//开始测试节点插入
	fmt.Println(prefix + "start insert node...")

	insertForEnd := 30
	//因为是有序的，所以不管从哪个方向，哪个点，或者哪个关键字，其内部都得重新计算位置，这里只需要测试一种插入即可
	for i:=0 ; i < insertForEnd ;i++{
		keyword := GetRandIntNum(insertForEnd)
		data := keyword
		linkedList.InsertNode(DIRECTION_FIRST,-1,0,keyword ,data)
	}
	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})
	nowLinkedListLen := linkedList.Len
	//开始测试 节点删除
	fmt.Println(prefix +"start delete--============================")
	delNodeCnt := 0//统计一共删除了几个节点
	//尾部删除一个
	linkedList.DelEndNode()
	delNodeCnt++
	//头部删除一个
	linkedList.DelFirstNode()
	delNodeCnt++
	//固定位置删除N个
	delNodeCnt++;delNodeCnt++
	hasDeletedNode,_ := linkedList.DelOneNodeByLocation(DIRECTION_FIRST,4,2)
	fmt.Println("hasDeletedNode len:",len(hasDeletedNode))
	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})
	//获取 keyword 为重复的列表
	repeatKeywordList,_ := linkedList.NodeRepeatTotal()
	//随便从重复列表里拿出一个keyword ，用于下面删除操作
	var repeatKeywordOneNodeKeyword int
	for k,_ :=range repeatKeywordList{
		repeatKeywordOneNodeKeyword = k
	}
	fmt.Println("repeatKeywordList: len:",len(repeatKeywordList),repeatKeywordList)
	//根据关键字删除N个,这里测试删除一个keyword有重复的情况
	linkedList.DelNodeByKeyword(repeatKeywordOneNodeKeyword)
	fmt.Println("delNodeCnt:" ,delNodeCnt, " repeatKeywordOneNodeKeywordNum:",repeatKeywordList[repeatKeywordOneNodeKeyword])
	delNodeCnt = delNodeCnt + repeatKeywordList[repeatKeywordOneNodeKeyword]
	nowLinkedListLen =  nowLinkedListLen - delNodeCnt

	TestLinkedListGetAllPrint(linkedList,ListSearchCondition{})
	TestLinkedListCheckLength(prefix,nowLinkedListLen,linkedList)



}


func TestLinkedListGetAllPrint(linkedList *LinkedList,listSearchCondition  ListSearchCondition){
	empty,nodeList  := linkedList.GetAllByFirst(listSearchCondition)
	if empty {
		fmt.Print(" err: GetAllByFirst is empty")
		return
	}
	fmt.Println("len : ",linkedList.Len)
	for k,node :=  range nodeList{
		fmt.Println("test print one node ,",k," keyword :",node.Keyword)
	}
}

