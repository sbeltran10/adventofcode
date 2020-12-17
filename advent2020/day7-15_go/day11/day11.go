// https://adventofcode.com/2020/day/10

package day11

import (
	"bufio"
)

var seats [][]string
var totalOccupied int = 0

func Exec(scanner *bufio.Scanner, args []string) int {

	// build seats 2d array
	for scanner.Scan() {
		row := scanner.Text()

		var rowArray []string
		for _, character := range row {
			rowArray = append(rowArray, string(character))
		}

		seats = append(seats, rowArray)
	}
	OccupySeats()
	return totalOccupied
}

func OccupySeats() {
	var currentCount int = 0
	var newSeatLayout [][]string
	for indexRow, seatRow := range seats {
		var newSeatRow []string
		for indexCol, seat := range seatRow {
			if seat == "." {
				newSeatRow = append(newSeatRow, ".")
			} else {
				adyacentList := GetAdyacentList(indexRow, indexCol, len(seats), len(seatRow))

				adyacentOccupied := 0
				for _, adyacentCoords := range adyacentList {
					if seats[adyacentCoords[0]][adyacentCoords[1]] == "#" {
						adyacentOccupied++
					}
				}
				if seat == "L" && adyacentOccupied == 0 {
					newSeatRow = append(newSeatRow, "#")
					currentCount++
				} else if adyacentOccupied >= 4 {
					newSeatRow = append(newSeatRow, "L")
				} else {
					newSeatRow = append(newSeatRow, seat)
					if seat == "#" {
						currentCount++
					}
				}
			}
		}
		newSeatLayout = append(newSeatLayout, newSeatRow)
	}

	// If count changed, attempt new occupation
	if totalOccupied != currentCount {
		totalOccupied = currentCount
		seats = newSeatLayout
		OccupySeats()
	}
}

func GetAdyacentList(rowIndex int, colIndex int, rowLength int, colLength int) [][]int {
	var adyacentList [][]int
	if rowIndex != 0 {
		adyacentList = append(adyacentList, []int{rowIndex - 1, colIndex})

		if colIndex != 0 {
			adyacentList = append(adyacentList, []int{rowIndex - 1, colIndex - 1})
		}

		if colIndex != colLength-1 {
			adyacentList = append(adyacentList, []int{rowIndex - 1, colIndex + 1})
		}
	}

	if rowIndex != rowLength-1 {
		adyacentList = append(adyacentList, []int{rowIndex + 1, colIndex})

		if colIndex != 0 {
			adyacentList = append(adyacentList, []int{rowIndex + 1, colIndex - 1})
		}

		if colIndex != colLength-1 {
			adyacentList = append(adyacentList, []int{rowIndex + 1, colIndex + 1})
		}
	}

	if colIndex != 0 {
		adyacentList = append(adyacentList, []int{rowIndex, colIndex - 1})
	}

	if colIndex != colLength-1 {
		adyacentList = append(adyacentList, []int{rowIndex, colIndex + 1})
	}

	return adyacentList
}
