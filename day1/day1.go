package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// calculate fuel by mass
func calculateFuel(mass int) (result int) {
	result = int(mass/3) - 2
	return
}

func main() {
	// Read input file
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	// make sure to close file after end of function
	defer file.Close()

	// create slice for modules
	var modules []int
	scanner := bufio.NewScanner(file)

	// scan file by line and cast to integer
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		modules = append(modules, calculateFuel(int(number)))
	}
	// calculate sum
	var sum int
	for _, n := range modules {
		sum += n
	}
	fmt.Println(sum)
}
