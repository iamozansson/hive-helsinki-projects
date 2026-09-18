package sprint

import "fmt"

func PointText(p Point) Point {
	p.Text = fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return p
}
