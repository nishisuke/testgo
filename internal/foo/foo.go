package foo

import (
	kk "mylinter/internal/bar"
)

func Foo() int { return 5 }

func aa() {
	kk.Bar()
}
