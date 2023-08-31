package math

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		lon1 float64
		lat1 float64
		lon2 float64
		lat2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0距离测试",
			args: args{
				lon1: 113.264385,
				lat1: 23.129112,
				lon2: 113.264385,
				lat2: 23.129112,
			},
			want: 0,
		},

		{
			name: "广州塔与津安创意园距离测试",
			args: args{
				lon1: 113.307649675152,
				lat1: 23.1200491020762,
				lon2: 113.43441,
				lat2: 23.135518,
			},
			want: 10900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 判断与want的误差在100米以内
			if got := Distance(tt.args.lon1, tt.args.lat1, tt.args.lon2, tt.args.lat2); got-tt.want > 100 {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
