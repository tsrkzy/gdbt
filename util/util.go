package util

import (
	"fmt"
	"strconv"
)

func IntToStr(n int) string {
	return strconv.Itoa(n)
}

func Echo(str string) {
	fmt.Println(str)
}

func Report(err error) {
	fmt.Println(err)
}

func StrIndexOf(val string, arr []string) int {
	for i := 0; i < len(arr); i++ {
		a := arr[i]
		if val == a {
			return i
		}
	}
	return -1
}
