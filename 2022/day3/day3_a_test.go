package main

import (
	"testing"
)

func Test_convertCharToPoints(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"a", args{'a'}, 1},
		{"A", args{'A'}, 27},
		{"p", args{'p'}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertCharToPoints(tt.args.char); got != tt.want {
				t.Errorf("convertCharToPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitHalf(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
		{"abcd", args{"abcd"}, "ab", "cd"},
		{"aaaaaaaabbbbbbbb", args{"aaaaaaaabbbbbbbb"}, "aaaaaaaa", "bbbbbbbb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitHalf(tt.args.line)
			if got != tt.want {
				t.Errorf("splitHalf() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitHalf() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_calculateTotal(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"vJrwpWtwJgWrhcsFMMfFFhFp", args{"vJrwpWtwJgWrhcsFMMfFFhFp"}, 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", args{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}, 38},
		{"PmmdzqPrVvPwwTWBwg", args{"PmmdzqPrVvPwwTWBwg"}, 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", args{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"}, 22},
		{"ttgJtRGJQctTZtZT", args{"ttgJtRGJQctTZtZT"}, 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", args{"CrZsJsPPZsGzwwsLwLmpwMDw"}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTotal(tt.args.line); got != tt.want {
				t.Errorf("calculateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
