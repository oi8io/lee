package cn

import (
	"container/list"
	"strconv"
	"strings"
	"testing"

	. "github.com/oi8io/lee/cn/common"
)

func TestConstructor(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Get(1)
	c.Put(3, 3) //淘汰2
	c.Get(2)
	c.Put(4, 4)
	c.Get(1)
	c.Get(3)
	c.Get(4)
	Dump(c)
}
func TestConstructor1(t *testing.T) {
	c := Constructor(2)
	cmds := []string{"put", "put", "put", "put", "get", "get"}
	numbers := [][]int{{2, 1}, {1, 1}, {2, 3}, {4, 1}, {1}, {2}}
	want := strings.Split("null,null,null,null,-1,3", ",")
	for i, cmd := range cmds {
		switch cmd {
		case "put":
			c.Put(numbers[i][0], numbers[i][1])
		case "get":
			got := c.Get(numbers[i][0])
			want, err := strconv.Atoi(want[i])
			if err != nil {
				t.Fatal(err)
			}
			if got != want {
				t.Errorf("Get() = %v, want %v", got, want)
			}
		}
	}
}
func TestConstructor2(t *testing.T) {
	c := Constructor(2)
	cmds := []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"}
	numbers := [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	want := strings.Split("null,null,1,null,-1,null,-1,3,4", ",")
	for i, cmd := range cmds {
		switch cmd {
		case "put":
			c.Put(numbers[i][0], numbers[i][1])
		case "get":
			got := c.Get(numbers[i][0])
			want, err := strconv.Atoi(want[i])
			if err != nil {
				t.Fatal(err)
			}
			if got != want {
				t.Errorf("Get() = %v, want %v", got, want)
			}
		}
	}
}

func TestX(t *testing.T) {
	list.New()
}
