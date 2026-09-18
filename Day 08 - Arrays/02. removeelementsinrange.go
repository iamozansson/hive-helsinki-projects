package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	if from > to {
		from, to = to, from
	}
	var newArr []float64
	for i, num := range arr {
		if i < from || i >= to {
			newArr = append(newArr, num)
		}
	}
	return newArr
}
