package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	src, _ := inputReader.ReadString('\r')
	count := wordcount(src)
	fmt.Println(count)
}
