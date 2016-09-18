package main
import "fmt"

func main() {

	// Problem sheet 1, Question 1
	/* 1. The natural numbers below 10 that are multiples of 3 or 5 are: 3, 5, 6 and 9. The sum
	   of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000 [1]. */
	fmt.Println("The Result is: ", sumOfMultiples())

	// W.I.P
	// Problem sheet 1, Question 2
	/* 2. The first six prime numbers in order are 2, 3, 5, 7, 11, and 13. So, for instance, the
	   4th prime number is 7. Write a program to find the 10,001st prime number. [1].*/
	// const MAX  = 10001


	// Problem sheet 1, Question 3
	/* 3. Write a program that accepts a user inputted string and prints its reverse.*/
	fmt.Println("The Reverse of entered string is: ", reverseString("Will Hogan"))
}


// Problem sheet 1, Question 1
func sumOfMultiples() int{
	const MAX  = 1000 // Declare a constant value of 1000
	sum := 0

	for i := 0; i < MAX; i++ { // Loop over the MAX value
		if (i % 3 == 0 || i % 5 == 0) { // if any of the looped value modulus 3 or 5 is 0
			sum += i // add them together
		}
	}
	return sum //  Return the integer sum total
} // End sumOfMultiples function



// Problem sheet 1, Question 3
func reverseString(str string) string {
	runes := []rune(str) // Declare a Rune array that holds the numeric value of each character
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i] // Loop over the array, and reverse the values
	}
	return string(runes) // Return the string equivalent of the value enetered.
}