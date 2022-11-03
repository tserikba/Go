package main

import "fmt"

func main() {
	var n, max, count int
	max = 0

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max {
			max = n
			count = 1
		} else if n == max {
			count++
		}

	}
	fmt.Println(count)
}
