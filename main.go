package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

// Main function that reads the birthyear from the args, calculates the current age,
// calculates the next age that is prime and adds it to the birth year.
func main() {
	birthyear := getYearFromOs()
	age := getAge(birthyear)
	nextPrime := nextPrimeNumber(age)
	fmt.Println(birthyear + nextPrime)
}

// Read and parse the program arguments. Only one int argument is allowed.
func getYearFromOs() int {
	if len(os.Args) != 2 {
		log.Fatal("Exactly one command-line parameter required.")
	}
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Could not parse '%v' to an integer.\n", os.Args[1])
	}

	return year
}

// Calculate current age of someone born in the given year. Assert positive age.
func getAge(birthyear int) int {
	age := time.Now().Year() - birthyear
	if age < 0 {
		log.Fatal("You cannot have a negative age.")
	}

	return age
}

// Calculates the next positive prime number, after the given number.
// Return -1 if Berntrand's postulate did not hold.
func nextPrimeNumber(n int) int {
	// Edge cases of Bertrand's postulate.
	if n < 1 {
		return 2
	}

	next := n + 1
	// Bertrand's postulate n < p < 2n.
	for i := next; i < 2*next; i++ {
		if isPrime(i) {
			return i
		}
	}

	// Berntrand's postulate didn't hold.
	return -1
}

// Return true, if the given number is prime. False, if not.
func isPrime(n int) bool {
	// edge case
	if n < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		// If any number divides n properly, n is not a prime number.
		if n%i == 0 {
			return false
		}
	}
	return true
}
