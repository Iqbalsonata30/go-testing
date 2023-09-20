package gotesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultEngine_WithMaxSpeed(t *testing.T) {
	TestCases := []struct {
		name string
		want int
	}{
		{
			name: "Test default engine with max speed",
			want: 50,
		},
	}

	for _, tc := range TestCases {

		t.Run(tc.name, func(t *testing.T) {
			e := DefaultEngine{}
			assert.Equal(t, tc.want, e.MaxSpeed())
		})

	}

}

func TestTurboEngine_WithMaxSpeed(t *testing.T) {
	TestCases := []struct {
		name string
		want int
	}{
		{
			name: "Test turbo engine with max speed",
			want: 100,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			e := TurboEngine{}
			assert.Equal(t, tc.want, e.MaxSpeed())

		})

	}

}

func TestCar_WithSpeed(t *testing.T) {
	type fields struct {
		speeder Speeder
	}
	TestCases := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Test Car Default Engine with speed",
			fields: fields{speeder: &DefaultEngine{}},
			want:   50,
		},
		{
			name:   "Test Car Turbo Engine with speed",
			fields: fields{speeder: &TurboEngine{}},
			want:   80,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Car{
				Speeder: tc.fields.speeder,
			}
			assert.Equal(t, tc.want, c.Speed())
		})

	}

}
