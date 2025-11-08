package vec

import "math"

type Vec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (v1 Vec4) Add(v2 Vec4) Vec4 {
	return Vec4{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
		W: v1.W + v2.W,
	}
}

func (v1 Vec4) Sub(v2 Vec4) Vec4 {
	return Vec4{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
		W: v1.W - v2.W,
	}
}

func (v1 Vec4) Mul(v2 Vec4) Vec4 {
	return Vec4{
		X: v1.X * v2.X,
		Y: v1.Y * v2.Y,
		Z: v1.Z * v2.Z,
		W: v1.W * v1.W,
	}
}

func (v1 Vec4) Scale(i float64) Vec4 {
	return Vec4{
		X: v1.X * i,
		Y: v1.Y * i,
		Z: v1.Z * i,
		W: v1.W * i,
	}
}

func (v1 Vec4) Dot(v2 Vec4) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z + v1.W*v2.W
}

func (v1 Vec4) Length() float64 {
	return math.Sqrt(v1.Dot(v1))
}

func (v1 Vec4) Normalize() Vec4 {
	len := v1.Length()
	if len == 0 {
		return Vec4{0, 0, 0, 0}
	}
	return Vec4{
		X: v1.X / len,
		Y: v1.Y / len,
		Z: v1.Z / len,
		W: v1.W / len,
	}
}
