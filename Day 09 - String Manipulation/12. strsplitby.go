package sprint

func StrSplitBy(s, sep string) []string {
	total := []string{}
	start := 0

	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			total = append(total, s[start:i])
			i += len(sep) - 1
			start = i + 1
		}
	}
	if start < len(s) {
		total = append(total, s[start:])
	}
	return total
}
