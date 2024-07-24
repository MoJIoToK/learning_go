package task33101

import "testing"

func TestCountLetters(t *testing.T) {
	type args struct {
		s string
		l byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Пример 1",
			args: args{
				s: "Hello, world!",
				l: 'l',
			},
			want: 3,
		},
		{
			name: "Пример 2",
			args: args{
				s: "",
				l: 'l',
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLetters(tt.args.s, tt.args.l); got != tt.want {
				t.Errorf("CountLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
