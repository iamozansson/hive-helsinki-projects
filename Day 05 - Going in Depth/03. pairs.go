package sprint

import "fmt"

var x string
func Pairs() string {
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if x != ""{
			x += ", "
			}
			x += fmt.Sprintf("%02d %02d", i, j)
	}
	}
	return x
}

