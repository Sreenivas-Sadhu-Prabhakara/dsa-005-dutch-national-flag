package main

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{2, 2, 2}, []int{2, 2, 2}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		in := append([]int{}, c.in...)
		SortColors(in)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("SortColors(%v) = %v, want %v", c.in, in, c.want)
		}
	}
}
