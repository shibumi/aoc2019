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

func intersection(wirepath1 map[point]int, wirepath2 map[point]int) {
	defer util.Elapsed("intersection")()
	intersections := make(map[point]int)
	var result int
	for k, v := range wirepath1 {
		if val, ok := wirepath2[k]; ok {
			if _, fine := intersections[k]; !fine {
				log.Println(k, v, val, v+val)
				intersections[k] = val + v
			}
		}
	}
	for _, v := range intersections {
		if result == 0 {
			result = v
		} else if v < result {
			result = v
		}
	}
	log.Println("Solution: ", result)
}

func splitOpcode(s string) (op int, err error) {
	op, err = strconv.Atoi(s[1:])
	if err != nil {
		log.Println("Received invalid opcode")
	}
	return
}

func calculateWirePositions(input [][]string) {
	// this is for measuring runtime of this function
	defer util.Elapsed("calculateWirePositions")()
	// we use hashmaps, with the point as key and the distance as value
	// with this approach we can easily search for an insection and compare the distance afterwards
	wirepath1 := make(map[point]int)
	wirepath2 := make(map[point]int)
	for index, wire := range input {
		var x int
		var y int
		var sum int
		for _, opcode := range wire {
			switch string(opcode[0]) {
			case "R":
				op, _ := splitOpcode(opcode)
				for i := 0; i < op; i++ {
					x++
					sum++
					if index == 0 {
						if _, ok := wirepath1[point{x: x, y: y}]; !ok {
							wirepath1[point{x: x, y: y}] = sum
						}
					} else if index == 1 {
						if _, ok := wirepath2[point{x: x, y: y}]; !ok {
							wirepath2[point{x: x, y: y}] = sum
						}
					} else {
						log.Fatal("Received more than two wirepaths")
					}
				}
			case "L":
				op, _ := splitOpcode(opcode)
				for i := 0; i < op; i++ {
					x--
					sum++
					if index == 0 {
						if _, ok := wirepath1[point{x: x, y: y}]; !ok {
							wirepath1[point{x: x, y: y}] = sum
						}
					} else if index == 1 {
						if _, ok := wirepath2[point{x: x, y: y}]; !ok {
							wirepath2[point{x: x, y: y}] = sum
						}
					} else {
						log.Fatal("Received more than two wirepaths")
					}
				}
			case "U":
				op, _ := splitOpcode(opcode)
				for i := 0; i < op; i++ {
					y++
					sum++
					if index == 0 {
						if _, ok := wirepath1[point{x: x, y: y}]; !ok {
							wirepath1[point{x: x, y: y}] = sum
						}
					} else if index == 1 {
						if _, ok := wirepath2[point{x: x, y: y}]; !ok {
							wirepath2[point{x: x, y: y}] = sum
						}
					} else {
						log.Fatal("Received more than two wirepaths")
					}
				}
			case "D":
				op, _ := splitOpcode(opcode)
				for i := 0; i < op; i++ {
					y--
					sum++
					if index == 0 {
						if _, ok := wirepath1[point{x: x, y: y}]; !ok {
							wirepath1[point{x: x, y: y}] = sum
						}
					} else if index == 1 {
						if _, ok := wirepath2[point{x: x, y: y}]; !ok {
							wirepath2[point{x: x, y: y}] = sum
						}
					} else {
						log.Fatal("Received more than two wirepaths")
					}
				}
			default:
				log.Println("We've read an invalid opcode")
			}
		}
	}
	intersection(wirepath1, wirepath2)
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
