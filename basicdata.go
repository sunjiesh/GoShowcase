package main

import (
	"fmt"
)

type subscriber struct {
	name string
	rate float64
}

type MyType string

func (m MyType) sayHi(name string) {
	fmt.Println("Hello", m, name)
}
func (m MyType) sayHi2(name string) {
	fmt.Println("Hello", m, name)
}

func demoMap() {
	var myMap1 map[string]float64
	myMap1 = make(map[string]float64)
	myMap2 := map[string]float64{"a": 1.1, "b": 1.2}
	keyArr := []string{"a", "b", "c"}
	myMap1["a"] = 1.1
	myMap1["b"] = 1.2
	fmt.Println(myMap1)
	fmt.Println(myMap2)
	for _, keyItem := range keyArr {
		value, ok := myMap2[keyItem]
		if ok {
			fmt.Printf("Found %s,key:%.2f\n", keyItem, value)
		} else {
			fmt.Printf("Not fund %s\n", keyItem)

		}
	}
	for key, value := range myMap1 {
		fmt.Printf("Key:%s,value:%.2f\n", key, value)
	}
}

func changeStuct(s *subscriber, newRate float64) {
	s.rate = newRate
}

func demoStruct() {
	var myStruct struct {
		number float64
		word   string
	}
	myStruct.number = 1
	myStruct.word = "a"
	fmt.Println(myStruct)
}

func demoCustType() {
	type car struct {
		brand string
		price float64
	}
	var car1 car
	var s1 subscriber
	s1.name = "s1"
	s1.rate = 100
	car1.brand = "brand1"
	car1.price = 100.2
	car2 := car{brand: "brand2", price: 200}

	changeStuct(&s1, 99)
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Printf("s1 rate:%.2f\n", s1.rate)
}

func demoCustType2() {
	golan := MyType("Go,")
	javalan := MyType("Java,")
	golan.sayHi("World")
	javalan.sayHi2("World")
}

func main() {
	demoMap()
	demoStruct()
	demoCustType()
	demoCustType2()
}
