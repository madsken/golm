package mat

import (
	"github.com/madsken/golm/vec"
)

type Mat4 struct {
	M [4][4]float64
}

func (m1 Mat4) Add(m2 Mat4) Mat4 {
	res := Mat4{}
	for i := range 4 {
		for j := range 4 {
			res.M[i][j] = m1.M[i][j] + m2.M[i][j]
		}
	}
	return res
}

func (m1 Mat4) Sub(m2 Mat4) Mat4 {
	res := Mat4{}
	for i := range 4 {
		for j := range 4 {
			res.M[i][j] = m1.M[i][j] - m2.M[i][j]
		}
	}
	return res
}

func (m1 Mat4) Mul(m2 Mat4) Mat4 {
	var res Mat4

	// Row 0
	res.M[0][0] = m1.M[0][0]*m2.M[0][0] + m1.M[0][1]*m2.M[1][0] + m1.M[0][2]*m2.M[2][0] + m1.M[0][3]*m2.M[3][0]
	res.M[0][1] = m1.M[0][0]*m2.M[0][1] + m1.M[0][1]*m2.M[1][1] + m1.M[0][2]*m2.M[2][1] + m1.M[0][3]*m2.M[3][1]
	res.M[0][2] = m1.M[0][0]*m2.M[0][2] + m1.M[0][1]*m2.M[1][2] + m1.M[0][2]*m2.M[2][2] + m1.M[0][3]*m2.M[3][2]
	res.M[0][3] = m1.M[0][0]*m2.M[0][3] + m1.M[0][1]*m2.M[1][3] + m1.M[0][2]*m2.M[2][3] + m1.M[0][3]*m2.M[3][3]

	// Row 1
	res.M[1][0] = m1.M[1][0]*m2.M[0][0] + m1.M[1][1]*m2.M[1][0] + m1.M[1][2]*m2.M[2][0] + m1.M[1][3]*m2.M[3][0]
	res.M[1][1] = m1.M[1][0]*m2.M[0][1] + m1.M[1][1]*m2.M[1][1] + m1.M[1][2]*m2.M[2][1] + m1.M[1][3]*m2.M[3][1]
	res.M[1][2] = m1.M[1][0]*m2.M[0][2] + m1.M[1][1]*m2.M[1][2] + m1.M[1][2]*m2.M[2][2] + m1.M[1][3]*m2.M[3][2]
	res.M[1][3] = m1.M[1][0]*m2.M[0][3] + m1.M[1][1]*m2.M[1][3] + m1.M[1][2]*m2.M[2][3] + m1.M[1][3]*m2.M[3][3]

	// Row 2
	res.M[2][0] = m1.M[2][0]*m2.M[0][0] + m1.M[2][1]*m2.M[1][0] + m1.M[2][2]*m2.M[2][0] + m1.M[2][3]*m2.M[3][0]
	res.M[2][1] = m1.M[2][0]*m2.M[0][1] + m1.M[2][1]*m2.M[1][1] + m1.M[2][2]*m2.M[2][1] + m1.M[2][3]*m2.M[3][1]
	res.M[2][2] = m1.M[2][0]*m2.M[0][2] + m1.M[2][1]*m2.M[1][2] + m1.M[2][2]*m2.M[2][2] + m1.M[2][3]*m2.M[3][2]
	res.M[2][3] = m1.M[2][0]*m2.M[0][3] + m1.M[2][1]*m2.M[1][3] + m1.M[2][2]*m2.M[2][3] + m1.M[2][3]*m2.M[3][3]

	// Row 3
	res.M[3][0] = m1.M[3][0]*m2.M[0][0] + m1.M[3][1]*m2.M[1][0] + m1.M[3][2]*m2.M[2][0] + m1.M[3][3]*m2.M[3][0]
	res.M[3][1] = m1.M[3][0]*m2.M[0][1] + m1.M[3][1]*m2.M[1][1] + m1.M[3][2]*m2.M[2][1] + m1.M[3][3]*m2.M[3][1]
	res.M[3][2] = m1.M[3][0]*m2.M[0][2] + m1.M[3][1]*m2.M[1][2] + m1.M[3][2]*m2.M[2][2] + m1.M[3][3]*m2.M[3][2]
	res.M[3][3] = m1.M[3][0]*m2.M[0][3] + m1.M[3][1]*m2.M[1][3] + m1.M[3][2]*m2.M[2][3] + m1.M[3][3]*m2.M[3][3]

	return res
}

