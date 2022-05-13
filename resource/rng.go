// A helper program that helps compare the reference Go PRNG with the translated
// Solidity version

package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	src := rand.New(rand.NewSource(911))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))
	fmt.Printf("%d\n", src.Intn(35))

	fmt.Println("Now bigints")
	// https://etherscan.io/address/0x2ed251752da7f24f33cfbd38438748bb8eeb44e1#readContract
	// getSeed(0x87e738a3d5e5345d6212d8982205a564289e6324, 32175)
	// = 105779926529366228504990970003713286107530024193944566341142813727459338091771
	seed := new(big.Int)
	seed, ok := seed.SetString("105779926529366228504990970003713286107530024193944566341142813727459338091771", 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	src = rand.New(rand.NewSource(seed.Int64()))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
	fmt.Printf("%d\n", src.Intn(50000000))
}
