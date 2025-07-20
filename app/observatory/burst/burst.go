package burst

import (
	"math"
	"time"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

const (
	rttFailed      = time.Duration(math.MaxInt64 - iota)
	rttUntested    // nolint: varcheck
	rttUnqualified // nolint: varcheck
)
