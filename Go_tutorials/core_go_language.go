package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println


func main() {

pl(reflect.TypeOf(23))

}