package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shibumi/aoc2019/util"
)

type point struct {
	x int
	y int
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(intersections map[point]int) {
	defer util.Elapsed("distance")()
	var result int
	for _, v := range intersections {
		if result == 0 {
			result = v
		} else if v < result {
			result = v
		}
	}
	log.Println(result)
}

func calculateWirePositions(input [][]string) {
	defer util.Elapsed("calculateWirePositions")()
	coordinates := make(map[point]int)
	intersections := make(map[point]int)
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
						if val, ok := coordinates[point{x: x, y: y}]; ok {
							intersections[point{x: x, y: y}] = val
						}
					} else {
						coordinates[point{x: x, y: y}] = absInt(x) + absInt(y)
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
						if val, ok := coordinates[point{x: x, y: y}]; ok {
							intersections[point{x: x, y: y}] = val
						}
					} else {
						coordinates[point{x: x, y: y}] = absInt(x) + absInt(y)
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
						if val, ok := coordinates[point{x: x, y: y}]; ok {
							intersections[point{x: x, y: y}] = val
						}
					} else {
						coordinates[point{x: x, y: y}] = absInt(x) + absInt(y)
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
						if index == 1 {
							if val, ok := coordinates[point{x: x, y: y}]; ok {
								intersections[point{x: x, y: y}] = val
							}
						} else {
							coordinates[point{x: x, y: y}] = absInt(x) + absInt(y)
						}
					}
				}
			default:
				log.Println("We've read an invalid opcode")
			}
		}
	}
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
