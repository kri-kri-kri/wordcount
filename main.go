package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Trim(os.Args[1], " ")
	var list []string
	if len(input) != 0 {
		list = strings.Split(input, " ")
	}
	fmt.Println(len(list))
}
