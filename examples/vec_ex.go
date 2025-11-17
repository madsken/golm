package main

import (
	"fmt"
	"github.com/madsken/golm"
)

func vec3Examples() {
	v1 := golm.NewVec3(1, 2, 3)
	v2 := golm.NewVec3(4, 5, 6)

	fmt.Println("v1")
	fmt.Println(v1)

	fmt.Println("v2")
	fmt.Println(v2)

	//functions
	fmt.Println("Vec add")
	fmt.Println(v1.Add(v2))

	fmt.Println("v1 sub v2")
	fmt.Println(v1.Sub(v2))

	fmt.Println("v2 sub v1")
	fmt.Println(v2.Sub(v1))

	fmt.Println("mul")
	fmt.Println(v1.Mul(v2))

	fmt.Println("scale")
	fmt.Println(v1.Scale(2))

	fmt.Println("dot")
	fmt.Println(v1.Dot(v2))

	fmt.Println("length")
	fmt.Println(v1.Length())

	fmt.Println("normalize")
	fmt.Println(v1.Normalize())

	fmt.Println("cross")
	fmt.Println(v1.Cross(v2))
}
