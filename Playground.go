package main

import (
	"fmt"
	"reflect"
)

func main() {
	m:=3.0
	n:= 1 + 2 + 3
	fmt.Println(n)
	fmt.Println(reflect.TypeOf(m))
}
