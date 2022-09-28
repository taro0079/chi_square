package main

import "testing"

func TestJudgement(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want bool
	}{
		{
			name: "when chi2 is under the threshold",
			arg:  1,
			want: false,
		},
		{
			name: "when chi2 is over the threshold",
			arg:  4,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgement(tt.arg); got != tt.want {
				t.Errorf("judgement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOriginalNonConversion(t *testing.T) {
	type args struct {
		originalParameter  float64
		originalConversion float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "when originalParameter is 1000 and originalConversion is 100",
			args: args{originalParameter: 1000.0, originalConversion: 100.0},
			want: 900.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := originalNonConversion(tt.args.originalParameter, tt.args.originalConversion); got != tt.want {
				t.Errorf("originalNonConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestNonConversion(t *testing.T) {
	type args struct {
		testParameter  float64
		testConversion float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "when testParameter is 1000 and testConversion is 100",
			args: args{testParameter: 1000.0, testConversion: 100.0},
			want: 900.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testNonConversion(tt.args.testParameter, tt.args.testConversion); got != tt.want {
				t.Errorf("testNonConversion() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestSumVar(t *testing.T) {
	tests := []struct {
		name string
		arg  [][]float64
		want float64
	}{
		{
			name: "sum all",
			arg:  [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}},
			want: 21.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumvar(tt.arg); got != tt.want {
				t.Errorf("sumvar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcVar(t *testing.T) {
	type args struct {
		real  float64
		ideal float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "when real is 200 and ideal is 100",
			args: args{real: 200.0, ideal: 100.0},
			want: 100.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcvar(tt.args.real, tt.args.ideal); got != tt.want {
				t.Errorf("calcvar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeIdeal(t *testing.T) {
	type args struct {
		originalConversion    float64
		originalNonConversion float64
		testConversion        float64
		testNonConversion     float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "when originalConversion is 100, originalNonConversion is 900, testConversion is 100, testNonConversion is 900",
			args: args{originalConversion: 100.0, originalNonConversion: 900.0, testConversion: 200.0, testNonConversion: 2000.0},
			want: [][]float64{{93.75, 906.25}, {206.25, 1993.75}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeideal(tt.args.originalConversion, tt.args.originalNonConversion, tt.args.testConversion, tt.args.testNonConversion)

			if got[0][0] != tt.want[0][0] {
				t.Errorf("makevar() = %v, want %v", got, tt.want)

			}

			if got[0][1] != tt.want[0][1] {
				t.Errorf("makevar() = %v, want %v", got, tt.want)
			}

			if got[1][0] != tt.want[1][0] {
				t.Errorf("makevar() = %v, want %v", got, tt.want)
			}

			if got[1][1] != tt.want[1][1] {
				t.Errorf("makevar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testMakeArray(t *testing.T) {
	type args struct {
		oc  float64
		onc float64
		tc  float64
		tnc float64
	}
	a := args{oc: 100.0, onc: 900.0, tc: 200.0, tnc: 1000.0}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "when oc is 100, onc is 900, tc is 200, tnc is 1000",
			args: a,
			want: [][]float64{{100.0, 900.0, 1000.0}, {200.0, 1000.0, 1200.0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makearray(tt.args.oc, tt.args.onc, tt.args.tc, tt.args.tnc)

			if got[0][0] != tt.want[0][0] {
				t.Errorf("makearray() = %v, want %v", got, tt.want)
			}
			if got[0][1] != tt.want[0][1] {
				t.Errorf("makearray() = %v, want %v", got, tt.want)
			}
			if got[0][2] != tt.want[0][2] {
				t.Errorf("makearray() = %v, want %v", got, tt.want)
			}
			if got[1][0] != tt.want[1][0] {
				t.Errorf("makearray() = %v, want %v", got, tt.want)
			}
			if got[1][1] != tt.want[1][1] {
				t.Errorf("makearray() = %v, want %v", got, tt.want)
			}
			if got[1][2] != tt.want[1][2] {
				t.Errorf("makearray() = %v, want %v", got, tt.want)
			}

		})
	}

}
