package main

import "testing"

func Test_calculateGroupBadgeValue(t *testing.T) {
	type args struct {
		elfGroup []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"group 1", args{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}}, 18},
		{"group 2", args{[]string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateGroupBadgeValue(tt.args.elfGroup); got != tt.want {
				t.Errorf("calculateGroupBadgeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
