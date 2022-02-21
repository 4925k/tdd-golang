package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	FINALWORD = "Go!"
	COUNTDOWN = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	configurableSleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, configurableSleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := COUNTDOWN; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(out, FINALWORD)
}
