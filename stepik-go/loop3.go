package main

import "fmt"

func main() {
	var a, sum int
	fmt.Scan(&a)
	sum = 0

	for i := 0; i <= a; i++ {
		var b int
		fmt.Scan(&b)
		if b < 100 && b >= 10 && b%8 == 0 {
			sum += b
		}

	}
	fmt.Println(sum)
}
