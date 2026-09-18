package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	result := make([][]int, 0)
	for i := range arr {
		sum := 0
		for j := range arr[i] {
			sum += arr[i][j]
		}
		if limit <= sum {
			result = append(result, arr[i])
		}
	}
	return result
}
