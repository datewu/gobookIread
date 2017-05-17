package main

type swimmer interface {
	swim()
}
type trainer interface {
	train()
}

type swimmerImpl struct{}

func (s swimmerImpl) swim() {
	println("swimming")
}

type compositeSwimmerB struct {
	trainer
	swimmer
}

type athlete struct{}

func (athlete) train() {
	println("training")

}

type animal struct{}

func (animal) eat() {
	println("eating")
}

type compositeSwimmerC struct {
	animal
	swimmer
}

func main() {
	sw := compositeSwimmerB{
		athlete{},
		swimmerImpl{},
	}
	sw.train()
	sw.swim()

	fish := compositeSwimmerC{
		animal{},
		swimmerImpl{},
	}
	fish.eat()
	fish.swim()
}
