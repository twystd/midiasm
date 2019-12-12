package metaevent

import (
	"fmt"
	"github.com/twystd/midiasm/midi/context"
	"io"
)

type TimeSignature struct {
	MetaEvent
	Numerator               uint8
	Denominator             uint8
	TicksPerClick           uint8
	ThirtySecondsPerQuarter uint8
}

func NewTimeSignature(event *MetaEvent, r io.ByteReader) (*TimeSignature, error) {
	if event.Type != 0x58 {
		return nil, fmt.Errorf("Invalid TimeSignature event type (%02x): expected '58'", event.Type)
	}

	data, err := read(r)
	if err != nil {
		return nil, err
	} else if len(data) != 4 {
		return nil, fmt.Errorf("Invalid TimeSignature length (%d): expected '4'", len(data))
	}

	numerator := data[0]
	ticksPerClick := data[2]
	thirtySecondsPerQuarter := data[3]

	denominator := uint8(1)
	for i := uint8(0); i < data[1]; i++ {
		denominator *= 2
	}

	return &TimeSignature{
		MetaEvent:               *event,
		Numerator:               numerator,
		Denominator:             denominator,
		TicksPerClick:           ticksPerClick,
		ThirtySecondsPerQuarter: thirtySecondsPerQuarter,
	}, nil
}

func (e *TimeSignature) Render(ctx *context.Context, w io.Writer) {
	fmt.Fprintf(w, "%s %-16s %d:%d, %d ticks-per-click, %d/32-per-quarter", e.MetaEvent, "TimeSignature", e.Numerator, e.Denominator, e.TicksPerClick, e.ThirtySecondsPerQuarter)
}