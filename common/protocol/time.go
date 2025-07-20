package protocol

import (
	"time"

	"github.com/v4fly/v4ray-core/v0/common/dice"
)

type Timestamp int64

type TimestampGenerator func() Timestamp

func NowTime() Timestamp {
	return Timestamp(time.Now().Unix())
}

func NewTimestampGenerator(base Timestamp, delta int) TimestampGenerator {
	return func() Timestamp {
		rangeInDelta := dice.Roll(delta*2) - delta
		return base + Timestamp(rangeInDelta)
	}
}
