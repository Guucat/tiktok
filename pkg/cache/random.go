package cache

import (
	"math/rand"
	"sync/atomic"
	"time"
)

var i int64

// RandExpiredTimeSec 根据固定有效期时间和随机有效期时间，计算出过期时间
func RandExpiredTimeSec(baseHour int, randHour int) time.Duration {

	s := rand.NewSource(atomic.AddInt64(&i, 1))
	r := rand.New(s)
	return time.Duration(time.Duration(baseHour)*time.Hour) + time.Duration(time.Duration(r.Intn(randHour*60*60))*time.Second)
}
