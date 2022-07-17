package weekly295

import "testing"

func Test_discountPrices(t *testing.T) {
	type args struct {
		sentence string
		discount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{"there are $1 $2 and $5$ candies in the shop", 50}, "there are $0.50 $1.00 and 5$ candies in the shop"},
		{"", args{"there are $1 $2 and 5$ candies in the shop", 50}, "there are $0.50 $1.00 and 5$ candies in the shop"},
		{"", args{"1 2 $3 4 $5 $6 7 8$ $9 $10$", 100}, "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"},
		{"", args{"706hzu76jjh7yufr5x9ot60v149k5 $7651913186 pw2o $6", 28}, "706hzu76jjh7yufr5x9ot60v149k5 $5509377493.92 pw2o $4.32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := discountPrices(tt.args.sentence, tt.args.discount); got != tt.want {
				t.Errorf("discountPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
