package util

import (
	"crypto/md5"
	"fmt"
)

func Password(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
func shoujihao(shoujihao string) bool {
	return false
}
