package main

import (
	"strings"
)

func toSnake(s string) (res string) {
	begin := 0
	rs := []rune(s)
	for i, r := range rs {
		if begin == 0 && i == 0 || r < 'A' || r > 'Z' {
			continue
		}
		res += string(rs[begin:i]) + "_"
		begin = i
	}

	res += string(rs[begin:])
	res = strings.ToLower(res)
	return
}
