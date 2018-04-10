package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	hw := strings.Split(nextLine(), " ")
	h, _ := strconv.Atoi(hw[0])

	for i := 0; i < h; i++ {
		c := nextLine()
		fmt.Println(c)
		fmt.Println(c)
	}
	 

}
