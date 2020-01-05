package main

import "testing"

func Test_nps(t *testing.T) {
	type args struct {
		scores [5]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nps(tt.args.scores); got != tt.want {
				t.Errorf("nps() = %v, want %v", got, tt.want)
			}
		})
	}
}