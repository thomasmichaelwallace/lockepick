// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import "encoding/json"

type (
	Lore        uint32
	LoreMastery uint24
)

const (
	Doom Lore = 1 << iota
	Roulette
	Tsunami
	AquaBreath
	Aero
	OneThousandNeedles
	MightyGuard
	RevengeBlast
	WhiteWind
	L5Death
	L4Flare
	L3Confuse
	ReflectLore
	LHoly
	Traveler
	ForceField
	Dischord
	BadBreath
	Transfusion
	Rippler
	StoneLore
	Quasar
	GrandDelta
	SelfDestruct
)

func (l *LoreMastery) UnmarshalJSON(data []byte) error {
	d := make(map[string]bool)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	var i Lore
	for k, v := range lorenames {
		if d[v] {
			i |= k
		}
	}
	(*uint24)(l).Set(uint32(i))
	return nil
}

func (l LoreMastery) MarshalJSON() ([]byte, error) {
	d := make(map[string]bool)
	for k, v := range lorenames {
		d[v] = (uint24(l).Uint() & uint32(k)) != 0
	}
	return json.Marshal(d)
}

var lorenames = map[Lore]string{
	Doom:               "Doom",
	Roulette:           "Roulette",
	Tsunami:            "Tsunami",
	AquaBreath:         "Aqua Breath",
	Aero:               "Aero",
	OneThousandNeedles: "1000 Needles",
	MightyGuard:        "Mighty Guard",
	RevengeBlast:       "Revenge Blast",
	WhiteWind:          "White Wind",
	L5Death:            "Lv.5 Death",
	L4Flare:            "Lv.4 Flare",
	L3Confuse:          "Lv.3 Confuse",
	ReflectLore:        "Reflect???",
	LHoly:              "Lv.? Holy",
	Traveler:           "Traveler",
	ForceField:         "Force Field",
	Dischord:           "Dischord",
	BadBreath:          "Bad Breath",
	Transfusion:        "Transfusion",
	Rippler:            "Rippler",
	StoneLore:          "Stone",
	Quasar:             "Quasar",
	GrandDelta:         "Grand Delta",
	SelfDestruct:       "Self-Destruct",
}

//var a = Roulette + Tsunami + AquaBreath + Aero + OneThousandNeedles + RevengeBlast + L5Death + ForceField + Dischord + Stone + SelfDestruct

//01111101
//128+32+16+8+4+2
