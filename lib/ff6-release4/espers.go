// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
)

type (
	Esper       uint8
	EspersOwned uint32
)

const (
	Ramuh Esper = iota
	Ifrit
	Shiva
	Siren
	Midgardsormr
	Maduin
	Catoblepas
	Bismark
	CaitSith
	Quetzalli
	Valigarmanda
	Odin
	Raiden
	Bahamut
	Alexandr
	Crusader
	RagnarokEsper
	Kirin
	ZonaSeeker
	Carbunkl
	Phantom
	Seraph
	Golem
	Unicorn
	Fenrir
	Lakshmi
	Phoenix
	Leviathan
	Cactuar
	Diabolos
	Gilgamesh
	NoEsper Esper = 255
)

var ename = map[Esper]string{
	Ramuh:         "Ramuh",
	Ifrit:         "Ifrit",
	Shiva:         "Shiva",
	Siren:         "Siren",
	Midgardsormr:  "Midgardsormr",
	Maduin:        "Maduin",
	Catoblepas:    "Catoblepas",
	Bismark:       "Bismark",
	CaitSith:      "Cait Sith",
	Quetzalli:     "Quetzalli",
	Valigarmanda:  "Valigarmanda",
	Odin:          "Odin",
	Raiden:        "Raiden",
	Bahamut:       "Bahamut",
	Alexandr:      "Alexandr",
	Crusader:      "Crusader",
	RagnarokEsper: "Ragnarok",
	Kirin:         "Kirin",
	ZonaSeeker:    "Zona Seeker",
	Carbunkl:      "Carbunkl",
	Phantom:       "Phantom",
	Seraph:        "Seraph",
	Golem:         "Golem",
	Unicorn:       "Unicorn",
	Fenrir:        "Fenrir",
	Lakshmi:       "Lakshmi",
	Phoenix:       "Phoenix",
	Leviathan:     "Leviathan",
	Gilgamesh:     "Gilgamesh",
	Cactuar:       "Cactuar",
	Diabolos:      "Diabolos",
}

func (e Esper) MarshalJSON() ([]byte, error) {
	return json.Marshal(ename[e])
}

func (e *Esper) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		*e = NoEsper
		return nil
	}
	for k, v := range ename {
		if v == s {
			*e = k
			return nil
		}
	}
	return fmt.Errorf("\"%s\" isn't a known Esper", s)
}

func (e EspersOwned) MarshalJSON() ([]byte, error) {
	d := make(map[string]bool)
	for k, v := range ename {
		d[v] = e&(1<<uint32(k)) != 0
	}
	return json.Marshal(d)
}

func (e *EspersOwned) UnmarshalJSON(data []byte) error {
	d := make(map[string]bool)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	for k, v := range ename {
		if d[v] {
			*e |= (1 << uint32(k))
		}
	}
	return nil
}
