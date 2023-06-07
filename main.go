package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// Main function that reads the birthyear from the args, calculates the current age,
// calculates the next age that is prime and adds it to the birth year.
func main() {
	fmt.Println("Enter a birthyear. (q to quit)") // Only prompt once.
	for {
		birthyear, quit := getYearFromConsole()
		if quit {
			return
		}

		age := getAge(birthyear)
		nextPrime := nextPrimeNumber(age)
		fmt.Printf("%v\n\n", birthyear+nextPrime)
	}
}

// Input loop for the user to put in multiple inputs. Returns only on valid inputs (integer, q for quit).
func getYearFromConsole() (int, bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input. Please enter only one number.\n\n")
			continue
		}
		input = strings.TrimRight(input, "\n")

		if input == "q" {
			return 0, true
		}

		ret, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Could not parse '%v' to an integer.\n\n", input)
			continue
		}
		return ret, false
	}
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
