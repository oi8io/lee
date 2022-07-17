package effective

import "testing"

func TestLoop(t *testing.T) {
	type args struct {
		src          []int
		validateOnly bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 5, 0, 3}, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Loop(tt.args.src, tt.args.validateOnly); (err != nil) != tt.wantErr {
				t.Errorf("Loop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckType(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckType()
		})
	}
}
