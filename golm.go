package golm

import (
	"github.com/madsken/golm/internal/constraints"
	"github.com/madsken/golm/mat"
	"github.com/madsken/golm/quat"
	"github.com/madsken/golm/vec"
)

type (
	Vec2 = vec.Vec2
	Vec3 = vec.Vec3
	Vec4 = vec.Vec4
	Quat = quat.Quat
	Mat3 = mat.Mat3
	Mat4 = mat.Mat4
)

func NewVec3[T constraints.FloatConvertible](x, y, z T) Vec3    { return vec.NewVec3(x, y, z) }
func NewVec2[T constraints.FloatConvertible](x, y T) Vec2       { return vec.NewVec2(x, y) }
func NewVec4[T constraints.FloatConvertible](x, y, z, w T) Vec4 { return vec.NewVec4(x, y, z, w) }
func NewQuat[T constraints.FloatConvertible](w, x, y, z T) Quat { return quat.NewQuat(w, x, y, z) }

func CreateTransformationMatrix(rot Mat3, translation Vec3) Mat4 {
	baseM4 := Mat4{}
	baseM4 = baseM4.Identity()

	//Input translation components
	baseM4.M[0][3] = translation.X
	baseM4.M[1][3] = translation.Y
	baseM4.M[2][3] = translation.Z

	//Input rotation components
	baseM4.M[0][0] = rot.M[0][0]
	baseM4.M[0][1] = rot.M[0][1]
	baseM4.M[0][2] = rot.M[0][2]
	baseM4.M[1][0] = rot.M[1][0]
	baseM4.M[1][1] = rot.M[1][1]
	baseM4.M[1][2] = rot.M[1][2]
	baseM4.M[2][0] = rot.M[2][0]
	baseM4.M[2][1] = rot.M[2][1]
	baseM4.M[2][2] = rot.M[2][2]

	return baseM4
}
