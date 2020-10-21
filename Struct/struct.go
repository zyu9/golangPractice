package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int //Number
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[1])

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
