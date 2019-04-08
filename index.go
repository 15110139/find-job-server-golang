package main

import (
	"fmt"
	"reflect"
)

type User struct {
	user string
}

func main() {
	var name string = "Carl Johannes"
	fmt.Println("name is a type of: ", reflect.TypeOf(name))

}
