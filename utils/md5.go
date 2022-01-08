package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func GetMD5Hash(text string) string {
	h := md5.New()
	_, err := io.WriteString(h, text)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
