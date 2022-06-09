package weekly

func divisorSubstrings(num int, k int) int {
	var answers = 0
	var cal = 1
	for i := 1; i <= k; i++ {
		cal *= 10
	}
	var x = num
	for {
		if x < cal/10 {
			break
		}
		cur := x % cal
		if cur != 0 && num%cur == 0 {
			answers++
		}
		x = x / 10
	}
	return answers
}
