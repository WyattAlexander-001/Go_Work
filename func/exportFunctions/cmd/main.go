// Need this at exportFunctions folder level
// go mod init exportFunctions
// go run cmd/main.go
//go build -o mycompiler.exe ./cmd/main.go
//^^ I chose the name compiler to so I can see something unique

package main

import (
	"fmt"

	"exportFunctions/mathlib"
)

func main() {
	sum := mathlib.Add(5, 7)
	fmt.Println("Sum:", sum)
	fmt.Scanln()
	// The following line would raise an error because `subtract` is not exported:
	// difference := mathlib.subtract(10, 3)
}
