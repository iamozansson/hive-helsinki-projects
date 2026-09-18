package sprint

func ShiftBy(r rune, step int) rune {
	return rune((int(r-'a')+step)%26 + 'a')
}
