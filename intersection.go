package main

import (
	"fmt"
)

/*
iki noktası bilinen bir doğru üzerinde "X" koordinatı verilen bir noktanın "Y"
koordinatını bulur.

iki noktası bilinen bir doğrunun dekleminden türetilmiştir.

y2-y1    y-y2
----- = ------
x2-x1    x-x2
*/
func findX(x1, y1, x2, y2, y int) int {
	return ((x2 * y) - (x1 * y) + (x1 * y2) - (y1 * x2)) / (y2 - y1)
}

/*
iki noktası bilinen bir doğru üzerinde "Y" koordinatı verilen bir noktanın "X"
koordinatını bulur.

iki noktası bilinen bir doğrunun dekleminden türetilmiştir.

y2-y1    y-y2
----- = ------
x2-x1    x-x2
*/
func findY(x1, y1, x2, y2, x int) int {
	return ((y2 * x) - (y1 * x) + (y1 * x2) - (x1 * y2)) / (x2 - x1)
}

func main() {
	//x := findX(2, 2, 20, 14, 6)
	//y := findY(2, 2, 20, 14, 8)
	//fmt.Println("x=", x, " y=", y)

	//x := findX(5, 21, 11, 13, 6)
	//y := findY(5, 21, 11, 13, 11)
	y := findY(11, 13, 5, 21, 11)
	fmt.Println("y=", y)

	// x ekseninen paralen bir doğu denemesi
	y = findY(0, 0, 9, 9, 11)
	fmt.Println("y=", y)

}
