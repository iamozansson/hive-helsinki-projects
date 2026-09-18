package sprint

func ToCapitalCase(s string) string {
	newWord := true
	var result []rune
	for _, v := range s {
		if v >= 'a' && v <= 'z' ||
			v >= 'A' && v <= 'Z' ||
			v >= '0' && v <= '9' {
			if newWord {
				if v >= 'a' && v <= 'z' {
					v = v - 'a' + 'A'
				}
			} else if v >= 'A' && v <= 'Z' {
				v = v + 'a' - 'A'
			}
			result = append(result, v)
			newWord = false
		} else {
			result = append(result, v)
			newWord = true
		}
	}
	return string(result)
}
