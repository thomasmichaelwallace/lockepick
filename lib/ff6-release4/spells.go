// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
)

type (
	Spell          uint8
	SpellMastery   [64]uint8
	SpellMasteries [13]SpellMastery
)

const (
	Fire Spell = iota
	Blizzard
	Thunder
	PoisonSpell
	Drain
	Fira
	Blizzara
	Thundara
	Bio
	Firaga
	Blizzaga
	Thundaga
	Break
	Death
	Holy
	Flare
	Gravity
	Graviga
	Banish
	Meteor
	Ultima
	Quake
	Tornado
	Meltdown
	Libra
	Slow
	Rasp
	Silence
	Protect
	Sleep
	Confuse
	Haste
	Stop
	Berserk
	Float
	ImpSpell
	Reflect
	Shell
	Vanish
	Hastega
	Slowga
	Osmose
	Teleport
	Quick
	Dispel
	Cure
	Cura
	Curaga
	Raise
	Arise
	Poisona
	Esuna
	Regen
	Reraise
	Flood
	Gravija
	Valor
)

var spellname = map[Spell]string{
	Fire:        "Fire",
	Blizzard:    "Blizzard",
	Thunder:     "Thunder",
	PoisonSpell: "Poison",
	Drain:       "Drain",
	Fira:        "Fira",
	Blizzara:    "Blizzara",
	Thundara:    "Thundara",
	Bio:         "Bio",
	Firaga:      "Firaga",
	Blizzaga:    "Blizzaga",
	Thundaga:    "Thundaga",
	Break:       "Break",
	Death:       "Death",
	Holy:        "Holy",
	Flare:       "Flare",
	Gravity:     "Gravity",
	Graviga:     "Graviga",
	Banish:      "Banish",
	Meteor:      "Meteor",
	Ultima:      "Ultima",
	Quake:       "Quake",
	Tornado:     "Tornado",
	Meltdown:    "Meltdown",
	Libra:       "Libra",
	Slow:        "Slow",
	Rasp:        "Rasp",
	Silence:     "Silence",
	Protect:     "Protect",
	Sleep:       "Sleep",
	Confuse:     "Confuse",
	Haste:       "Haste",
	Stop:        "Stop",
	Berserk:     "Berserk",
	Float:       "Float",
	ImpSpell:    "Imp",
	Reflect:     "Reflect",
	Shell:       "Shell",
	Vanish:      "Vanish",
	Hastega:     "Hastega",
	Slowga:      "Slowga",
	Osmose:      "Osmose",
	Teleport:    "Teleport",
	Quick:       "Quick",
	Dispel:      "Dispel",
	Cure:        "Cure",
	Cura:        "Cura",
	Curaga:      "Curaga",
	Raise:       "Raise",
	Arise:       "Arise",
	Poisona:     "Poisona",
	Esuna:       "Esuna",
	Regen:       "Regen",
	Reraise:     "Reraise",
	Flood:       "Flood",
	Gravija:     "Gravija",
	Valor:       "Valor",
}

func (sm *SpellMasteries) UnmarshalJSON(data []byte) error {
	d := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	for i := range sm {
		if err := json.Unmarshal(d[cname[CharId(i)]], &sm[i]); err != nil {
			return err
		}
	}
	return nil
}

func (sm SpellMasteries) MarshalJSON() ([]byte, error) {
	d := make(map[string]*SpellMastery)
	for i := range sm {
		d[cname[CharId(i)]] = &sm[i]
	}
	return json.Marshal(d)
}

func (sm *SpellMastery) UnmarshalJSON(data []byte) error {
	d := make(map[string]uint8)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	for k, v := range spellname {
		val := d[v]
		if val >= 100 {
			val = 0xff
		}
		sm[k] = val
	}
	return nil
}

func (sm SpellMastery) MarshalJSON() ([]byte, error) {
	d := make(map[string]uint8)
	for k, v := range spellname {
		val := sm[k]
		if val > 100 {
			val = 100
		}
		d[v] = val
	}
	return json.Marshal(d)
}
