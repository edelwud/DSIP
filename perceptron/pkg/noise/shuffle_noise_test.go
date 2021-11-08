package noise

import (
	"io/ioutil"
	"perceptron/internal/utils"
	"strconv"
	"strings"
	"testing"
)

func TestShuffleNoise_Run(t *testing.T) {
	dir, err := ioutil.ReadDir("../../resources/training/")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range dir {
		shape, err := utils.ReadShape("../../resources/training/" + file.Name())
		if err != nil {
			t.Fatal(err)
		}

		for i := 0.0; i <= 1; i += 0.1 {
			shuffle := CreateShuffleNoise(shape, i)
			shuffled := shuffle.Run()

			percent := int(i * 100)

			err = utils.WriteShape(shuffled, "../../resources/shuffle/"+strings.Split(file.Name(), ".")[0]+"_"+strconv.Itoa(percent)+".txt")
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
