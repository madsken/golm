package quat

import "math"

type Quat struct {
	W, X, Y, Z float64
}

func (q1 Quat) Add(q2 Quat) Quat {
	return Quat{
		W: q1.W + q2.W,
		X: q1.X + q2.X,
		Y: q1.Y + q2.Y,
		Z: q1.Z + q2.Z,
	}
}

func (q1 Quat) Sub(q2 Quat) Quat {
	return Quat{
		W: q1.W - q2.W,
		X: q1.X - q2.X,
		Y: q1.Y - q2.Y,
		Z: q1.Z - q2.Z,
	}
}

func (q1 Quat) Scale(i float64) Quat {
	return Quat{
		W: q1.W * i,
		X: q1.X * i,
		Y: q1.Y * i,
		Z: q1.Z * i,
	}
}

func (q1 Quat) Dot(q2 Quat) float64 {
	return q1.W*q2.W + q1.X*q2.X + q1.Y*q2.Y + q1.Z*q2.Z
}

func (q1 Quat) Mul(q2 Quat) Quat {
	return Quat{
		W: q1.W*q2.W - q1.X*q2.X - q1.Y*q2.Y - q1.Z*q2.Z,
		X: q1.W*q2.X + q1.X*q2.W + q1.Y*q2.Z - q1.Z*q2.Y,
		Y: q1.W*q2.Y - q1.X*q2.Z + q1.Y*q2.W + q1.Z*q2.X,
		Z: q1.W*q2.Z + q1.X*q2.Y - q1.Y*q2.X + q1.Z*q2.W,
	}
}

func (q1 Quat) Length() float64 {
	return math.Sqrt(q1.Dot(q1))
}

func (q1 Quat) Normalize() Quat {
	length := q1.Length()
	if length == 0 {
		return Quat{0, 0, 0, 0}
	}
	return q1.Scale(1 / length) // norm a vec/quat => divide each element with length, hence scale each element with 1/length
}

func (q1 Quat) Conjugate() Quat {
	return Quat{
		W: q1.W,
		X: -q1.W,
		Y: -q1.Y,
		Z: -q1.Z,
	}
}

func (q1 Quat) Inverse() Quat {
	len2 := q1.Dot(q1)
	if len2 == 0 {
		return Quat{0, 0, 0, 0}
	}
	conj := q1.Conjugate()
	return conj.Scale(1 / len2)
}
