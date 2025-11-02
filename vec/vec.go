package vec

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) Add(i Vec3) Vec3 {
	return Vec3{
		X: v.X + i.X,
		Y: v.Y + i.Y,
		Z: v.Z + i.Z,
	}
}

func (v Vec3) Sub(i Vec3) Vec3 {
	return Vec3{
		X: v.X - i.X,
		Y: v.Y - i.Y,
		Z: v.Z - i.Z,
	}
}

func (v Vec3) Scale(i float64) Vec3 {
	return Vec3{
		X: v.X * i,
		Y: v.Y * i,
		Z: v.Z * i,
	}
}

func (v Vec3) Dot(i Vec3) float64 {
	return v.X*i.X + v.Y*i.Y + v.Z*v.Z
}
