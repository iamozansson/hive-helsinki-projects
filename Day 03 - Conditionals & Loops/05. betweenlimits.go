package sprint

func BetweenLimits(from, to rune) string {
	var text string
	if from > to {
		from, to = to, from
	}
	for i := from +1; i < to ; i++ {
	text += string(i)
	}
	return text
}

