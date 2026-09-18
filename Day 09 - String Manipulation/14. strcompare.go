package sprint

func StrCompare(a, b string) int {
	mLen := len(a)
	if len(b) < mLen {
		mLen = len(b)
	}
	for i := 0; i < mLen; i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}
