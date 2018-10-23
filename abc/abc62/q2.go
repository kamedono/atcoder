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

func nextLine() string {
    sc.Scan()
    return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	h := nextInt() + 2
	w := nextInt() + 2
	// for i := 0; i < h; i++ {
	// 	a.
	for i := 0; i < h; i++ {
		if i == 0 || i == h - 1 {
			// 最初と最後
			for j := 0; j < w; j++ {
				fmt.Print("*")
			}	
		}else {
			fmt.Print("*")
			fmt.Print(nextLine())
			fmt.Print("*")
		}
		fmt.Println()
	}
}
