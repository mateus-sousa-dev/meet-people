package global

import "time"

var now NowFunc = time.Now

type NowFunc func() time.Time

func Now() time.Time {
	return now()
}
