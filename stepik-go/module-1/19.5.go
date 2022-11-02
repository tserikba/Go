package main

import "fmt"

func main() {
	var leap int
	fmt.Scan(&leap)

	if leap%400 == 0 || (leap%4 == 0 && leap%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
