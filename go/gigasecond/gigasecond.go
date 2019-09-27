// Package gigasecond provides functions for calculating when someone
// turns 10^9 seconds old.
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to input
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
