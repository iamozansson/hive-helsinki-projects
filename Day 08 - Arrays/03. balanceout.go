package sprint

func BalanceOut(arr []bool) []bool {
	tTrue := 0
	fFalse := 0
	for _, counter := range arr {
		if counter == true {
			tTrue += 1
		}
		if counter == false {
			fFalse += 1
		}
	}
	var added bool
	if tTrue < fFalse {
		added = true
	} else {
		added = false
	}
	for i := min(tTrue, fFalse); i < max(tTrue, fFalse); i++ {
		arr = append(arr, added)
	}
	return arr
}
