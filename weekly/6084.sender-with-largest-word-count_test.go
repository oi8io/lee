package weekly

import "testing"

func Test_largestWordCount(t *testing.T) {
	type args struct {
		messages []string
		senders  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{[]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"}, []string{"Alice", "userTwo", "userThree", "Alice"}}, "Alice"},
		{"", args{[]string{"How is leetcode for everyone", "Leetcode is useful for practice"}, []string{"Bob", "Charlie"}}, "Charlie"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestWordCount(tt.args.messages, tt.args.senders); got != tt.want {
				t.Errorf("largestWordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
