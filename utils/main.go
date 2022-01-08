package utils

import (
	"fmt"
	"time"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
	Key             = "key"
	Value           = "value"
	Perm            = 0644
	TimeFormat      = "2006-01-02 15:04:05"
	DateFormat      = "20060102"
	HttpLogFormat   = "%s %s %s %q %s %s"
	RequestID       = "X-Request-ID"
	Unknown         = "unknown"
)

func GenerateFilePath() string {
	return fmt.Sprintf("/tmp/%s-data.json", time.Now().Format(DateFormat))
}
