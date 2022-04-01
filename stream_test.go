package stream

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {

	data := make([]int, 100)

	for i := 0; i < 100; i++ {
		data[i] = i
	}

	fmt.Println("---------ForEach---------")

	StreamOfSlice(data).
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

	anyMatch := StreamOfSlice(data).
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

	fmt.Println("---------AllMatch---------")

	allMatch := StreamOfSlice(data).
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
		AllMatch(func(i int) bool {
			return i > 10
		})

	fmt.Printf("Result: %v\n", allMatch)

	fmt.Println("---------FirstMatch---------")

	firstMatch := StreamOfSlice(data).
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

	fmt.Println("---------Count---------")

	count := StreamOfSlice(data).
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
		Count()

	fmt.Printf("Result: %d\n", count)

	fmt.Println("---------FindFirst---------")

	first := StreamOfSlice(data).
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
		FindFirst()

	fmt.Printf("Result: %d\n", first)

	fmt.Println("---------Collect---------")

	slice := StreamOfSlice(data).
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
		Collect()

	fmt.Printf("Result: %v\n", slice)

}
