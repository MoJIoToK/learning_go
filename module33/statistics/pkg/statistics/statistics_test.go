package statistics

import "testing"

func TestAvg(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "непустой массив",
			args: args{
				nums: []float64{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "пустой массив",
			args: args{
				nums: []float64{},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Avg(tt.args.nums); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Задание 33.4.1
func TestMaxN(t *testing.T) {
	nums := []float64{1, 2, 3}
	want := 3.0
	got := MaxN(nums)
	if got != want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}

	nums = []float64{-21.5, -2.6, -3}
	want = -2.6
	got = MaxN(nums)
	if got != want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}

}
