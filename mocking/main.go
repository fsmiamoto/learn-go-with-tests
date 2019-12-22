package main

import (
	"os"
	"time"
)

func main() {
	slepper := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}

	Countdown(os.Stdout, slepper)
}
