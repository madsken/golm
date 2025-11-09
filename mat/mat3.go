package mat

import (
	"github.com/madsken/golm/vec"
)

type Mat3 struct {
	M [3][3]float64
}

func (m1 Mat3) Add(m2 Mat3) Mat3 {
	res := Mat3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res.M[i][j] = m1.M[i][j] + m2.M[i][j]
		}
	}
	return res
}

func (m1 Mat3) Sub(m2 Mat3) Mat3 {
	res := Mat3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res.M[i][j] = m1.M[i][j] - m2.M[i][j]
		}
	}
	return res
}

func (m1 Mat3) Mul(m2 Mat3) Mat3 {
	res := Mat3{}

	res.M[0][0] = m1.M[0][0]*m2.M[0][0] + m1.M[0][1]*m2.M[1][0] + m1.M[0][2]*m2.M[2][0]
	res.M[0][1] = m1.M[0][0]*m2.M[0][1] + m1.M[0][1]*m2.M[1][1] + m1.M[0][2]*m2.M[2][1]
	res.M[0][2] = m1.M[0][0]*m2.M[0][2] + m1.M[0][1]*m2.M[1][2] + m1.M[0][2]*m2.M[2][2]

	res.M[1][0] = m1.M[1][0]*m2.M[0][0] + m1.M[1][1]*m2.M[1][0] + m1.M[1][2]*m2.M[2][0]
	res.M[1][1] = m1.M[1][0]*m2.M[0][1] + m1.M[1][1]*m2.M[1][1] + m1.M[1][2]*m2.M[2][1]
	res.M[1][2] = m1.M[1][0]*m2.M[0][2] + m1.M[1][1]*m2.M[1][2] + m1.M[1][2]*m2.M[2][2]

	res.M[2][0] = m1.M[2][0]*m2.M[0][0] + m1.M[2][1]*m2.M[1][0] + m1.M[2][2]*m2.M[2][0]
	res.M[2][1] = m1.M[2][0]*m2.M[0][1] + m1.M[2][1]*m2.M[1][1] + m1.M[2][2]*m2.M[2][1]
	res.M[2][2] = m1.M[2][0]*m2.M[0][2] + m1.M[2][1]*m2.M[1][2] + m1.M[2][2]*m2.M[2][2]

	return res
}

func (m1 Mat3) Scale(factor float64) Mat3 {
	res := Mat3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res.M[i][j] = m1.M[i][j] * factor
		}
	}
	return res
}

func (m1 Mat3) Transpose() Mat3 {
	res := Mat3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res.M[i][j] = m1.M[j][i]
		}
	}
	return res
}

func (m1 Mat3) Determinant() float64 {
	// using notation from https://en.wikipedia.org/wiki/Determinant
	aei := m1.M[0][0] * m1.M[1][1] * m1.M[2][2]
	bfg := m1.M[0][1] * m1.M[1][2] * m1.M[2][0]
	cdh := m1.M[0][2] * m1.M[1][0] * m1.M[2][1]

	ceg := m1.M[0][2] * m1.M[1][1] * m1.M[2][0]
	bdi := m1.M[0][1] * m1.M[1][0] * m1.M[2][2]
	afh := m1.M[0][0] * m1.M[1][2] * m1.M[2][1]

	return aei + bfg + cdh - ceg - bdi - afh
}

func (m1 Mat3) Inverse() Mat3 {
	det := m1.Determinant()
	if det == 0 {
		return Mat3{} //Consider returning an error in the cases it can fail
	}
	return m1.Scale(1 / det)
}

func (m1 Mat3) MulVec(v2 vec.Vec3) vec.Vec3 {
	return vec.Vec3{
		X: m1.M[0][0]*v2.X + m1.M[0][1]*v2.Y + m1.M[0][2]*v2.Z,
		Y: m1.M[1][0]*v2.X + m1.M[1][1]*v2.Y + m1.M[1][2]*v2.Z,
		Z: m1.M[2][0]*v2.X + m1.M[2][1]*v2.Y + m1.M[2][2]*v2.Z,
	}
}

func (m1 Mat3) Identity() Mat3 {
	return Mat3{
		M: [3][3]float64{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
}
