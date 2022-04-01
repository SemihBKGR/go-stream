package main

import (
	"fmt"
	"go-stream"
)

func main() {

	data := make([]int, 100)

	for i := 0; i < 100; i++ {
		data[i] = i
	}

	stream.StreamOf(data).Filter(func(i int) bool {
		return i%2 != 0
	}).ForEach(func(i int) {
		fmt.Println(i)
	})

}
