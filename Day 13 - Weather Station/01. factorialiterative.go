package sprint

func FactorialIterative(n int) int {
	num := 1
	for i := 1; i <= n; i++ {
		num *= i
	}
	return num
}
