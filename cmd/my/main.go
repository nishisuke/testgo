package main

import (
	"fmt"
	h "net/http"

	bb "github.com/nishisuke/testgo/internal/bar"
	"github.com/nishisuke/testgo/internal/foo"
)

func main() {
	var r h.Request
	fmt.Println(r)
	foo.Foo()
	bb.Bar()
}
