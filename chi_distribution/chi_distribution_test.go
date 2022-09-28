package chi_distribution

import (
	"math"
	"testing"
)

func TestChiDist(t *testing.T) {
	type args struct {
		dof float64
		x   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "when dof = 1 and x = 1",
			args: args{dof: 1, x: 1},
			want: 0.24197072451914337, // pythonのscipyで計算した
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChiDist(tt.args.dof, tt.args.x)
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("ChiDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
