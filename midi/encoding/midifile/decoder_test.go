package midifile

import (
	"github.com/twystd/midiasm/midi"
	"github.com/twystd/midiasm/midi/events"
	"github.com/twystd/midiasm/midi/events/meta"
	"github.com/twystd/midiasm/midi/events/midi"
	"github.com/twystd/midiasm/midi/types"
	"reflect"
	"testing"
)

var SMF0 = []byte{
	0x4d, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x60,

	0x4d, 0x54, 0x72, 0x6b, 0x00, 0x00, 0x00, 0x6f,
	0x00, 0xff, 0x51, 0x03, 0x07, 0xa1, 0x20,
	0x00, 0xff, 0x54, 0x05, 0x4d, 0x2d, 0x3b, 0x07, 0x27,
	0x00, 0xff, 0x00, 0x02, 0x00, 0x17,
	0x00, 0xff, 0x01, 0x0d, 0x54, 0x68, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x54, 0x68, 0x61, 0x74,
	0x00, 0xff, 0x02, 0x04, 0x54, 0x68, 0x65, 0x6d,
	0x00, 0xff, 0x03, 0x0f, 0x41, 0x63, 0x6f, 0x75, 0x73, 0x74, 0x69, 0x63, 0x20, 0x47, 0x75, 0x69, 0x74, 0x61, 0x72,
	0x00, 0xff, 0x04, 0x0a, 0x44, 0x69, 0x64, 0x67, 0x65, 0x72, 0x69, 0x64, 0x6f, 0x6f,
	0x00, 0xff, 0x59, 0x02, 0x00, 0x01,
	0x00, 0xff, 0x7f, 0x06, 0x00, 0x00, 0x3b, 0x3a, 0x4c, 0x5e,
	0x00, 0x91, 0x31, 0x48,
	0x00, 0x3c, 0x4c,
	0x00, 0x81, 0x31, 0x64,
	0x00, 0xff, 0x2f, 0x00,
}

var SMF1 = []byte{
	0x4d, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x02, 0x00, 0x60,

	0x4d, 0x54, 0x72, 0x6b, 0x00, 0x00, 0x00, 0x21,
	0x00, 0xff, 0x03, 0x09, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x31,
	0x00, 0xff, 0x51, 0x03, 0x07, 0xa1, 0x20,
	0x00, 0xff, 0x54, 0x05, 0x4d, 0x2d, 0x3b, 0x07, 0x27,
	0x00, 0xff, 0x2f, 0x00,

	0x4d, 0x54, 0x72, 0x6b, 0x00, 0x00, 0x00, 0x5f,
	0x00, 0xff, 0x00, 0x02, 0x00, 0x17,
	0x00, 0xff, 0x01, 0x0d, 0x54, 0x68, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x54, 0x68, 0x61, 0x74,
	0x00, 0xff, 0x02, 0x04, 0x54, 0x68, 0x65, 0x6d,
	0x00, 0xff, 0x03, 0x0f, 0x41, 0x63, 0x6f, 0x75, 0x73, 0x74, 0x69, 0x63, 0x20, 0x47, 0x75, 0x69, 0x74, 0x61, 0x72,
	0x00, 0xff, 0x04, 0x0a, 0x44, 0x69, 0x64, 0x67, 0x65, 0x72, 0x69, 0x64, 0x6f, 0x6f,
	0x00, 0xff, 0x59, 0x02, 0x00, 0x01,
	0x00, 0xff, 0x7f, 0x06, 0x00, 0x00, 0x3b, 0x3a, 0x4c, 0x5e,
	0x00, 0x91, 0x31, 0x48,
	0x00, 0x3c, 0x4c,
	0x00, 0x81, 0x31, 0x64,
	0x00, 0xff, 0x2f, 0x00,
}

var MTHD0 = midi.MThd{
	Tag:      "MThd",
	Length:   6,
	Format:   0,
	Tracks:   1,
	Division: 96,
	PPQN:     96,
	Bytes:    []byte{0x4d, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x60},
}

var MTHD1 = midi.MThd{
	Tag:      "MThd",
	Length:   6,
	Format:   1,
	Tracks:   2,
	Division: 96,
	PPQN:     96,
	Bytes:    []byte{0x4d, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x02, 0x00, 0x60},
}

