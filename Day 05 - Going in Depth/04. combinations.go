package sprint

import "fmt"


func Combinations() string {
x := ""
	for i := 0; i < 10; i++{
		for j := i + 1 ; j < 10; j++{
			for k := j + 1; k < 10; k++{
			if x != ""{
			x += ", "
			}
			x += fmt.Sprintf("%d%d%d", i, j, k) 
			}
		}
	}
	return x
}
