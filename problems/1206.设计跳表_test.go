package problems

import (
	"reflect"
	"testing"
)

func TestSkiplistConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Skiplist
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SkiplistConstructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SkiplistConstructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_Add(t *testing.T) {
	type fields struct {
		down  *Skiplist
		next  *Skiplist
		data  int
		level int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Skiplist{
				down:  tt.fields.down,
				next:  tt.fields.next,
				data:  tt.fields.data,
				level: tt.fields.level,
			}
			this.Add(tt.args.num)
		})
	}
}

func TestSkiplist_Erase(t *testing.T) {
	type fields struct {
		down  *Skiplist
		next  *Skiplist
		data  int
		level int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Skiplist{
				down:  tt.fields.down,
				next:  tt.fields.next,
				data:  tt.fields.data,
				level: tt.fields.level,
			}
			if got := this.Erase(tt.args.num); got != tt.want {
				t.Errorf("Erase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_Search(t *testing.T) {
	type fields struct {
		down  *Skiplist
		next  *Skiplist
		data  int
		level int
	}
	type args struct {
		target int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Skiplist{
				down:  tt.fields.down,
				next:  tt.fields.next,
				data:  tt.fields.data,
				level: tt.fields.level,
			}
			if got := this.Search(tt.args.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
