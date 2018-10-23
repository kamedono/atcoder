package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func solve1() int {
	// init
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()
	c := nextInt()

	ab := a + b
	ac := a + c
	bc := b + c

	dataInt := []int{ab, ac, bc}
	sort.Ints(dataInt)
	return dataInt[0]
}

func main() {
	fmt.Println(solve1())
}
