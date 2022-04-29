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
	if (b.MaxRange == 0) || (b.MaxDomain == 0 || corpusSize == 0) {
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

func conatins(arr []CellElement, ele CellElement) bool {
	for _, element := range arr {
		if ele.Data == element.Data {
			return true
		}
	}
	return false
}

func (b *Board) ValidateBoard() error {
	corpusSize := len(b.Corpus)
	if (b.MaxRange == 0) || (b.MaxDomain == 0 || corpusSize == 0) {
		return errors.New("board has not been initialized correctly | corpus, domain, or range may not be initialized")
	}

	for i := 0; i < b.MaxRange; i++ {
		for j := 0; j < b.MaxDomain; j++ {
			if len(b.data[i][j]) == 1 {
				for k := j; k < j+b.MaxDomain; k++ {
					index := k % b.MaxDomain
					if len(b.data[i][index]) == 1 && conatins(b.data[i][j], b.data[i][index][0]) {
						return errors.New("board has is unvalidated | there is a misplaced cell")
					}
				}

				for k := i; k < i+b.MaxRange; k++ {
					index := k % b.MaxRange
					if len(b.data[index][j]) == 1 && conatins(b.data[i][j], b.data[index][j][0]) {
						return errors.New("board has is unvalidated | there is a misplaced cell")
					}
				}
			}
		}
	}
	return nil
}

func (b *Board) InsertAt(x, y int) error {

	return nil
}
