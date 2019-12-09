package main

import (
	"fmt"
	"strconv"
)

func checkNeighborship(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func checkNeverDecrease(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

// ruleOfTwo hehe.. I like this star wars metapher.
func ruleOfTwo(n int) bool {
	var exit bool
	var exception bool
	s := strconv.Itoa(n)
	for _, e := range s {
		var counter uint
		for i, d := range s {
			if e == d {
				counter++
			}
			if counter > 2 {
				exit = true
			}
			if i == len(s)-1 && counter == 2 {
				exception = true
			}
		}
	}
	if exception {
		return true
	}
	if exit {
		return false
	}
	return true
}

func main() {
	// Contraints for this puzzle:
	// It is a six-digit number.
	// The value is within the range given in your puzzle input.
	// Two adjacent digits are the same (like 22 in 122345).
	// Going from left to right, the digits never decrease;
	// they only ever increase or stay the same (like 111123 or 135679).

	// first constraint: the passwords is in the range 183564-657474.
	// this should also meet the constraint regarding 6 digits.
	var result uint
	for i := 183564; i <= 657474; i++ {
		if checkNeighborship(i) {
			if checkNeverDecrease(i) {
				if ruleOfTwo(i) {
					result++
				}
			}
		}
	}
	fmt.Println(result)
}
