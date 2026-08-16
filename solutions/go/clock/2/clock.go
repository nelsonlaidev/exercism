package clock

import "fmt"

// Define the Clock type here.
type Clock int

func New(h, m int) Clock {
	total := (h*60 + m) % (24 * 60)
	if total < 0 {
		total += 24 * 60
	}
	return Clock(total)
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
