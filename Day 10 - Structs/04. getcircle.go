package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	return Circle{
		Radius:    r,
		Diameter:  2 * r,
		Area:      3.14 * r * r,
		Perimeter: 2 * 3.14 * r,
	}
}
