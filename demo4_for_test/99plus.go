package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int {
	return a + b
}

func multiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			var ret string
			if i*j < 10 && j != 1 {
				ret = " " + strconv.Itoa(i*j)
			} else {
				ret = strconv.Itoa(i * j)
			}

			fmt.Print(j, " * ", i, " = ", ret, "   ")
		}
		fmt.Print("\n")
	}
}

func main() {
	multiplicationTable()
}
