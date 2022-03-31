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

	stream.StreamOf(data).
		Filter(func(e *int) bool {
			return *e%2 != 0
		}).
		Peek(func(e *int) {
			fmt.Println(*e)
		}).
		Limit(5).
		ForEach(func(e *int) {
			fmt.Println(*e)
		})
}
