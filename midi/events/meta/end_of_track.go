package metaevent

import (
	"fmt"
	"github.com/twystd/midiasm/midi/types"
	"io"
)

type EndOfTrack struct {
	Tag    string
	Status types.Status
	Type   types.MetaEventType
}

func NewEndOfTrack(r io.ByteReader, status types.Status, eventType types.MetaEventType) (*EndOfTrack, error) {
	if eventType != 0x2f {
		return nil, fmt.Errorf("Invalid EndOfTrack event type (%02x): expected '2f'", eventType)
	}

	data, err := read(r)
	if err != nil {
		return nil, err
	} else if len(data) != 0 {
		return nil, fmt.Errorf("Invalid EndOfTrack length (%d): expected '0'", len(data))
	}

	return &EndOfTrack{
		Tag:    "EndOfTrack",
		Status: status,
		Type:   eventType,
	}, nil
}
