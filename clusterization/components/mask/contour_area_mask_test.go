package mask

import (
	"fmt"
	"testing"
)

func TestContourAreaMask_Generate(t *testing.T) {
	contourMask := CreateContourAreaMask()
	mask := contourMask.Generate(3)
	fmt.Println(mask)
}
