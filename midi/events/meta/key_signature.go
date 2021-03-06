package metaevent

import (
	"fmt"
	"github.com/twystd/midiasm/midi/context"
	"github.com/twystd/midiasm/midi/types"
)

type KeyType uint8

func (k KeyType) String() string {
	if k == 1 {
		return "minor"
	}

	return "major"
}

type KeySignature struct {
	Tag         string
	Status      types.Status
	Type        types.MetaEventType
	Accidentals int8
	KeyType     KeyType
	Key         string
}

var major_keys = map[int8]string{
	0:  "C major",
	1:  "G major",
	2:  "D major",
	3:  "A major",
	4:  "E major",
	5:  "B major",
	6:  "F\u266f major",
	-1: "F major",
	-2: "B\u266d major",
	-3: "E\u266d major",
	-4: "A\u266d major",
	-5: "D\u266d major",
	-6: "G\u266d major",
}

var minor_keys = map[int8]string{
	0:  "A minor",
	1:  "E minor",
	2:  "B minor",
	3:  "F\u266f minor",
	4:  "C\u266f minor",
	5:  "G\u266f minor",
	6:  "D\u266f minor",
	-1: "D minor",
	-2: "G minor",
	-3: "C minor",
	-4: "F minor",
	-5: "B\u266d minor",
	-6: "E\u266d minor",
}

func NewKeySignature(ctx *context.Context, bytes []byte) (*KeySignature, error) {
	if len(bytes) != 2 {
		return nil, fmt.Errorf("Invalid KeySignature length (%d): expected '2'", len(bytes))
	}

	accidentals := int8(bytes[0])
	keyType := bytes[1]
	if keyType != 0 && keyType != 1 {
		return nil, fmt.Errorf("Invalid KeySignature key type (%d): expected a value in the interval [0,1]", keyType)
	}

	key := ""
	switch keyType {
	case 0:
		if signature, ok := major_keys[accidentals]; !ok {
			return nil, fmt.Errorf("Invalid major key signature (%d accidentals): expected a value in the interval [-6,0]", accidentals)
		} else {
			key = signature
		}
	case 1:
		if signature, ok := minor_keys[accidentals]; !ok {
			return nil, fmt.Errorf("Invalid minor key signature (%d accidentals): expected a value in the interval [-6,0]", accidentals)
		} else {
			key = signature
		}
	}

	if accidentals < 0 {
		ctx.UseFlats()
	} else {
		ctx.UseSharps()
	}

	return &KeySignature{
		Tag:         "KeySignature",
		Status:      0xff,
		Type:        0x59,
		Accidentals: accidentals,
		KeyType:     KeyType(keyType),
		Key:         key,
	}, nil
}
