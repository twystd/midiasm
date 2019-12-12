package metaevent

import (
	"fmt"
	"github.com/twystd/midiasm/midi/context"
	"io"
)

type Lyric struct {
	MetaEvent
	Lyric string
}

func NewLyric(event *MetaEvent, r io.ByteReader) (*Lyric, error) {
	if event.Type != 0x05 {
		return nil, fmt.Errorf("Invalid Lyric event type (%02x): expected '05'", event.Type)
	}

	lyric, err := read(r)
	if err != nil {
		return nil, err
	}

	return &Lyric{
		MetaEvent: *event,
		Lyric:     string(lyric),
	}, nil
}

func (e *Lyric) Render(ctx *context.Context, w io.Writer) {
	fmt.Fprintf(w, "%s %-16s %s", e.MetaEvent, "Lyric", e.Lyric)
}