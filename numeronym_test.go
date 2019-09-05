package numeronym

import "testing"

func TestNumeronym(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "Andreessen Horowitz",
			},
			want: "a16z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Numeronym(tt.args.s); got != tt.want {
				t.Errorf("Numeronym() = %v, want %v", got, tt.want)
			}
		})
	}
}
