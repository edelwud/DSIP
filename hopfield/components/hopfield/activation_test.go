package hopfield

import "testing"

func TestSignFunction(t *testing.T) {
	value := SignFunction(-100)
	if value >= 0 {
		t.Fatal("unexpected value: ", value)
	}
}
