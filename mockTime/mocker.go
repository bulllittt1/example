package mocktime

import (
	"context"
	"time"
)

type key int

const (
	_ key = iota
	clientTimeKey
)

var defaultCtx = context.Background()

var timeNowFunc = func() uint {
	return uint(time.Now().Unix())
}

func mockTimeNow(timeToReturn uint) func() {
	timeNowFunc = func() uint { return timeToReturn }
	defaultCtx = context.WithValue(defaultCtx, clientTimeKey, int64(timeToReturn))
	return func() {
		timeNowFunc = func() uint { return uint(time.Now().Unix()) }
		defaultCtx = context.Background()
	}
}
