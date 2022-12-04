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
	//Number int
	Parent     int
	LeftChild  int
	RightChild int
}

type Tree map[int]Node

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

	(*t)[node] = Node{
		Parent:     parent,
		LeftChild:  lv,
		RightChild: rv,
	}

}

func (t *Tree) FindRoot(v int) (root int) {
	if (*t)[v].Parent != 0 {
		v = t.FindRoot((*t)[v].Parent)
	}
	return v
}

func (t *Tree) ReplaceNode(vNode int) {
	v := (*t)[vNode]
	p := (*t)[v.Parent]
	pp := (*t)[p.Parent]

	// fmt.Printf("\nv=%d\n", v)
	// fmt.Printf("key_v_before: %d, node: %+v\n", v, node)
	// fmt.Printf("key_p_before: %d, node: %+v\n", node.Parent, parent)
	// fmt.Printf("pp_before: %d, node: %+v\n", parent.Parent, pp)

	if p.Parent != 0 {
		switch v.Parent {
		case pp.LeftChild:
			{
				(*t)[p.Parent] = Node{
					Parent:     pp.Parent,
					LeftChild:  vNode,
					RightChild: pp.RightChild,
				}
			}
		case pp.RightChild:
			{
				(*t)[p.Parent] = Node{
					Parent:     pp.Parent,
					LeftChild:  pp.LeftChild,
					RightChild: vNode,
				}
			}
		}
	}

	switch vNode {
	case p.LeftChild:
		{

			(*t)[v.Parent] = Node{
				Parent:     vNode,
				LeftChild:  v.LeftChild,
				RightChild: p.RightChild,
			}

			(*t)[vNode] = Node{
				Parent:     p.Parent,
				LeftChild:  v.Parent,
				RightChild: v.RightChild,
			}

			if v.LeftChild != 0 {
				(*t)[v.LeftChild] = Node{
					Parent:     v.Parent,
					LeftChild:  (*t)[v.LeftChild].LeftChild,
					RightChild: (*t)[v.LeftChild].RightChild,
				}
			}

		}
	case p.RightChild:
		{
			(*t)[v.Parent] = Node{
				Parent:     vNode,
				LeftChild:  p.LeftChild,
				RightChild: v.RightChild,
			}

			(*t)[vNode] = Node{
				Parent:     p.Parent,
				LeftChild:  v.LeftChild,
				RightChild: v.Parent,
			}

			if v.RightChild != 0 {
				(*t)[v.RightChild] = Node{
					Parent:     v.Parent,
					LeftChild:  (*t)[v.RightChild].LeftChild,
					RightChild: (*t)[v.RightChild].RightChild,
				}
			}

		}
	}

	// fmt.Printf("key_v: %d, node: %+v\n", v, (*t)[v])
	// fmt.Printf("key_p: %d, node: %+v\n", node.Parent, (*t)[node.Parent])
	// fmt.Printf("pp: %d, node: %+v\n", parent.Parent, (*t)[parent.Parent])
	// fmt.Printf("lChild: %d, node: %+v\n", node.LeftChild, (*t)[node.LeftChild])
	// fmt.Printf("rChild: %d, node: %+v\n", node.RightChild, (*t)[node.RightChild])
}

func (t *Tree) PrintLVR(root int, res *string) {
	if lc := (*t)[root]; lc.LeftChild != 0 {
		t.PrintLVR(lc.LeftChild, res)
	}

	*res = fmt.Sprintf("%s %s", *res, strconv.Itoa(root))
	//fmt.Printf(" %d", root)

	if rc := (*t)[root]; rc.RightChild != 0 {
		t.PrintLVR(rc.RightChild, res)
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

	//Q, err := strconv.Atoi(strArr[1])
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fileScanner.Scan()
	repls := strings.Split(fileScanner.Text(), " ")

	tree := make(Tree, N)
	tree.AddNodes(0, 1, N)

	//****************************************************************************************************************************************************************************************************************
	// for key, node := range tree {
	// 	fmt.Printf("key: %d, node: %+v\n", key, node)
	// }
	//****************************************************************************************************************************************************************************************************************

	//fmt.Printf("Replace request: %+v\n", repls)
	for _, replaceNode := range repls {

		node, err := strconv.Atoi(replaceNode)
		if err != nil {
			log.Fatal("Error: ", err)
		}

		tree.ReplaceNode(node)

		// if i >= Q {
		// 	break
		// }
	}

	//****************************************************************************************************************************************************************************************************************
	// for key, node := range tree {
	// 	fmt.Printf("key: %d, node: %+v\n", key, node)
	// }
	//****************************************************************************************************************************************************************************************************************

	root := tree.FindRoot(1)
	//fmt.Printf("root = %d\n", root)

	var result string
	tree.PrintLVR(root, &result)
	//fmt.Printf("res = %s\n", result)

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
