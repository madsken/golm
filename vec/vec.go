package vec

type Vec[T any] interface {
	Add(T) T
	Sub(T) T
	Mul(T) T
	Scale(float64) T
	Dot(T) float64
	Length() float64
	Normalize() T
}
