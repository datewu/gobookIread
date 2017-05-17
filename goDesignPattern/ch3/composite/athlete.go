package main

import "fmt"

type athlete struct{}

func (a *athlete) train() {
	fmt.Println("Training")
}

type compositeSwimmerA struct {
	MyAthlete athlete
	MySwim    func()
}

func swim() {
	fmt.Println("Swimming")
}

type animal struct{}

func (a *animal) eat() {
	println("Eating")
}

type shark struct {
	animal
	Swim func()
}

func main() {
	swimmer := compositeSwimmerA{MySwim: swim}
	swimmer.MyAthlete.train()
	swimmer.MySwim()

	fish := shark{Swim: swim}
	fish.eat()
	fish.Swim()
}
