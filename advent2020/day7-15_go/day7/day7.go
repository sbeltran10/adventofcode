// https://adventofcode.com/2020/day/7

package day7

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	Color  string
	Amount int
}

var bagsrules = map[string][]Rule{}

func Exec(scanner *bufio.Scanner, args []string) int {

	// Build bagrules map first
	for scanner.Scan() {
		ruleStr := scanner.Text()
		splitRule := strings.Split(ruleStr, "bags contain")
		ruleColor := strings.Trim(splitRule[0], " ")

		// Convert the rules into an array
		splitBagRules := strings.Split(splitRule[1], ",")
		var currentBagRules []Rule

		for _, rule := range splitBagRules {
			// Remove unnecesary info
			noContainsIndex := strings.Index(rule, "no other")
			// Skip if bag contains no more items
			if noContainsIndex == -1 {
				endIndex := strings.Index(rule, "bag") - 1
				sanitizedRule := strings.TrimSpace(rule[:endIndex])
				var numberOfTimes int
				x, err := strconv.ParseInt(string(sanitizedRule[0]), 10, 64)
				if err != nil {
					fmt.Println(err)
				}
				numberOfTimes = int(x)
				ruleName := strings.TrimSpace(sanitizedRule[1:])

				currentBagRules = append(currentBagRules, Rule{
					Color:  ruleName,
					Amount: numberOfTimes,
				})
			}
			bagsrules[ruleColor] = currentBagRules
		}
	}

	var inputBagColor string = strings.Replace(args[1], "_", " ", -1)

	// If part 1
	if args[0] == "part1" {
		return part1Count(inputBagColor)
	} else {
		return part2Count(inputBagColor)
	}
}

// TODO: Recursiveness can be improved
func part1Count(inputBagColor string) int {
	// Go over each rule checking if the input bag can be reached at least once
	// Replace underscore in bag color argument with space
	var count int = 0
	for _, rules := range bagsrules {
		for _, rule := range rules {
			// Break when target bag color is found at least once
			if CheckForColor(rule.Color, inputBagColor) {
				count++
				break
			}
		}
	}
	return count
}

// Recursive func to check the map for the desired color
func CheckForColor(color string, targetRule string) bool {
	if color == targetRule {
		return true
	} else {
		if val, ok := bagsrules[color]; ok {
			for _, rule := range val {
				// Check own rules if any
				// Return only inside the if so the loop doesn't break if no result is found
				if CheckForColor(rule.Color, targetRule) {
					return true
				}
			}
			return false
		} else {
			return false
		}
	}
}

func part2Count(inputBagColor string) int {
	// Start counting all of the bags from the target bag
	var count int = CountBags(inputBagColor)
	return count
}

func CountBags(color string) int {
	var count int = 0
	if val, ok := bagsrules[color]; ok {
		for _, rule := range val {
			// Add own amount first
			count += rule.Amount

			// Add children amounts multiplied by my own amount
			count += CountBags(rule.Color) * rule.Amount
		}
	}
	return count
}
