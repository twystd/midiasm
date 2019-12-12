package metaevent

import (
	"bytes"
	"github.com/twystd/midiasm/midi/context"
	"github.com/twystd/midiasm/midi/events"
	"strings"
	"testing"
)

func TestMIDIPortRender(t *testing.T) {
	ctx := context.Context{
		Scale: context.Sharps,
	}

	metaevent := MetaEvent{
		events.Event{76, 12, 0xff, []byte{0x00, 0xff, 0x21, 0x01, 57}},
		0x21,
	}

	var eventlist = []struct {
		event    events.IEvent
		expected string
	}{
		{&MIDIPort{metaevent, 57}, "MIDIPort         57"},
	}

	for _, v := range eventlist {
		w := new(bytes.Buffer)

		v.event.Render(&ctx, w)

		if !strings.HasSuffix(w.String(), v.expected) {
			t.Errorf("MIDIPort rendered incorrectly\nExpected: '%s'\ngot:      '%s'", v.expected, w.String())
		}
	}
}