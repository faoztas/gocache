package utils

import "testing"

func TestGetMD5Hash(t *testing.T) {
	value := GetMD5Hash("hello world!")
	if value == "" {
		t.Errorf("md5 generation failed")
	}
}
