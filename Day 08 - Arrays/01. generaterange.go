package sprint

func GenerateRange(min, max int) []int {
	length := max - min
	if min >= max {
		return nil
	}
	mySlice := make([]int, length)
	for i := 0; i < length; i++ {
		mySlice[i] += i + min
	}
	return mySlice
}
