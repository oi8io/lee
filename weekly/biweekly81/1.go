/**

 */

package weekly

import (
	"strings"
)

func countAsterisks(s string) int {
	x := strings.Split(s, "|")
	var answer int
	for i := 0; i < len(x)-1; i = i + 2 {
		for j := 0; j < len(x[i]); j++ {
			if x[i][j] == '*' {
				answer++
			}
		}
	}
	return answer
}
