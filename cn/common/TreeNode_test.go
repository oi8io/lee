package cn

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	s := NewTreeCodec()
	//deserialize := s.deserialize("1,null,2,3")
	deserialize := s.deserialize("3,9,20,null,null,15,7")
	//deserialize := s.deserialize("3,9,20,null,null,15,7,null,null,null,null")
	s1 := NewTreeCodec()
	serialize := s1.serialize(deserialize)
	fmt.Println(serialize)
}
