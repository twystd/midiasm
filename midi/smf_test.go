package midi

import (
	"reflect"
	"testing"
)

func TestUnmarshalSMF(t *testing.T) {
	bytes := []byte{
		0x4d, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x60,
	}

	mthd := MThd{
		Tag:      "MThd",
		Length:   6,
		Format:   1,
		Tracks:   0,
		Division: 96,
		PPQN:     96,
		Bytes:    []byte{0x4d, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x60},
	}

	smf := SMF{}
	if err := smf.UnmarshalBinary(bytes); err != nil {
		t.Fatalf("Unexpected error unmarshaling MThd: %v", err)
	}

	if !reflect.DeepEqual(*smf.MThd, mthd) {
		t.Errorf("MThd incorrectly unmarshaled\n   expected:%v\n   got:     %v", mthd, *smf.MThd)
	}
}
