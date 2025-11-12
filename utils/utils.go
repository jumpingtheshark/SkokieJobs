package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func LoadFile(fpath string) string {
	data, err := os.ReadFile(fpath)
	if err != nil {
		panic(err)
	}
	rval := string(data)
	return rval
}
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

func Date2string(t time.Time) string {
	s := t.Format("2006-01-02")
	return s
}

func TodayDate() time.Time {
	today := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		0,
		0,
		0,
		0,
		time.UTC)
	return today

}
