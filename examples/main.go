package main

import (
	"github.com/madsken/golm"
	"fmt"
)

func main() {
	v1 := golm.Vec3{1, 2, 3}
	v2 := golm.Vec3{4, 5, 6}

	fmt.Println("v1")
	printVec(v1)
	
	fmt.Println("v2")
	printVec(v2)


	//functions
	fmt.Println("Vec add")
	printVec(v1.Add(v2))

	fmt.Println("v1 sub v2")
	printVec(v1.Sub(v2))
	
	fmt.Println("v2 sub v1")
	printVec(v2.Sub(v1))

	fmt.Println("mul")
	printVec(v1.Mul(v2))

	fmt.Println("scale")
	printVec(v1.Scale(2))
	
	fmt.Println("dot")
	fmt.Println(v1.Dot(v2))

	fmt.Println("length")
	fmt.Println(v1.Length())


	fmt.Println("normalize")
	fmt.Println(v1.Normalize())

	fmt.Println("cross")
	fmt.Println(v1.Cross(v2))
}

func printVec(vec golm.Vec3)  {
	fmt.Printf("X: %f\n", vec.X)
	fmt.Printf("Y: %f\n", vec.Y)
	fmt.Printf("Z: %f\n", vec.Z)
}
