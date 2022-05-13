// A helper program that helps compare the reference Go PRNG with the translated
// Solidity version

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	src := rand.New(rand.NewSource(911))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(44))
	fmt.Printf("%d\n", src.Intn(45))
}
