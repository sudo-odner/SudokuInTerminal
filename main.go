package main

import (
	"fmt"
)

type SudokuArea [9][9]SudokuCell

type SudokuCell struct {
	data   int
	access bool
}

func (a *SudokuArea) swichTranspositionArea() {
	copyArea := a
	for y := range a {
		for x := range a {
			a[x][y] = copyArea[y][x]
		}
	}
}

func createSudokuBase() SudokuArea {
	var area SudokuArea
	shapeMiniArea := 3
	for y := 0; y < (shapeMiniArea * shapeMiniArea); y++ {
		for x := 0; x < (shapeMiniArea * shapeMiniArea); x++ {
			area[y][x] = SudokuCell{
				data:   ((y*shapeMiniArea+y/shapeMiniArea+x)%(shapeMiniArea*shapeMiniArea) + 1),
				access: false,
			}
		}
	}
	return area
}

func newGameSudoku() {
	game := createSudokuBase()
	fmt.Println(game[0][:3])
	fmt.Println(game[1][:3])
	fmt.Println(game[2][:3])
	game.swichTranspositionArea()
	fmt.Println()
	fmt.Println(game[0][:3])
	fmt.Println(game[1][:3])
	fmt.Println(game[2][:3])
}

func main() {
	newGameSudoku()
}
