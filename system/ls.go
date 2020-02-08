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
	if len(dirs) == 0 {
		printFiles("./")
	} else {
		for _, dir := range dirs {
			fmt.Println(dir + ":")
			printFiles(dir)
		}
	}
}
