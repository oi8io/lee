package lang

import "testing"

func Test_nextInt(t *testing.T) {
	type args struct {
		b []byte
		i int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{"", args{[]byte{'1', 'A', 'a', '2', '3', 'A', '9'}, 5}, 9, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := nextInt(tt.args.b, tt.args.i)
			if got != tt.want {
				t.Errorf("nextInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("nextInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefer(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		// TODO: Add test cases.
		{"", args{1, 3}, 4000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := Defer(tt.args.a, tt.args.b); gotC != tt.wantC {
				t.Errorf("Defer() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
func TestDeferb(t *testing.T) {
	b()
}
