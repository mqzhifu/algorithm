package container

import (
	"errors"
	"fmt"
)

type Graph struct {
	Debug int
	Pool []*GraphNode
	Matrix [][]int
}

type GraphNode struct {
	Id int
	Name string
}
func NewGraph(Debug int)*Graph{
	graph := new(Graph)
	graph.Debug = Debug

	return graph
}
//计算两点间的 路径
func  (graph *Graph)SearchNodePath(node1Name string ,node2Name string)error{
	if node1Name == node2Name{
		return graph.makeError("node1Name == node2Name")
	}

	empty , node1 := graph.GetNodeByName(node1Name)
	if empty{
		err := graph.makeError("node1Name search empty~")
		return err
	}

	empty , node2 := graph.GetNodeByName(node2Name)
	if empty{
		err := graph.makeError("node2Name search empty~")
		return err
	}

	graph.SearchNodePathRecursive(node1,node2,[]*GraphNode{})
	//node1Index := node1.Id
	//for k,_ := range graph.Matrix{
	//	if graph.Matrix[node1Index][k] == 1{
	//
	//	}
	//}

	return nil
}

func  (graph *Graph)SearchNodePathRecursive(node1 *GraphNode,node2 *GraphNode,historyBack []*GraphNode){
	//graph.Print("SearchNodePathRecursive:",node1.Name)
	for _,v:=range historyBack{
		if v == node1{
			//graph.Print("repeat node.")
			return
		}
	}

	historyBack = append(historyBack,node1)

	if node1 == node2{
		str := ""
		for _,y := range historyBack{
			str = str+ y.Name + " "
		}
		graph.Print("node1 :",node1.Name , " node2 :", node2.Name, "search ok~history len:",len(historyBack)," " ,str)
		return
	}

	for k,_ := range graph.Matrix{
		if graph.Matrix[k][node1.Id] == 1{
			nextNode := graph.Pool[k]
			graph.SearchNodePathRecursive(nextNode,node2,historyBack)
		}
	}
}

func (graph *Graph)IsEmpty()bool{
	//if  heap.GetLength() <= 0 {
	//	return true
	//}
	return false
}

//输出信息，用于debug
func (graph *Graph) Print(a ...interface{}) (n int, err error) {
	if graph.Debug > 0{
		return fmt.Println(a)
	}
	return
}

func (graph *Graph)CreateGraphNode(name string)error{
	if graph.IsEmpty(){
		newGraphNode := GraphNode{Name: name,Id: graph.GetLength()}
		graph.Pool = append(graph.Pool,&newGraphNode)
		return nil
	}

	for _,v:=range graph.Pool{
		if name == v.Name{
			err := graph.makeError("name repeat")
			return err
		}
	}

	newGraphNode := GraphNode{Name: name,Id: graph.GetLength()}
	graph.Pool = append(graph.Pool,&newGraphNode)
	return nil
}

//创建一个error,统一管理
func (graph *Graph)makeError(msg string)error{
	graph.Print("[errors] " + msg)
	return errors.New(msg)
}

func (graph *Graph)GetLength()int{
	return len(graph.Pool)
}

//func (graph *Graph)InsertNode(){
//
//}

func (graph *Graph) InitMatrix(){
	graph.Matrix = make([][]int,graph.GetLength())
	for i:=0;i<graph.GetLength();i++{
		graph.Matrix[i] = make([]int,graph.GetLength())
	}
}
func (graph *Graph) ShowMatrix(){
	fmt.Print("  ")
	for i:=0;i<graph.GetLength();i++{
		fmt.Print(graph.Pool[i].Name," ")
	}
	fmt.Println()
	for i:=0;i<graph.GetLength();i++{
		fmt.Print(graph.Pool[i].Name,":")
		for j:=0;j<graph.GetLength();j++{
			fmt.Print(graph.Matrix[i][j] , " ")
		}
		fmt.Println()
	}
}

func (graph *Graph)GetNodeByName(name string)(empty bool,node *GraphNode){
	if graph.IsEmpty(){
		return true,node
	}

	for _,v:=range graph.Pool {
		if v.Name == name{
			return false ,v
		}
	}

	return true,node
}

func (graph *Graph) SetLinkBorder(node1Name string,linkNodes []string)error{
	if graph.GetLength() <= 1{
		return graph.makeError("len <= 1")
	}

	if len(linkNodes) <=0 {
		return graph.makeError("len(linkNodes) <=0")
	}


	empty , node1 := graph.GetNodeByName(node1Name)
	if empty{
		err := graph.makeError("node1Name search empty~")
		return err
	}

	for _,v:= range linkNodes{
		if v == node1Name{
			return graph.makeError("node1Name == node2Name")
		}

		empty , _ := graph.GetNodeByName(v)
		if empty{
			err := graph.makeError("node2Name : "+v+" search empty~")
			return err
		}

		//if graph.Matrix[node1.id][node2.id] == 1{
		//	return graph.makeError("has link ,do not repeat opt.")
		//}
	}

	for _,v:= range linkNodes{
		_,node2 := graph.GetNodeByName(v)

		graph.Matrix[node1.Id][node2.Id] = 1
		graph.Matrix[node2.Id][node1.Id] = 1

		graph.Print("set "+node1.Name + " linked "+ node2.Name)
	}

	return nil
}

func TestGraph(){
	graph := NewGraph(1)
	graph.CreateGraphNode("J")
	graph.CreateGraphNode("I")
	graph.CreateGraphNode("E")
	graph.CreateGraphNode("D")
	graph.CreateGraphNode("G")
	graph.CreateGraphNode("F")
	graph.CreateGraphNode("A")
	graph.CreateGraphNode("H")
	graph.CreateGraphNode("B")
	graph.CreateGraphNode("C")

	graph.InitMatrix()

	graph.SetLinkBorder("J",[]string{"I","E","D"})
	graph.SetLinkBorder("I",[]string{"H","G"})
	graph.SetLinkBorder("E",[]string{"J","D","F","A"})
	graph.SetLinkBorder("D",[]string{"J","E","C"})
	graph.SetLinkBorder("G",[]string{"I","F","H"})
	graph.SetLinkBorder("F",[]string{"E","G","A"})
	graph.SetLinkBorder("A",[]string{"E","F","B","C"})
	graph.SetLinkBorder("H",[]string{"I","G","B"})
	graph.SetLinkBorder("B",[]string{"A","H","C"})
	graph.SetLinkBorder("C",[]string{"D","A","B"})

	graph.ShowMatrix()
	graph.SearchNodePath ("J","I")
}