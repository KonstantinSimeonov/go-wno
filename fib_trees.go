package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Node struct {
	Value    int    `json:"value"`
	Children []Node `json:"children"`
}

func Treebonacci(n int) Node {
	fibs := make([]Node, n+1)

	fibs[0] = Node{0, []Node{}}

	if n >= 1 {
		fibs[1] = Node{1, []Node{}}
	}

	for i := 2; i <= n; i++ {
		n2 := fibs[i-2]
		n1 := fibs[i-1]
		fibs[i] = Node{n1.Value + n2.Value, []Node{n1, n2}}
	}

	return fibs[n]
}

func GenTree(r *rand.Rand, depth, max_children, d int) Node {
	node := Node{r.Int() % 1000, []Node{}}
	if d >= depth {
		return node
	}

	children_count := r.Int() % max_children
	for children_count > 0 {
		node.Children = append(node.Children, GenTree(r, depth, max_children, d+1))
		children_count--
	}

	return node
}

func ToJsonArray(n *Node) string {
	result := "[" + strconv.Itoa(n.Value) + ", ["
	for i, c := range n.Children {
		result += ToJsonArray(&c)
		if i < len(n.Children)-1 {
			result += ", "
		}
	}
	result += "]]"

	return result
}

func main() {
	depth := 5
	max_children := 5
	tree_type := "normal"
	format := "json"

	for i, arg := range os.Args {
		switch arg {
		case "--fib":
			tree_type = "fib"
		case "--depth":
			d, _ := strconv.Atoi(os.Args[i+1])
			depth = d
		case "--max-children":
			mc, _ := strconv.Atoi(os.Args[i+1])
			max_children = mc
		case "--format":
			format = os.Args[i+1]
		case "--help":
			fmt.Println(`Generate trees in json format
  --help                      Show ze help
  --fib                       Generate a fibonacci tree
  --depth [INT]               The depth of the tree
  --max-children [INT]        Limits how many children a non-fibonacci tree could have
  --format [json|json-array]  Output the tree in one of the following formats:
      json        type Node = { value: number, children: Node[] }
      json-array  type Node = [number, Node[]]`)
			return
		default:
		}
	}

	var tree Node
	if tree_type == "fib" {
		tree = Treebonacci(depth)
	} else {
		tree = GenTree(rand.New(rand.NewSource(42)), depth, max_children, 0)
	}

	switch format {
	case "json":
		json_tree, _ := json.Marshal(tree)
		fmt.Println(string(json_tree))
	case "json-array":
		fmt.Println(ToJsonArray(&tree))
	}
}
