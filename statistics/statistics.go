package statistics

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"time"
)

const (
	statisticsChannelSize = 1000
)

// TODO define a Statistics struct with an uint8 chan, an uint32 counter, a start time and logging period duration
// Statistics is the worker to persist the request statistics
type Statistics struct {
	statistics    chan uint8
	counter       uint32
	start         time.Time
	loggingPeriod time.Duration
}

// NewStatistics creates a new statistics structure and launches its worker routine
func NewStatistics(loggingPeriod time.Duration) *Statistics {
	// TODO build a new Statistics object with a sized channel, initialized counter and start time
	// TODO and logging period as param
	sw := Statistics{
		statistics:    make(chan uint8, statisticsChannelSize),
		counter:       0,
		start:         time.Now(),
		loggingPeriod: loggingPeriod,
	}

	// TODO launch the run in a separate Go routine in background
	go sw.run()

	// TODO return the initialized and started object
	return &sw
}

// PlusOne is used to send a statistics hit increment
func (sw *Statistics) PlusOne() {
	// TODO push a hit in the statistics channel
	sw.statistics <- uint8(1) // need to make sure the value has the right type
}

func (sw *Statistics) run() {
	// TODO build a new time Ticker from the logging period
	ticker := time.NewTicker(sw.loggingPeriod)

	// TODO build a infinite loop and the channel selection inside
	for {
		select {
		// TODO build a first select case from the statistics channel
		// TODO add the hit count to the counter and log it as debug level
		case stat := <-sw.statistics:
			logger.
				WithField("stat", stat).
				Debug("new count received")

		// TODO build a second case on the time Ticker chan
		case <-ticker.C:
			// TODO retrieve the elapsed time since start
			elapsed := time.Since(sw.start)
			// TODO log the hit/sec rate
			logger.
				WithField("elapsed time", elapsed).
				WithField("count", sw.counter).
				Warn("monitoring request")
			// TODO reset the counter and the start time
			sw.counter = 0
			sw.start = time.Now()
			fmt.Println("inside the run")
		}
	}
}
