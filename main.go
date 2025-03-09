package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"

	"github.com/sammanadh/sam_compression/pkg/ds"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the file to encode!")
	}

	file_path := os.Args[1]
	file, err := os.Open(file_path)

	if err != nil {
		fmt.Printf("File %s does not exist!\n", file_path)
		os.Exit(1)
	}

	charmap := make(map[byte]int)

	buffer := make([]byte, 8)

	for {
		n, err := file.Read(buffer)

		if n > 0 {
			for _, char := range buffer[:n] { // only looping through the processed bytes and ignoring null bytes
				charmap[char] += 1
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
	}

	// creating a priority queue using a simple slice
	var chars []byte = []byte{}
	var counts []int = []int{}

	for char, count := range charmap {
		counts = append(counts, count)
		chars = append(chars, char)
		l := len(counts) - 2
		for ; l >= 0 && counts[l] < count; l-- {
			swap(chars, l, l+1)
			swap(counts, l, l+1)
		}
	}

	// var maxHeap *ds.BinaryMaxHeap[byte] = ds.NewBinaryMaxHeap[byte]()
	var orderedNodes []*ds.Node[byte] = []*ds.Node[byte]{}
	for i, count := range counts {
		newNode := &ds.Node[byte]{Value: chars[i], Weight: count, Left: nil, Right: nil}
		orderedNodes = append(orderedNodes, newNode)
	}

	hoffmanTree := createHuffmanTree(orderedNodes)
	// printBinaryTree(hoffmanTree)

	prefixCodeTable := generatePrefixCodeTable(hoffmanTree)
	for k, v := range prefixCodeTable {
		fmt.Println(string(k), v)
	}
}

func swap[T any](s []T, idx1 int, idx2 int) {
	temp := s[idx1]
	s[idx1] = s[idx2]
	s[idx2] = temp
}

func createHuffmanTree[T any](nodes []*ds.Node[T]) *ds.Node[T] {
	nodesClone := slices.Clone(nodes)
	for len(nodesClone) > 1 {
		l := len(nodesClone) - 1
		combinedWeight := nodesClone[l].Weight + nodesClone[l-1].Weight
		var zero T // zero value for type "T"
		newSmallestNode := &ds.Node[T]{Value: zero, Weight: combinedWeight, Left: nodesClone[l], Right: nodesClone[l-1]}
		nodesClone = nodesClone[:l]

		smallestNodeIdx := l - 1
		nodesClone[smallestNodeIdx] = newSmallestNode

		// insertion sift
		for i := smallestNodeIdx; i > 0; i-- {
			if nodesClone[i-1].Weight < nodesClone[i].Weight {
				swap(nodesClone, i-1, i)
			} else {
				break
			}
		}
	}
	return nodesClone[0]
}

func generatePrefixCodeTable(ht *ds.Node[byte]) map[byte]string {
	preffixCodes := make(map[byte]string)
	var traverseTree func(*ds.Node[byte], string)
	var zero byte // zero value for type "byte"
	traverseTree = func(currNode *ds.Node[byte], currCode string) {
		if currNode == nil {
			return
		} else if currNode.Value != zero {
			preffixCodes[currNode.Value] = currCode
		} else {
			traverseTree(currNode.Left, currCode+"0")
			traverseTree(currNode.Right, currCode+"1")
		}
	}
	traverseTree(ht, "")
	return preffixCodes
}

func printBinaryTree[T any](root *ds.Node[T]) {
	height := getHeight(root)
	level := 1
	levelNodes := []*ds.Node[T]{root}
	strigifiedTree := ""
	for ; level < height; level++ {
		spaceBeforeAndAfter := int(math.Pow(2, float64(height-level)) - 1)
		spaceBetweenNodes := int(math.Pow(2, float64(height-level+1)) - 1)

		// add the string for this row in strigifiedTree
		for idx, node := range levelNodes {
			var strNodeWeight string
			if node == nil {
				strNodeWeight = " "
			} else {
				strNodeWeight = fmt.Sprintf("%d", node.Weight)
			}
			if idx == 0 {
				// strigifiedTree += fmt.Sprintf("%*s", spaceBeforeAndAfter, strNodeWeight)
				for i := 0; i < spaceBeforeAndAfter; i++ {
					strigifiedTree += "_"
				}
				strigifiedTree += fmt.Sprintf("%s", strNodeWeight)
			} else {
				// strigifiedTree += fmt.Sprintf("%*s", spaceBetweenNodes, strNodeWeight)
				for i := 0; i < spaceBetweenNodes; i++ {
					strigifiedTree += "_"
				}
				strigifiedTree += fmt.Sprintf("%s", strNodeWeight)
			}
		}
		// strigifiedTree += "\n"
		for i := 0; i < spaceBeforeAndAfter; i++ {
			strigifiedTree += "_"
		}
		strigifiedTree += "\n"

		newLevelNodes := []*ds.Node[T]{}
		// update levelNodes for the next row
		for _, node := range levelNodes {
			if node == nil {
				newLevelNodes = append(newLevelNodes, nil)
			} else {
				newLevelNodes = append(newLevelNodes, node.Left, node.Right)
			}
		}
		levelNodes = newLevelNodes
	}

	fmt.Printf(strigifiedTree)
}

func getHeight[T any](root *ds.Node[T]) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

type HuffmanTreeNode struct {
	weight int
	left   *HuffmanTreeNode
	right  *HuffmanTreeNode
}
