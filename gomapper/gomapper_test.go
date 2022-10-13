package gomapper

import "testing"

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"CapitalizeEveryThirdAlphanumericChar - success",
			args{
				"Aspiration.com",
			},
			"asPirAtiOn.cOm",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CapitalizeEveryThirdAlphanumericChar(tt.args.s); got != tt.want {
				t.Errorf("CapitalizeEveryThirdAlphanumericChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
