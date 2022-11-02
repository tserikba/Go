package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a > 9 && a < 100 {
		fmt.Println(a / 10)
	} else if a > 99 && a < 1000 {
		fmt.Println(a / 100)
	} else if a > 999 && a < 10000 {
		fmt.Println(a / 1000)
	} else if a > 1 && a <= 9 {
		fmt.Println(a)
	} else if a <= 10000 {
		fmt.Println(a / 10000)
	}
}
