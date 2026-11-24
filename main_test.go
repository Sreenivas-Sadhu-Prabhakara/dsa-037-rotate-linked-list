package main

import (
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}}, // k wraps past the length
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1}, 99, []int{1}},
		{[]int{}, 3, []int{}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}}, // full rotation is identity
	}
	for _, c := range cases {
		got := toSlice(RotateRight(fromSlice(c.in), c.k))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RotateRight(%v,%d) = %v, want %v", c.in, c.k, got, c.want)
		}
	}
}
