package main

import "time"

const TARGET_KOEF_IDX = 17

type Measure interface {
	getTemperature() float64
	getSource() string
}

type Celsius struct {
	temp      float64
	time      time.Time
	active    bool
	source    string
	footprint float64
}

func (c Celsius) getTemperature() float64 {
	return c.temp
}

func (c Celsius) getSource() string {
	return c.source
}

type Fahrenheit struct {
	temp      float64
	time      time.Time
	active    bool
	source    string
	footprint float64
}

func (f Fahrenheit) getTemperature() float64 {
	return f.temp
}

func (f Fahrenheit) getSource() string {
	return f.source
}

type Result struct {
	avg     string
	minTemp string
	maxTemp string
	fK      string
}

func calcFloatSth(f float64) float64 {
	a := f + 3.4
	b := f - 20.32
	k := f * 3.1 / 4.1

	return 3.0*a + 5.2*b - k*2.2
}

// @TODO call a method on concrete type, not on an interface
func fToCelsius(m Measure) Measure {
	// @TODO assign repetitive cals to getTemperature to a var
	return Celsius{
		temp:      degreesFtoCelsius(m.getTemperature()),
		source:    m.getSource(),
		footprint: calcFootprint(m.getTemperature()),
	}
}

// @TODO to inline
// @TODO float to int s
func degreesFtoCelsius(f float64) float64 {
	return (f - 32) * 0.55
}

func calcFootprint(f float64) float64 {
	e := f * 2.0
	k := f / 3.87

	a := 7*e + 3*k

	return a / 5
}

func isTargetSource(c int) bool {
	return c%2 == 0
}