var MTRK0 = []*midi.MTrk{
	&midi.MTrk{
		Tag:         "MTrk",
		TrackNumber: 0,
		Length:      111,
		Bytes:       []byte{0x4d, 0x54, 0x72, 0x6b, 0x00, 0x00, 0x00, 0x6f},
		Events: []*events.Event{
			&tempo,
			&smpteOffset,
			&sequenceNumber,
			&text,
			&copyright,
			&acousticGuitar,
			&didgeridoo,
			&aMinor,
			&motu,
			&noteOnCS2,
			&noteOnC3,
			&noteOffCS2,
			&endOfTrack,
		},
	},
}

var MTRK1 = []*midi.MTrk{
	&midi.MTrk{
		Tag:         "MTrk",
		TrackNumber: 0,
		Length:      33,
		Bytes:       []byte{0x4d, 0x54, 0x72, 0x6b, 0x00, 0x00, 0x00, 0x21},
		Events: []*events.Event{
			&example1,
			&tempo,
			&smpteOffset,
			&endOfTrack,
		},
	},

	&midi.MTrk{
		Tag:         "MTrk",
		TrackNumber: 1,
		Length:      95,
		Bytes:       []byte{0x4d, 0x54, 0x72, 0x6b, 0x00, 0x00, 0x00, 0x5f},
		Events: []*events.Event{
			&sequenceNumber,
			&text,
			&copyright,
			&acousticGuitar,
			&didgeridoo,
			&aMinor,
			&motu,
			&noteOnCS2,
			&noteOnC3,
			&noteOffCS2,
			&endOfTrack,
		},
	},
}

func TestDecodeFormat0(t *testing.T) {
	testDecode(t, SMF0, &MTHD0, MTRK0)
}

func TestDecodeFormat1(t *testing.T) {
	testDecode(t, SMF1, &MTHD1, MTRK1)
}

func testDecode(t *testing.T, bytes []byte, mthd *midi.MThd, tracks []*midi.MTrk) {
	decoder := NewDecoder()

	smf, err := decoder.Decode(bytes)
	if err != nil {
		t.Fatalf("unexpected error decoding valid MIDI file: %v", err)
	}

	if smf == nil {
		t.Fatalf("decoder returned a 'nil' result for MIDI file")
	}

	if !reflect.DeepEqual(*smf.MThd, *mthd) {
		t.Errorf("MThd incorrectly unmarshaled\n   expected:%v\n   got:     %v", *mthd, *smf.MThd)
	}

	for i, track := range tracks {
		mtrk := smf.Tracks[i]

		if mtrk.Tag != track.Tag {
			t.Errorf("MTrk[%d]: incorrectly unmarshaled 'Tag'\n   expected:%v\n   got:     %v", i, track.Tag, mtrk.Tag)
		}

		if mtrk.TrackNumber != tracks[i].TrackNumber {
			t.Errorf("MTrk[%d]: incorrectly unmarshaled 'TrackNumber'\n   expected:%v\n   got:     %v", i, track.TrackNumber, mtrk.TrackNumber)
		}

		if mtrk.Length != tracks[i].Length {
			t.Errorf("MTrk[%d]: incorrectly unmarshaled 'Length'\n   expected:%v\n   got:     %v", i, track.Length, mtrk.Length)
		}

		if !reflect.DeepEqual(mtrk.Bytes[0:8], tracks[i].Bytes[0:8]) {
			t.Errorf("MTrk[%d]: incorrectly unmarshaled 'Bytes'\n   expected:%v\n   got:     %v", i, track.Bytes[0:8], mtrk.Bytes[0:8])
		}

		if len(mtrk.Events) != len(tracks[i].Events) {
			t.Errorf("MTrk[%d]: incorrectly unmarshaled 'Events'\n   expected:%v\n   got:     %v", i, len(track.Events), len(mtrk.Events))
		} else {
			for j, e := range mtrk.Events {
				if !reflect.DeepEqual(e, tracks[i].Events[j]) {
					t.Errorf("MTrk[%d]: incorrectly unmarshaled event\n   expected:%#v\n   got:     %#v", i, track.Events[j], e)
				}
			}
		}
	}
}

// TEST EVENTS

var sequenceNumber = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x00, 0x02, 0x00, 0x17},
	Event: &metaevent.SequenceNumber{
		Tag:            "SequenceNumber",
		Status:         0xff,
		Type:           types.MetaEventType(0x00),
		SequenceNumber: 23,
	},
}

var text = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x01, 0x0d, 0x54, 0x68, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x54, 0x68, 0x61, 0x74},
	Event: &metaevent.Text{
		Tag:    "Text",
		Status: 0xff,
		Type:   types.MetaEventType(0x01),
		Text:   "This and That",
	},
}

