package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// get next int value
func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// main calc
func calc() int {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	max := nextInt()
	min := max

	for i := 1; i < n; i++ {
		val := nextInt()
		if val > max {
			max = val
		} else if val < min {
			min = val
		}
	}

	return max - min

}

// main
func main() {
	fmt.Println(calc())
}
