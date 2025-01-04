package weight

import "testing"

func TestConvertWeight(t *testing.T) {
	result, err := ConvertWeight(1, "kilogram", "gram")
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	expected := 1000.000
	if result != expected {
		t.Errorf("result = %.3f; want %.3f\n", result, expected)
	}
}
