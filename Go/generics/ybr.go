package main

import (
	"fmt"
)

type Cat int8
type Lion int16
type Tiger int32
type Bear int64

// type constraint (...s are composite types?)
type Animals interface {
	Cat | Lion | Tiger | Bear
}


type ScaryAnimals interface {
	 Lions | Tigers | Bears
}

//func Screm[K comparable, V ScaryAnimals](m map[K]V ) V {
//	var scanimals V
//	for _, scaryone := range m {
//		fmt.Printf("%v", scaryone)
//		scanimals += scaryone
//	}
//	return scanimals
//}

//Lions.Screm


func main(){
	// cats
	lions := map[string]Lion{
        "swipe":  10,
    }

	// more cats
	tigers := map[string]Tiger{
        "swipe":  35.0,
    }

	// init sgt grumbles
	bears := map[string]Bear{
        "swipe":  79.00,
    }

	cats := map[string]Cat{
        "swipe":  1,
    }

	fmt.Printf("%v:%v",ScaryAnimals(bears))
	fmt.Printf("%v:%v",ScaryAnimals(lions))
	fmt.Printf("%v:%v",ScaryAnimals(tigers))
	fmt.Printf("%v:%v",ScaryAnimals(cats))
}



