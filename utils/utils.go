package utils

import (
	"fmt"
	"strconv"
)

func LoadFile() {}
func String2int(s string) int {
	num, err := strconv.Atoi(s) // converts string to int
	if err != nil {
		fmt.Println("Error:", err)
		panic("failed to convert string to int")
	}
	return num
}

func Int2string(i int) string {

	str := strconv.Itoa(i)
	return str
}

func Bigint2string(i int64) string {

	str := strconv.FormatInt(i, 10)
	return str
}
