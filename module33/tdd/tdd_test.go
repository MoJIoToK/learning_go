package tdd

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		a Point
		b Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Пример 1",
			args: args{
				a: Point{X: 1, Y: 1},
				b: Point{X: 4, Y: 5},
			},
			want: 5,
		},
		{
			name: "Пример с отрицательными координатами",
			args: args{
				a: Point{X: 2, Y: -2},
				b: Point{X: 4, Y: 5},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
