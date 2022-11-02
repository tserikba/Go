package main

import "fmt"

func main() {
	var lucky int
	fmt.Scan(&lucky)

	first_half := lucky / 1000
	second_half := lucky % 1000

	first := first_half / 100
	second := first_half / 10 % 10
	third := first_half % 100 % 10
	fourth := second_half / 100
	fifth := second_half / 10 % 10
	sixth := second_half % 100 % 10

	if first+second+third == fourth+fifth+sixth {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
