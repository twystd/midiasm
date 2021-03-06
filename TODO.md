## v0.0

*Disassembler*

- [x] Rework MIDI event parser
- [x] Rework META event parser
- [x] Extract notes
- [x] Use microseconds as integer time base
- [x] Ellipsize too long hex
- [x] Log errors/warning to stderr
- [x] Write to file
- [x] Split tracks to separate files
- [x] Validate (missing end of track, tempo events)
- [x] --debug
- [x] --verbose
- [x] Print note name + octave
- [x] Decode note in context of current scale
- [x] Keep NoteOff name to be NoteOn name if note is tied across KeySignature
- [x] Unglobalize manufacturer list (e.g. see UnmarshalSMF unit tests which keep the last value loaded from 'conf')
- [x] Unit test for invalid in SMFUnmarshal running status 
- [x] Configurable formats
- [x] Identify manufacturer for SequencerSpecificEvent (http://www.somascape.org/midi/tech/spec.html#sysexnotes)
- [x] Identify manufacturer for SysEx (http://www.somascape.org/midi/tech/spec.html#sysexnotes)
- [x] Check SMTPEOffset only in track 0 for format 1
- [x] Running status
- [x] Identify controller numbers (http://www.somascape.org/midi/tech/spec.html#ctrlnums)
- [x] MTrk unit test for running status
- [x] Identify bank for Program Change
- [x] Format 0
- [x] Format 1
- [ ] Format 2
- [ ] Add outstanding events to TestDecode

### Notes 

- [x] Print note name + octave
- [x] Rework as SMF processor
- [ ] Check loss of precision
- [ ] Unit tests for tempo map to time conversion
- [ ] Configurable formats
- [ ] Pretty print
- [ ] Format 0
- [ ] Format 2
- [ ] NoteOn with 0 velocity -> NoteOff

### MIDI events

- [x] 8n/Note Off
- [x] 9n/Note On
- [x] An/Polyphonic Pressure
- [x] Bn/Controller
- [x] Cn/Program Change
- [x] Dn/Channel Pressure
- [x] En/Pitch Bend

### META events

- [x] 00/Sequence Number
- [x] 01/Text
- [x] 02/Copyright
- [x] 03/Track Name
- [x] 04/Instrument Name
- [x] 05/Lyric
- [x] 06/Marker
- [x] 07/Cue Point
- [x] 08/Program Name
- [x] 09/Device Name
- [x] 20/MIDI Channel Prefix
- [x] 21/MIDI Port
- [x] 2F/End of Track
- [x] 51/Tempo
- [x] 54/SMPTE Offset
- [x] 58/Time Signature
- [x] 59/Key Signature
- [x] 7F/Sequencer Specific Event

- [x] TimeSignature: [Unicode fractions](http://unicodefractions.com)
- [x] KeySignature:  [Unicode symbols](https://unicode-table.com/en/blocks/musical-symbols/)

### SysEx events

- [x] Single messages
- [x] Continuation events
- [x] Escape sequences
- [x] Casio sequences

## TODO

### Disassembler

1. Rework decoder using tags/reflection/grammar+packrat-parser/kaitai/binpac/somesuch
2. Reference files
   - Format 0
   - Format 1
   - Format 2

### Other

1.  Assembler
2.  TSV
3.  Export to JSON
4.  Export to S-expressions
5.  VSCode plugin
    -  [Language Server Protocol Tutorial: From VSCode to Vim](https://www.toptal.com/javascript/language-server-protocol-tutorial)
6.  Convert between formats 0, 1 and 2
7.  [Manufacturer ID's](https://www.midi.org/specifications-old/category/reference-tables) (?)
8.  Check against reference files from [github:nfroidure/midifile](https://github.com/nfroidure/midifile)
9.  [How to use a field of struct or variable value as template name?](https://stackoverflow.com/questions/28830543/how-to-use-a-field-of-struct-or-variable-value-as-template-name)
10. Online/Javascript version
12. https://github.com/go-interpreter/chezgo
13. SDK (?)
14. mmap
15. REST/GraphQL interface

