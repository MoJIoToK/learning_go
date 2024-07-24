package task3415

import "testing"

func TestNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Пример 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "Пример 2",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "Пример 3",
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Num(tt.args.n); got != tt.want {
				t.Errorf("Num() = %v, want %v", got, tt.want)
			}
		})
	}
}
