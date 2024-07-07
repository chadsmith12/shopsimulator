package simulation_test

import (
	"testing"
	"time"

	"github.com/chadsmith12/coffeeshop/pkgs/simulation"
)

type dayTestCase struct {
    name string
    day simulation.SimulatedDay
    expected time.Duration
}

func TestDaySimulation(t *testing.T) {
    testCases := []dayTestCase{
	{
	    name: "24 seconds in a day gives 1 second for hour",
	    day: simulation.NewSimulatedDay(time.Second * 24),
	    expected: time.Second,
	},
    }
    for _, testCase := range testCases {
	t.Run(testCase.name, func (t *testing.T) {
	    actual := testCase.day.Hour()
	    if actual != testCase.expected {
		t.Errorf("SimulatedDay.Hour() = %d wanted %d\n", actual, testCase.expected)
	    }
	})	
    }
}
