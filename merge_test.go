package mergesorted

import (
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"normal descending", []int{5, 3, 1}, []int{1, 3, 5}},
		{"single element", []int{7}, []int{7}},
		{"empty", []int{}, []int{}},
		{"negatives", []int{0, -2, -5}, []int{-5, -2, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseSlice(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseSlice(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMergeTwoAscending(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{"both non-empty", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"a empty", []int{}, []int{1, 2}, []int{1, 2}},
		{"b empty", []int{1, 2}, []int{}, []int{1, 2}},
		{"both empty", []int{}, []int{}, []int{}},
		{"duplicates", []int{1, 2, 2}, []int{2, 3}, []int{1, 2, 2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoAscending(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoAscending(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name                         string
		collection1, collection2, collection3 []int
		want                         []int
	}{
		{
			name:        "basic case",
			collection1: []int{9, 5, 1},
			collection2: []int{2, 6, 8},
			collection3: []int{3, 4, 7},
			want:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:        "duplicates across slices",
			collection1: []int{5, 3, 1},
			collection2: []int{1, 3, 5},
			collection3: []int{1, 3, 5},
			want:        []int{1, 1, 1, 3, 3, 3, 5, 5, 5},
		},
		{
			name:        "all empty",
			collection1: []int{},
			collection2: []int{},
			collection3: []int{},
			want:        []int{},
		},
		{
			name:        "collection1 empty",
			collection1: []int{},
			collection2: []int{1, 3},
			collection3: []int{2, 4},
			want:        []int{1, 2, 3, 4},
		},
		{
			name:        "collection2 empty",
			collection1: []int{4, 2},
			collection2: []int{},
			collection3: []int{1, 3},
			want:        []int{1, 2, 3, 4},
		},
		{
			name:        "collection3 empty",
			collection1: []int{4, 2},
			collection2: []int{1, 3},
			collection3: []int{},
			want:        []int{1, 2, 3, 4},
		},
		{
			name:        "single elements",
			collection1: []int{3},
			collection2: []int{1},
			collection3: []int{2},
			want:        []int{1, 2, 3},
		},
		{
			name:        "negative numbers",
			collection1: []int{0, -3, -7},
			collection2: []int{-5, -1, 2},
			collection3: []int{-6, -2, 1},
			want:        []int{-7, -6, -5, -3, -2, -1, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.collection1, tt.collection2, tt.collection3)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
