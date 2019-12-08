package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func equal(x point, y point) bool {
	if x.x == y.x && x.y == y.y {
		return true
	}
	return false
}

func contains(coordinate point, coordinates []point) bool {
	for _, e := range coordinates {
		if equal(e, coordinate) {
			return true
		}
	}
	return false
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(intersections []point) {
	var distances []int
	for _, p := range intersections {
		result := absInt(p.x) + absInt(p.y)
		distances = append(distances, result)
	}
	var solution int
	for i, p := range distances {
		if i == 0 {
			solution = p
		}
		if p < solution {
			solution = p
		}
	}
	log.Println(solution)
}

func calculateWirePositions(input [][]string) {
	var coordinates []point
	var intersections []point
	for index, wire := range input {
		x := 0
		y := 0
		for _, opcode := range wire {
			switch string(opcode[0]) {
			case "R":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				for i := 0; i <= op; i++ {
					x++
					if index == 1 {
						if contains(point{x: x, y: y}, coordinates) {
							intersections = append(intersections, point{x: x, y: y})
						}
					} else {
						coordinates = append(coordinates, point{x: x, y: y})
					}
				}
			case "L":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				for i := 0; i <= op; i++ {
					x--
					if index == 1 {
						if contains(point{x: x, y: y}, coordinates) {
							intersections = append(intersections, point{x: x, y: y})
						}
					} else {
						coordinates = append(coordinates, point{x: x, y: y})
					}
				}
			case "U":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				for i := 0; i <= op; i++ {
					y++
					if index == 1 {
						if contains(point{x: x, y: y}, coordinates) {
							intersections = append(intersections, point{x: x, y: y})
						}
					} else {
						coordinates = append(coordinates, point{x: x, y: y})
					}
				}
			case "D":
				op, err := strconv.Atoi(opcode[1:])
				if err != nil {
					log.Println("Received invalid opcode")
				}
				for i := 0; i <= op; i++ {
					y--
					if index == 1 {
						if contains(point{x: x, y: y}, coordinates) {
							intersections = append(intersections, point{x: x, y: y})
						}
					} else {
						coordinates = append(coordinates, point{x: x, y: y})
					}
				}
			default:
				log.Println("We've read an invalid opcode")
			}
		}
	}
	log.Println(intersections)
	distance(intersections)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// initialize a 2D slice
	// This solution is limited to two lines output
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
