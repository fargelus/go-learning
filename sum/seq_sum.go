package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func seqSum(sequence string) int {
	nums := strings.Split(sequence, " ")
	var sum int
	for _, strNum := range nums {
		num, _ := strconv.Atoi(strNum)
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("Input numbers for calculating sum")

	var scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Println(seqSum(line))
	}
}
