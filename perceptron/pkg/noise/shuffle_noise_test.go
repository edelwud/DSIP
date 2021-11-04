package noise

import (
	"perceptron/internal/utils"
	"strconv"
	"testing"
)

func TestShuffleNoise_Run(t *testing.T) {
	shape, err := utils.ReadShape("../../resources/training/1.txt")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0.0; i <= 1; i += 0.1 {
		shuffle := CreateShuffleNoise(shape, i)
		shuffled := shuffle.Run()

		percent := int(i * 100)
		err = utils.WriteShape(shuffled, "../../resources/shuffle/1_"+strconv.Itoa(percent)+".txt")
		if err != nil {
			t.Fatal(err)
		}
	}
}
