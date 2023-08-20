package calc

import "testing"

func TestAdd(t *testing.T) {
	// t.Run("1 + 1 = 2", func(t *testing.T) {
	// 	result := Add(1, 1)
	// 	if result != 2 {
	// 		t.Errorf("expect 1 + 1 = 2, but got %d", result)
	// 	}
	// })

	// t.Run("1 + 2 = 3", func(t *testing.T) {
	// 	result := Add(1, 2)
	// 	if result != 3 {
	// 		t.Errorf("expect 1 + 2 = 3, but got %d", result)
	// 	}
	// })

	type args struct {
		x int
		y int
	}

	type testcase struct {
		name string
		args args
		want int
	}

	tests := []testcase{
		{
			"1 + 1 = 2",
			args{
				1,
				1,
			},
			2,
		},
		{
			"2 + 2 = 4",
			args{
				2,
				2,
			},
			4,
		},
		{
			"3 + 3 = 6",
			args{
				3,
				3,
			},
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
