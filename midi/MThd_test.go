package midi

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMThdUnmarshal(t *testing.T) {
	bytes := []byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x11, 0x00, 0x60}
	expected := MThd{
		Tag:           "MThd",
		Length:        6,
		Format:        1,
		Tracks:        17,
		Division:      0x0060,
		PPQN:          96,
		SMPTETimeCode: false,
		SubFrames:     0,
		FPS:           0,
		Bytes:         []byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x11, 0x00, 0x60},
	}

	mthd := MThd{}
	if err := mthd.UnmarshalBinary(bytes); err != nil {
		t.Fatalf("Unexpected error unmarshaling MThd: %v", err)
	}

	if !reflect.DeepEqual(mthd, expected) {
		t.Errorf("MThd incorrectly unmarshaled\n   expected:%+v\n   got:     %+v", expected, mthd)
	}
}

func TestMThdUnmarshalSMTPE(t *testing.T) {
	bytes := []byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x11, 0xe7, 0x28}
	expected := MThd{
		Tag:           "MThd",
		Length:        6,
		Format:        1,
		Tracks:        17,
		Division:      0xe728,
		PPQN:          0,
		SMPTETimeCode: true,
		SubFrames:     40,
		FPS:           25,
		Bytes:         []byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x11, 0xe7, 0x28},
	}

	mthd := MThd{}
	if err := mthd.UnmarshalBinary(bytes); err != nil {
		t.Fatalf("Unexpected error unmarshaling MThd: %v", err)
	}

	if !reflect.DeepEqual(mthd, expected) {
		t.Errorf("MThd incorrectly unmarshaled\n   expected:%+v\n   got:     %+v", expected, mthd)
	}
}

func TestMThdUnmarshalInvalidBytes(t *testing.T) {
	mthd := MThd{}
	bytes := [][]byte{
		[]byte{0x4D, 0x54, 0x72, 0x6b, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x11, 0x00, 0x60},
		[]byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x11, 0x00, 0x60},
		[]byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x03, 0x00, 0x11, 0x00, 0x60},
		[]byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x11, 0x00},
	}

	for _, b := range bytes {
		if err := mthd.UnmarshalBinary(b); err == nil {
			t.Fatalf("Expected error unmarshaling MThd: got %v", err)
		}
	}
}

func TestMThdPrint(t *testing.T) {
	expected := "4D 54 68 64 00 00 00 06 00 01 00 11 00 60   MThd length:6, format:1, tracks:17, metrical time:96 ppqn"

	mthd := MThd{
		Tag:      "MThd",
		Length:   6,
		Format:   1,
		Tracks:   17,
		Division: 96,
		PPQN:     96,
		Bytes:    []byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x11, 0x00, 0x60},
	}

	w := new(bytes.Buffer)

	mthd.Print(w)

	if w.String() != expected {
		t.Errorf("%s rendered incorrectly\nExpected: '%s'\ngot:      '%s'", "MThd", expected, w.String())
	}
}

func TestMThdSMTPEPrint(t *testing.T) {
	expected := "4D 54 68 64 00 00 00 06 00 01 00 11 E7 28   MThd length:6, format:1, tracks:17, SMPTE:25 fps,40 sub-frames"

	mthd := MThd{
		Tag:           "MThd",
		Length:        6,
		Format:        1,
		Tracks:        17,
		Division:      0xe728,
		SMPTETimeCode: true,
		SubFrames:     40,
		FPS:           25,
		Bytes:         []byte{0x4D, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x11, 0xe7, 0x28},
	}

	w := new(bytes.Buffer)

	mthd.Print(w)

	if w.String() != expected {
		t.Errorf("%s rendered incorrectly\nExpected: '%s'\ngot:      '%s'", "MThd", expected, w.String())
	}
}
