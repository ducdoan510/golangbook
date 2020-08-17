package ex01_03

import "strings"

func echoWithJoin(args []string) string {
	return strings.Join(args, " ")
}

func echoWithConcat(args []string) string {
	ans, sep := "", ""
	for _, arg := range args {
		ans += sep + arg
		sep = " "
	}
	return ans
}
