package vec

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 Vec2) Mul(v2 Vec2) Vec2 {
	return Vec2{
		X: v1.X * v2.X,
		Y: v1.Y * v2.Y,
	}
}

func (v1 Vec2) Scale(i float64) Vec2 {
	return Vec2{
		X: v1.X * i,
		Y: v1.Y * i,
	}
}

func (v1 Vec2) Dot(v2 Vec2) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (v1 Vec2) Length() float64 {
	return math.Sqrt(v1.Dot(v1))
}

func (v1 Vec2) Normalize() Vec2 {
	len := v1.Length()
	if len == 0 {
		return Vec2{0, 0}
	}
	return Vec2{
		X: v1.X / len,
		Y: v1.Y / len,
	}
}
