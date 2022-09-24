package utils

func ValueInsideOfSlice(value interface{}, slice interface{}) bool {
	for _, v := range slice.([]string) {

		if value == v {
			return true
		}

	}

	return false
}
