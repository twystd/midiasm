package metaevent

import (
	"fmt"
	"io"
)

type InstrumentName struct {
	Tag string
	MetaEvent
	Name string
}

func NewInstrumentName(event *MetaEvent, r io.ByteReader) (*InstrumentName, error) {
	if event.Type != 0x04 {
		return nil, fmt.Errorf("Invalid InstrumentName event type (%02x): expected '04'", event.Type)
	}

	name, err := read(r)
	if err != nil {
		return nil, err
	}

	return &InstrumentName{
		Tag:       "InstrumentName",
		MetaEvent: *event,
		Name:      string(name),
	}, nil
}
