package clock

import (
	"fmt"
	"time"
)

const dayInMin = 1440

type Clock int

func calc(add int, h int, m int) (int, int) {
	now := time.Now()

	// Neues Datum mit Stunde, Minute, Sekunde, Nanosekunde = 0
	t := time.Date(now.Year(), now.Month(), now.Day(), h, m, 0, 0, now.Location())
	t = t.Add(time.Duration(add) * time.Minute)
	return t.Hour(), t.Minute()

}
func set(h int, m int) (int, int) {
	now := time.Now()

	// Neues Datum mit Stunde, Minute, Sekunde, Nanosekunde = 0
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	t = t.Add(time.Duration(h) * time.Hour)
	t = t.Add(time.Duration(m) * time.Minute)
	return t.Hour(), t.Minute()

}

func New(h, m int) Clock {
	h, m = set(h, m)

	c := Clock((h*60 + m) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)

}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
