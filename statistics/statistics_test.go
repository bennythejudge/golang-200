package statistics

import (
	"testing"
	"time"
)

const (
	statPeriod    = 500 * time.Millisecond
	expectedCount = 2 // what is this value?
	emptyCount    = 0
)

func TestStatistics(t *testing.T) {

	statistics := NewStatistics(2 * statPeriod)

	//TODO: implement this test to test the statistics middleware (test the value of its counter)
	go func() {
		statistics.PlusOne()
	}()

	go func() {
		statistics.PlusOne()
	}()

	//TODO: you will need time.Sleep(statPeriod) because of the statistics ticker
	time.Sleep(statPeriod)

	if statistics.counter != expectedCount {
		t.Errorf("stats count %d differs from expected %d", statistics.counter, expectedCount)
	}

	time.Sleep(2 * statPeriod)

	if statistics.counter != emptyCount {
		t.Errorf("stats counter %d differs from expected emptyCount %d", statistics.counter, emptyCount)
	}

}
