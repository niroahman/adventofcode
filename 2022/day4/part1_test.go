package main

import (
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"testing"
)

func fetchInput() string {
	input_bytes, err := ioutil.ReadFile("input_test")
	if err != nil {
		log.Fatal("Failed to read input file.")
	}

	input := string(input_bytes)
	input = strings.Trim(input, "\n")
	return input

}

func Test_solvePart1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		// TODO: Add test cases.
		{"example 1", args{fetchInput()}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.input); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"example 1", args{1, 3}, []int{1, 2, 3}},

		{"example 2", args{1, 1}, []int{1}},

		{"example 3", args{1, 2}, []int{1, 2}},
		// test range 55-80
		{"example 4", args{55, 80}, []int{55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRange(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertSectionToNumbers(t *testing.T) {
	type args struct {
		section string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"example 1", args{"1-3"}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertSectionToNumbers(tt.args.section); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertSectionToNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
