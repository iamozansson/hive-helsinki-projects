package sprint

func AlphabetMastery(n int) string {
	counter := ""
	for i := 0; i < n; i++ {
		counter += string(rune('a' + i))
	}
	return counter
}
