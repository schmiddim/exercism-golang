package meteorology

import "fmt"

type TemperatureUnit int

func (t TemperatureUnit) String() string {
	fmtString := "°F"
	if t == Celsius {
		fmtString = "°C"
	}

	return fmtString
}

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

func (s SpeedUnit) String() string {
	fmtString := "km/h"

	if s == MilesPerHour {
		fmtString = "mph"
	}
	return fmtString
}

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {

	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {

	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity",
		m.location,
		m.temperature,
		m.windDirection,
		m.windSpeed,
		m.humidity,
	)
}
