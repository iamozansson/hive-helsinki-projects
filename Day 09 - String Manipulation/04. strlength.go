package sprint

func StrLength(s string) []int {
	runes := []rune(s)
	return []int{len(runes), len(s)}
}
