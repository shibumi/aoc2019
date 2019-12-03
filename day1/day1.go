package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateFuel(mass int) (result int) {
	result = int(mass/3) - 2
	return
}

func calculateFuelOfFuel(fuel int, fuels *[]int) (result int) {
	fuel = calculateFuel(fuel)
	if fuel <= 0 {
		return 0
	}
	*fuels = append(*fuels, fuel)
	return calculateFuelOfFuel(fuel, fuels)
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var modules []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		modules = append(modules, calculateFuel(int(number)))
	}

	cpy := make([]int, len(modules))
	copy(cpy, modules)
	for _, fuel := range cpy {
		calculateFuelOfFuel(fuel, &modules)
	}
	var sum int
	for _, n := range modules {
		sum += n
	}
	fmt.Println(sum)
}
