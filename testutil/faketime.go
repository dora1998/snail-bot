package testutil

import (
	"github.com/dora1998/snail-bot/clock"
	"time"
)

func SetFakeTime(t time.Time) {
	clock.Now = func() time.Time {
		return t
	}
}

func ResetFake() {
	clock.Now = time.Now
}
