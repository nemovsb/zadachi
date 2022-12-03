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

func (t *Tree) ReplaceNode(v int) {
	node := (*t)[v]
	parent := (*t)[node.Parent]
	pp := (*t)[parent.Parent]

	fmt.Printf("\nv=%d\n", v)
	fmt.Printf("key_v_before: %d, node: %+v\n", v, node)
	fmt.Printf("key_p_before: %d, node: %+v\n", node.Parent, parent)
	fmt.Printf("pp_before: %d, node: %+v\n", parent.Parent, pp)

	switch node.Parent {
	case pp.LeftChild:
		{
			(*t)[parent.Parent] = Node{
				Parent:     pp.Parent,
				LeftChild:  v,
				RightChild: pp.RightChild,
			}
		}
	case pp.RightChild:
		{
			(*t)[parent.Parent] = Node{
				Parent:     pp.Parent,
				LeftChild:  pp.LeftChild,
				RightChild: v,
			}
		}
	}

	switch v {
	case parent.LeftChild:
		{

			(*t)[node.Parent] = Node{
				Parent:     v,
				LeftChild:  node.LeftChild,
				RightChild: parent.RightChild,
			}

			(*t)[v] = Node{
				Parent:     parent.Parent,
				LeftChild:  node.Parent,
				RightChild: node.RightChild,
			}

			(*t)[node.LeftChild] = Node{
				Parent:     node.Parent,
				LeftChild:  (*t)[node.LeftChild].LeftChild,
				RightChild: (*t)[node.LeftChild].RightChild,
			}

		}
	case parent.RightChild:
		{
			(*t)[node.Parent] = Node{
				Parent:     node.Parent,
				LeftChild:  parent.LeftChild,
				RightChild: node.RightChild,
			}

			(*t)[v] = Node{
				Parent:     parent.Parent,
				LeftChild:  node.LeftChild,
				RightChild: node.Parent,
			}

			(*t)[node.RightChild] = Node{
				Parent:     v,
				LeftChild:  (*t)[node.RightChild].LeftChild,
				RightChild: (*t)[node.RightChild].RightChild,
			}

		}
	}

	// (*t)[node.LeftChild] = Node{
	// 	Parent:     v,
	// 	LeftChild:  (*t)[node.LeftChild].LeftChild,
	// 	RightChild: (*t)[node.LeftChild].RightChild,
	// }

	// (*t)[node.RightChild] = Node{
	// 	Parent:     v,
	// 	LeftChild:  (*t)[node.RightChild].LeftChild,
	// 	RightChild: (*t)[node.RightChild].RightChild,
	// }

	fmt.Printf("key_v: %d, node: %+v\n", v, (*t)[v])
	fmt.Printf("key_p: %d, node: %+v\n", node.Parent, (*t)[node.Parent])
	fmt.Printf("pp: %d, node: %+v\n", parent.Parent, (*t)[parent.Parent])
	fmt.Printf("lChild: %d, node: %+v\n", node.LeftChild, (*t)[node.LeftChild])
	fmt.Printf("rChild: %d, node: %+v\n", node.RightChild, (*t)[node.RightChild])
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
	for key, node := range tree {
		fmt.Printf("key: %d, node: %+v\n", key, node)
	}
	//****************************************************************************************************************************************************************************************************************

	fmt.Printf("Replace request: %+v\n", repls)
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
	for key, node := range tree {
		fmt.Printf("key: %d, node: %+v\n", key, node)
	}
	//****************************************************************************************************************************************************************************************************************

	// output, err := os.OpenFile(`output.txt`, os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }
	// defer output.Close()

	// err = output.Truncate(0)
	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	// output.WriteString(strings.Trim(result, " "))

}
