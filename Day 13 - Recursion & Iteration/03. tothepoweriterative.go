package sprint

func ToThePowerIterative(n int, power int) int {
	result := 1
	if n <= 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		result *= n
	}
	return result
}
