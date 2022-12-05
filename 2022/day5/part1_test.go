package main

import (
	"reflect"
	"testing"
)

func Test_readContainerStacks(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"example 1", args{"    [C]         [Q]         [V]    "}, [][]string{{}, {"C"}, {}, {}, {"Q"}, {}, {}, {"V"}, {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readContainerStacks(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readContainerStacks() = %v, want %v", got, tt.want)
			}
		})
	}
}
