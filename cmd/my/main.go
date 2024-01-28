package main

import (
	"fmt"
	bb "mylinter/internal/bar"
	"mylinter/internal/foo"
	h "net/http"
)

func main() {
	var r h.Request
	fmt.Println(r)
	foo.Foo()
	bb.Bar()
}
