package main

import (
	"time"

	"github.com/v4fly/v4ray-core/v0/main/v2binding"
)

func main() {
	v2binding.StartAPIInstance(".")
	for {
		time.Sleep(time.Minute)
	}
}
