package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseOpcodes(opcodes []int) bool {
	for i := 0; i <= len(opcodes); i = i + 4 {
		switch opcodes[i] {
		case 1:
			x := opcodes[i+1]
			y := opcodes[i+2]
			result := opcodes[i+3]
			opcodes[result] = opcodes[x] + opcodes[y]
		case 2:
			x := opcodes[i+1]
			y := opcodes[i+2]
			result := opcodes[i+3]
			opcodes[result] = opcodes[x] * opcodes[y]
		case 99:
			if opcodes[0] == 19690720 {
				return true
			}
			return false
		default:
			// if we are here, we got an invalid opcode
			return false
		}
	}
	return false
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var opcodes []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, e := range strings.Split(line, ",") {
			n, err := strconv.Atoi(e)
			if err != nil {
				log.Fatal(err)
			}
			opcodes = append(opcodes, n)
		}
	}
	// modify opcodes
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			opcodesCopy := make([]int, len(opcodes))
			copy(opcodesCopy, opcodes)
			opcodesCopy[1] = i
			opcodesCopy[2] = j
			if parseOpcodes(opcodesCopy) {
				fmt.Println("We've found a solution: ")
				fmt.Println(i*100 + j)
			}
		}
	}
}
