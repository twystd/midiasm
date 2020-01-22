package operations

import (
	"fmt"
	"github.com/twystd/midiasm/midi"
	"github.com/twystd/midiasm/midi/types"
	"io"
	"strings"
	"text/template"
)

const document string = `
>>>>>>>>>>>>>>>>>>>>>>>>>
{{pad (ellipsize .MThd.Bytes 0 14) 42}}  {{template "MThd" .MThd}}

{{range .Tracks}}{{pad (ellipsize      .Bytes 0 8)  42}}  {{template "MTrk" .}}
{{range .Events}}{{pad (ellipsize      .Bytes 0 14) 42}}  {{template "event" .}}{{end}}
{{end}}
>>>>>>>>>>>>>>>>>>>>>>>>>

`

var templates = map[string]string{
	"MThd": `{{.Tag}} length:{{.Length}}, format:{{.Format}}, tracks:{{.Tracks}}, {{if not .SMPTETimeCode }}metrical time:{{.PPQN}} ppqn{{else}}SMPTE:{{.FPS}} fps,{{.SubFrames}} sub-frames{{end}}`,
	"MTrk": `{{.Tag}} {{.TrackNumber}} length:{{.Length}}`,

	"event": `tick:{{pad .Tick.String 9}}  delta:{{pad .Delta.String 9}}  {{template "events" .}}`,
	"events": `{{if eq .Tag "SequenceNumber"}}{{template "sequenceno" .}}
{{else if eq .Tag "Text"                   }}{{template "text"                   .}}
{{else if eq .Tag "Copyright"              }}{{template "copyright"              .}}
{{else if eq .Tag "TrackName"              }}{{template "trackname"              .}}
{{else if eq .Tag "InstrumentName"         }}{{template "instrumentname"         .}}
{{else if eq .Tag "Lyric"                  }}{{template "lyric"                  .}}
{{else if eq .Tag "Marker"                 }}{{template "marker"                 .}}
{{else if eq .Tag "CuePoint"               }}{{template "cuepoint"               .}}
{{else if eq .Tag "ProgramName"            }}{{template "programname"            .}}
{{else if eq .Tag "DeviceName"             }}{{template "devicename"             .}}
{{else if eq .Tag "MIDIChannelPrefix"      }}{{template "midichannelprefix"      .}}
{{else if eq .Tag "MIDIPort"               }}{{template "midiport"               .}}
{{else if eq .Tag "EndOfTrack"             }}{{template "endoftrack"             .}}
{{else if eq .Tag "Tempo"                  }}{{template "tempo"                  .}}
{{else if eq .Tag "SMPTEOffset"            }}{{template "smpteoffset"            .}}
{{else if eq .Tag "TimeSignature"          }}{{template "timesignature"          .}}
{{else if eq .Tag "KeySignature"           }}{{template "keysignature"           .}}
{{else if eq .Tag "SequencerSpecificEvent" }}{{template "sequencerspecificevent" .}}
{{else if eq .Tag "NoteOff"                }}{{template "noteoff"                .}}
{{else if eq .Tag "NoteOn"                 }}{{template "noteon"                 .}}
{{else if eq .Tag "PolyphonicPressure"     }}{{template "polyphonicpressure"     .}}
{{else if eq .Tag "Controller"             }}{{template "controller"             .}}
{{else if eq .Tag "ProgramChange"          }}{{template "programchange"          .}}
{{else if eq .Tag "ChannelPressure"        }}{{template "channelpressure"        .}}
{{else if eq .Tag "PitchBend"              }}{{template "pitchbend"              .}}
{{else if eq .Tag "SysExMessage"           }}{{template "sysexmessage"           .}}
{{else                                     }}{{template "unknown"                .}}
{{end}}`,

	"sequenceno":             `{{.Type}} {{pad .Tag 22}} {{.SequenceNumber}}`,
	"text":                   `{{.Type}} {{pad .Tag 22}} {{.Text}}`,
	"copyright":              `{{.Type}} {{pad .Tag 22}} {{.Copyright}}`,
	"trackname":              `{{.Type}} {{pad .Tag 22}} {{.Name}}`,
	"instrumentname":         `{{.Type}} {{pad .Tag 22}} {{.Name}}`,
	"lyric":                  `{{.Type}} {{pad .Tag 22}} {{.Lyric}}`,
	"marker":                 `{{.Type}} {{pad .Tag 22}} {{.Marker}}`,
	"cuepoint":               `{{.Type}} {{pad .Tag 22}} {{.CuePoint}}`,
	"programname":            `{{.Type}} {{pad .Tag 22}} {{.Name}}`,
	"devicename":             `{{.Type}} {{pad .Tag 22}} {{.Name}}`,
	"midichannelprefix":      `{{.Type}} {{pad .Tag 22}} {{.Channel}}`,
	"midiport":               `{{.Type}} {{pad .Tag 22}} {{.Port}}`,
	"endoftrack":             `{{.Type}} {{    .Tag   }}`,
	"tempo":                  `{{.Type}} {{pad .Tag 22}} tempo:{{.Tempo}}`,
	"smpteoffset":            `{{.Type}} {{pad .Tag 22}} {{.Hour}} {{.Minute}} {{.Second}} {{.FrameRate}} {{.Frames}} {{.FractionalFrames}}`,
	"timesignature":          `{{.Type}} {{pad .Tag 22}} {{.Numerator}}/{{.Denominator}}, {{.TicksPerClick }} ticks per click, {{.ThirtySecondsPerQuarter}}/32 per quarter`,
	"keysignature":           `{{.Type}} {{pad .Tag 22}} {{.Key}}`,
	"sequencerspecificevent": `{{.Type}} {{pad .Tag 22}} {{.Data}}`,

	"noteoff":            `{{.Status}} {{pad .Tag 22}} channel:{{pad .Channel 2}} note:{{.Note.Name}}, velocity:{{.Velocity}}`,
	"noteon":             `{{.Status}} {{pad .Tag 22}} channel:{{pad .Channel 2}} note:{{.Note.Name}}, velocity:{{.Velocity}}`,
	"polyphonicpressure": `{{.Status}} {{pad .Tag 22}} channel:{{pad .Channel 2}} pressure:{{.Pressure}}`,
	"controller":         `{{.Status}} {{pad .Tag 22}} channel:{{pad .Channel 2}} controller:{{.Controller}}, value:{{.Value}}`,
	"programchange":      `{{.Status}} {{pad .Tag 22}} channel:{{pad .Channel 2}} program:{{.Program }}`,
	"channelpressure":    `{{.Status}} {{pad .Tag 22}} channel:{{pad .Channel 2}} pressure:{{.Pressure}}`,
	"pitchbend":          `{{.Status}} {{pad .Tag 22}} channel:{{pad .Channel 2}} bend:{{.Bend}}`,

	"sysexmessage": `{{.Status}} {{pad .Tag 22}} {{.Data}}`,

	"unknown": `?? {{.Tag}}`,
}

