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
	c := nextLine()
	ans := ""
	if c == "a" || c == "i" || c == "u"|| c == "e"|| c == "o" {
		ans = "vowel"
	} else {
		ans = "consonant"
	}

	fmt.Println(ans)
}
