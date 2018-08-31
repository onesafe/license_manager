package views

import (
	"time"
)

func nowMs() int64 {
	ret := int64(time.Now().UnixNano() / 1000000)
	return ret
}
