package cn

import (
	"github.com/davecgh/go-spew/spew"
	"io"
)

func Dump(a ...interface{}) {
	spew.Dump(a)
}

func Sdump(a ...interface{}) {
	spew.Sdump(a)
}

func Fdump(w io.Writer, a ...interface{}) {
	spew.Fdump(w, a)
}
