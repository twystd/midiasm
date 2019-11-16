package metaevent

import (
	"fmt"
	"io"
)

type TimeSignature struct {
	MetaEvent
	numerator               uint8
	denominator             uint8
	ticksPerClick           uint8
	thirtySecondsPerQuarter uint8
}

func NewTimeSignature(event MetaEvent, data []byte) (*TimeSignature, error) {
	if event.status != 0xff {
		return nil, fmt.Errorf("Invalid TimeSignature status (%02x): expected 'ff'", event.status)
	}

	if event.eventType != 0x58 {
		return nil, fmt.Errorf("Invalid TimeSignature event type (%02x): expected '58'", event.eventType)
	}

	if event.length != 4 {
		return nil, fmt.Errorf("Invalid TimeSignature length (%d): expected '3'", event.length)
	}

	numerator := data[0]
	denominator := data[1]
	ticksPerClick := data[2]
	thirtySecondsPerQuarter := data[3]

	return &TimeSignature{
		MetaEvent:               event,
		numerator:               numerator,
		denominator:             denominator,
		ticksPerClick:           ticksPerClick,
		thirtySecondsPerQuarter: thirtySecondsPerQuarter,
	}, nil
}

func (e *TimeSignature) Render(w io.Writer) {
	fmt.Fprintf(w, "   ")
	for _, b := range e.bytes {
		fmt.Fprintf(w, "%02x ", b)
	}
	fmt.Fprintf(w, "                            ")

	fmt.Fprintf(w, "%02x/%-16s delta:%-10d numerator:%d denominator:%d ticks/click:%d 1/32-per-quarter:%d\n", e.eventType, "TimeSignature", e.delta, e.numerator, e.denominator, e.ticksPerClick, e.thirtySecondsPerQuarter)
}
