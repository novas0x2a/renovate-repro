package main

import (
	"fmt"

	"go.mozilla.org/sops/v3"
)

func main() {
	fmt.Println("use some ", sops.Comment{})
}
