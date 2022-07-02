package cn

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestConstructor2(t *testing.T) {
	c := Constructor(2)
	cmds := []string{"put", "put", "get", "put", "get", "get", "put", "get", "get", "get"}
	numbers := [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}, {4, 4}, {1}, {3}, {4}}
	want := strings.Split("null, null, 1, null, -1, 3, null, -1, 3, 4", ", ")
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

func TestConstructor3(t *testing.T) {
	c := Constructor(10)
	cmds := []string{"LFUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	numbers := [][]int{{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}}
	want := strings.Split("null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null", ",")
	for i, cmd := range cmds {
		switch cmd {
		case "put":
			c.Put(numbers[i][0], numbers[i][1])
			if numbers[i][0] == 12 {
				fmt.Println(c.list, c.ele, lv(c.ele[numbers[i][0]]))
			}
		case "get":
			got := c.Get(numbers[i][0])
			want, err := strconv.Atoi(want[i])
			if err != nil {
				t.Fatal(err)
			}
			if got != want {
				t.Errorf("[ERR]Get(%d) = %v, want %v", numbers[i][0], got, want)
			} else {
				if numbers[i][0] == 12 {
					t.Logf("[INF]Get(%d) = %v, want %v", numbers[i][0], got, want)
				}
			}
		}
	}
}
func TestConstructor1(t *testing.T) {
	c := Constructor(0)
	cmds := []string{"LFUCache", "put", "get"}
	numbers := [][]int{{0}, {0, 0}, {0}}
	want := strings.Split("null, null, 1, null, -1, 3, null, -1, 3, 4", ", ")
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

func TestLFUCache_findFreq(t *testing.T) {
	var c = Constructor(10)
	c.f = []int{6, 8, 11, 13, 15, 16}
	for i := 4; i < 28; i++ {
		c.insertFreq(i)
	}
	fmt.Println(c.f)
	c.deleteFreq(28)
	c.deleteFreq(4)
	c.deleteFreq(7)
	c.deleteFreq(26)
	c.deleteFreq(27)
	fmt.Println(c.f)
}
