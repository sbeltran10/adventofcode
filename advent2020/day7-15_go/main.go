package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"puzzle/day10"
	"puzzle/day11"
	"puzzle/day7"
	"puzzle/day8"
	"puzzle/day9"
)

var puzzleFnMap = map[string]interface{}{
	"day7":  day7.Exec,
	"day8":  day8.Exec,
	"day9":  day9.Exec,
	"day10": day10.Exec,
	"day11": day11.Exec,
}

func main() {
	// First arg is the name of the day to exec
	args := os.Args[1:]
	dayArg := args[0]

	// Create scanner to read input
	f, err := os.Open(fmt.Sprintf("%s/input.txt", dayArg))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := puzzleFnMap[dayArg].(func(*bufio.Scanner, []string) int)(scanner, args[1:])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
