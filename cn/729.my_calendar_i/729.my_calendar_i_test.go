package cn

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	b3 := c.Book(10, 20)
	if !b3 {
		t.Errorf("want true but %t", b3)
	}
	book := c.Book(15, 25)
	if book {
		t.Errorf("want false but %t", book)
	}
	fmt.Println(c.start, c.end)
	b := c.Book(20, 30)
	if !b {
		t.Errorf("want true but %t", b)
	}
	fmt.Println(c.start, c.end)
	b2 := c.Book(3, 6)
	if !b2 {
		t.Errorf("want true but %t", b2)
	}
	fmt.Println(c.start, c.end)
	c.Book(3, 9)
	c.Book(6, 10)
	c.Book(30, 40)

	for i := 0; i < c.n; i++ {
		fmt.Println(c.start[i], c.end[i])
	}
	//cn.Dump(c.start)
	//cn.Dump(c.end)
}

func TestMyCalendar_Book(t *testing.T) {
	var x = [][]int{{37, 50}, {33, 50}, {4, 17}, {35, 48}, {8, 25}}
	c := Constructor()
	for i := 0; i < len(x); i++ {
		c.Book(x[i][0], x[i][1])
	}
}
func TestMyCalendar_Book2(t *testing.T) {
	var x = [][]int{{10, 20}, {15, 25}, {20, 30}}
	c := Constructor()
	for i := 0; i < len(x); i++ {
		c.Book(x[i][0], x[i][1])
	}
}

func TestMyCalendar_Book1(t *testing.T) {

	x := [][]int{{47, 50}, {33, 41}, {39, 45}, {33, 42}, {25, 32}, {26, 35}, {19, 25}, {3, 8}, {8, 13}, {18, 27}}
	want := strings.Split("true,true,false,false,true,false,true,true,true,false", ",")
	c := Constructor()
	for i := 0; i < len(x); i++ {
		book := c.Book(x[i][0], x[i][1])
		parseBool, _ := strconv.ParseBool(want[i])
		if parseBool != book {
			t.Errorf("push [%d,%d] want %t but %t %+v %+v", x[i][0], x[i][1], parseBool, book, c.start, c.end)
		}
	}
}

func TestMyCalendar_BookX(t *testing.T) {

	var x = [][]int{{10, 20}, {15, 25}, {20, 30}}
	want := strings.Split("true,false,true", ",")
	c := Constructor()
	for i := 0; i < len(x); i++ {
		book := c.Book(x[i][0], x[i][1])
		parseBool, _ := strconv.ParseBool(want[i])
		if parseBool != book {
			t.Errorf("push [%d,%d] want %t but %t %+v %+v", x[i][0], x[i][1], parseBool, book, c.start, c.end)
		}
	}
}

func Test_insert(t *testing.T) {
	arr := insert([]int{1, 2, 3, 5, 6}, 0, 10)
	fmt.Println(arr)
}

func Test_findNumber(t *testing.T) {
	arr := []int{5, 7, 9}
	for i := 1; i < 10; i++ {
		number := findNumber(arr, i)
		arr = insert(arr, number, i)
	}
	fmt.Println(arr)
}
