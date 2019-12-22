package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Prints 3 to Go!", func(t *testing.T) {
		buf := &bytes.Buffer{}
		spy := &SpySleeper{}

		Countdown(buf, spy)

		got := buf.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spy.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spy.Calls)
		}
	})

	t.Run("Sleep before every print", func(t *testing.T) {
		spy := &CountdownOpsSpy{}
		Countdown(spy, spy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spy.Calls) {
			t.Errorf("wanted calls %v but got %v", want, spy.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
