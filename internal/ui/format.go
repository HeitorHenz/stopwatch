package ui

import (
	"fmt"
	"time"
)

func format(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	d = d.Round(10 * time.Microsecond)

	h := d / time.Hour
	m := d / time.Minute % 60
	s := d / time.Second % 60
	cs := d / (10 * time.Millisecond) % 100

	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d.%02d", h, m, s, cs)
	}
	return fmt.Sprintf("%02d:%02d.%02d", m, s, cs)
}
