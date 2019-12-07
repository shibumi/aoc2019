package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
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
	fmt.Println(intersections)
}

func calculateWirePositionss(input [][]string) {
	var wireOpcodes [][]coordinates
	var savedOpcodes []coordinates
	for _, wire := range input {
		x := 0
		y := 0
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
	calculateWirePositionss(opcodes)
}
