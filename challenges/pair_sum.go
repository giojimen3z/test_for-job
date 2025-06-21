package challenges

func CheckPairSum(asc []int, desc []int, target int) int {
	count := 0
	i, j := 0, 0

	for i < len(asc) && j < len(desc) {
		sum := asc[i] + desc[j]

		switch {
		case sum == target:
			count++
			i++
			j++
		case sum < target:
			i++
		default:
			j++
		}
	}

	return count
}
