package main

import (
	"fmt"
	"time"
)

func showMessage(msg string, done chan bool) {
	fmt.Println(msg)
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}

	done <- true
}

func fibonacci(c, done chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-done:
			return
		}
	}
}

func fibonacciRec(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	// 8 - 1 = 7
	// 7- 1 = 6
	// 3- 1 = 2

	return fibonacciRec(n-1) + fibonacciRec(n-2)
}

func fatorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * fatorial(n-1)
}

func sum(n int) int {
	if n == 1 {
		return 1
	}

	return n + sum(n-1)
}

func isPalindrome(c string) bool {
	if len(c) <= 1 {
		return true
	}

	if c[0] != c[len(c)-1] {
		return false
	}

	fmt.Println("c[1 : len(c)-1])", c[1:len(c)-1])
	return isPalindrome(c[1 : len(c)-1])
}

func pow(x int, y int) int {
	if y == 0 {
		return 1
	}

	return x * pow(x, y-1)
}

func invertString(s string) string {
	if len(s) <= 1 {
		return s
	}

	fmt.Println("OI", string(s[1:]), string(s[0]))
	return invertString(s[1:]) + string(s[0])
}

func main() {

	// c := make(chan int)
	// done := make(chan int)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	done <- 0
	// }()
	// fibonacci(c, done)

	// fmt.Println("hasa", fibonacciRec(5))

	fmt.Println("Fat", fatorial(5))
	fmt.Println("Fat", sum(5))
	fmt.Println("Fat", isPalindrome("radar"))
	fmt.Println("Fat", pow(2, 6))
	fmt.Println("Fat", invertString("hello"))

}
