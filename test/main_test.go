package test

import (
	"testing"
)

func Test_t1(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "t1",
			args: args{
				src: []string{"w", "w", "w", ".", "b", "y", "t", "e", "d", "a", "n", "c", "e", ".", "c", "o", "m"},
			},
			want: []string{"c", "o", "m", ".", "b", "y", "t", "e", "d", "a", "n", "c", "e", ".", "w", "w", "w"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1(tt.args.src)
			for i, s := range tt.args.src {
				if s != tt.want[i] {
					panic("err")
				}
			}
		})
	}
}
