package gomapper

import (
	"reflect"
	"testing"
)

var ss *skipString = NewSkipString(3, "Aspiration.com")

func Test_skipString_GetValueAsRuneSlice(t *testing.T) {
	tests := []struct {
		name string
		want []rune
	}{
		{
			"success",
			[]rune("Aspiration.com"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ss.GetValueAsRuneSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("skipString.GetValueAsRuneSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPositionMapper(t *testing.T) {
	type args struct {
		pos int
		arr []rune
	}
	tests := []struct {
		name string
		args args
		want map[int]rune
	}{
		{
			"success",
			args{
				3,
				ss.arr,
			},
			map[int]rune{2: rune('P'), 5: rune('A'), 8: rune('O'), 12: rune('O')},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPositionMapper(tt.args.pos, tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPositionMapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_skipString_String(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"success",
			"capitalized string is: asPirAtiOn.cOm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MapString(ss)
			if got := ss.String(); got != tt.want {
				t.Errorf("skipString.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
