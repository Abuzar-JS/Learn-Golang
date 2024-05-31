package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm Walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is ", d.first, "and I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func MethodsPoint() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2)
}
