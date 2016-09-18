package main

import "fmt"

func main() {

	// Practice Question 1

	sum := 0;

	for i := 0; i < 1000; i++ {
		if i % 5 == 0 {// || i % 5 = 0) {
			sum += i
		}
	}
	fmt.Println(sum)
}