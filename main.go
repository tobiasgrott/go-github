package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hi, DEV World! 👍")
	fmt.Printf("The sum of %d and %d is %d", 5, 5, Sum(5, 5))
}
