package vec

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	const eps = 1e-9
	return math.Abs(a-b) < eps
}

func vecEquals(a, b Vec3) bool {
	return floatEquals(a.X, b.X) && floatEquals(a.Y, b.Y) && floatEquals(a.Z, b.Z)
}

func TestAdd(t *testing.T) {
	v1 := Vec3{1.0, 2.0, 3.0}
	v2 := Vec3{4.0, 5.0, 6.0}
	expected := Vec3{5.0, 7.0, 9.0}
	result := v1.Add(v2)
	if !vecEquals(result, expected) {
		t.Errorf("Add: expected %v, got %v", expected, result)
	}
}

func TestSub(t *testing.T) {
	v1 := Vec3{1.0, 2.0, 3.0}
	v2 := Vec3{4.0, 5.0, 6.0}
	expected := Vec3{-3.0, -3.0, -3.0}
	result := v1.Sub(v2)
	if !vecEquals(result, expected) {
		t.Errorf("Sub: expected %v, got %v", expected, result)
	}
}

func TestScale(t *testing.T) {
	v := Vec3{1.5, -2.0, 0.5}
	scale := 2.0
	expected := Vec3{3.0, -4.0, 1.0}
	result := v.Scale(scale)
	if !vecEquals(result, expected) {
		t.Errorf("Scale: expected %v, got %v", expected, result)
	}
}

func TestDot(t *testing.T) {
	v1 := Vec3{1.0, 3.0, -5.0}
	v2 := Vec3{4.0, -2.0, -1.0}
	// 1*4 + 3*(-2) + (-5)*(-1) = 4 - 6 + 5 = 3
	expected := 3.0
	result := v1.Dot(v2)
	if !floatEquals(result, expected) {
		t.Errorf("Dot: expected %v, got %v", expected, result)
	}
}
