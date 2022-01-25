package spiralmmatrix2

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "t1",
			args: args{
				n: 3,
			},
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			name: "t2",
			args: args{
				n: 2,
			},
			want: [][]int{{1, 2}, {4, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
