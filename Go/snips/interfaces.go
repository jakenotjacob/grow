package main

import "fmt"

type Martian struct {
	Blorp string
}

type Human struct {
	Boogie string
}

type Dancer interface {
	Dance() string
}

func (h *Human) Dance() string {
	return h.Boogie
}

func (m *Martian) Dance() string {
	return m.Blorp
}

func main() {
	h := &Human{"yeeehaw"}
	m := &Martian{"basdfasf"}
	for _, d := range []Dancer{h, m} {
		fmt.Println(d.Dance())
	}
}
