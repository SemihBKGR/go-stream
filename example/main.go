package main

import (
	"fmt"
	"go-stream"
)

func main() {

	data := make([]int, 5)
	data[0] = 0
	data[1] = 1
	data[2] = 2
	data[3] = 3
	data[4] = 4
	s := stream.StreamOf(data)
	s.ForEach(func(i int) {
		fmt.Println(i)
	})
}
