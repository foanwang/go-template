package util

import "strconv"

func StrToInt(input string) int {
	i, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic(err)
	}
	result := int(i)
	return result
}

func IntToStr(input int) string {
	result := strconv.Itoa(input)
	return result
}
