package sprint

func ToUpperCase(s string) string {
	var result string
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			v = v - 'a' + 'A'
		}
		result += string(v)
	}
	return result
}
