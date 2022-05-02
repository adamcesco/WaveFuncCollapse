package wfcp

import (
	"errors"
	"fmt"
)

type CellElementer interface {
	EqualFlags(CellElementer) bool
	Print()
} // each type that enacts cellElement needs have a unique flag that can be compared

type CellElement struct {
	RowExclusions []CellElementer
	ColExclusions []CellElementer
	Data          CellElementer
}

type Board struct {
	MaxRange  int
	MaxDomain int
	data      [][][]CellElement //D1 = X, D2 = Y, D3, = options of that cell
	Corpus    []CellElement
}

func (b *Board) EmptyBoard() error {
	corpusSize := len(b.Corpus)
	if (b.MaxRange == 0) || (b.MaxDomain == 0) || (corpusSize == 0) {
		return errors.New("board has not been initialized correctly | corpus, domain, or range may not be initialized")
	}

	b.data = make([][][]CellElement, b.MaxRange)
	for i := 0; i < b.MaxRange; i++ {
		b.data[i] = make([][]CellElement, b.MaxDomain)
		for j := 0; j < b.MaxDomain; j++ {
			b.data[i][j] = make([]CellElement, len(b.Corpus))
			copy(b.data[i][j], b.Corpus)
		}
	}
	return nil
}

func contains(arr *[]CellElementer, ele *CellElement) bool {
	for _, element := range *arr {
		if ele.Data.EqualFlags(element) {
			return true
		}
	}
	return false
}

func (b *Board) ValidateBoard() error {
	corpusSize := len(b.Corpus)
	if (b.MaxRange == 0) || (b.MaxDomain == 0) || (corpusSize == 0) {
		return errors.New("board has not been initialized correctly | corpus, domain, or range may not be initialized")
	}

	for i := 0; i < b.MaxRange; i++ {
		for j := 0; j < b.MaxDomain; j++ {
			if len(b.data[i][j]) == 1 {
				for k := j + 1; k < j+b.MaxDomain; k++ {
					index := k % b.MaxDomain
					if len(b.data[i][index]) == 1 && contains(&b.data[i][j][0].RowExclusions, &b.data[i][index][0]) {
						return errors.New("board is unvalidated | there is a misplaced cell")
					} else if len(b.data[i][index]) < 1 {
						return errors.New("board is unvalidated | there is a cell with no options")
					}
				}

				for k := i + 1; k < i+b.MaxRange; k++ {
					index := k % b.MaxRange
					if len(b.data[index][j]) == 1 && contains(&b.data[i][j][0].ColExclusions, &b.data[index][j][0]) {
						return errors.New("board is unvalidated | there is a misplaced cell")
					} else if len(b.data[index][j]) < 1 {
						return errors.New("board is unvalidated | there is a cell with no options")
					}
				}
			}
		}
	}
	return nil
}

func remove(s []CellElement, i int) []CellElement {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func at(arr *[]CellElement, ele *CellElementer) (int, error) {
	for i, element := range *arr {
		if element.Data.EqualFlags(*ele) {
			return i, nil
		}
	}
	return -1, errors.New("passed CellElementer was not found in passed []CellElement")
}

func (b *Board) InsertAt(x, y int, ele CellElement) error {
	corpusSize := len(b.Corpus)
	if (y > b.MaxRange) || (x > b.MaxDomain) || (corpusSize == 0) {
		return errors.New("incorrect x & y values, or corpus may not be initialized")
	}

	//validating
	for i := x + 1; i < x+b.MaxDomain; i++ {
		xIndex := i % b.MaxDomain
		if len(b.data[y][xIndex]) == 1 && (contains(&ele.RowExclusions, &b.data[y][xIndex][0]) || contains(&b.data[y][xIndex][0].RowExclusions, &ele)) {
			return errors.New("passed cell was not placed | horizontal-misplacement of the passed cell")
		}
	}
	for i := y + 1; i < y+b.MaxRange; i++ {
		yIndex := i % b.MaxRange
		if len(b.data[yIndex][x]) == 1 && (contains(&ele.ColExclusions, &b.data[yIndex][x][0]) || contains(&b.data[yIndex][x][0].ColExclusions, &ele)) {
			return errors.New("passed cell was not placed | vertical-misplacement of the passed cell")
		}
	}

	b.data[y][x] = []CellElement{ele}

	//placing and updating options
	for i := x + 1; i < x+b.MaxDomain; i++ {
		xIndex := i % b.MaxDomain
		if len(b.data[y][xIndex]) > 1 {
			//remove all row exclusions from b.data[y][xIndex]
			for _, it := range ele.RowExclusions {
				subscr, er := at(&b.data[y][xIndex], &it)
				if er == nil {
					b.data[y][xIndex] = remove(b.data[y][xIndex], subscr)
				}
			}
		}

	}
	for i := y + 1; i < y+b.MaxRange; i++ {
		yIndex := i % b.MaxRange
		if len(b.data[yIndex][x]) > 1 {
			for _, it := range ele.ColExclusions {
				subscr, er := at(&b.data[yIndex][x], &it)
				if er == nil {
					b.data[yIndex][x] = remove(b.data[yIndex][x], subscr)
				}
			}
		}
	}

	return nil
}

func (b *Board) Print() {
	for _, i2 := range b.data {
		for _, j3 := range i2 {
			//if len(j3) > 1 {
			fmt.Print(len(j3), " ")
			//} else {
			//	j3[0].Data.Print()
			//	fmt.Print(" ")
			//}
		}
		fmt.Println()
	}
}
