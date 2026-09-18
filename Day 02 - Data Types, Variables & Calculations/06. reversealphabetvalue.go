package sprint

func ReverseAlphabetValue(ch rune) rune{
	reverse := int(ch) - 'a'
	return rune(25 - reverse) + 97
}

