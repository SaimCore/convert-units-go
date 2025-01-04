package temperature

import "testing"

func TestConverTemprerature(t *testing.T) {
	result, err := ConvertTemperature(0, "celsius", "fahrenheit")
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	expected := 32.000
	if result != expected {
		t.Errorf("result = %.3f; want %.3f\n", result, expected)
	}
}
