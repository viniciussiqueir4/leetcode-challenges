package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	rows := strings.Split(s, "")
	split := make([][]string, numRows)
	b, l := 0, numRows-1
	c, dir := 0, 1

	for _, v := range rows {
		split[c] = append(split[c], v)

		if c == b {
			dir = 1
		}

		if c == l {
			dir = -1
		}

		if dir == 1 {
			c++
		} else if dir == -1 && c != 0 {
			c--
		}
	}

	result := ""

	for _, s := range split {
		result = result + strings.Join(s, "")
	}

	return result
}

func main() {
	fmt.Println(convert("AB", 1))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 3))

}
