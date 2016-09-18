package main
import "fmt"

func main() {

	// Practice Question 1
	/* 1. The natural numbers below 10 that are multiples of 3 or 5 are: 3, 5, 6 and 9. The sum
	   of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000 [1]. */
	fmt.Println("Result is: ", sumOfMultiples())
}

func sumOfMultiples() int{
	var sum int = 0
	const MAX  = 1000

	for i := 1; i < MAX; i++ {
		if (i % 3 == 0 || i % 5 == 0) {
			sum += i
		}
	}
	return sum
} // End sumOfMultiples function