package main

import (
	"strings"
)

func wordcount(in string) int {
	a := strings.Split(in, " ")
	return (len(a))
}
