package config

import (
	"strconv"
)

func ConvertStringToInt(s string) (i int, err error) {
	i, err = strconv.Atoi(s)
	if err != nil {
		return
	}
	return
}
