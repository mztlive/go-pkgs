package structure

import "testing"

type DemoStruct struct {
	OrderNo string `json:"order_no"`
	OpenID  string `json:"open_id"`
}

func TestBaseFulihuiEvent_FindVar(t *testing.T) {
	event := &DemoStruct{
		OrderNo: "123456",
		OpenID:  "openid",
	}

	tests := []struct {
		name    string
		varName string
		want    interface{}
	}{
		{
			name:    "Find order_no",
			varName: "order_no",
			want:    "123456",
		},
		{
			name:    "Find payment_time",
			varName: "open_id",
			want:    "openid",
		},
		{
			name:    "Find non-existent var",
			varName: "non_existent_var",
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTagVar(event, tt.varName, "json"); got != tt.want {
				t.Errorf("FindVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
