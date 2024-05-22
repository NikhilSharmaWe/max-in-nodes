package main

func alreadyPresent(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}
