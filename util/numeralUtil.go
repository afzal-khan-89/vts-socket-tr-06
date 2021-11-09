package util

import (
	"fmt"
	"strconv"
)

func Hex2Int(hexStr string) (int64, error) {
	intValue, err := strconv.ParseInt(hexStr, 16, 0)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func Bin2Int(binStr string) (int64, error) {
	intValue, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func hex2Bin(hexStr string) (string, error) {
	ui, err := strconv.ParseUint(hexStr, 16, 64)
	if err != nil {
		return "", err
	}

	format := fmt.Sprintf("%%0%db", len(hexStr)*4)
	return fmt.Sprintf(format, ui), nil
}

func ReverseByte(val byte) byte {
	var rval byte = 0
	for i := uint(0); i < 8; i++ {
		if val&(1<<i) != 0 {
			rval |= 0x80 >> i
		}
	}
	return rval
}

func ReverseUint8(val uint8) uint8 {
	return ReverseByte(val)
}

func ReverseUint16(val uint16) uint16 {
	var rval uint16 = 0
	for i := uint(0); i < 16; i++ {
		if val&(uint16(1)<<i) != 0 {
			rval |= uint16(0x8000) >> i
		}
	}
	return rval
}
