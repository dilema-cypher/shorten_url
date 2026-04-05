package utils

import (
	"time"
)

var location *time.Location

func init() {
	var err error
	location, err = time.LoadLocation("America/Cuiaba")
	if err != nil {
		location = time.UTC
	}
}

func TimeNow() time.Time {
	return time.Now().In(location)
}