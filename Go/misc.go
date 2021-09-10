package main

import "fmt"

const (
	topL = `┌`
	topR = `┐`
	btmL = `└`
	btmR = `┘`
	brkC = `─`
	brkW = `│`
)

// Loc represents the location cartesian coordinate pair
// Pos is the func fetching the pair
type Loc [2]int
type Pos func() Loc

// grid is a
type Grid map[Loc][]*Tile

type Tile struct {
	Pos []int
	//position
	//contents
}

func (t *Tile) Draw() {
	fmt.Print(topL)
	fmt.Println(topR)
	fmt.Print(btmL)
	fmt.Print(btmR)
}

func main() {

	t := &Tile{Pos: []int{1, 2}}
	t.Draw()
	//top
}

//var ErrServerClosed = errors.New("$pkgname: Server closed")

//type Header map[string][]string
// header = map[string][]string{
//			"Accept-Encoding": {"gzip, deflate"},
//			"Accept-Language": {"en-us"},
//}
//
//type Request struct

//type HandlerFunc func(ResponseWriter, *Request) error

//type ConsBundle [2]ConsFunc
//
//type ConsFunc func() (interface{}, error)
//
//type Cons struct {
//	a, b ConsFunc
//}
//
//type HandlerFunc func(string, string) error
//
////
//func (c *Cons) Car() (ConsFunc, error) {
//	return c.a, nil
//}
//
//func (c *Cons) Cdr() (ConsFunc, error) {
//	return c.b, nil
//}
//
//
