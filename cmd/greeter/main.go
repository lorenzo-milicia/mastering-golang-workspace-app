package main

import (
	"fmt"

	"go.lorenzomilicia.dev/go-master/workspace-lib/pkg/greeting"
)

func main() {
	g := greeting.Greet(greeting.Polish)
	fmt.Println(g)
}
