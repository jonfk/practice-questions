package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
)

/*

BFS: use a queue

*/

type Graph struct {
	Nodes []*Node
}

type Node struct {
	Value    int
	Visited  bool
	Adjacent []*Node
}

func (a *Graph) AddNode(x *Node) {
	a.Nodes = append(a.Nodes, x)
}

func (a *Graph) GetNode(x int) *Node {
	for i := range a.Nodes {
		if a.Nodes[i].Value == x {
			return a.Nodes[i]
		}
	}
	return nil
}

func (n *Node) addAdj(x *Node) {
	n.Adjacent = append(n.Adjacent, x)
	x.Adjacent = append(x.Adjacent, n)
}

func main() {
	fmt.Printf("graph-base")
	graph := &Graph{}
	graph.AddNode(&Node{Value: 1})
	graph.AddNode(&Node{Value: 2})
	graph.AddNode(&Node{Value: 3})
	graph.AddNode(&Node{Value: 4})
	graph.AddNode(&Node{Value: 5})

	x1 := graph.GetNode(1)
	x2 := graph.GetNode(2)
	x3 := graph.GetNode(3)
	x4 := graph.GetNode(4)
	x5 := graph.GetNode(5)
	x1.addAdj(x4)
	x1.addAdj(x2)
	x2.addAdj(x3)
	x2.addAdj(x5)
	x3.addAdj(x5)

	spew.Dump(graph)

	//found := BFS(graph, 6)
	found := DFS(x1, 5)
	fmt.Println("found")
	spew.Dump(found)
}

type Queue struct {
	Value []*Node
}

func (a *Queue) Enqueue(x *Node) error {
	if a == nil {
		return fmt.Errorf("enqueuing on nil queue")
	}
	a.Value = append(a.Value, x)
	return nil
}
func (a *Queue) Dequeue() (*Node, error) {
	if a == nil || len(a.Value) < 1 {
		return nil, fmt.Errorf("Dequeuing on nil or empty queue")
	}
	var x *Node
	x, a.Value = a.Value[0], a.Value[1:]
	return x, nil
}

func (a *Queue) IsEmpty() bool {
	if len(a.Value) > 0 {
		return false
	}
	return true
}

func BFS(g *Graph, searching int) *Node {
	for i := range g.Nodes {
		g.Nodes[i].Visited = false
	}
	root := g.Nodes[0]
	queue := &Queue{}
	err := queue.Enqueue(root)
	if err != nil {
		log.Fatal(err)
	}
	for !queue.IsEmpty() {
		x, err := queue.Dequeue()
		if err != nil {
			log.Fatal(err)
		}
		x.Visited = true
		if x.Value == searching {
			return x
		}

		for i := range x.Adjacent {
			if !x.Adjacent[i].Visited {
				err = queue.Enqueue(x.Adjacent[i])
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return nil
}

// recursive
func DFS(root *Node, searching int) *Node {
	if root.Visited {
		return nil
	} else {
		if root.Value == searching {
			return root
		}
		root.Visited = true
		for i := range root.Adjacent {
			found := DFS(root.Adjacent[i], searching)
			if found != nil {
				return found
			}
		}
		return nil
	}
}

// stack-based
