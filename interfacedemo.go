package main

import (
	"fmt"
)

type NoiseMaker interface {
	MakeSound()
}

type Whistle string
type Horn string
type ErrorDemo string

func (w Whistle) MakeSound() {
	fmt.Println("make sound,", w)
}

func (h Horn) MakeSound() {
	fmt.Println("make sound,", h)
}
func (h Horn) MakeSound2() {
	fmt.Println("make sound,", h)
}

func (e ErrorDemo) Error() string {
	return string(e)
}

func Play(n NoiseMaker) {
	n.MakeSound()
}

func demoInterface() {
	var toy NoiseMaker
	horn := Horn("Horn Demo2")
	toy = Whistle("Whistle Demo")
	toy.MakeSound()
	toy = Horn("Horn Demo")
	Play(toy)
	horn.MakeSound2()
}

func demoError() {
	var err error
	err = ErrorDemo("Test error")
	fmt.Println(err)
}

func main() {
	demoInterface()
	demoError()
}
