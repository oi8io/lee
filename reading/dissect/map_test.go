package dissect_test

import (
	cn "github.com/oi8io/lee/cn/common"
	"math"
	"testing"
)

func TestMap(t *testing.T) {
	var x = make(map[int]int, 10000)
	cn.Dump(x)
	for i := 0; i < 10000; i++ {
		x[i] = i
	}
	cn.Dump(x)
}

func TestMapAssignmentNan(t *testing.T) {
	m := make(map[float64]int, 0)
	nan := math.NaN()

	// Test assignment.
	m[nan] = 1
	m[nan] = 2
	m[nan] = 4

	testMapNan(t, m)
}

func testMapNan(t *testing.T, m map[float64]int) {
	if len(m) != 3 {
		t.Error("length wrong")
	}
	s := 0
	for k, v := range m {
		if k == k {
			t.Error("nan disappeared")
		}
		if (v & (v - 1)) != 0 {
			t.Error("value wrong")
		}
		s |= v
	}
	if s != 7 {
		t.Error("values wrong")
	}
}
