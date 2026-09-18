package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}
	counter := ""
	for i := 'z'; i >= 'a'; i -= rune(step) {
		counter += string(i)
	}
	return counter
}
