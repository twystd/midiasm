package metaevent

import (
	"fmt"
	"github.com/twystd/midiasm/midi/event"
	"io"
)

type MetaEvent struct {
	event.Event
	eventType byte
	length    uint32
	bytes     []byte
}

func Parse(e event.Event, x []byte, r io.ByteReader) (event.IEvent, error) {
	if e.Status != 0xff {
		return nil, fmt.Errorf("Invalid MetaEvent tag (%02x): expected 'ff'", e.Status)
	}

	bytes := make([]byte, 0)
	bytes = append(bytes, x...)

	b, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, b)
	eventType := b & 0x7f

	l, m, err := vlq(r)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, m...)

	d := make([]byte, l)
	for i := uint32(0); i < l; i++ {
		if b, err = r.ReadByte(); err != nil {
			return nil, err
		} else {
			d[i] = b
		}
	}
	bytes = append(bytes, d...)

	event := MetaEvent{
		Event:     e,
		eventType: eventType,
		length:    l,
		bytes:     bytes,
	}

	switch eventType {
	case 0x03:
		return NewTrackName(event, d)

	case 0x2f:
		return NewEndOfTrack(event, d)

	case 0x51:
		return NewTempo(event, d)

	case 0x58:
		return NewTimeSignature(event, d)

	case 0x59:
		return NewKeySignature(event, d)
	}

	return nil, fmt.Errorf("Unrecognised META event: %02x", eventType)
}

func vlq(r io.ByteReader) (uint32, []byte, error) {
	l := uint32(0)
	bytes := make([]byte, 0)

	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, nil, err
		}
		bytes = append(bytes, b)

		l <<= 8
		l += uint32(b & 0x7f)

		if b&0x80 == 0 {
			break
		}
	}

	return l, bytes, nil
}
