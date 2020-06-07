package abstract_factory

import (
	"errors"
	"fmt"
)

// 抽象工厂模式
type Vehicle interface {
	NumWheels() int
	NumSeats() int
	GetVehicle()
}

type Car interface {
	NumDoors() int
	Vehicle
}

type Motorbike interface {
	GetMotorbikeType() int
	Vehicle
}

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}
type Cars struct{}
type LuxuryCar struct {
	Cars
}
type FamilyCar struct {
	Cars
}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Invalide CarType: %d", v))
	}
}

func (c *Cars) NumDoors() int {
	return 4
}

func (c *Cars) NumWheels() int {
	return 4
}

func (c *Cars) NumSeats() int {
	return 5
}

func (c *Cars) GetVehicle() {
	doors := c.NumDoors()
	wheels := c.NumWheels()
	seats := c.NumWheels()
	fmt.Printf("doors: %d, wheels: %d, seats: %d", doors, wheels, seats)
	fmt.Println("")
}

func (fc *FamilyCar) NumDoors() int {
	return 5
}

const (
	SportMotorBikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}
type Motorbikes struct{}
type SportMotorBike struct {
	Motorbikes
}
type CruiseMotorbike struct {
	Motorbikes
}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorBikeType:
		return new(SportMotorBike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Invalide MotorType: %d", v))
	}
}

func (m *Motorbikes) NumWheels() int {
	return 2
}

func (m *Motorbikes) NumSeats() int {
	return 1
}

func (m *Motorbikes) GetMotorbikeType() int {
	return SportMotorBikeType
}

func (m *Motorbikes) GetVehicle() {
	wheels := m.NumWheels()
	seats := m.NumSeats()
	motorType := m.GetMotorbikeType()
	fmt.Printf("Wheels: %d, Seats: %d, MotorType: %d", wheels, seats, motorType)
	fmt.Println("")
}

func (cm *CruiseMotorbike) NumSeats() int {
	return 2
}

func (cm *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Invalide VehicleFactoryType: %d", f))
	}
}
