package metaevent

import (
	"fmt"
	"io"
)

type Text struct {
	Tag string
	MetaEvent
	Text string
}

func NewText(event *MetaEvent, r io.ByteReader) (*Text, error) {
	if event.Type != 0x01 {
		return nil, fmt.Errorf("Invalid Text event type (%02x): expected '01'", event.Type)
	}

	data, err := read(r)
	if err != nil {
		return nil, err
	}

	return &Text{
		Tag:       "Text",
		MetaEvent: *event,
		Text:      string(data),
	}, nil
}
