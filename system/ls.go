package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printFiles(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {
	var dirs = os.Args[1:]
	switch len(dirs) {
	case 0:
		printFiles("./")
	case 1:
		printFiles(dirs[0])
	default:
		for i, dir := range dirs {
			if i > 0 {
				fmt.Println()
			}
			fmt.Println(dir + ":")
			printFiles(dir)
		}
	}
}
