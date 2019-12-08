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
	log.Println("start calculating intersections")
	var intersections []coordinates
	log.Println("len(input[0]): ", len(input[0]))
	log.Println("len(input[1]): ", len(input[1]))
	for _, coordinate := range input[0] {
		for _, comparedCoord := range input[1] {
			if reflect.DeepEqual(coordinate, comparedCoord) {
				intersections = append(intersections, coordinate)
			}
		}
	}
	log.Println("Found intersections")
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
				for i := 0; i <= op; i++ {
					x++
					savedOpcodes = append(savedOpcodes, coordinates{x: x, y: y})
				}
			case "L":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				for i := 0; i <= op; i++ {
					x--
					savedOpcodes = append(savedOpcodes, coordinates{x: x, y: y})
				}
			case "U":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				for i := 0; i <= op; i++ {
					y++
					savedOpcodes = append(savedOpcodes, coordinates{x: x, y: y})
				}
			case "D":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				for i := 0; i <= op; i++ {
					y--
					savedOpcodes = append(savedOpcodes, coordinates{x: x, y: y})
				}
			default:
				log.Println("We've read an invalid opcode")
			}
		}
		wireOpcodes = append(wireOpcodes, savedOpcodes)
	}
	log.Println(wireOpcodes)
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
