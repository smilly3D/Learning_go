package mypackage

import (
	"fmt"
	"publicAndPrivate/mypackage/internal/foo"
)

var Bar string = "hello, Bar"

func PrintMinha() {
	fmt.Println(foo.Minha)
}
