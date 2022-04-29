package main

import (
	"fmt"
	wfcp2 "wavefunc.go/wfcp"
)

func main() {
	fmt.Println("Hello World")

	v1 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 1}},
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

	vn1 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 1}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -1}},
		Data: CellWrap{
			Data: -1,
		},
	}

	vn2 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 2}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -2}},
		Data: CellWrap{
			Data: -2,
		},
	}

	vn3 := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: 3}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: -3}},
		Data: CellWrap{
			Data: -3,
		},
	}

	board := wfcp2.Board{
		MaxRange:  10,
		MaxDomain: 10,
		Corpus:    []wfcp2.CellElement{v1, v2, v3, vn1, vn2, vn3},
	}
	board.EmptyBoard()
	er := board.ValidateBoard()
	fmt.Printf("Error 1: %v\n", er)

	er = board.InsertAt(5, 5, v1)
	fmt.Printf("Error 2: %v\n", er)

	er = board.InsertAt(5, 4, vn1)
	fmt.Printf("Error 3: %v\n", er)

	er = board.ValidateBoard()
	fmt.Printf("Error 4: %v\n", er)
}

type CellWrap struct {
	Data int
}

func (i CellWrap) EqualFlags(p wfcp2.CellElementer) bool {
	return i == p
}
