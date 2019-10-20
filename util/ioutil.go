package util

func JoinBytesInWord(bytes []uint8) uint16 {
	return (uint16(bytes[1]) << 8) | uint16(bytes[0])
}

func BreakWordIntoBytes(word uint16) []uint8 {
	return []uint8{uint8(word), uint8(word >> 8)}
}
