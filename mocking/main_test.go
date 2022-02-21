package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	SLEEP = "sleep"
	WRITE = "write"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, SLEEP)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, WRITE)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("mocking buffer and sleeper", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"

		if got != want {
			t.Errorf("got %v want %v\n", got, want)
		}
	})

	t.Run("tracking call and sleep sequence", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperations{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{WRITE, SLEEP, WRITE, SLEEP, WRITE, SLEEP, WRITE}

		if !reflect.DeepEqual(spySleeperPrinter.Calls, want) {
			t.Errorf("got %v want %v\n", spySleeperPrinter.Calls, want)
		}
	})

	t.Run("configurable sleeper", func(t *testing.T) {
		sleepTime := 4 * time.Second
		spyTime := &SpyTime{}
		configfurableSleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		configfurableSleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("got %v want %v\n", spyTime.durationSlept, sleepTime)
		}

	})
}
