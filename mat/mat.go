package mat

type Mat[T any] interface {
	Add(T) T
	Sub(T) T
	Mul(T) T
	Scale(float64) T
	Transpose() T
	Determinant() float64
	Inverse() T
	MulVec(T) T  //Vec input, Vec output
	Identity() T //Returns identity matrix
}
