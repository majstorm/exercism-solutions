package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	hrs := h + m/60
	m %= 60
	if m < 0 {
		m = 60 + m
		hrs--
	}
	if hrs < 0 {
		hrs = 24 + hrs%24
	}
	return Clock{
		hour:   hrs % 24,
		minute: m,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
