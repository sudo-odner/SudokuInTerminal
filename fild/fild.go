package fild

import (
	"math"
	"math/rand"
	"sudoku/fild/cell"
)

type Fild struct {
	fild      [9][9]cell.Cell
	different float64
}

func (f Fild) Init(different float64) (bool, Fild) {
	if different < 0 || different > 1 {
		return false, Fild{}
	}
	f.createBaseFild()
	f.createMixFild()
	f.createPuzzleFild(different)
	f.different = different
	return true, f
}

func (f Fild) Copy() Fild {
	var newFild Fild

	for i := range f.fild {
		copy(newFild.fild[i][:], f.fild[i][:])
	}
	return newFild
}

func (f *Fild) createBaseFild() {
	var (
		shapeMiniFild = 3
		cellFild      cell.Cell
		num           int8
	)

	for y := 0; y < (shapeMiniFild * shapeMiniFild); y++ {
		for x := 0; x < (shapeMiniFild * shapeMiniFild); x++ {
			cellFild = cell.Cell{}
			num = int8((y*shapeMiniFild+y/shapeMiniFild+x)%(shapeMiniFild*shapeMiniFild) + 1)
			f.fild[y][x] = cellFild.Init(num, true)
		}
	}
}

func (f *Fild) createMixFild() {
	for i := 0; i < 10; i++ {
		switch rand.Intn(5) {

		case 0:
			f.transposition()
		case 1:
			f.swapRowsArea()
		case 2:
			f.swapRowsSmall()
		case 3:
			f.swapColumsSmall()
		case 4:
			f.swapColumsArea()
		}
	}
}

func (f *Fild) transposition() {
	var (
		copyFild     = f.Copy()
		copyFildFild = copyFild.fild
		fildFild     = f.fild
	)

	for y := range copyFildFild {
		for x := range copyFildFild {
			fildFild[y][x] = copyFildFild[x][y]
		}
	}
}

func (f *Fild) swapRowsSmall() {
	var (
		fild                  = f.fild
		indexArea             = rand.Intn(3)
		indexFirstLineInArea  = rand.Intn(3)
		indexSecondLineInArea = rand.Intn(3)
	)

	for indexFirstLineInArea == indexSecondLineInArea {
		indexSecondLineInArea = rand.Intn(3)
	}

	indexFirstLine, indexSecondLine := indexArea*3+indexFirstLineInArea, indexArea*3+indexSecondLineInArea

	fild[indexFirstLine], fild[indexSecondLine] = fild[indexSecondLine], fild[indexFirstLine]
}

func (f *Fild) swapColumsSmall() {
	f.transposition()
	f.swapRowsSmall()
	f.transposition()
}

func (f *Fild) swapRowsArea() {
	var (
		fild            = f.fild
		indexFirstArea  = rand.Intn(3)
		indexSecondArea = rand.Intn(3)
	)

	for indexFirstArea == indexSecondArea {
		indexSecondArea = rand.Intn(3)
	}

	indexFirstArea, indexSecondArea = indexFirstArea*3, indexSecondArea*3

	for i := 0; i < 3; i++ {
		fild[indexFirstArea+i], fild[indexSecondArea+i] = fild[indexSecondArea+i], fild[indexFirstArea+i]
	}
}

func (f *Fild) swapColumsArea() {
	f.transposition()
	f.swapRowsArea()
	f.transposition()
}

func (f *Fild) createPuzzleFild(difficult float64) {
	var (
		countFild        = float64(81.0)
		percentEmptyCell = int(math.Round(countFild * difficult))
		fild             = f.fild
		randX            int
		randY            int
	)

	for percentEmptyCell != 0 {
		randX, randY = rand.Intn(9), rand.Intn(9)
		oldValueCell, _ := fild[randY][randX].GetCell()

		fild[randY][randX].SetCell(0, true)

		flag, _ := f.Soluve()
		if flag {
			difficult--
		} else {
			fild[randY][randX].SetCell(oldValueCell, false)
		}
	}
}

func (f Fild) checkInsertValue(x, y, value int8) bool {
	return f.checkInsetValueInLine(x, y, value) && f.checkInsertValueInColum(x, y, value) && f.checkInsertValueInAMiniFild(x, y, value)
}

func (f Fild) checkInsetValueInLine(x, y, value int8) bool {
	var (
		counter = 0
		fild    = f.fild
	)

	for i := range fild {
		valueCell, _ := fild[y][i].GetCell()
		if valueCell == value {
			counter++
		}
	}
	return counter == 0
}

func (f Fild) checkInsertValueInColum(x, y, value int8) bool {
	f.transposition()

	var (
		counter = 0
		fild    = f.fild
	)

	for i := range fild {
		valueCell, _ := fild[x][i].GetCell()
		if valueCell == value {
			counter++
		}
	}
	return counter == 0
}

func (f Fild) checkInsertValueInAMiniFild(x, y, value int8) bool {
	var (
		miniFildY = int(math.Floor(float64(y)/3) * 3)
		miniFildX = int(math.Floor(float64(x)/3) * 3)
		counter   = 0
		fild      = f.fild
	)

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			valueCell, _ := fild[y+miniFildY][x+miniFildX].GetCell()
			if valueCell == value {
				counter++
			}
		}
	}
	return counter == 0
}

func (f Fild) Soluve() (bool, Fild) {
	var (
		soluve     = f.Copy()
		soluveFild = soluve.fild
	)

	for y := range soluveFild {
		for x := range soluveFild[y] {
			_, accessCell := soluveFild[y][x].GetCell()
			if !accessCell {
				soluveFild[y][x].SetCell(0, true)
			}
		}
	}

	var err = soluve.generateSoluve()
	return err, soluve
}

func (f *Fild) generateSoluve() bool {
	var fild = f.fild
	for y := range fild {
		for x := range fild[y] {
			valueCell, _ := fild[y][x].GetCell()
			if valueCell == 0 {
				for num := 1; num <= 9; num++ {
					if f.checkInsertValue(int8(x), int8(y), int8(num)) {
						fild[y][x].SetCell(int8(num), true)

						if f.generateSoluve() {
							return true
						} else {
							fild[y][x].SetCell(0, true)
						}
					}
				}
				return false
			}
		}
	}
	return true
}
