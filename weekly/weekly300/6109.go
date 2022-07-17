package weekly

import "fmt"

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const N = 1000000007
	var know, forgot = make([]int, n), make([]int, n)
	know[0] = 1
	forgot[0] = 0
	for i := 1; i < n; i++ {
		var s = 0
		if i >= forget {
			forgot[i] = know[i-forget]
		}
		for j := i - forget + 1; j < i; j++ {
			if j < 0 {
				continue
			}
			//
			if i-j == delay { //告诉其他人
				s = (s + (know[j]-forgot[i])*2) % N
			} else {
				s = (s + know[j] - forgot[j]) % N
			}
		}
		know[i] = s
	}
	fmt.Println(forgot)
	fmt.Println(know)
	return know[n-1]
}
