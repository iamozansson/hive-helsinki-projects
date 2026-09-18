package sprint

func Countdown(n int) string {
	x := ""
	if n < 1 {
	return "0!"
	}
	for i := n; i > 0; i -=2{
	if x != ""{
		x += ", "
	}
	x += string(rune(i + '0'))
	}
	x += ", 0!"
	return x
}
