package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var maxlist = []int{}
var minlist = []int{}

var RANK = 400

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
func calc() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	for i := 0; i < n; i++ {
		a := nextInt()
		rank := a / 400
		if rank > 7 {
			maxlist = append(maxlist, len(maxlist)+8)
		} else {
			if assoc(rank, maxlist) == -1 {
				maxlist = append(maxlist, rank)
				minlist = append(minlist, rank)
			}
		}
	}
}

// list finder
func assoc(i int, list []int) int {
	ans := -1

	for _, v := range list {
		if v == i {
			ans = v
		}
	}
	return ans
}

// main
func main() {
	calc()
	if len(minlist) == 0 {
		minlist = append(minlist, 1)
	}
	fmt.Print(len(minlist), len(maxlist))
}
