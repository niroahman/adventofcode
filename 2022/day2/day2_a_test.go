package main

import (
	"testing"
)

func Test_pointsForGame2(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A X", args{"A X"}, 4},
		{"B Y", args{"B Y"}, 5},
		{"C Z", args{"C Z"}, 6},
		{"A Y", args{"A Y"}, 7},
		{"B Z", args{"B Z"}, 8},
		{"C X", args{"C X"}, 9},
		{"A Z", args{"A Z"}, 0},
		{"B X", args{"B X"}, 0},
		{"C Y", args{"C Y"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointsForGame2(tt.args.row); got != tt.want {
				t.Errorf("pointsForGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
