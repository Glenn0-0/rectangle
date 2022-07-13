package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int

	fmt.Print("Enter base of the rectangle: ")
	fmt.Scanln(&a)

	fmt.Print("Enter side of the rectangle: ")
	fmt.Scanln(&b)

	var rec1 rectangle
	rec1 = rectangle.createRec(rec1, a, b)

	switchOptions(rec1)

	fmt.Println(rec1)
}

func switchOptions(rec1 rectangle) {
	option := chooseOption(rec1)

	switch option {

	case 1:
		fmt.Println(rec1.getArea())

	case 2:
		var n int
		fmt.Print("\nEnter a number by which to enlarge by: ")
		fmt.Scanln(&n)
		rec1.enlargeBy(n)

	case 3:
		var way int
		way = determineWay()
		rec1.recPrint(way)

	case 4:
		var a int
		var sq1 square
		fmt.Print("\nEnter side of the square: ")
		fmt.Scanln(&a)
		sq1 = square.createSq(sq1, a)
		fmt.Println(rec1.hasSquares(sq1))

	case 5:
		var n, m int
		fmt.Print("\nEnter base of the rectangle: ")
		fmt.Scanln(&n)

		fmt.Print("Enter side of the rectangle: ")
		fmt.Scanln(&m)
		rec1 = rec1.createRec(n, m)

	case 6:
		fmt.Println("\nNow, please, enter another rectangle's dimensions to compare with")

		var n, m int
		fmt.Print("Enter base of a new rectangle: ")
		fmt.Scanln(&n)

		fmt.Print("Enter side of a new rectangle: ")
		fmt.Scanln(&m)

		rec1.compareArea(n, m)

	case 7:
		fmt.Println("\nThank you for your time. Program exited.")
		os.Exit(0)
		
	}

	switchOptions(rec1)
}

func chooseOption(rec1 rectangle) (int) {
	var option int

	fmt.Println("\nChoose an operation:\n1 - Get rectangle's area\n2 - Enlarge the rectangle\n3 - Print the rectangle\n4 - Count the number of squares in the rectangle\n5 - Create new rectangle\n6 - Compare rectangle's area\n7 - Exit program")
	
	fmt.Print("Enter number: ")
	fmt.Scanln(&option)

	if option < 1 || option > 7 {
		fmt.Println(fmt.Errorf("for the rectangle %v: validation failed: invalid option value: expected %v to equal 1, 2, 3, 4, 5, 6 or 7", rec1, option))
		chooseOption(rec1)
	}

	return option
}

func determineWay() int {
	fmt.Println("\nWhich way would you like to print th rectangle?\n0 - vertically\n1 - horizontally")

	var way int
	fmt.Print("Enter option:")
	fmt.Scanln(&way)

	if way != 0 && way != 1 {
		fmt.Println(fmt.Errorf("for rectangle: printRec: validation failed: invalid option value: expected %v to equal 0 or 1", way))
		determineWay()
	}

	return way
}