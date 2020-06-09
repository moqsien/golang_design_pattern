package composite

import (
	"fmt"
)

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training...")
}

type SwimmerImp1 struct{}

func (s *SwimmerImp1) Swim() {
	fmt.Println("Swimming...")
}

type CompositeSwimmer struct {
	Swimmer
	Trainer
}
