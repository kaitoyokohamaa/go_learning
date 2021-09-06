package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func pow(x,n,lim float64)float64{
	if v := math.Pow(x,n);v <lim {
		return v
	}else {
		fmt.Printf("%g >=%g\n",v,lim)
	}
	return lim
}

func main() {
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	// var i, j int =1,2
	// k:=3
	// c, python,java :=true,false,"no!"
	// fmt.Println(i,j,k,c,python,java)
	// i:=42
	// f:=float64(i)
	// u:=uint(f)
	// fmt.Println(i,f,u)
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

