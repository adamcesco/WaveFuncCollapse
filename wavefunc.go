package main

import (
	"fmt"
	"time"
	wfcp2 "wavefunc.go/wfcp"
)

func main() {
	airCell := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{},
		ColExclusions: []wfcp2.CellElementer{},
		Data: CellWrap{
			Data: " ",
		},
	}

	aCell := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: "a"}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: "z"}},
		Data: CellWrap{
			Data: "a",
		},
	}

	bCell := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: "b"}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: "y"}},
		Data: CellWrap{
			Data: "b",
		},
	}

	cCell := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: "c"}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: "x"}},
		Data: CellWrap{
			Data: "c",
		},
	}

	zCell := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: "a"}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: "z"}},
		Data: CellWrap{
			Data: "z",
		},
	}

	yCell := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: "b"}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: "y"}},
		Data: CellWrap{
			Data: "y",
		},
	}

	xCell := wfcp2.CellElement{
		RowExclusions: []wfcp2.CellElementer{CellWrap{Data: "c"}},
		ColExclusions: []wfcp2.CellElementer{CellWrap{Data: "x"}},
		Data: CellWrap{
			Data: "x",
		},
	}

	board := wfcp2.Board{
		MaxRange:  10,
		MaxDomain: 10,
		Corpus:    []wfcp2.CellElement{aCell, bCell, cCell, zCell, yCell, xCell, airCell},
	}
	board.EmptyBoard()
	er := board.ValidateBoard()
	fmt.Printf("Error 1: %v\n", er)

	er = board.InsertAt(5, 5, aCell)
	fmt.Printf("Error 2: %v\n", er)

	er = board.InsertAt(5, 4, zCell)
	fmt.Printf("Error 3: %v\n", er)

	er = board.ValidateBoard()
	fmt.Printf("Error 4: %v\n", er)

	er = board.ValidateBoard()
	fmt.Printf("Error 5: %v\n", er)

	board.Print()
	fmt.Println("----------------------------------")

	for i := 0; i < 10; i++ {
		board.EmptyBoard()
		board.Solve()
		board.Print()
		fmt.Println("----------------------------------")
		time.Sleep(2 * time.Second)
	}
}

type CellWrap struct {
	Data string
}

func (i CellWrap) EqualFlags(p wfcp2.CellElementer) bool {
	return i == p
}

func (i CellWrap) Print() {
	fmt.Print(" ", i.Data, " ")
}
