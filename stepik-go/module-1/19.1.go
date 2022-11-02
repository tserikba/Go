package main

import "fmt"

func main() {
	var c int32
	fmt.Scan(&c)
	switch {
	case c > 0:
		fmt.Println("the number is negative")
	case c < 0:
		fmt.Println("the number is positive")
	case c == 0:
		fmt.Println("null")
	}
}
