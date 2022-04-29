package main

import (
	"fmt"
	wfcp2 "wavefunc.go/wfcp"
)

func main() {
	fmt.Println("Hello World")
	v1 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: -1}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -1}},
		Data: CellWrap{
			Data: 1,
		},
	}

	v2 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 2}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -2}},
		Data: CellWrap{
			Data: 2,
		},
	}

	v3 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 3}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -3}},
		Data: CellWrap{
			Data: 3,
		},
	}

	v4 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 1}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -1}},
		Data: CellWrap{
			Data: -1,
		},
	}

	v5 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 2}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -2}},
		Data: CellWrap{
			Data: -2,
		},
	}

	v6 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 3}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -3}},
		Data: CellWrap{
			Data: -3,
		},
	}

	board := wfcp2.Board{
		MaxRange:  10,
		MaxDomain: 10,
		Corpus:    []wfcp2.CellElement{v1, v2, v3, v4, v5, v6},
	}
	er := board.EmptyBoard()
	fmt.Printf("Error: %v", er)
}

type CellWrap struct {
	Data int
}

func (i CellWrap) EqualFlags(p wfcp2.CellElementer) bool {
	return i == p
}
