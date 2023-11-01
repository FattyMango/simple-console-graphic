package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	// Define the original graphic as a 2D array of characters
	graphic := [][]rune{
		{'1', '"', '"','"','4'},
		{'"', '/', '"','/','"'},
		{'"', '"', '*','"','"'},
		{'2', '/', '"','/','3'},
	}

	const numFrames = 90
	const delay = 1000 * time.Millisecond

	for i := 0; i < numFrames; i++ {
		clearScreen()

		transposeMatrix(graphic)
		reverseColumns(graphic)
		printArray(graphic)

		time.Sleep(delay)
	}
}

func printArray(arr [][]rune) {
	for _, row := range arr {
		for _, val := range row {
			fmt.Printf("%c ", val)
		}
		fmt.Println()
	}}

func transposeMatrix(matrix [][]rune) [][]rune {
    if len(matrix) == 0 {
        return matrix
    }

    rows, cols := len(matrix), len(matrix[0])

    transposed := make([][]rune, cols)
    for i := range transposed {
        transposed[i] = make([]rune, rows)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            transposed[j][i] = matrix[i][j]
        }
    }

    return transposed
}
func reverseColumns(matrix [][]rune) {
    rows, cols := len(matrix), len(matrix[0])

    for j := 0; j < cols; j++ {
        for i := 0; i < rows/2; i++ {
            matrix[i][j], matrix[rows-1-i][j] = matrix[rows-1-i][j], matrix[i][j]
        }
    }
}
func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	default:
		// Unsupported platform; clear the screen by printing newlines
		fmt.Print("\033[H\033[2J")
	}
}