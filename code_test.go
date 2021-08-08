package services

import "testing"

func Test_grade(t *testing.T) {
	type args struct {
		nilai int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case return C",
			args: args{
				nilai: 40,
			},
			want: "C",
		},
		{
			name: "case return B",
			args: args{
				nilai: 90,
			},
			want: "B",
		},
		{
			name: "case return A",
			args: args{
				nilai: 80,
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grade(tt.args.nilai); got != tt.want {
				t.Errorf("grade() = %v, want %v", got, tt.want)
			}
		})
	}
}
