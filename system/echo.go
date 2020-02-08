package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var text = strings.Join(os.Args[1:], " ")
	fmt.Println(text)
}
