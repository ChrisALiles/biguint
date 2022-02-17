package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/ChrisALiles/biguint"
)

func main() {
	fmt.Println("Args:")
	for _, arg := range os.Args[1:] {
		fmt.Printf("   %v \n", arg)
	}
	f, err := os.Create("memprofile")
	if err != nil {
		fmt.Println(err)
		return
	}

	str, err := biguint.Add(os.Args[1:]...)
	fmt.Printf("Add - result %v: err %v\n", str, err)
	str, err = biguint.Subtract(os.Args[1], os.Args[2])
	fmt.Printf("Subtract - result %v: err %v\n", str, err)
	str, err = biguint.Multiply(os.Args[1], os.Args[2])
	fmt.Printf("Multiply - result %v: err %v\n", str, err)
	str, err = biguint.Divide(os.Args[1], os.Args[2])
	fmt.Printf("Divide - result %v: err %v\n", str, err)
	str, err = biguint.Exp(os.Args[1], os.Args[2])
	fmt.Printf("Exp - result %v: err %v\n", str, err)
	pprof.WriteHeapProfile(f)
	f.Close()

}
