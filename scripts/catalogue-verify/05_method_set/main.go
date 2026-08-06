package main

import (
	"fmt"
	"reflect"
)

type Stringer interface{ String() string }

type T struct{}

func (t *T) String() string { return "T" } // pointer receiver only

func main() {
	valueImplements := reflect.TypeOf(T{}).Implements(reflect.TypeOf((*Stringer)(nil)).Elem())
	ptrImplements := reflect.TypeOf(&T{}).Implements(reflect.TypeOf((*Stringer)(nil)).Elem())

	fmt.Println("T implements Stringer:", valueImplements)
	fmt.Println("*T implements Stringer:", ptrImplements)

	if !valueImplements && ptrImplements {
		fmt.Println("BUG REPRODUCED: method set mismatch — value type excludes pointer-receiver methods")
	} else {
		fmt.Println("not reproduced")
	}
}
