package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		intVar  int
		isError bool
	}{
		{
			name:    "Empty string throws error",
			str:     "",
			intVar:  0,
			isError: true,
		},
		{
			name:    "string as non int value",
			str:     "not an int",
			intVar:  0,
			isError: true,
		},
		{
			name:    "zero value",
			str:     "0",
			intVar:  0,
			isError: false,
		},
		{
			name:    "normal value",
			str:     "123456",
			intVar:  123456,
			isError: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			intVar, err := StringToInt(test.str)
			if test.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.intVar, intVar)
			}
		})
	}
}

func TestRuneToIntConversion(t *testing.T) {
	assert.Equal(t, 1, int('a')-96)
	assert.Equal(t, 26, int('z')-96)
	assert.Equal(t, 27, int('A')-38)
	assert.Equal(t, 52, int('Z')-38)
}

func TestSplitAfterNRunes(t *testing.T) {
	type args struct {
		str           string
		divEveryRunes int
	}
	tests := []struct {
		name  string
		param string
		divBy int
		want  []string
	}{
		{
			name:  "[Z] [M] [P]",
			param: "[Z] [M] [P]",
			divBy: 4,
			want:  []string{"[Z] ", "[M] ", "[P]"},
		},
		{
			name:  "[Z] [M]    ",
			param: "[Z] [M]    ",
			divBy: 4,
			want:  []string{"[Z] ", "[M] ", "   "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SplitAfterNRunes(tt.param, tt.divBy), "SplitAfterNRunes(%v, %v)", tt.param, tt.divBy)
		})
	}
}
