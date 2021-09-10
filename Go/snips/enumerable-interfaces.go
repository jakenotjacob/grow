package main

import "io"

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

type Seekeable interface{}
type Seeker struct {
	Seekable
}

func main() {

}
