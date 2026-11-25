package stack

import (
	"reflect"
	"testing"
)

func TestOperations(t *testing.T) {
	tests := []struct {
		name         string
		a, b         []int
		wantA, wantB []int
		run          func(*[]int, *[]int)
	}{
		{"sa", []int{3, 1, 2}, []int{5, 4}, []int{1, 3, 2}, []int{5, 4}, func(a, b *[]int) { Sa(*a) }},
		{"sb", []int{3, 1}, []int{5, 4, 6}, []int{3, 1}, []int{4, 5, 6}, func(a, b *[]int) { Sb(*b) }},
		{"ss", []int{3, 1}, []int{5, 4}, []int{1, 3}, []int{4, 5}, func(a, b *[]int) { Ss(*a, *b) }},
		{"pa", []int{3, 1}, []int{5, 4}, []int{5, 3, 1}, []int{4}, Pa},
		{"pb", []int{3, 1}, []int{5, 4}, []int{1}, []int{3, 5, 4}, Pb},
		{"ra", []int{3, 1, 2}, []int{5, 4}, []int{1, 2, 3}, []int{5, 4}, func(a, b *[]int) { Ra(a) }},
		{"rb", []int{3, 1}, []int{5, 4, 6}, []int{3, 1}, []int{4, 6, 5}, func(a, b *[]int) { Rb(b) }},
		{"rr", []int{3, 1, 2}, []int{5, 4, 6}, []int{1, 2, 3}, []int{4, 6, 5}, Rr},
		{"rra", []int{3, 1, 2}, []int{5, 4}, []int{2, 3, 1}, []int{5, 4}, func(a, b *[]int) { Rra(a) }},
		{"rrb", []int{3, 1}, []int{5, 4, 6}, []int{3, 1}, []int{6, 5, 4}, func(a, b *[]int) { Rrb(b) }},
		{"rrr", []int{3, 1, 2}, []int{5, 4, 6}, []int{2, 3, 1}, []int{6, 5, 4}, Rrr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := append([]int{}, tt.a...)
			b := append([]int{}, tt.b...)
			tt.run(&a, &b)
			if !reflect.DeepEqual(a, tt.wantA) || !reflect.DeepEqual(b, tt.wantB) {
				t.Errorf("got A=%v B=%v; want A=%v B=%v", a, b, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestEmptyAndSingleElement(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		run  func(*[]int, *[]int)
	}{
		{"sa empty", []int{}, []int{}, func(a, b *[]int) { Sa(*a) }},
		{"sb empty", []int{}, []int{}, func(a, b *[]int) { Sb(*b) }},
		{"ss empty", []int{}, []int{}, func(a, b *[]int) { Ss(*a, *b) }},
		{"pa empty B", []int{1}, []int{}, Pa},
		{"pb empty A", []int{}, []int{2}, Pb},
		{"ra single", []int{1}, []int{}, func(a, b *[]int) { Ra(a) }},
		{"rb single", []int{}, []int{2}, func(a, b *[]int) { Rb(b) }},
		{"rr single", []int{1}, []int{2}, Rr},
		{"rra single", []int{1}, []int{}, func(a, b *[]int) { Rra(a) }},
		{"rrb single", []int{}, []int{2}, func(a, b *[]int) { Rrb(b) }},
		{"rrr single", []int{1}, []int{2}, Rrr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := append([]int{}, tt.a...)
			b := append([]int{}, tt.b...)
			tt.run(&a, &b)
			if !reflect.DeepEqual(a, tt.a) || !reflect.DeepEqual(b, tt.b) {
				t.Errorf("unexpected change: got A=%v B=%v; want A=%v B=%v", a, b, tt.a, tt.b)
			}
		})
	}
}
