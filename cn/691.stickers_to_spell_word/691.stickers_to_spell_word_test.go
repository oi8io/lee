package cn

import (
	"testing"
)

func Test_minStickers(t *testing.T) {
	type args struct {
		stickers []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		//输入：stickers = ["notice","possible"], target = "basicbasic"

		{"", args{[]string{"notice", "possible"}, "basicbasic"}, -1},
		//输入： stickers = ["with","example","science"], target = "thehat"
		{"", args{[]string{"with", "example", "science"}, "thehat"}, 3},
		{"", args{[]string{"with", "example", "science"}, "ht"}, 1},
		{"", args{[]string{"these", "guess", "about", "garden", "him"}, "atomher"}, 3},
		{"", args{[]string{"these", "guess", "about", "garden", "him"}, "atomher"}, 3},
		{"", args{[]string{"a", "enemy", "material", "whose", "twenty", "describe", "magnet", "put", "hundred", "discuss"}, "separatewhich"}, 5},
		{"", args{[]string{"a", "enemy", "material", "whose", "twenty", "describe", "magnet", "put", "hundred", "discuss"}, "enemy"}, 1},
		{"", args{[]string{"a", "enemy", "material", "whose", "twenty", "describe", "magnet", "put", "hundred", "discuss"}, "a"}, 1},
		//{"", args{[]string{"a", "enemy", "material", "whose", "twenty", "describe", "magnet", "put", "hundred", "discuss"}, "hundredpdxxx"}, 1},
		//{"", args{[]string{"a", "enemy", "material", "whose", "twenty", "describe", "magnet", "put", "hundred", "discuss"}, "aawhh"}, 1},
		{"", args{[]string{"heavy", "claim", "seven", "set", "had", "it", "dead", "jump", "design", "question", "sugar", "dress", "any", "special", "ground", "huge", "use", "busy", "prove", "there", "lone", "window", "trip", "also", "hot", "choose", "tie", "several", "be", "that", "corn", "after", "excite", "insect", "cat", "cook", "glad", "like", "wont", "gray", "especially", "level", "when", "cover", "ocean", "try", "clean", "property", "root", "wing"}, "travelbell"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStickers(tt.args.stickers, tt.args.target); got != tt.want {
				t.Errorf("minStickers() = %v, want %v", got, tt.want)
			}
		})
	}
}
