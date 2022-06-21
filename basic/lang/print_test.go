package lang

import (
	"fmt"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	fprint, err := fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fprint)
}