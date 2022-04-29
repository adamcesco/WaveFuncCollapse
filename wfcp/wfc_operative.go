package wfcp

import "errors"

type CellElementer interface {
	EqualFlags(CellElementer) bool
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
			b.data[i][j] = b.Corpus
		}
	}
	return nil
}

func conatins(arr *[]CellElementer, ele *CellElement) bool {
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
					if len(b.data[i][index]) == 1 && conatins(&b.data[i][j][0].RowExclusions, &b.data[i][index][0]) {
						return errors.New("board is unvalidated | there is a misplaced cell")
					}
				}

				for k := i + 1; k < i+b.MaxRange; k++ {
					index := k % b.MaxRange
					if len(b.data[index][j]) == 1 && conatins(&b.data[i][j][0].ColExclusions, &b.data[index][j][0]) {
						return errors.New("board is unvalidated | there is a misplaced cell")
					}
				}
			}
		}
	}
	return nil
}

func (b *Board) InsertAt(x, y int, ele CellElement) error {
	corpusSize := len(b.Corpus)
	if (y > b.MaxRange) || (x > b.MaxDomain) || (corpusSize == 0) {
		return errors.New("incorrect x & y values, or corpus may not be initialized")
	}

	//validating
	for i := x + 1; i < x+b.MaxDomain; i++ {
		index := i % b.MaxDomain
		if len(b.data[y][index]) == 1 && (conatins(&ele.RowExclusions, &b.data[y][index][0]) || conatins(&b.data[y][index][0].RowExclusions, &ele)) {
			return errors.New("passed cell was not placed | horizontal-misplacement of the passed cell")
		}
	}
	for i := y + 1; i < y+b.MaxRange; i++ {
		index := i % b.MaxRange
		if len(b.data[index][x]) == 1 && (conatins(&ele.ColExclusions, &b.data[index][x][0]) || conatins(&b.data[index][x][0].ColExclusions, &ele)) {
			return errors.New("passed cell was not placed | vertical-misplacement of the passed cell")
		}
	}

	b.data[y][x] = []CellElement{ele}

	//placing and updating options
	for i := x + 1; i < x+b.MaxDomain; i++ {
		index := i % b.MaxDomain
		if len(b.data[index][x]) > 1 {
			//for j, _ := range b.data[y][index] {
			//	//b.data[y][index][j]
			//}
		}

	}
	for i := y + 1; i < y+b.MaxRange; i++ {
		index := i % b.MaxRange
		if len(b.data[index][x]) > 1 {
			//for j, _ := range b.data[index][x] {
			//
			//}
		}
	}

	return nil
}