type Print struct {
	Writer func(midi.Chunk) (io.Writer, error)
}

func (p *Print) Execute(smf *midi.SMF) error {
	if w, err := p.Writer(smf.MThd); err != nil {
		return err
	} else {
		smf.MThd.Print(w)
	}

	for _, track := range smf.Tracks {
		if w, err := p.Writer(track); err != nil {
			return err
		} else {
			track.Print(w)
		}
	}

	return nil
}

func (p *Print) PrintWithTemplate(smf *midi.SMF, w io.Writer) error {
	functions := template.FuncMap{
		"ellipsize": ellipsize,
		"pad":       pad,
	}

	tmpl, err := template.New("SMF").Funcs(functions).Parse(document)
	if err != nil {
		return err
	}

	for name, t := range templates {
		if _, err = tmpl.New(name).Parse(t); err != nil {
			return err
		}
	}

	return tmpl.Execute(w, smf)
}

func ellipsize(bytes types.Hex, offsets ...int) string {
	start := 0
	end := len(bytes)

	if len(offsets) > 0 && offsets[0] > 0 {
		start = offsets[0]
	}

	if len(offsets) > 1 && offsets[1] > start && offsets[1] < end {
		end = offsets[1]
	}

	hex := bytes[start:end].String()
	if end-start < len(bytes) {
		hex += `…`
	}

	return hex
}

func pad(v interface{}, width int) string {
	s := fmt.Sprintf("%v", v)
	if width < len([]rune(s)) {
		return s
	}

	return s + strings.Repeat(" ", width-len([]rune(s)))
}
