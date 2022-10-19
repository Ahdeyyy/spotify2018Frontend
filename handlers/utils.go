package handlers

import (
	"strconv"
)

func msToMin(ms int) string {
	min := ms/60000
	sec := (ms%60000)/1000
	return strconv.Itoa(min) + ":" + strconv.Itoa(sec)
}

func addOne(index int) int {
	return index + 1
}
