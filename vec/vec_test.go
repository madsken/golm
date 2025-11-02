package vec

import "testing"

func TestVec3(t *testing.T) {
	vec3 := Vec3{1, 1, 1}

	if vec3.X != 1.0 {
		t.Errorf("Test failed. Expected: %f, got: %f", 1.0, vec3.X)
	}
	if vec3.Y != 1.0 {
		t.Errorf("Test failed. Expected: %f, got: %f", 1.0, vec3.X)
	}
	if vec3.Z != 1.0 {
		t.Errorf("Test failed. Expected: %f, got: %f", 1.0, vec3.X)
	}
}
