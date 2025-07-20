package task

import "github.com/v4fly/v4ray-core/v0/common"

// Close returns a func() that closes v.
func Close(v interface{}) func() error {
	return func() error {
		return common.Close(v)
	}
}
