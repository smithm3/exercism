package gigasecond

import "time"

const testVersion = 4

const gigSecond = 1000000000

func AddGigasecond(t time.Time) time.Time {
	return time.Unix(t.Unix() + gigSecond, 0)
}
