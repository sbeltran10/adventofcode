// https://adventofcode.com/2020/day/9

package day9

import (
	"bufio"
	"strconv"
)

var xmasData []int
var preamble int
var current int = 0
var found int = -1

func Exec(scanner *bufio.Scanner, args []string) int {

	preambleVal, _ := strconv.Atoi(args[1])
	preamble = preambleVal

	for scanner.Scan() {
		numberVal, _ := strconv.Atoi(scanner.Text())
		xmasData = append(xmasData, numberVal)

		// Check if number needs to be verified
		if current >= preamble {
			sumFound := false
			for i, valuei := range xmasData[current-preamble : current] {
				for _, valuej := range xmasData[current-preamble+i+1 : current] {
					if valuei+valuej == numberVal {
						sumFound = true
						break
					}
				}
				if sumFound {
					break
				}
			}
			if !sumFound {
				found = numberVal
			}
		}
		if found != -1 {
			break
		}
		current++
	}

	// Find contiguous set of numbers
	if args[0] == "part2" {
		// TODO
	}

	return found
}
