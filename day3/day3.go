package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

func calculateDistance(input []coordinates) {
	var solution float64
	var results []float64
	for _, coordinate := range input {
		results = append(results, math.Abs(float64(coordinate.x+coordinate.y)))
	}
	for _, e := range results {
		if solution == 0 {
			solution = e
		} else if e < solution {
			solution = e
		}
	}
	log.Println(solution)
}

func findIntersection(input [][]coordinates) {
	var intersections []coordinates
	for _, coordinate := range input[0] {
		for _, comparedCoord := range input[1] {
			if reflect.DeepEqual(coordinate, comparedCoord) {
				intersections = append(intersections, coordinate)
			}
		}
	}
	calculateDistance(intersections)
}

func calculateWirePositions(input [][]string) {
	var wireOpcodes [][]coordinates
	for _, wire := range input {
		x := 0
		y := 0
		var savedOpcodes []coordinates
		for _, opcode := range wire {
			switch string(opcode[0]) {
			case "R":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				x += op
			case "L":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				x -= op
			case "U":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				y += op
			case "D":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				y -= op
			default:
				log.Println("We've read an invalid opcode")
			}
			savedOpcodes = append(savedOpcodes, coordinates{x: x, y: y})
		}
		wireOpcodes = append(wireOpcodes, savedOpcodes)
	}
	findIntersection(wireOpcodes)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// initialize a 2D slice
	opcodes := [][]string{
		{},
		{},
	}
	var counter uint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, e := range strings.Split(line, ",") {
			opcodes[counter] = append(opcodes[counter], e)
		}
		counter++
	}
	calculateWirePositions(opcodes)
}
