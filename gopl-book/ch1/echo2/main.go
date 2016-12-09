package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for index, arg := range os.Args[1:] {
		indexString := strconv.Itoa(index)
		fmt.Println(indexString + " " + arg)
	}
}
