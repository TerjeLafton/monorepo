package main

import (
	"fmt"

	external "github.com/terjelafton/monorepo/libs/real/pkg"
	"golang.org/x/example/hello/reverse"

	internal "github.com/terjelafton/monorepo/services/testing/pkg"
)

func main() {
	internal.Internal()
	external.External()

	string := "Hello, Bazel!"
	reverse := reverse.String(string)
	fmt.Println(string)
	fmt.Println(reverse)
}
