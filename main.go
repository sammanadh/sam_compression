package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file_path := os.Args[1]
	file, err := os.Open(file_path)

	if err != nil {
		fmt.Printf("File %s does not exist!\n", file_path)
		os.Exit(1)
	}

	charmap := make(map[byte]int)

	buffer := make([]byte, 8)

	for {
		_, err := file.Read(buffer)

		for _, char := range buffer {
			charmap[char] += 1
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
	}

	fmt.Println(charmap[[]byte("X")[0]])
}
