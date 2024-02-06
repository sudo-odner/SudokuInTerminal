package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Sudoku [9][9]SudokuCell

type SudokuCell struct {
	data   int
	access bool
}

func (a *Sudoku) print() {
	var mainLine = strings.Repeat("=", 39)
	var secondLine = "|---|---|---||---|---|---||---|---|---|"
	for y := range a {
		if y%3 == 0 {
			fmt.Println(mainLine)
		} else {
			fmt.Println(secondLine)
		}
		for x := range a[y] {
			cell := a[y][x]
			if x%3 == 0 {
				fmt.Print("|")
			}
			fmt.Printf(" %v |", cell.data)
		}
		fmt.Print("\n")
	}
	fmt.Print(mainLine, "\n\n")
}

func (a *Sudoku) createSudokuBase() {
	shapeMiniArea := 3
	for y := 0; y < (shapeMiniArea * shapeMiniArea); y++ {
		for x := 0; x < (shapeMiniArea * shapeMiniArea); x++ {
			a[y][x] = SudokuCell{
				data:   ((y*shapeMiniArea+y/shapeMiniArea+x)%(shapeMiniArea*shapeMiniArea) + 1),
				access: false,
			}
		}
	}
}

func (a *Sudoku) transposition() {
	var newArea Sudoku

	for i := range a {
		copy(newArea[i][:], a[i][:])
	}

	for y := range newArea {
		for x := range newArea {
			a[y][x] = newArea[x][y]
		}
	}
}

func (a *Sudoku) swapRowsSmall() {
	indexArea := rand.Intn(3)
	indexFirstLineInArea := rand.Intn(3)
	indexSecondLineInArea := rand.Intn(3)

	for indexFirstLineInArea == indexSecondLineInArea {
		indexSecondLineInArea = rand.Intn(3)
	}

	indexFirstLine, indexSecondLine := indexArea*3+indexFirstLineInArea, indexArea*3+indexSecondLineInArea

	a[indexFirstLine], a[indexSecondLine] = a[indexSecondLine], a[indexFirstLine]
}

func (a *Sudoku) swapColumsSmall() {
	a.transposition()
	a.swapRowsSmall()
	a.transposition()
}

func (a *Sudoku) swapRowsArea() {
	indexFirstArea := rand.Intn(3)
	indexSecondArea := rand.Intn(3)

	for indexFirstArea == indexSecondArea {
		indexSecondArea = rand.Intn(3)
	}

	indexFirstArea, indexSecondArea = indexFirstArea*3, indexSecondArea*3

	for i := 0; i < 3; i++ {
		a[indexFirstArea+i], a[indexSecondArea+i] = a[indexSecondArea+i], a[indexFirstArea+i]
	}
}

func (a *Sudoku) swapColumsArea() {
	a.transposition()
	a.swapRowsArea()
	a.transposition()
}

func newGameSudoku() {
	var game Sudoku
	game.createSudokuBase()
	game.print()
	game.swapColumsArea()
	game.print()
}

func main() {
	newGameSudoku()
}

/*
=======================================
| 5 | 2 | 5 || 0 | 7 | 0 || 0 | 0 | 0 |
|---|---|---||---|---|---||---|---|---|
| 6 | 0 | 0 || 1 | 9 | 5 || 0 | 0 | 0 |
|---|---|---||---|---|---||---|---|---|
| 0 | 9 | 8 || 0 | 0 | 0 || 0 | 6 | 0 |
=======================================
| 8 | 0 | 0 || 0 | 6 | 0 || 0 | 0 | 3 |
|---|---|---||---|---|---||---|---|---|
| 4 | 0 | 0 || 8 | 0 | 3 || 0 | 0 | 1 |
|---|---|---||---|---|---||---|---|---|
| 7 | 0 | 0 || 0 | 2 | 0 || 0 | 0 | 6 |
=======================================
| 8 | 0 | 0 || 0 | 6 | 0 || 0 | 0 | 3 |
|---|---|---||---|---|---||---|---|---|
| 4 | 0 | 0 || 8 | 0 | 3 || 0 | 0 | 1 |
|---|---|---||---|---|---||---|---|---|
| 7 | 0 | 0 || 0 | 2 | 0 || 0 | 0 | 6 |
=======================================
*/
