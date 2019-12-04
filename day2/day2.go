package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseOpcodes(opcodes []int) {
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
			fmt.Println(opcodes)
			return
		default:
			log.Fatalf("We shouldn't be here: index: %v, value: %v, opcodes: %v", i, opcodes[i], opcodes)
		}
	}
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
	opcodes[1] = 12
	opcodes[2] = 2
	parseOpcodes(opcodes)
}