var copyright = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x02, 0x04, 0x54, 0x68, 0x65, 0x6d},
	Event: &metaevent.Copyright{
		Tag:       "Copyright",
		Status:    0xff,
		Type:      types.MetaEventType(0x02),
		Copyright: "Them",
	},
}

var example1 = events.Event{
	Bytes: types.Hex{0x0, 0xff, 0x3, 0x9, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x31},
	Event: &metaevent.TrackName{
		Tag:    "TrackName",
		Status: 0xff,
		Type:   types.MetaEventType(0x03),
		Name:   "Example 1",
	},
}

var acousticGuitar = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x03, 0x0f, 0x41, 0x63, 0x6f, 0x75, 0x73, 0x74, 0x69, 0x63, 0x20, 0x47, 0x75, 0x69, 0x74, 0x61, 0x72},
	Event: &metaevent.TrackName{
		Tag:    "TrackName",
		Status: 0xff,
		Type:   types.MetaEventType(0x03),
		Name:   "Acoustic Guitar",
	},
}

var didgeridoo = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x04, 0x0a, 0x44, 0x69, 0x64, 0x67, 0x65, 0x72, 0x69, 0x64, 0x6f, 0x6f},
	Event: &metaevent.InstrumentName{
		Tag:    "InstrumentName",
		Status: 0xff,
		Type:   types.MetaEventType(0x04),
		Name:   "Didgeridoo",
	},
}

var aMinor = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x59, 0x02, 0x00, 0x01},
	Event: &metaevent.KeySignature{
		Tag:         "KeySignature",
		Status:      0xff,
		Type:        types.MetaEventType(0x59),
		Accidentals: 0,
		KeyType:     1,
		Key:         "A minor",
	},
}

var motu = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x7f, 0x06, 0x00, 0x00, 0x3b, 0x3a, 0x4c, 0x5e},
	Event: &metaevent.SequencerSpecificEvent{
		Tag:    "SequencerSpecificEvent",
		Status: 0xff,
		Type:   types.MetaEventType(0x7f),
		Manufacturer: types.Manufacturer{
			ID:     []byte{0x00, 0x00, 0x3b},
			Region: "American",
			Name:   "Mark Of The Unicorn (MOTU)",
		},
		Data: []byte{0x3a, 0x4c, 0x5e},
	},
}

var noteOnCS2 = events.Event{
	Bytes: types.Hex{0x00, 0x91, 0x31, 0x48},
	Event: &midievent.NoteOn{
		Tag:     "NoteOn",
		Status:  0x91,
		Channel: types.Channel(0x01),
		Note: types.Note{
			Value: 49,
			Name:  "C♯2",
			Alias: "C♯2",
		},
		Velocity: 72,
	},
}

var noteOnC3 = events.Event{
	Bytes: types.Hex{0x00, 0x3c, 0x4c},
	Event: &midievent.NoteOn{
		Tag:     "NoteOn",
		Status:  0x91,
		Channel: types.Channel(0x01),
		Note: types.Note{
			Value: 60,
			Name:  "C3",
			Alias: "C3",
		},
		Velocity: 76,
	},
}

var noteOffCS2 = events.Event{
	Bytes: types.Hex{0x00, 0x81, 0x31, 0x64},
	Event: &midievent.NoteOff{
		Tag:     "NoteOff",
		Status:  0x81,
		Channel: types.Channel(0x01),
		Note: types.Note{
			Value: 49,
			Name:  "C♯2",
			Alias: "C♯2",
		},
		Velocity: 100,
	},
}

var tempo = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x51, 0x03, 0x07, 0xa1, 0x20},
	Event: &metaevent.Tempo{
		Tag:    "Tempo",
		Status: 0xff,
		Type:   types.MetaEventType(0x51),
		Tempo:  500000,
	},
}

var smpteOffset = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x54, 0x05, 0x4d, 0x2d, 0x3b, 0x07, 0x27},
	Event: &metaevent.SMPTEOffset{
		Tag:              "SMPTEOffset",
		Status:           0xff,
		Type:             types.MetaEventType(0x54),
		Hour:             13,
		Minute:           45,
		Second:           59,
		FrameRate:        25,
		Frames:           7,
		FractionalFrames: 39,
	},
}

var endOfTrack = events.Event{
	Bytes: types.Hex{0x00, 0xff, 0x2f, 0x00},
	Event: &metaevent.EndOfTrack{
		Tag:    "EndOfTrack",
		Status: 0xff,
		Type:   types.MetaEventType(0x2f),
	},
}
