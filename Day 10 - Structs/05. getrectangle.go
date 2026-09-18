package sprint

type Coords struct {
	X int
	Y int
}

type Rectangle struct {
	Min       Coords
	Max       Coords
	Width     int
	Height    int
	Area      int
	Perimeter int
}

func GetRectangle(min, max Coords) Rectangle {
	width := max.X - min.X
	height := max.Y - min.Y
	area := width * height
	perimeter := 2 * (width + height)
	return Rectangle{
		Min:       min,
		Max:       max,
		Width:     width,
		Height:    height,
		Area:      area,
		Perimeter: perimeter,
	}
}
