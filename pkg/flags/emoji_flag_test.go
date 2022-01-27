package flags

import "testing"

func TestGetFlag(t *testing.T) {
	type args struct {
		nationality string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return correct flag",
			args: args{nationality: "Irish"},
			want: "🇮🇪",
		},
		{
			name: "Should return earth if nationality unknown",
			args: args{nationality: "Unknown"},
			want: "🌎",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFlag(tt.args.nationality); got != tt.want {
				t.Errorf("GetFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
