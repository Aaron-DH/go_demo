package main

import (
	"fmt"
)

func main() {
	var arrs = [5]int{1, 2, 3, 4, 5}
	for idx, value := range arrs {
		arrs[idx] = value * 2
		fmt.Printf("Array idx[%d] is %d!\n", idx, arrs[idx])
	}

	var arr1 = new([5]int)
	var arr2 [5]int
	var inta int = 3
	arr1[0] = inta
	arr2[0] = inta
	fmt.Printf("value: %d\n", arr1[0])
	fmt.Printf("value: %d\n", arr2[0])
}
