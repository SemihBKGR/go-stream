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

	fmt.Println("---------ForEach---------")

	stream.StreamOf(data).
		Filter(func(i int) bool {
			return i%2 != 0
		}).
		Peek(func(i int) {
			fmt.Printf("Peek-1: %d\n", i)
		}).
		Limit(10).
		Peek(func(i int) {
			fmt.Printf("Peek-2: %d\n", i)
		}).
		ForEach(func(i int) {
			fmt.Printf("ForEach: %d\n", i)
		})

	fmt.Println("---------AnyMatch---------")

	anyMatch := stream.StreamOf(data).
		Filter(func(i int) bool {
			return i%2 != 0
		}).
		Peek(func(i int) {
			fmt.Printf("Peek-1: %d\n", i)
		}).
		Limit(10).
		Peek(func(i int) {
			fmt.Printf("Peek-2: %d\n", i)
		}).
		AnyMatch(func(i int) bool {
			return i > 10
		})

	fmt.Printf("Result: %v\n", anyMatch)

	fmt.Println("---------FirstMatch---------")

	firstMatch := stream.StreamOf(data).
		Filter(func(i int) bool {
			return i%2 != 0
		}).
		Peek(func(i int) {
			fmt.Printf("Peek-1: %d\n", i)
		}).
		Limit(10).
		Peek(func(i int) {
			fmt.Printf("Peek-2: %d\n", i)
		}).
		FirstMatch(func(i int) bool {
			return i > 10
		})

	fmt.Printf("Result: %d\n", firstMatch)

}
