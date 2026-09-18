package sprint

func StrConcatWith(strs []string, sep string) string {
	newArr := ""
	for i, v := range strs {
		if i > 0 {
			newArr += sep
		}
		newArr += v
	}
	return newArr
}
