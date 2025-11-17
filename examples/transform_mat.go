package main

import (
	"fmt"

	"github.com/madsken/golm"
)

func TransformationMatExample() {
	v := golm.NewVec3(1, 2, 3)
	m := golm.Mat3{}
	m = m.Identity()

	transform := golm.CreateTransformationMatrix(m, v)
	fmt.Println(transform)
}
