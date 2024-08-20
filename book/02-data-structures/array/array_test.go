package array

import (
	"fmt"
	"testing"
)

func TestArrayCapacityAndLength(t *testing.T) {
	// arrayCapacityAndLength()
}

func arrayCapacityAndLength() {
	array := [5]int{}
	fmt.Printf("array := [5]int{} => %v\n", array)
	fmt.Println("array: capacity:", cap(array))
	fmt.Println("array: length:", len(array))

	slice := make([]int, 5)
	fmt.Printf("slice := make([]int, 5) => %v\n", slice)
	fmt.Println("slice: capacity:", cap(slice))
	fmt.Println("slice: length:", len(slice))
}
