package main

import (
	"fmt"
	"math"
)

func raytrace_integer(x0, y0, x1, y1 int) {
	dx := math.Abs(float64(x1 - x0))
	dy := math.Abs(float64(y1 - y0))
	x := x0
	y := y0
	n := 0 + dx + dy

	fmt.Println(n)

	x_inc := 0
	y_inc := 0

	if x1 > x0 {
		x_inc = 1
	} else {
		x_inc = -1
	}

	if y1 > y0 {
		y_inc = 1
	} else {
		y_inc = -1
	}

	error_value := dx - dy

	dx *= 2
	dy *= 2

	for n > 0 {
		n--
		fmt.Printf("x = %d, y= %d\n", x, y)

		if error_value > 0 {
			x += x_inc
			error_value -= dy
		} else {
			y += y_inc
			error_value += dx
		}
	}
}

func raytrace(x0, y0, x1, y1 float64) {
	dx := math.Abs(x1 - x0)
	dy := math.Abs(y1 - y0)

	x := int(math.Floor(x0))
	y := int(math.Floor(y0))

	var n int = 1
	var x_inc, y_inc int

	var error_value float64

	if dx == 0 {
		x_inc = 0
		error_value = math.Inf(1)
		//error_value = std::numeric_limits<double>::infinity()
	} else if x1 > x0 {
		x_inc = 1
		n += int(math.Floor(x1)) - x
		error_value = (math.Floor(x0) + 1 - x0) * dy
	} else {
		x_inc = -1
		n += x - int(math.Floor(x1))
		error_value = (x0 - math.Floor(x0)) * dy
	}

	if dy == 0 {
		y_inc = 0
		error_value = math.Inf(1)
		//error_value -= std::numeric_limits<double>::infinity();
	} else if y1 > y0 {
		y_inc = 1
		n += int(math.Floor(y1)) - y
		error_value -= (math.Floor(y0) + 1 - y0) * dx
	} else {
		y_inc = -1
		n += y - int(math.Floor(y1))
		error_value -= (y0 - math.Floor(y0)) * dx
	}

	for n > 0 {
		//fmt.Println("error:", error_value)
		n--
		fmt.Println(x, y)

		if error_value > 0 {
			y += y_inc
			error_value -= dx
		} else {
			x += x_inc
			error_value += dy
		}
	}
}
func main() {
	raytrace_integer(0, 0, 10, 5)
	raytrace(30.5, 40.5, 35.5, 43.25)
	fmt.Println("...")
	raytrace(30.0, 40.0, 40, 50)
}
