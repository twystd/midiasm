package main

import (
	"flag"
	"fmt"
	"github.com/twystd/midiasm/midi"
	"github.com/twystd/midiasm/midi/encoding/midifile"
	"github.com/twystd/midiasm/midi/types"
	"io/ioutil"
	"os"
)

type command interface {
	flagset() *flag.FlagSet
	config() string
	Execute(*midi.SMF)
}

var cli = map[string]command{
	"print": &Print{},
	"notes": &Notes{},
}

func main() {
	cmd, filename, err := parse()
	if err != nil {
		fmt.Printf("Error: unable to parse command line (%v)\n", err)
		return
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if conf := cmd.config(); conf != "" {
		f, err := os.Open(conf)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		manufacturers, err := types.LoadManufacturers(f)

		f.Close()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		types.AddManufacturers(manufacturers)
	}

	decoder := midifile.NewDecoder()

	smf, err := decoder.Decode(bytes)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if smf == nil {
		fmt.Printf("Error: Mysteriously failed to decode MIDI file\n")
		return
	}

	if errors := smf.Validate(); len(errors) > 0 {
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "WARNING: there are validation errors:\n")
		for _, e := range errors {
			fmt.Fprintf(os.Stderr, "         ** %v\n", e)
		}
		fmt.Fprintln(os.Stderr)
	}

	cmd.Execute(smf)
}

func parse() (command, string, error) {
	cmd := &Print{}
	if len(os.Args) > 1 {
		c, ok := cli[os.Args[1]]
		if ok {
			flagset := c.flagset()
			if err := flagset.Parse(os.Args[2:]); err != nil {
				return cmd, "", err
			}

			return c, flagset.Arg(0), nil
		}
	}

	flagset := cmd.flagset()
	if err := flagset.Parse(os.Args[1:]); err != nil {
		return cmd, "", err
	}

	return cmd, flagset.Arg(0), nil
}
