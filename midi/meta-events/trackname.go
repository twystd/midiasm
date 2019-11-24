package metaevent

import (
	"fmt"
	"io"
)

type TrackName struct {
	MetaEvent
	Name string
}

func NewTrackName(event *MetaEvent, r io.ByteReader) (*TrackName, error) {
	if event.eventType != 0x03 {
		return nil, fmt.Errorf("Invalid TrackName event type (%02x): expected '03'", event.eventType)
	}

	name, err := read(r)
	if err != nil {
		return nil, err
	}

	return &TrackName{
		MetaEvent: *event,
		Name:      string(name),
	}, nil
}

func (e *TrackName) Render(w io.Writer) {
	fmt.Fprintf(w, "%s %-16s name:%s", e.MetaEvent, "TrackName", e.Name)
}
