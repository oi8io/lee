package cn

import (
	"strconv"
	"strings"
	"testing"
)

func TestConstructor3(t *testing.T) {
	c := []string{"put", "put", "get", "get", "put", "get", "remove", "get"}
	d := [][]int{{1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2}}
	es := "null,null,1,-1,null,1,null,-1"
	e := strings.Split(es, ",")
	m := Constructor()
	for i, _ := range c {
		exp, _ := strconv.Atoi(e[i])
		k := d[i]
		switch c[i] {
		case "put":
			m.Put(k[0], k[1])
		case "remove":
			m.Remove(k[0])
		case "get":
			got := m.Get(k[0])
			if got != exp {
				t.Errorf("GET %d want %d but %d", k[0], exp, got)
			}
		}
	}
}
