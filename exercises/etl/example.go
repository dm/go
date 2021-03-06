package etl

import "strings"

const testVersion = 1

type legacy map[int][]string

func Transform(in legacy) (out map[string]int) {
	out = make(map[string]int)
	for key, values := range in {
		for _, val := range values {
			out[strings.ToLower(val)] = key
		}
	}
	return
}
