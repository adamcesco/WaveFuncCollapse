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

func (b *Board) ValidateBoard() error {
	for i := 0; i < b.MaxRange; i++ {
		for j := 0; j < b.MaxDomain; j++ {

		}
	}
	return nil
}

func (b *Board) InsertAt(x, y int) error {

	return nil
}
