package util

import "os"

func ExpandText(targetString string, targetExpand string, expandValue string) (fs string, err error) {
	m := map[string] string {
		targetExpand: expandValue,
	}

	fs = os.Expand(targetString, func(s string) string { return m[s] })
	return fs, nil
}
