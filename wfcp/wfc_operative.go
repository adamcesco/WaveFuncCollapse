package wfcp

type CellElementer interface {
	EqualFlags(CellElementer) bool
} // each type that enacts cellElement needs have a unique flag that can be compared

type CellElement struct {
	rowOptions    []CellElementer
	colOptions    []CellElementer
	RowExclusions []CellElementer
	ColExclusions []CellElementer
	Data          CellElementer
}

type Board struct {
	Data   [][]CellElement
	Corpus []CellElement
}

func (b *Board) EmptyBoard(maxDomain, maxRange int) *Board {
	b.Data = make([][]CellElement, maxRange)
	for i := 0; i < maxRange; i++ {
		b.Data[i] = make([]CellElement, maxDomain)
	}
	return b
}

func (b *Board) ValidateBoard() *Board {

	return b
}
