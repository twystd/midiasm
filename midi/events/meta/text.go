package metaevent

import (
	"fmt"
	"github.com/twystd/midiasm/midi/events"
	"github.com/twystd/midiasm/midi/types"
)

type Text struct {
	Tag    string
	Status types.Status
	Type   types.MetaEventType
	Text   string
}

func NewText(r events.EventReader, status types.Status, eventType types.MetaEventType) (*Text, error) {
	if eventType != 0x01 {
		return nil, fmt.Errorf("Invalid Text event type (%02x): expected '01'", eventType)
	}

	data, err := r.ReadVLF()
	if err != nil {
		return nil, err
	}

	return &Text{
		Tag:    "Text",
		Status: status,
		Type:   eventType,
		Text:   string(data),
	}, nil
}
