package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseOnlyLetters(s string) string {
	strs := strings.Split(s, "")
	i, e := 0, len(strs)-1

	// unicode.IsLetter()
	for i < e {
		f_aux := strs[i]
		s_aux := strs[e]
		if unicode.IsLetter(rune(f_aux[0])) && unicode.IsLetter(rune(s_aux[0])) {
			strs[i] = s_aux
			strs[e] = f_aux
		} else if !unicode.IsLetter(rune(f_aux[0])) {
			e++
		} else {
			i--
		}
		i++
		e--
	}

	return strings.Join(strs, "")
}

func main() {
	a := reverseOnlyLetters("Test1ng-Leet=code-Q!")
	fmt.Println("AAA", a)
}
