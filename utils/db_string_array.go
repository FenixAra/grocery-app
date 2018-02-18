package utils

import "strings"

func ArrayDBFormat(s []string) string {
	return "{" + strings.Join(s, ",") + "}"
}
