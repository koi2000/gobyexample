package main

import "fmt"

func main() {
	if 7%2 == 1 {
		fmt.Println("yes")
	}

	if num := 9; num < 0 {
		fmt.Println("less than 9")
	} else if num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println(num, "m")
	}
}
