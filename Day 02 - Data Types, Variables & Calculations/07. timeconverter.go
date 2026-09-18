package sprint

func TimeConverter(totalSeconds int) (int, int, int) {
	h := totalSeconds / 3600
	m := (totalSeconds % 3600) / 60
	s := totalSeconds % 60
	return h, m, s
}



