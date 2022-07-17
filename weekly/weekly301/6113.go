package weekly

import (
	"math"
)

type SmallestInfiniteSet struct {
	data []uint64
}

func Constructor() SmallestInfiniteSet {
	data := make([]uint64, 1000/64+1)
	for i := 0; i < len(data); i++ {
		data[i] = math.MaxUint64
	}
	return SmallestInfiniteSet{data: data}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	i := 0
	for ; i < 1000/64+1; i++ {
		if this.data[i] > 0 {
			break
		}
	}
	for x := 1; x <= 64; x++ {
		m := uint64(1 << (64 - x))
		//fmt.Printf("xxxxx %064b\n", m)
		if this.data[i]&m > 0 {
			//fmt.Sprintf("....... %064b", ^(1 << x))
			this.data[i] = this.data[i] ^ m
			return i*64 + x
		}
	}
	return 0
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	n := num - 1
	this.data[n/64] = this.data[n/64] | uint64(1<<(63-n%64))
}
