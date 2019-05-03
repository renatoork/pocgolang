package Timer

import (
	_ "fmt"
	"time"
)

var timeStart time.Time

func TimeStart() {
	timeStart = time.Time{}
	timeStart = time.Now()
}

func TimeEnd() string {
	var duration time.Duration
	duration = time.Since(timeStart)
	return duration.String()
}
