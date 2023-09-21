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

type FakeEngine struct{}

func (e FakeEngine) MaxSpeed() int {
	return 5
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
		{
			name:   "speed should be 20 when use maxspeed less than 10",
			fields: fields{speeder: &FakeEngine{}},
			want:   20,
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
