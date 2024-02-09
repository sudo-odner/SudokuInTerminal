package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/fatih/color"
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
			noAccess := color.New(color.FgCyan).SprintFunc()
			if cell.access {
				fmt.Printf(" %v |", cell.data)
			} else {
				fmt.Printf(" %v |", noAccess(cell.data))
			}
		}
		fmt.Print("\n")
	}
	fmt.Print(mainLine, "\n\n")
}

func (a *Sudoku) returnSudokuByCordsPlayer(posX, posY int) {
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
			noAccess := color.New(color.FgCyan).SprintFunc()
			if y == posY && x == posX {
				if cell.access {
					fmt.Printf("[%v]|", cell.data)
				} else {
					fmt.Printf("[%v]|", noAccess(cell.data))
				}
			} else {
				if cell.access {
					fmt.Printf(" %v |", cell.data)
				} else {
					fmt.Printf(" %v |", noAccess(cell.data))
				}
			}
		}
		fmt.Print("\n")
	}
	fmt.Print(mainLine, "\n\n")
}

func (a *Sudoku) setCellUser(x, y, value int) bool {
	if a[y][x].access {
		a[y][x].data = value
		return false
	}

	return true
}

func (a Sudoku) copy() Sudoku {
	var newSudoku Sudoku

	for i := range a {
		copy(newSudoku[i][:], a[i][:])
	}
	return newSudoku
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
	copySudoku := a.copy()

	for y := range copySudoku {
		for x := range copySudoku {
			a[y][x] = copySudoku[x][y]
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

func (a Sudoku) checkValueColum(x, y, value int) bool {
	a.transposition()

	counter := 0
	for i := range a {
		if a[x][i].data == value {
			counter++
		}
	}
	return counter == 0
}

func (a Sudoku) checkValueLine(x, y, value int) bool {
	counter := 0
	for i := range a {
		if a[y][i].data == value {
			counter++
		}
	}
	return counter == 0
}

func (a Sudoku) checkValueArea(x, y, value int) bool {
	yArea, xArea := int(math.Floor(float64(y)/3)*3), int(math.Floor(float64(x)/3)*3)

	counter := 0
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if a[yArea+y][xArea+x].data == value {
				counter++
			}
		}
	}
	return counter == 0
}
func (a Sudoku) checkValue(x, y, value int) bool {
	if a.checkValueColum(x, y, value) && a.checkValueLine(x, y, value) && a.checkValueArea(x, y, value) {
		return true
	}
	return false
}

func (a *Sudoku) generationSolveSudoku() bool {
	for y := range a {
		for x := range a[y] {
			if a[y][x].data == 0 {
				for num := 1; num <= 9; num++ {
					if a.checkValue(x, y, num) {
						a[y][x].data = num

						if a.generationSolveSudoku() {
							return true
						} else {
							a[y][x].data = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func (a Sudoku) solveSudoku() (bool, Sudoku) {
	answerSudoku := a.copy()
	if !answerSudoku.generationSolveSudoku() {
		return false, answerSudoku
	}
	return true, answerSudoku
}

func (a *Sudoku) createBaseGameSudoku() {
	a.createSudokuBase()
	for i := 0; i < 10; i++ {
		switch rand.Intn(5) {

		case 0:
			a.transposition()
		case 1:
			a.swapRowsArea()
		case 2:
			a.swapRowsSmall()
		case 3:
			a.swapColumsSmall()
		case 4:
			a.swapColumsArea()
		}
	}
}

func (a *Sudoku) generationGameSudoku(difficult int) {
	a.createBaseGameSudoku()
	for difficult != 0 {
		x, y := rand.Intn(9), rand.Intn(9)
		old_data := a[y][x].data
		a[y][x] = SudokuCell{
			data:   0,
			access: true,
		}
		flag, _ := a.solveSudoku()
		if flag {
			difficult--
		} else {
			a[y][x] = SudokuCell{
				data:   old_data,
				access: false,
			}
		}
	}
}

func newGameSudoku(difficult int) Sudoku {
	var game Sudoku
	switch difficult {
	case 0:
		difficult = 41
	case 1:
		difficult = 56
	case 2:
		difficult = 66
	}
	game.generationGameSudoku(difficult)
	return game
}

func main() {
	flag := true
	flagInsert := false
	cordPlayer := []int{0, 0}
	var game Sudoku
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("       Press n for a start game       ")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			return true, nil // Stop listener by returning true on Ctrl+C
		}
		if flag {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("       Press n for a start game       ")
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
		if key.String() == "n" {
			flagInsert = false
			flag = false
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			game = newGameSudoku(2)
			game.returnSudokuByCordsPlayer(cordPlayer[0], cordPlayer[1])
			fmt.Println("         insert mode off          ")
			fmt.Println("          n - New game            ")
			fmt.Println("i - Insert number 1-9, 0 - nothing")
			fmt.Println("        ←, →, ↑, ↓ - moving       ")

		}

		if key.String() == "up" && !flag && !flagInsert {
			cordPlayer[1]--
			if cordPlayer[1] < 0 {
				cordPlayer[1] += 9
			}
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			game.returnSudokuByCordsPlayer(cordPlayer[0], cordPlayer[1])
			fmt.Println("         insert mode off          ")
			fmt.Println("          n - New game            ")
			fmt.Println("i - Insert number 1-9, 0 - nothing")
			fmt.Println("        ←, →, ↑, ↓ - moving       ")
		}
		if key.String() == "down" && !flag && !flagInsert {
			cordPlayer[1]++
			if cordPlayer[1] > 8 {
				cordPlayer[1] -= 9
			}
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			game.returnSudokuByCordsPlayer(cordPlayer[0], cordPlayer[1])
			fmt.Println("         insert mode off          ")
			fmt.Println("          n - New game            ")
			fmt.Println("i - Insert number 1-9, 0 - nothing")
			fmt.Println("        ←, →, ↑, ↓ - moving       ")
		}
		if key.String() == "left" && !flag && !flagInsert {
			cordPlayer[0]--
			if cordPlayer[0] < 0 {
				cordPlayer[0] += 9
			}
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			game.returnSudokuByCordsPlayer(cordPlayer[0], cordPlayer[1])
			fmt.Println("         insert mode off          ")
			fmt.Println("          n - New game            ")
			fmt.Println("i - Insert number 1-9, 0 - nothing")
			fmt.Println("        ←, →, ↑, ↓ - moving       ")
		}
		if key.String() == "right" && !flag && !flagInsert {
			cordPlayer[0]++
			if cordPlayer[0] > 8 {
				cordPlayer[0] -= 9
			}
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			game.returnSudokuByCordsPlayer(cordPlayer[0], cordPlayer[1])
			fmt.Println("         insert mode off          ")
			fmt.Println("          n - New game            ")
			fmt.Println("i - Insert number 1-9, 0 - nothing")
			fmt.Println("        ←, →, ↑, ↓ - moving       ")
		}

		if key.String() == "i" && !flag {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			game.returnSudokuByCordsPlayer(cordPlayer[0], cordPlayer[1])
			if !flagInsert {
				flagInsert = true
				fmt.Println("         insert mode on           ")
			} else {
				flagInsert = false
				fmt.Println("         insert mode off          ")
			}
			fmt.Println("          n - New game            ")
			fmt.Println("i - Insert number 1-9, 0 - nothing")
			fmt.Println("        ←, →, ↑, ↓ - moving       ")
			return false, nil
		}

		if (key.String() == "0" || key.String() == "1" || key.String() == "2" || key.String() == "3" || key.String() == "4" || key.String() == "5" || key.String() == "6" || key.String() == "7" || key.String() == "8" || key.String() == "9") && flagInsert {
			value := key.String()
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			num, _ := strconv.Atoi(value)

			if game.setCellUser(cordPlayer[0], cordPlayer[1], num) {
				fmt.Println()
				fmt.Println()
				fmt.Println()
				fmt.Println()
				fmt.Println()
				fmt.Println("     The number cannot be changed     ")
				fmt.Println("             press enter              ")
				fmt.Println()
				fmt.Println()
				fmt.Println()
				fmt.Println()
			} else {
				game.returnSudokuByCordsPlayer(cordPlayer[0], cordPlayer[1])
				fmt.Println("         insert mode off          ")
				fmt.Println("          n - New game            ")
				fmt.Println("i - Insert number 1-9, 0 - nothing")
				fmt.Println("        ←, →, ↑, ↓ - moving       ")
				flagInsert = false
			}
		}

		if key.String() == "enter" && !flag && flagInsert {

			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			game.returnSudokuByCordsPlayer(cordPlayer[0], cordPlayer[1])
			fmt.Println("          n - New game            ")
			fmt.Println("i - Insert number 1-9, 0 - nothing")
			fmt.Println("        ←, →, ↑, ↓ - moving       ")
		}
		// fmt.Println("\r" + key.String()) // Print every key press
		return false, nil // Return false to continue listening
	})
}
