package main

import "strconv"

func IntToHex(number int64) string {
	return strconv.FormatInt(number, 16)
}

func IntToDec(number int64) string {
	return strconv.FormatInt(number, 10)
}
