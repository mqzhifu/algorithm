package container

import "fmt"

func TestBinaryTree(){
	binaryTree := NewBinaryTree(100,false,1)
	forEndNum := 10
	insertSuccessNum := 0
	for i:=0;i< forEndNum ;i++{
		keyword := GetRandIntNum(100)
		err := binaryTree.InsertOneNode(keyword,nil)
		if err == nil{
			insertSuccessNum++
		}
	}
	deep := binaryTree.GetDeep(1)
	fmt.Println("deep level :",deep)

	binaryTree.ShowPrintTree()
	leftNode := binaryTree.GetTheLeftNode()
	fmt.Println("left node :",leftNode.Keyword)

	rightNode := binaryTree.GetTheRightNode()
	fmt.Println("right node :",rightNode.Keyword)

	fmt.Println("root node :",binaryTree.GetRootNode().Keyword)

	list := binaryTree.EachByFirst()
	fmt.Println("EachByFirst:" )
	for k,v := range list{
		fmt.Println(k,v)
	}

	list = binaryTree.EachByMiddle()
	fmt.Println("EachByMiddle:" )
	for k,v := range list{
		fmt.Println(k,v)
	}

	list = binaryTree.EachByAfter()
	fmt.Println("EachByAfter:" )
	for k,v := range list{
		fmt.Println(k,v)
	}

}
