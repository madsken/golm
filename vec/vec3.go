package vec

import (
	"github.com/madsken/golm/internal/constraints"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func NewVec3[T constraints.FloatConvertible](x, y, z T) Vec3 {
	return Vec3{
		X: float64(x),
		Y: float64(y),
		Z: float64(z),
	}
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func (v1 Vec3) Mul(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X * v2.X,
		Y: v1.Y * v2.Y,
		Z: v1.Z * v2.Z,
	}
}

func (v1 Vec3) Scale(i float64) Vec3 {
	return Vec3{
		X: v1.X * i,
		Y: v1.Y * i,
		Z: v1.Z * i,
	}
}

func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vec3) Length() float64 {
	return math.Sqrt(v1.Dot(v1))
}

func (v1 Vec3) Normalize() Vec3 {
	len := v1.Length()
	if len == 0 {
		return Vec3{0, 0, 0}
	}
	return Vec3{
		X: v1.X / len,
		Y: v1.Y / len,
		Z: v1.Z / len,
	}
}

func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}
