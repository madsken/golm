package main

import (
	"fmt"
	"github.com/madsken/golm"
)

func mat4Examples() {
	m4 := golm.Mat4{}
	m4.M[0][0] = 0.0
	m4.M[0][1] = 0.1
	m4.M[0][2] = 0.2
	m4.M[0][3] = 0.3
	m4.M[1][0] = 1.0
	m4.M[1][1] = 1.1
	m4.M[1][2] = 1.2
	m4.M[1][3] = 1.3
	m4.M[2][0] = 2.0
	m4.M[2][1] = 2.1
	m4.M[2][2] = 2.2
	m4.M[2][3] = 2.3
	m4.M[3][0] = 3.0
	m4.M[3][1] = 3.1
	m4.M[3][2] = 3.2
	m4.M[3][3] = 3.3
	fmt.Println(m4)
}
