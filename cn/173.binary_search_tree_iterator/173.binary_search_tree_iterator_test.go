package cn

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	tr := BuildTreeWithString("1")
	constructor := BSTIteratorConstructor(tr)

	cmd := []string{"next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"}
	values := strings.Split("1, 7, true, 9, true, 15, true, 20, false", ", ")
	for i, s := range cmd {
		switch s {
		case "next":
			next := constructor.Next()
			atoi, err := strconv.Atoi(values[i])
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("NEXT", next)
			if atoi != next {
				t.Fatal(atoi, next, "NOT EQUAL")
			}
		case "hasNext":
			next := constructor.HasNext()
			parseBool, err := strconv.ParseBool(values[i])
			if err != nil {
				t.Fatal(err)
			}
			if next != parseBool {
				t.Fatal("NOT EQUAL")
			}
		}
	}
}

func TestConstructor1(t *testing.T) {
	tr := BuildTreeWithString("7,3,15,null,null,9,20")
	constructor := BSTIteratorConstructor(tr)

	cmd := []string{"next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"}
	values := strings.Split("3, 7, true, 9, true, 15, true, 20, false", ", ")
	for i, s := range cmd {
		switch s {
		case "next":
			next := constructor.Next()
			atoi, err := strconv.Atoi(values[i])
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("NEXT", next)
			if atoi != next {
				t.Fatal(atoi, next, "NOT EQUAL")
			}
		case "hasNext":
			next := constructor.HasNext()
			parseBool, err := strconv.ParseBool(values[i])
			if err != nil {
				t.Fatal(err)
			}
			if next != parseBool {
				t.Fatal("NOT EQUAL ", "GOT ", next, "WANT ", parseBool)
			}
		}
	}
}

func TestConstructor2(t *testing.T) {

	tr := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}
	constructor := BSTIteratorConstructor(tr)
	cmd := []string{"hasNext", "next", "hasNext", "next", "hasNext"}
	values := strings.Split("true,1,true,2,false", ",")
	for i, s := range cmd {
		switch s {
		case "next":
			next := constructor.Next()
			atoi, err := strconv.Atoi(values[i])
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("NEXT", next)
			if atoi != next {
				t.Fatal(atoi, next, "NOT EQUAL")
			}
		case "hasNext":
			next := constructor.HasNext()
			parseBool, err := strconv.ParseBool(values[i])
			if err != nil {
				t.Fatal(err)
			}
			if next != parseBool {
				t.Fatal("NOT EQUAL ", "GOT ", next, "WANT ", parseBool)
			}
		}
	}
}
