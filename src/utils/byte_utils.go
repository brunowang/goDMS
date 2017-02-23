package utils

func ByteToInt16(arr []byte, offset int) int16 {
	if arr == nil || len(arr)-offset < 2 {
		panic("byte array should not be nil, and must have 2 bits")
	}
	return int16(((arr[offset+2] & 0xff) << 8) | (arr[offset+3] & 0xff))
}

func Int16ToByte(sum int16) [2]byte {
	arr := [2]byte{}
	arr[0] = byte(sum >> 8)
	arr[1] = byte(sum & 0xff)
	return arr
}

func ByteToInt(arr []byte, offset int) int {
	if arr == nil || len(arr)-offset < 4 {
		panic("byte array should not be nil, and must have 4 bits")
	}
	return int(((arr[offset] & 0xff) << 24) | ((arr[offset+1] & 0xff) << 16) | ((arr[offset+2] & 0xff) << 8) | (arr[offset+3] & 0xff))
}

func IntToByte(sum int) [4]byte {
	arr := [4]byte{}
	arr[0] = byte(sum >> 24)
	arr[1] = byte(sum >> 16)
	arr[2] = byte(sum >> 8)
	arr[3] = byte(sum & 0xff)
	return arr
}

func ByteToInt64(arr []byte, offset int) int64 {
	if arr == nil || len(arr)-offset < 8 {
		panic("byte array should not be nil, and must have 8 bits")
	}
	return int64(((arr[offset] & 0xff) << 56) | ((arr[offset+1] & 0xff) << 48) | ((arr[offset+2] & 0xff) << 40) | ((arr[offset+3] & 0xff) << 32) | ((arr[offset] & 0xff) << 24) | ((arr[offset+1] & 0xff) << 16) | ((arr[offset+2] & 0xff) << 8) | (arr[offset+3] & 0xff))
}

func Int64ToByte(sum int64) [8]byte {
	arr := [8]byte{}
	arr[0] = byte(sum >> 56)
	arr[1] = byte(sum >> 48)
	arr[2] = byte(sum >> 40)
	arr[3] = byte(sum >> 32)
	arr[4] = byte(sum >> 24)
	arr[5] = byte(sum >> 16)
	arr[6] = byte(sum >> 8)
	arr[7] = byte(sum & 0xff)
	return arr
}

func ByteToFloat(arr []byte, offset int) float32 {
	if arr == nil || len(arr)-offset < 4 {
		panic("byte array should not be nil, and must have 4 bits")
	}
	i := ByteToInt(arr, offset)
	return float32(i)
}

func FloatToByte(fl float32) [4]byte {
	i := int(fl)
	return IntToByte(i)
}

func ByteToFloat64(arr []byte, offset int) float64 {
	if arr == nil || len(arr)-offset < 8 {
		panic("byte array should not be nil, and must have 8 bits")
	}
	i := ByteToInt64(arr, offset)
	return float64(i)
}

func Float64ToByte(fl float64) [8]byte {
	l := int64(fl)
	return Int64ToByte(l)
}
