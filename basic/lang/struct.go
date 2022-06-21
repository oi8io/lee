package lang

import "fmt"

type tester struct {
	s string
}

func (t tester) Print(s string) {
	fmt.Println(t.s, s)
}

func (t *tester) Print1(s string) {
	fmt.Println(t.s, s)
}
