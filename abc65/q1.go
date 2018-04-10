package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func calc() int {
	// init
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		a[i] = nextInt() - 1
	}

	// calc
	for i := 0; i < n; i++ {
		if a[a[i]] != -1 {
			ans = a[a[i]]
			a[a[i]] = -1
		}else {
			return -1
		}
	}
	return ans + 1


}

func main() {
	fmt.Println(calc())
}
