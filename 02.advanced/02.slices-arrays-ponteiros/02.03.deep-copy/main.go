package main

import (
	"fmt"
	"strconv"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	copyMatrix := deepCopy(matrix)
	copyMatrix[0][0] = 10

	printMatrix("Matriz Original:", matrix)
	printMatrix("Matriz Copiada:", copyMatrix)
}

func deepCopy(matrix [][]int) [][]int {
	copyMatrix := make([][]int, len(matrix))
	for i, row := range matrix {
		copyMatrix[i] = append([]int(nil), row...)
	}
	return copyMatrix
}

func printMatrix(title string, matrix [][]int) {
	fmt.Println(title)
	if len(matrix) == 0 {
		fmt.Println("[]")
		return
	}

	widths := columnWidths(matrix)

	for rowIndex, row := range matrix {
		leftDelimiter, rightDelimiter := matrixDelimiters(rowIndex, len(matrix))

		fmt.Print(leftDelimiter)
		for colIndex, width := range widths {
			fmt.Print(" ")

			if colIndex < len(row) {
				fmt.Printf("%*d", width, row[colIndex])
				continue
			}

			fmt.Printf("%*s", width, "")
		}
		fmt.Printf(" %s\n", rightDelimiter)
	}
}

func columnWidths(matrix [][]int) []int {
	var widths []int
	for _, row := range matrix {
		for len(widths) < len(row) {
			widths = append(widths, 0)
		}
		for colIndex, val := range row {
			if width := len(strconv.Itoa(val)); width > widths[colIndex] {
				widths[colIndex] = width
			}
		}
	}
	return widths
}

func matrixDelimiters(rowIndex, rowCount int) (string, string) {
	if rowCount == 1 {
		return "[", "]"
	}
	if rowIndex == 0 {
		return "⎡", "⎤"
	}
	if rowIndex == rowCount-1 {
		return "⎣", "⎦"
	}
	return "⎢", "⎥"
}
