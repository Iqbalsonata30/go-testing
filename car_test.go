package gotesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type MockEngine struct {
	mock.Mock
}

func (m *MockEngine) MaxSpeed() int {
	args := m.Called()
	return args.Get(0).(int)
}

func TestCar_Speed_WithMock(t *testing.T) {
	t.Run("Just called once", func(t *testing.T) {
		mock := new(MockEngine)
		car := Car{
			Speeder: mock,
		}
		mock.On("MaxSpeed").Return(9).Once()
		speed := car.Speed()
		assert.Equal(t, 20, speed)
		mock.AssertExpectations(t)

	})

	t.Run("Just called 3 times", func(t *testing.T) {
		mock := new(MockEngine)
		car := Car{
			Speeder: mock,
		}
		mock.On("MaxSpeed").Return(60).Times(3)
		speed := car.Speed()
		assert.Equal(t, 60, speed)
		mock.AssertExpectations(t)
	})
}
