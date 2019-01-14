package lib

import (
	"fmt"
	"strconv"
)

func intToStr(n int) string {
	return strconv.Itoa(n)
}

func echo(str string) {
	fmt.Println(str)
}

func report(err error) {
	fmt.Println(err)
}

func strIndexOf(val string, arr []string) int {
	for i := 0; i < len(arr); i++ {
		a := arr[i]
		if val == a {
			return i
		}
	}
	return -1
}
