// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/robmerrell/comandante"
	"io/ioutil"
	"log"
	"os"
)

var (
	savfile  = "filesRockDat0.sav"
	jsonfile = "filesRockDat0.json"
)

func tojson() error {
	data, err := ioutil.ReadFile(savfile)
	if err != nil {
		return err
	}
	
	d, err := tojsonbytes(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(jsonfile, d, 0644)
}

func fromjson() error {
	data, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		return err
	}

	s, err := fromjsonbytes(data)
	if err != nil {
		return err
	}
	_, err = os.Stat(savfile)
	if err == nil {
		_, err = os.Stat(savfile + ".bak")
		if err != nil {
			os.Rename(savfile, savfile+".bak")
		}
	}
	return ioutil.WriteFile(savfile, s, 0644)
}

func main() {
	bin := comandante.New("ff6", "Manage Final Fantasy VI mobile save files")
	bin.IncludeHelp()
	tj := comandante.NewCommand("tojson", "Convert .sav save file to .json", tojson)
	tj.FlagInit = func(fs *flag.FlagSet) {
		fs.StringVar(&savfile, "i", savfile, "Input .sav file")
		fs.StringVar(&jsonfile, "o", jsonfile, "Output .json file")
	}
	bin.RegisterCommand(tj)
	fj := comandante.NewCommand("fromjson", "Convert .json save file to .sav", fromjson)
	fj.FlagInit = func(fs *flag.FlagSet) {
		fs.StringVar(&jsonfile, "i", jsonfile, "Input .json file")
		fs.StringVar(&savfile, "o", savfile, "Output .sav file")
	}
	bin.RegisterCommand(fj)

	flag.Parse()

	if err := bin.Run(); err != nil {
		log.Fatalln(err)
	}
}
