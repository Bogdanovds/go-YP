package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name   string
		values float64
		want   float64
	}{
		{
			name:   "simple test #1",
			values: 3.1,
			want:   3.1,
		},
		{
			name:   "simple test #2",
			values: -3.14,
			want:   3.14,
		},
		{
			name:   "simple test #3",
			values: -0,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Abs(tt.values)
			assert.Equal(t, tt.want, v) // меняем на функцию Equal из пакета assert
		})
	}
}
