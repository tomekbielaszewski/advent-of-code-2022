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
