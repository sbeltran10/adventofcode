// https://adventofcode.com/2020/day/10

package day10

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

var adapters []int

// 1 and 3 start at 1 because of the outlet and the device itself
var differences = map[int]int{
	1: 1,
	2: 0,
	3: 1,
}

var adapterCombinations = map[string]bool{}

func Exec(scanner *bufio.Scanner, args []string) int {

	// build array of adapters
	for scanner.Scan() {
		numberVal, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, numberVal)
	}

	// Sort so it is easier to find jolt differences
	sort.Ints(adapters)

	if args[0] == "part1" {
		return GetDifferences()
	} else {
		return GetDistinctWays(adapters)
	}

}

func GetDifferences() int {
	// Get jolt differences between all adapters
	for i, adapter := range adapters[1:] {
		differences[adapter-adapters[i]]++
	}

	return differences[1] * differences[3]
}

// EXTREMELY INNEFICENT, MIGHT NOT EVEN WORK ATM
func GetDistinctWays(modifiedArray []int) int {
	fmt.Println(modifiedArray)
	// Add current combination to the list of tried combinations
	currentUnique := fmt.Sprint(modifiedArray)

	if !adapterCombinations[currentUnique] {
		adapterCombinations[currentUnique] = true

		var waysArrangeSub int = 1

		// Start finding elements in the array that can be removed
		for i, adapter := range modifiedArray[:len(modifiedArray)-2] {
			// fmt.Println(i)

			if modifiedArray[i+2]-adapter <= 3 {
				// Can remove element i+1
				// Create copy to not modify the original
				var arrayCopy []int = make([]int, len(modifiedArray))
				copy(arrayCopy, modifiedArray)
				waysArrangeSub += GetDistinctWays(append(arrayCopy[:i+1], arrayCopy[i+2:]...))
			}
		}

		return waysArrangeSub
	}
	return 0
}
