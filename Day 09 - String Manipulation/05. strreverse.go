package sprint

func StrReverse(s string) string {
	runes := []rune(s)
	newS := ""
	for i := len(runes) - 1; i >= 0; i-- {
		newS += string(runes[i])
	}
	return newS
}
