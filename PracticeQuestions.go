// Author: Will Hogan
// Date: 29/09/2016
// Module: Emerging Technologies
// Details: Problem Sheet 1, Programming language 'GO'

package main

import (
	"fmt"
	"strconv"
	// "math/rand"
	// "time"
)

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



	// Problem sheet 1, Question 4
	/* 4. Write a program that takes as input a positive integer and applies the following
	   operations until the sequence begins to repeat: if the number is even, divide it by 2,
	   but if the number is odd, multiply it by 3 and add 1. The program should print the
	   generated sequence to the screen. You might want to consider whether the program
	   always terminates, and what will happen should the program encounter a 0. */

	oddEvenArray() // My Version
	collatz(200) // Ian's Version from Lab



	// W.I.P
	// Problem sheet 1, Question 5
	/* 5. Write a program that accepts four characters as input, and outputs all permutations
       of those four characters. What should your program do if two or more of the
       characters are the same? */

	const MAXCHARS =  4
	var charVals []string
	singleChar := ""
	count := 0

	for count < MAXCHARS {
		fmt.Println("Enter a Char") // Enter a value
		fmt.Scan(&singleChar) // Initialise string with scanned value

		if _, err := strconv.Atoi(singleChar); err == nil { //  If not a character....
			fmt.Println(singleChar, " Is a number, Enter characters only") // ...throw error
		}else if len(singleChar) > 1 { // If char size is more than one....
			fmt.Println("Enter only one Character") // Throw error
		}else if (contains(charVals, singleChar) ==  true ){ // Check if the current value entered already exists in the array
			fmt.Println("You've already entered that character") // If so throw error message
		}else {
			charVals = append(charVals, singleChar) // Append the char to the charVals array
			count++ // Increment the counter
		}
	}
	fmt.Println("You Entered 4 characters ", charVals) // Display the Char array in console

	// Smaller Array Shuffle, need to use Rand Seed, See below shuffle function
		/*arr := []string{"A", "B", "C", "D"}
		fmt.Println(arr)
		shuffle(arr)
		fmt.Println(arr)*/
}


// Function to shuffle array, W.I.P
/*func shuffle(arr []string) {
	// rand.Seed(time.Nanosecond) // Not working...
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Int31n(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}*/


// Function to check if value already exists in the array in Question 5
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


// Problem sheet 1, Question 1
func sumOfMultiples() int{
	const MAX = 1000 // Declare a constant value of 1000
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
} // End Reverse String



// Problem sheet 1, Question 4 (Collatz Conjecture)
func oddEvenArray() {
	const MAX = 15 // Create a Max number entty amount
	count := 0 // Count
	num := 0 // The Number to be entered
	var scanValues bool = true
	var valueArray []int

	for scanValues && MAX > count { // While The Max Amount has not been enterd into the Array, keep going
		fmt.Println("Enter a Number") //  Enter a value
		fmt.Scan(&num) // Initialise num with scanned value

		if (num <= 0) {
			fmt.Println("Please enter a valid number") // If num is 0 or less, display error
		} else if (num % 2 == 0) { // If Even..
			num = (num / 2) // Do this...
			valueArray = append(valueArray, num) // And append to the Array Result
			count++ // Increment the count
		} else if (num % 2 != 0) { // If odd...
			num = (num * 3 + 1) // Do this....
			valueArray = append(valueArray, num)// And append to the Array Result
			count++ // Increment the count
		}
	}
	fmt.Println("Filled Int Array Result: ", valueArray) // Display the result in console
} // End oddEvenArray


// This function iterates over a postive 'n' until sequesnce repeats
// At each iteration n is divided by 2 if it's even otherwise it's X 3 and 1 is added
func collatz(n uint) {
	for ; n != 1; {
		fmt.Print(n, " ")
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}
	fmt.Println(n)
}