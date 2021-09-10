package main

import (
	"fmt"
	"io"
)

// thinking on scrollable buffers / "paging"

// interfaces are (value, type)
// - value: holding underlying concrete type
// - type:  underlying method holder ('calling a method on an
//
// "Calling a method on an interface value executes the method of the same name on its underlying type."
//
// fmt.Printf("(%v, %T)\n", i, i)

type Pageable interface {
	io.Reader
}

// pager itself is pageable (to view nested ones?) WE CAN GO DEEPER
type Pager struct {
	Pageable
	Page []*Pageable
}

type Seekable interface {
	Pageable
}

// aggregate interface; behavior container
type Seeky interface {
	Sneek()
	Snoop()
	Poke()
	Prod()
}

// youre a seeker if you're seekable and have something to seek through (a page)
type Seeker struct {
	Seekable       // interface
	Pageable       // interface
	Seeky          // interface
	Page     Pager // struct!
}

// seekers are seekable
type Seekers interface {
	Seekable
}

func main() {
	var s Seeky

	fmt.Printf("(%v, %T)\n", s, s)

	s = &Seeker{}

	fmt.Printf("(%v, %T)\n", s, s)

}
