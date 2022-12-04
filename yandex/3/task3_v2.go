package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent     int
	LeftChild  int
	RightChild int
}

type Tree struct {
	Data map[int]Node
	Root int
}

func NewTree(n int) Tree {
	data := make(map[int]Node, n)
	tree := Tree{
		Data: data,
		Root: 1,
	}
	tree.AddNodes(0, 1, n)

	return tree
}

func (t *Tree) AddNodes(parent, node int, nMax int) {

	lv := node * 2
	if lv > nMax {
		lv = 0
	} else {
		t.AddNodes(node, lv, nMax)
	}

	rv := node*2 + 1
	if rv > nMax {
		rv = 0
	} else {
		t.AddNodes(node, rv, nMax)
	}

	(*t).Data[node] = Node{
		Parent:     parent,
		LeftChild:  lv,
		RightChild: rv,
	}

}

func (t *Tree) ReplaceNode(vNode int) {
	v := (*t).Data[vNode]
	p := (*t).Data[v.Parent]
	pp := (*t).Data[p.Parent]

	//Если p - не был корнем, то меняем в pp потомка
	if p.Parent != 0 {
		switch v.Parent {
		//Если p был левым потомком pp, то v становится левым потомком pp
		case pp.LeftChild:
			{
				(*t).Data[p.Parent] = Node{
					Parent:     pp.Parent,
					LeftChild:  vNode,
					RightChild: pp.RightChild,
				}
			}
		//Если p был правым потомком pp, то v становится правым потомком pp
		case pp.RightChild:
			{
				(*t).Data[p.Parent] = Node{
					Parent:     pp.Parent,
					LeftChild:  pp.LeftChild,
					RightChild: vNode,
				}
			}
		}
	} else { //Если p был корнем, то v становится новым корнем
		(*t).Root = vNode
	}

	switch vNode {
	//Если v был левым потомком, то :
	case p.LeftChild:
		{
			// p
			(*t).Data[v.Parent] = Node{
				Parent:     vNode,        //родителем p становится v
				LeftChild:  v.LeftChild,  //левым потомком p становится левый потомок v
				RightChild: p.RightChild, //правый потомок не изменяется
			}

			//v
			(*t).Data[vNode] = Node{
				Parent:     p.Parent,     //родителем v становится pp
				LeftChild:  v.Parent,     //левым потомком v становится p
				RightChild: v.RightChild, //правый потомок не изменяется
			}

			//Если левый потомок v (vl) существует
			if v.LeftChild != 0 {

				//vl
				(*t).Data[v.LeftChild] = Node{
					Parent:     v.Parent,                          //предком vl становится p
					LeftChild:  (*t).Data[v.LeftChild].LeftChild,  //левый потомок не изменяется
					RightChild: (*t).Data[v.LeftChild].RightChild, //правый потомок не изменяется
				}
			}

		}
	case p.RightChild:
		{
			//p
			(*t).Data[v.Parent] = Node{
				Parent:     vNode,        //родителем p становится v
				LeftChild:  p.LeftChild,  //левый потомок не изменяется
				RightChild: v.RightChild, //правым потомком p становится правый потомок v
			}

			//v
			(*t).Data[vNode] = Node{
				Parent:     p.Parent,    //родителем v становится pp
				LeftChild:  v.LeftChild, //левый потомок не изменяется
				RightChild: v.Parent,    //правым потомком v становится p
			}

			//Если правый потомок v (vr) существует
			if v.RightChild != 0 {

				//vr
				(*t).Data[v.RightChild] = Node{
					Parent:     v.Parent,                           //предком vr становится p
					LeftChild:  (*t).Data[v.RightChild].LeftChild,  //левый потомок не изменяется
					RightChild: (*t).Data[v.RightChild].RightChild, //левый потомок не изменяется
				}
			}
		}
	}
}

func (t *Tree) PrintLVR(res *string) {
	t.printLVR(t.Root, res)
}

func (t *Tree) printLVR(root int, res *string) {
	if lc := (*t).Data[root]; lc.LeftChild != 0 {
		t.printLVR(lc.LeftChild, res)
	}

	*res = fmt.Sprintf("%s %s", *res, strconv.Itoa(root))
	//fmt.Printf(" %d", root)

	if rc := (*t).Data[root]; rc.RightChild != 0 {
		t.printLVR(rc.RightChild, res)
	}
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Scan()
	strArr := strings.Split(fileScanner.Text(), " ")

	N, err := strconv.Atoi(strArr[0])
	if err != nil {
		log.Fatal("Error: ", err)
	}

	tree := NewTree(N)

	Q, err := strconv.Atoi(strArr[1])
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fileScanner.Scan()
	repls := strings.Split(fileScanner.Text(), " ")

	for i := 0; i < Q; i++ {

		node, err := strconv.Atoi(repls[i])
		if err != nil {
			log.Fatal("Error: ", err)
		}

		tree.ReplaceNode(node)
	}

	var result string
	tree.PrintLVR(&result)

	output, err := os.OpenFile(`output.txt`, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer output.Close()

	err = output.Truncate(0)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	output.WriteString(strings.Trim(result, " "))

}