func (m1 Mat4) Scale(factor float64) Mat4 {
	res := Mat4{}
	for i := range 4 {
		for j := range 4 {
			res.M[i][j] = m1.M[i][j] * factor
		}
	}
	return res
}

func (m1 Mat4) Transpose() Mat4 {
	res := Mat4{}
	for i := range 4 {
		for j := range 4 {
			res.M[i][j] = m1.M[j][i]
		}
	}
	return res
}

func (m1 Mat4) Determinant() float64 {
	a := m1.M[0][0]
	b := m1.M[0][1]
	c := m1.M[0][2]
	d := m1.M[0][3]

	// Minor 0
	minor0 := Mat3{
		M: [3][3]float64{
			{m1.M[1][1], m1.M[1][2], m1.M[1][3]},
			{m1.M[2][1], m1.M[2][2], m1.M[2][3]},
			{m1.M[3][1], m1.M[3][2], m1.M[3][3]},
		},
	}.Determinant()

	// Minor 1
	minor1 := Mat3{
		M: [3][3]float64{
			{m1.M[1][0], m1.M[1][2], m1.M[1][3]},
			{m1.M[2][0], m1.M[2][2], m1.M[2][3]},
			{m1.M[3][0], m1.M[3][2], m1.M[3][3]},
		},
	}.Determinant()

	// Minor 2
	minor2 := Mat3{
		M: [3][3]float64{
			{m1.M[1][0], m1.M[1][1], m1.M[1][3]},
			{m1.M[2][0], m1.M[2][1], m1.M[2][3]},
			{m1.M[3][0], m1.M[3][1], m1.M[3][3]},
		},
	}.Determinant()

	// Minor 3
	minor3 := Mat3{
		M: [3][3]float64{
			{m1.M[1][0], m1.M[1][1], m1.M[1][2]},
			{m1.M[2][0], m1.M[2][1], m1.M[2][2]},
			{m1.M[3][0], m1.M[3][1], m1.M[3][2]},
		},
	}.Determinant()

	return a*minor0 - b*minor1 + c*minor2 - d*minor3
}

