package main

import (
	//"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Binode struct {
	Node1, Node2 *Binode
	Data         int
}

func sortedArrayToBST(a []int) *Binode {
	if len(a) == 0 {
		return nil
	}
	if len(a) < 2 {
		return &Binode{Node1: nil, Data: a[0], Node2: nil}
	}
	mid := len(a) / 2
	left := sortedArrayToBST(a[:mid])
	right := sortedArrayToBST(a[mid+1:])
	return &Binode{Node1: left, Data: a[mid], Node2: right}
}

func revConvert(root *Binode, left bool) *Binode {
	if root == nil {
		return nil
	} else if root.Node1 == nil && root.Node2 == nil {
		return root
	} else {
		nodeBefore := revConvert(root.Node1, true)
		if nodeBefore != nil {
			nodeBefore.Node2 = root
		}
		root.Node1 = nodeBefore
		if root.Node2 != nil {
			nodeAfter := revConvert(root.Node2, false)
			nodeAfter.Node1 = root
			root.Node2 = nodeAfter
		}

		if !left {
			result := root
			for result.Node1 != nil {
				result = result.Node1
			}
			return result
		} else {
			result := root
			for result.Node2 != nil {
				result = result.Node2
			}
			return result
		}
	}
}

func convBST2LL(root *Binode) *Binode {
	var left, right, result *Binode
	if root.Node1 != nil {
		left = revConvert(root.Node1, true)
	}

	if root.Node2 != nil {
		right = revConvert(root.Node2, false)
	}
	left.Node2 = root
	right.Node1 = root
	root.Node1 = left
	root.Node2 = right
	result = root
	for result.Node1 != nil {
		result = result.Node1
	}
	return result

}

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	testBst := sortedArrayToBST(test)

	result := convBST2LL(testBst)
	spew.Dump(result)
}
