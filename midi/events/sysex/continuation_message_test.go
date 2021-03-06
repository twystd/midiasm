package sysex

import (
	"bufio"
	"bytes"
	"github.com/twystd/midiasm/midi/context"
	"github.com/twystd/midiasm/midi/types"
	"reflect"
	"testing"
)

func TestParseContinuationMessage(t *testing.T) {
	ctx := context.NewContext()
	ctx.Casio = true

	r := bufio.NewReader(bytes.NewReader([]byte{0x05, 0x7e, 0x00, 0x09, 0x01, 0xf7}))

	event, err := Parse(reader{r}, 0xf7, ctx)
	if err != nil {
		t.Fatalf("Unexpected SysEx continuation message parse error: %v", err)
	}

	if event == nil {
		t.Fatalf("Unexpected SysEx continuation message parse error - returned %v", event)
	}

	message, ok := event.(*SysExContinuationMessage)
	if !ok {
		t.Fatalf("SysEx continuation message parse error - returned %T", event)
	}

	expected := types.Hex{0x7e, 0x00, 0x09, 0x01}
	if !reflect.DeepEqual(message.Data, expected) {
		t.Errorf("Invalid SysEx continuation message data - expected:%v, got: %v", expected, message.Data)
	}

	if ctx.Casio {
		t.Errorf("context Casio flag not reset")
	}
}
