package hopfield

func SignFunction(value int) (result int) {
	if value >= 0 {
		result = 1
	} else {
		result = -1
	}
	return
}
