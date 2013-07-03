package core

func booltou(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func utobool(i uint8) bool {
	if i != 0 {
		return true
	}
	return false
}
