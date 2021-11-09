package container

import (
	"fmt"
	"os"
)

func TestBinaryTree(){
	TestBinaryBalance()
	os.Exit(11)

	binaryTree := NewBinaryTree(100,FLAG_NORMAL,1)
	forEndNum := 10
	insertSuccessNum := 0
	for i:=0;i< forEndNum ;i++{
		keyword := GetRandIntNum(100)
		_,err := binaryTree.InsertOneNode(keyword,nil)
		if err == nil{
			insertSuccessNum++
		}
	}
	breadthFirst := binaryTree.GetDeep(1)
	fmt.Println("breadthFirst level :",breadthFirst)

	deepFirst := binaryTree.GetDeep(2)
	fmt.Println("deepFirst level :",deepFirst)

	binaryTree.ShowPrintTree()
	leftNode := binaryTree.GetTheLeftNode()
	fmt.Println("left node :",leftNode.Keyword)

	rightNode := binaryTree.GetTheRightNode()
	fmt.Println("right node :",rightNode.Keyword)

	fmt.Println("root node :",binaryTree.GetRootNode().Keyword)

	list := binaryTree.EachByFirst()
	fmt.Println("EachByFirst:" )
	for k,v := range list{
		fmt.Println(k,"keyword:",v.Keyword , " leftDeep:",v.LeftDeep, " rightDeep:",v.RightDeep , " deepDesc:",v.DeepDesc , " deepAsc:",v.DeepAsc)
	}

	list = binaryTree.EachByMiddle()
	fmt.Println("EachByMiddle:" )
	for k,v := range list{
		fmt.Println(k,"keyword:",v.Keyword , " leftDeep:",v.LeftDeep, " rightDeep:",v.RightDeep, " deepDesc:",v.DeepDesc, " deepAsc:",v.DeepAsc)
	}

	list = binaryTree.EachByAfter()
	fmt.Println("EachByAfter:" )
	for k,v := range list{
		fmt.Println(k,"keyword:",v.Keyword , " leftDeep:",v.LeftDeep, " rightDeep:",v.RightDeep, " deepDesc:",v.DeepDesc, " deepAsc:",v.DeepAsc)
	}

}

func TestBinaryBalance(){
	binaryTree := NewBinaryTree(100,FLAG_BALANCE,1)

	binaryTree.InsertOneNode(100,"10")
	binaryTree.ShowPrintTree()

	binaryTree.InsertOneNode(80,"7")
	binaryTree.ShowPrintTree()

	binaryTree.InsertOneNode(60,"5")
	binaryTree.ShowPrintTree()

	binaryTree.InsertOneNode(40,"4")
	binaryTree.ShowPrintTree()

	binaryTree.InsertOneNode(20,"3")
	binaryTree.ShowPrintTree()

	binaryTree.InsertOneNode(45,"3")
	binaryTree.ShowPrintTree()

	binaryTree.InsertOneNode(70,"3")
	binaryTree.ShowPrintTree()

	binaryTree.InsertOneNode(50,"3")
	binaryTree.ShowPrintTree()


	binaryTree.DelOneByKeyword(60)
	binaryTree.ShowPrintTree()

	binaryTree.DelOneByKeyword(40)
	binaryTree.ShowPrintTree()


	//
	//binaryTree.InsertOneNode(25,"2")
	//binaryTree.ShowPrintTree()
	//
	//binaryTree.InsertOneNode(22,"2")
	//binaryTree.ShowPrintTree()



	//binaryTree.InsertOneNode(10,"10")
	//binaryTree.ShowPrintTree()
	//
	//binaryTree.InsertOneNode(20,"20")
	//binaryTree.ShowPrintTree()
	//
	//binaryTree.InsertOneNode(30,"30")
	//binaryTree.ShowPrintTree()
	//
	//binaryTree.InsertOneNode(40,"40")
	//binaryTree.ShowPrintTree()
	//
	//binaryTree.InsertOneNode(50,"40")
	//binaryTree.ShowPrintTree()

}
