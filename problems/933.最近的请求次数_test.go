package problems

import (
	"reflect"
	"testing"
)

func TestNewRecentCounter(t *testing.T) {
	tests := []struct {
		name string
		want RecentCounter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRecentCounter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecentCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecentCounter_Ping(t *testing.T) {
	/** ["RecentCounter", "ping", "ping", "ping", "ping"]
	 * [[], [1], [100], [3001], [3002]]
	 * 输出：
	 * [null, 1, 2, 3, 3]
	 *
	 * 解释：
	 * RecentCounter recentCounter = new RecentCounter();
	 * recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
	 * recentCounter.ping(100);   // requests = [1, 100]，范围是 [-2900,100]，返回 2
	 * recentCounter.ping(3001);  // requests = [1, 100, 3001]，范围是 [1,3001]，返回 3
	 * recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002]，范围是
	 * [2,3002]，返回 3
	 */
	recentCounter := NewRecentCounter()
	recentCounter.Ping(1)
	recentCounter.Ping(100)
	recentCounter.Ping(3001)
	recentCounter.Ping(3002)
	recentCounter.Ping(3008)
	recentCounter.Ping(3009)
}

func TestRecentCounter_Ping1(t *testing.T) {
	type args struct {
		t int
	}
	var tests []struct {
		name string
		args args
		want int
	}
	for i := 1; i < 301; i++ {
		tests = append(tests, struct {
			name string
			args args
			want int
		}{"aaa", args{i}, i})
	}
	for i := 3001; i < 3010; i++ {
		tests = append(tests, struct {
			name string
			args args
			want int
		}{"aaa", args{i}, 301})
	}
	this := NewRecentCounter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := this.Ping(tt.args.t); got != tt.want {
				t.Errorf("Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}
