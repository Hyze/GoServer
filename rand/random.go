package rand

import (
	"math/rand"
	"time"
)

// randomDuration returns a random duration up to max, at intervals of max/10.
func RandomDuration(max time.Duration) time.Duration {
	return time.Duration(1+int64(rand.Intn(10))) * (max / 10)
}

