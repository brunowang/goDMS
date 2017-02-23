package utils

func Hash(str string) uint {
	var ret uint = 0x811c9dc5
	for _, c := range str {
		ret ^= uint(c)
		ret *= 0x01000193
	}
	return ret
}

func Hash64(str string) uint64 {
	var ret uint64 = 0xCBF29CE484222325
	for _, c := range str {
		ret ^= uint64(c)
		ret *= 0x100000001B3
	}
	return ret
}
