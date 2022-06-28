package adapter

import "time"

type RealClock struct {
}

func (r *RealClock) Now() time.Time {
	return time.Now()
}
