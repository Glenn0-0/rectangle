package main

import "fmt"

type rectangle struct {
	base int
	side int
}

type square struct {
	side int
}

type recGetter interface {
	getArea() int
	enlargeBy(int)
	recPrint()
	hasSquares(square) int
	create(int, int)
}

func (rec rectangle) getArea() int {
	return rec.base * rec.side
}

func (rec *rectangle) enlargeBy(n int) {
	rec.base *= n
	rec.side *= n
}

func (rec rectangle) recPrint(way int) {
	if (way == 1 && rec.side > rec.base) || (way == 0 && rec.side < rec.base) {
		temp := rec.side
		rec.side = rec.base
		rec.base = temp
	}
	for i := 0; i < rec.side; i++ {
		for k := 0; k < rec.base; k++ {
			fmt.Print("O")
		}
		fmt.Print("\n")
	}
}

func (rec rectangle) hasSquares(sq square) (num int) {
	num = (rec.base - (sq.side - 1))*(rec.side - (sq.side - 1))
	return num
}

func (rec rectangle) createRec(a, b int) (figure rectangle) {
	figure = rectangle{
		base: a,
		side: b,
	}
	return figure
}

func (sq square) createSq(a int) (figure square) {
	figure = square {
		side: a,
	}
	return figure
}