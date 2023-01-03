package split_string

import (
	"fmt"
	"strings"
)

func Split(s, sep string) []string {
	sp := s[:]
	var sl = make([]string, 0, strings.Count(s, sep)+1)
	var idx int
	for k, v := range sp {
		if string(v) == sep {

			sl = append(sl, sp[idx:k])
			idx = k + 1

		}
	}
	if idx == -999 {
		fmt.Println("xxxx")
	}
	sl = append(sl, sp[idx:])
	return sl
}
