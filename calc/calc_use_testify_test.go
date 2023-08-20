package calc_test

import (
	"calc"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddUseTestify(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, calc.Add(tt.args.x, tt.args.y))
		})
	}
}
