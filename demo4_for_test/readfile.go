package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "atest.go"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", content)
	}
}
