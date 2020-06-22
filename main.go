package main

import (
	"fmt"
	"golang_design_pattern/abstract_factory"
	"golang_design_pattern/adapter"
	"golang_design_pattern/bridge"
	"golang_design_pattern/builder"
	"golang_design_pattern/chain"
	"golang_design_pattern/command"
	"golang_design_pattern/composite"
	"golang_design_pattern/decorator"
	"golang_design_pattern/factory"
	"golang_design_pattern/flyweight"
	"golang_design_pattern/prototype"
	"golang_design_pattern/proxy"
	"golang_design_pattern/single"
	"golang_design_pattern/strategy"
	"golang_design_pattern/templ"
	"math/rand"
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

func compositeTest() {
	compositeSwimmer := &composite.CompositeSwimmer{
		&composite.SwimmerImp1{},
		&composite.Athlete{},
	}
	compositeSwimmer.Swim()
	compositeSwimmer.Train()
}

func adapterTest() {
	newPrinter := &adapter.PrinterAdapter{nil, ">>New Printer<<"}
	newRes := newPrinter.PrintSorted()
	fmt.Println(newRes)
	oldPrinter := &adapter.PrinterAdapter{new(adapter.LegacyPrinterImp), ">>Old Printer<<"}
	oldRes := oldPrinter.PrintSorted()
	fmt.Println(oldRes)
}

type writer struct {
	Msg string
}

func (w *writer) Write(p []byte) (n int, err error) {
	fmt.Println(w.Msg)
	return
}

func bridgeTest() {
	normalPrinter := &bridge.NormalPrinter{"Normal", new(bridge.PrinterImpl1)}
	normalPrinter.Print()
	w := &writer{"Writer"}
	printer := &bridge.PrinterImpl2{w}
	packetPrinter := &bridge.PacketPrinter{"Packet", printer}
	packetPrinter.Print()
}

func proxyTest() {
	someDatabase := proxy.UserList{}
	rand.Seed(2342342)
	for i := 0; i < 200; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, proxy.User{ID: n})
	}
	myProxy := proxy.UserListProxy{
		SomeDatabase:  someDatabase,
		StackCache:    proxy.UserList{},
		StackCapacity: 3,
	}
	user, _ := myProxy.FindUser(1)
	fmt.Printf("user: %s\n", user)
	user2, _ := myProxy.FindUser(1)
	fmt.Printf("user2: %s\n", user2)
	fmt.Printf("cache: %s\n", myProxy.StackCache)
}

func decoratorTest() {
	pizza := &decorator.Onion{&decorator.Meat{&decorator.PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err == nil {
		fmt.Println(pizzaResult)
	} else {
		fmt.Println(err)
	}

}

func flyweightTest() {
	flFactory := flyweight.NewTeamFactory()
	team1 := flFactory.GetTeam(1)
	fmt.Printf("team1: %s\n", team1)
	team2 := flFactory.GetTeam(3)
	fmt.Printf("team2: %s\n", team2)
}

func strategyTest() {
	strategy.Use("console")
	strategy.Use("image")
}

func chainTest() {
	logger2 := chain.SecondLogger{nil}
	logger1 := chain.FirstLogger{
		NextChain: &logger2,
	}
	logger1.Next("hello")
}

func commandTest() {
	queue := command.CommandQueue{}
	queue.AddCommand(command.CreateCommand("1st"))
	queue.AddCommand(command.CreateCommand("2nd"))
	queue.AddCommand(command.CreateCommand("3rd"))
	queue.AddCommand(command.CreateCommand("4th"))
	queue.AddCommand(command.CreateCommand("5th"))
}

type retriever struct {
	Msg string
}

func (r *retriever) Message() string {
	return r.Msg
}

func templateTest() {
	r := &retriever{"test"}
	t := &templ.Template{}
	res := t.ExecuteAlgorithm(r)
	fmt.Printf("result: %s\n", res)
}

func main() {
	singleTest()
	builderTest()
	factoryTest()
	abstractFactoryTest()
	prototypeTest()
	compositeTest()
	adapterTest()
	bridgeTest()
	proxyTest()
	decoratorTest()
	flyweightTest()
	strategyTest()
	chainTest()
	commandTest()
	templateTest()
}
