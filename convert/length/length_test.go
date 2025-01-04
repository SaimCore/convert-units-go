package length

import "testing"

func TestConvertLength(t *testing.T) {
	result, err := ConvertLength(1, "kilometer", "meter")
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	expected := 1000.000
	if result != expected {
		t.Errorf("result = %.3f; want = %.3f\n", result, expected)
	}
}
