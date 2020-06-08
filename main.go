package main

import (
	"fmt"
	"golang_design_pattern/abstract_factory"
	"golang_design_pattern/builder"
	"golang_design_pattern/factory"
	"golang_design_pattern/prototype"
	"golang_design_pattern/single"
)

func singleTest() {
	instance := single.GetInstance()
	count1 := instance.AddOne()
	fmt.Printf("Singleton, first value:%d", count1)
	fmt.Printf("\n")
	instance2 := single.GetInstance()
	count := instance2.AddOne()
	fmt.Printf("Singleton2, second value:%d", count)
	fmt.Printf("\n")
}

func builderTest() {
	director := builder.ManufacturingDirector{}
	carBuilder := &builder.CarBuilder{}
	director.SetBuilder(carBuilder)
	director.Constructor()
	car := carBuilder.GetVehicle()
	fmt.Printf("builder, car: %s", car)
	fmt.Printf("\n")
}

func factoryTest() {
	method1, _ := factory.GetPaymentMethod(1)
	res1 := method1.Pay(1.11)
	fmt.Println(res1)
	method2, _ := factory.GetPaymentMethod(2)
	res2 := method2.Pay(2.22)
	fmt.Println(res2)
}

func abstractFactoryTest() {
	vehicleFactory, _ := abstract_factory.BuildFactory(abstract_factory.CarFactoryType)
	luxuryCar, _ := vehicleFactory.NewVehicle(abstract_factory.LuxuryCarType)
	luxuryCar.GetVehicle()
	vehicleFactory2, _ := abstract_factory.BuildFactory(abstract_factory.MotorbikeFactoryType)
	sportMotor, _ := vehicleFactory2.NewVehicle(abstract_factory.SportMotorBikeType)
	sportMotor.GetVehicle()
}

func prototypeTest() {
	shirtCache := &prototype.ShirtCache{}
	item1, _ := shirtCache.GetClone(prototype.White)
	info1 := item1.GetInfo()
	fmt.Printf(info1)
	item2, _ := shirtCache.GetClone(prototype.Black)
	info2 := item2.GetInfo()
	fmt.Printf(info2)
}

func main() {
	singleTest()
	builderTest()
	factoryTest()
	abstractFactoryTest()
	prototypeTest()
}
