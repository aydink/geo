package main

// Generalized with integer

type Point struct {
	X, Y int
}

func Bresenham(x1, y1, x2, y2 int) []Point {

	points := make([]Point, 0)

	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		//img.Set(x1, y1, col)
		//fmt.Printf("(%d, %d)", x1, y1)
		points = append(points, Point{x1, y1})

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			//img.Set(x1, y1, col)
			//fmt.Printf("(%d, %d)", x1, y1)
			points = append(points, Point{x1, y1})
			x1++
		}
		//img.Set(x1, y1, col)
		//fmt.Printf("(%d, %d)", x1, y1)
		points = append(points, Point{x1, y1})

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			//img.Set(x1, y1, col)
			//fmt.Printf("(%d, %d)", x1, y1)
			points = append(points, Point{x1, y1})
			y1++
		}
		//img.Set(x1, y1, col)
		//fmt.Printf("(%d, %d)", x1, y1)
		points = append(points, Point{x1, y1})

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				//img.Set(x1, y1, col)
				//fmt.Printf("(%d, %d)", x1, y1)
				points = append(points, Point{x1, y1})
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				//img.Set(x1, y1, col)
				//fmt.Printf("(%d, %d)", x1, y1)
				points = append(points, Point{x1, y1})
				x1++
				y1--
			}
		}
		//img.Set(x1, y1, col)
		//fmt.Printf("(%d, %d)", x1, y1)
		points = append(points, Point{x1, y1})

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				//img.Set(x1, y1, col)
				//fmt.Printf("(%d, %d)", x1, y1)
				points = append(points, Point{x1, y1})
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				//img.Set(x1, y1, col)
				//fmt.Printf("(%d, %d)", x1, y1)
				points = append(points, Point{x1, y1})
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		//img.Set(x2, y2, col)
		//fmt.Printf("(%d, %d)", x2, y2)
		points = append(points, Point{x2, y2})

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				//img.Set(x1, y1, col)
				//fmt.Printf("(%d, %d)", x1, y1)
				points = append(points, Point{x1, y1})
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				//img.Set(x1, y1, col)
				//fmt.Printf("(%d, %d)", x1, y1)
				points = append(points, Point{x1, y1})
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		//img.Set(x2, y2, col)
		//fmt.Printf("(%d, %d)", x2, y2)
		points = append(points, Point{x2, y2})
	}

	return points
}
