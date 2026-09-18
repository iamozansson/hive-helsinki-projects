package sprint

func Accumulate(n int) int {
	var count int
	for i := 0; i <= n; i++ {
		count += i
	}
	return count
}
