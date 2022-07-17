package weekly

import (
	"fmt"
	"strings"
)

func decodeMessage(key string, message string) string {
	var rem = make(map[byte]byte)
	key = strings.ReplaceAll(key, " ", "")
	fmt.Println(key)
	var x = 0
	for i := 0; i < len(key); i++ {
		if _, ok := rem[key[i]]; !ok {
			rem[key[i]] = byte(x%26 + 'a')
			x++
		}
	}
	var answer = make([]byte, 0)
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			answer = append(answer, message[i])
		} else {
			answer = append(answer, rem[message[i]])
		}
	}
	return string(answer)
}
