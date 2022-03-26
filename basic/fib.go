package basic

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
func Fib1(n int) int {
	var pre, result int

	for i := 0; i <= n; i++ {
		if i <= 1 {
			result = i
		} else {
			tmp := result
			result = tmp + pre
			pre = tmp
		}
	}
	return result
}
