package simulation 

import (
	"time"
)

const dayDuration time.Duration = 24 * time.Hour 

// A SimulatedDay allows you to say how long a day should be in a simulation.
// This alows you to work with duration as you would a normal time without having to scale the time.
// Example: A SimulatedDay of time.Second * 24 says that a day  is 24 seconds long.
// Meaning that an Hour in the simulated day is 1 second.
type SimulatedDay struct {
    timeLength time.Duration
}

// Creates a new SimulatedDay that is the length of the duration specified.
func NewSimulatedDay(dayLength time.Duration) SimulatedDay {
    return SimulatedDay{
    	timeLength: dayLength,
    }
}

// Gets how long a millisecond is inside of this simulated day as a time.Duration
func (simulatedDay SimulatedDay) Milliseconds() time.Duration {
    return simulatedDay.timeLength / time.Duration(dayDuration.Milliseconds())
}

// Gets how long a second is inside of this simulated day as a time.Duration
func (simulatedDay SimulatedDay) Seconds() time.Duration {
    return simulatedDay.timeLength / time.Duration(dayDuration.Seconds())
}

// Gets how long a minute is inside of this simulated day as a time.Duration
func (simulatedDay SimulatedDay) Minute() time.Duration {
    return simulatedDay.timeLength / time.Duration(dayDuration.Minutes())
}

// Gets how long a hour is inside of this simulated day as a time.Duration
func (simulatedDay SimulatedDay) Hour() time.Duration {
    return simulatedDay.timeLength / time.Duration(dayDuration.Hours()) 
}