func (m1 Mat4) Inverse() Mat4 {
	det := m1.Determinant()
	if det == 0 {
		panic("Mat4 is singular and cannot be inverted")
	}
	invDet := 1.0 / det

	//No loops here, we going for speed
	// Precompute all 3x3 minors for cofactors
	cof00 := Mat3{M: [3][3]float64{
		{m1.M[1][1], m1.M[1][2], m1.M[1][3]},
		{m1.M[2][1], m1.M[2][2], m1.M[2][3]},
		{m1.M[3][1], m1.M[3][2], m1.M[3][3]},
	}}.Determinant()

	cof01 := Mat3{M: [3][3]float64{
		{m1.M[1][0], m1.M[1][2], m1.M[1][3]},
		{m1.M[2][0], m1.M[2][2], m1.M[2][3]},
		{m1.M[3][0], m1.M[3][2], m1.M[3][3]},
	}}.Determinant()

	cof02 := Mat3{M: [3][3]float64{
		{m1.M[1][0], m1.M[1][1], m1.M[1][3]},
		{m1.M[2][0], m1.M[2][1], m1.M[2][3]},
		{m1.M[3][0], m1.M[3][1], m1.M[3][3]},
	}}.Determinant()

	cof03 := Mat3{M: [3][3]float64{
		{m1.M[1][0], m1.M[1][1], m1.M[1][2]},
		{m1.M[2][0], m1.M[2][1], m1.M[2][2]},
		{m1.M[3][0], m1.M[3][1], m1.M[3][2]},
	}}.Determinant()

	cof10 := Mat3{M: [3][3]float64{
		{m1.M[0][1], m1.M[0][2], m1.M[0][3]},
		{m1.M[2][1], m1.M[2][2], m1.M[2][3]},
		{m1.M[3][1], m1.M[3][2], m1.M[3][3]},
	}}.Determinant()

	cof11 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][2], m1.M[0][3]},
		{m1.M[2][0], m1.M[2][2], m1.M[2][3]},
		{m1.M[3][0], m1.M[3][2], m1.M[3][3]},
	}}.Determinant()

	cof12 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][1], m1.M[0][3]},
		{m1.M[2][0], m1.M[2][1], m1.M[2][3]},
		{m1.M[3][0], m1.M[3][1], m1.M[3][3]},
	}}.Determinant()

	cof13 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][1], m1.M[0][2]},
		{m1.M[2][0], m1.M[2][1], m1.M[2][2]},
		{m1.M[3][0], m1.M[3][1], m1.M[3][2]},
	}}.Determinant()

	cof20 := Mat3{M: [3][3]float64{
		{m1.M[0][1], m1.M[0][2], m1.M[0][3]},
		{m1.M[1][1], m1.M[1][2], m1.M[1][3]},
		{m1.M[3][1], m1.M[3][2], m1.M[3][3]},
	}}.Determinant()

	cof21 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][2], m1.M[0][3]},
		{m1.M[1][0], m1.M[1][2], m1.M[1][3]},
		{m1.M[3][0], m1.M[3][2], m1.M[3][3]},
	}}.Determinant()

	cof22 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][1], m1.M[0][3]},
		{m1.M[1][0], m1.M[1][1], m1.M[1][3]},
		{m1.M[3][0], m1.M[3][1], m1.M[3][3]},
	}}.Determinant()

	cof23 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][1], m1.M[0][2]},
		{m1.M[1][0], m1.M[1][1], m1.M[1][2]},
		{m1.M[3][0], m1.M[3][1], m1.M[3][2]},
	}}.Determinant()

	cof30 := Mat3{M: [3][3]float64{
		{m1.M[0][1], m1.M[0][2], m1.M[0][3]},
		{m1.M[1][1], m1.M[1][2], m1.M[1][3]},
		{m1.M[2][1], m1.M[2][2], m1.M[2][3]},
	}}.Determinant()

	cof31 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][2], m1.M[0][3]},
		{m1.M[1][0], m1.M[1][2], m1.M[1][3]},
		{m1.M[2][0], m1.M[2][2], m1.M[2][3]},
	}}.Determinant()

	cof32 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][1], m1.M[0][3]},
		{m1.M[1][0], m1.M[1][1], m1.M[1][3]},
		{m1.M[2][0], m1.M[2][1], m1.M[2][3]},
	}}.Determinant()

	cof33 := Mat3{M: [3][3]float64{
		{m1.M[0][0], m1.M[0][1], m1.M[0][2]},
		{m1.M[1][0], m1.M[1][1], m1.M[1][2]},
		{m1.M[2][0], m1.M[2][1], m1.M[2][2]},
	}}.Determinant()

	// Build the adjugate (transpose of cofactor matrix) with alternating signs
	return Mat4{
		M: [4][4]float64{
			{+cof00 * invDet, -cof10 * invDet, +cof20 * invDet, -cof30 * invDet},
			{-cof01 * invDet, +cof11 * invDet, -cof21 * invDet, +cof31 * invDet},
			{+cof02 * invDet, -cof12 * invDet, +cof22 * invDet, -cof32 * invDet},
			{-cof03 * invDet, +cof13 * invDet, -cof23 * invDet, +cof33 * invDet},
		},
	}
}

func (m1 Mat4) MulVec(v2 vec.Vec4) vec.Vec4 {
	return vec.Vec4{
		X: m1.M[0][0]*v2.X + m1.M[0][1]*v2.Y + m1.M[0][2]*v2.Z + m1.M[0][3]*v2.W,
		Y: m1.M[1][0]*v2.X + m1.M[1][1]*v2.Y + m1.M[1][2]*v2.Z + m1.M[1][3]*v2.W,
		Z: m1.M[2][0]*v2.X + m1.M[2][1]*v2.Y + m1.M[2][2]*v2.Z + m1.M[2][3]*v2.W,
		W: m1.M[3][0]*v2.X + m1.M[3][1]*v2.Y + m1.M[3][2]*v2.Z + m1.M[3][3]*v2.W,
	}
}

func (m1 Mat4) Identity() Mat4 {
	return Mat4{
		M: [4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}
