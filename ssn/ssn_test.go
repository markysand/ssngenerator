package ssn

import "testing"

func TestGetCheckSum(t *testing.T) {
	type args struct {
		n SSN
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				// 19750930-1938
				n: SSN{1, 9, 7, 5, 0, 9, 3, 0, 1, 9, 3, 8},
			},
			want: 8,
		},
		{
			name: "Test 1",
			args: args{
				// 20110530-4933
				n: SSN{2, 0, 1, 1, 0, 5, 3, 0, 4, 9, 3, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.n.GetCheckSum(); got != tt.want {
				t.Errorf("GetCheckSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
