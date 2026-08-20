package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	hours   int
	minutes int
}

func mod(a, m int) int {
	return ((a % m) + m) % m
}

func New(h, m int) Clock {
	h, m = resolveClock(h, m)

	return Clock{
		hours:   h,
		minutes: m,
	}
}

func (c Clock) Add(m int) Clock {
	h, m := resolveClock(c.hours, c.minutes+m)

	c.hours = h
	c.minutes = m

	return c
}

func (c Clock) Subtract(m int) Clock {
	h, m := resolveClock(c.hours, c.minutes-m)

	c.hours = h
	c.minutes = m

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

func resolveClock(h, m int) (int, int) {
	if m >= 60 {
		h += int(m / 60)
		m = mod(m, 60)
	}

	if m < 0 {
		m = 60 + m
		if m < 0 {
			h += int(m/60) - 2
			m = mod(m, 60)
		} else {
			h -= 1
		}
	}

	h = mod(h, 24)

	return h, m
}
