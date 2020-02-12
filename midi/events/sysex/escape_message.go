package sysex

import (
	"fmt"
	"github.com/twystd/midiasm/midi/context"
	"github.com/twystd/midiasm/midi/types"
	"io"
)

type SysExEscapeMessage struct {
	Tag    string
	Status types.Status
	Data   types.Hex
}

func NewSysExEscapeMessage(r io.ByteReader, status types.Status, ctx *context.Context) (*SysExEscapeMessage, error) {
	if status != 0xf7 {
		return nil, fmt.Errorf("Invalid SysExEscapeMessage event type (%02x): expected 'F7'", status)
	}

	if ctx.Casio() {
		return nil, fmt.Errorf("F7 is not valid for SysExEscapeMessage event in Casio mode")
	}

	data, err := read(r)
	if err != nil {
		return nil, err
	}

	return &SysExEscapeMessage{
		Tag:    "SysExEscape",
		Status: status,
		Data:   data,
	}, nil
}
