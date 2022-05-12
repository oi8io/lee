package problems

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"AAAAAAAA", "AAAAAAGT", []string{"AAAAAAAA", "AACAAAAA", "ADCAAAAA", "TDCAAAAA", "AAAAAACT", "AAAAACCT", "AAAAGCCT", "AAAAGCAT", "AAAAAAGT", "AAAAGCAA", "AAAAGAAA", "AAAAAAAT"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
