package main

import (
	"fmt"
	"io"
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
		for l >= 0 && counts[l] < count {
			swap(chars, l, l+1)
			swap(counts, l, l+1)
		}
	}

	var orderedNodes []*ds.Node[byte]
	for i, count := range counts {
		char := chars[i]
		orderedNodes = append(orderedNodes, ds.NewNode(char, count))
	}

	hoffmanTree := createHuffmanTree(orderedNodes)
	fmt.Println(hoffmanTree.Weight)
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
		newSmallestNode := &ds.Node[T]{Value: zero, Weight: combinedWeight, Left: nodesClone[l-1], Right: nodesClone[l]}
		nodesClone = nodesClone[:l]
		nodesClone[l-1] = newSmallestNode
	}
	return nodesClone[0]
}

type HuffmanTreeNode struct {
	weight int
	left   *HuffmanTreeNode
	right  *HuffmanTreeNode
}
