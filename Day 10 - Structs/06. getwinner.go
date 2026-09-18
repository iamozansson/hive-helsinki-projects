package sprint

type Contestant struct {
	Name   string
	Scores []int
}

func GetWinner(c1, c2 Contestant) string {

	sum1 := 0
	sum2 := 0

	for _, score := range c1.Scores {
		sum1 += score
	}
	for _, score := range c2.Scores {
		sum2 += score
	}
	if sum1 > sum2 {
		return c1.Name
	}
	return c2.Name
}
