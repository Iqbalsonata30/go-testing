package gotesting

type Speeder interface {
	MaxSpeed() int
}

type Car struct {
	Speeder
}

func NewCar(speeder Speeder) *Car {
	return &Car{
		Speeder: speeder,
	}

}

func (c Car) Speed() int {
	defaultSpeed := 80

	if defaultSpeed > c.Speeder.MaxSpeed() {
		return c.Speeder.MaxSpeed()
	}

	if c.Speeder.MaxSpeed() < 10 {
		return 10
	}

	return defaultSpeed

}

type DefaultEngine struct{}

func (e *DefaultEngine) MaxSpeed() int {
	return 50
}

type TurboEngine struct{}

func (e *TurboEngine) MaxSpeed() int {
	return 100
}
