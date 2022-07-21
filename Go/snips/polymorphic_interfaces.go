package main

// This is a little snippet testing out embedded/polymorphic interfaces
// after learning about embedded structs and interfaces.
// (Ruby ¯\_(ツ)_/¯)

import "fmt"

type Base interface {
	Name() string
}

type Doer interface {
	//Name() string
	Base
	Do() string
}

type Shooer interface {
	//Name() string
	Base
	Shoo() string
}

type Thing struct {
	//	Base
	name string
}

type DoThing struct {
	Thing
}

type ShooThing struct {
	Thing
}

func (t *Thing) Name() string {
	return t.name
}

func (t *DoThing) do() string {
	t.name = "donethang"
	return t.name
}

func (t *DoThing) Do() string {
	return t.do()
}

func (t *ShooThing) shoo() string {
	t.name = "shoothang"
	return "shoo"
}

func (t *ShooThing) Shoo() string {
	return t.shoo()
}

type Tester interface {
	Do() string
}

func main() {
	var t Doer
	t = &DoThing{}

	var s Shooer
	s = &ShooThing{}

	// Bail out if Thing (struct) or Shoer(interface) no longer implement Base int
	var _ Base = (*Thing)(nil)
	var _ Base = (Shooer)(nil)

	//kinds := []inteface{}

	//var people = []*Person{
	//	{"foo"},
	//	{"bar"},
	//}

	for _, i := range []Base{s, t} {
		switch x := i.(type) {
		case Doer:
			fmt.Println("itsadoer")
		case Shooer:
			fmt.Println("itsashooer")
		default:
			fmt.Printf("It was neither! %T\n", x)
		}
	}
}
