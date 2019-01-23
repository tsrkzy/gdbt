package util

import (
	"fmt"
	"regexp"
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

func StripNewLine(val string) string {
	re := regexp.MustCompile(`[\r\n]`)
	out := re.ReplaceAllString(val, "")

	return out
}

func RemoveTrailEmptyLines(str string) string {
	reTrailEmptyLines := regexp.MustCompile(`(?m)(\s*)*\z`)
	replaced := reTrailEmptyLines.ReplaceAllString(str, "")

	return replaced
}

func RemoveCommentLines(str string) string {
	reCommentLines := regexp.MustCompile(`(?m)^#\![^\n]*\n`)
	replaced := reCommentLines.ReplaceAllString(str, "")

	return replaced
}
