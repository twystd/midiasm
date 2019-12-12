package midievent

import (
	"fmt"
	"github.com/twystd/midiasm/midi/context"
	"io"
)

type ChannelPressure struct {
	MidiEvent
	Pressure byte
}

func NewChannelPressure(event *MidiEvent, r io.ByteReader) (*ChannelPressure, error) {
	if event.Status&0xF0 != 0xD0 {
		return nil, fmt.Errorf("Invalid ChannelPressure status (%02x): expected 'D0'", event.Status&0x80)
	}

	pressure, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	return &ChannelPressure{
		MidiEvent: *event,
		Pressure:  pressure,
	}, nil
}

func (e *ChannelPressure) Render(ctx *context.Context, w io.Writer) {
	fmt.Fprintf(w, "%s %-16s channel:%d pressure:%d", e.MidiEvent, "ChannelPressure", e.Channel, e.Pressure)
}