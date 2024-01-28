package foo

import (
	kk "github.com/nishisuke/testgo/internal/bar"
)

func Foo() int { return 5 }

func aa() {
	kk.Bar()
}
