package main

import "fmt"

func GetYangHuiNextLint(lines []int) []int {
	var out []int
	out = append(out, 1)

	linelength := len(lines)
	if linelength == 0 {
		return out
	}
	for i := 0; i < linelength-1; i++ {
		out = append(out, lines[i]+lines[i+1])
	}

	out = append(out, 1)
	return out
}

func main() {
	nums := []int{}
	var i int
	for i = 0; i < 10; i++ {
		nums = GetYangHuiNextLint(nums)
		fmt.Println(nums)
	}
}
