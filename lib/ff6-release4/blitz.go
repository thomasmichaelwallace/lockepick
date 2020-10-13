// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import "encoding/json"

type (
	Blitz        uint8
	BlitzMastery uint8
)

const (
	RagingFist Blitz = 1 << iota
	AuraCannon
	MeteorStrike
	RisingPhoenix
	Chakra
	RazorGale
	SoulSpiral
	PhantomRush
)

func (b *BlitzMastery) UnmarshalJSON(data []byte) error {
	d := make(map[string]bool)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	var i Blitz
	for k, v := range blitznames {
		if d[v] {
			i |= k
		}
	}
	*b = BlitzMastery(i)
	return nil
}

func (b BlitzMastery) MarshalJSON() ([]byte, error) {
	d := make(map[string]bool)
	for k, v := range blitznames {
		d[v] = Blitz(b)&k != 0
	}
	return json.Marshal(d)
}

var blitznames = map[Blitz]string{
	RagingFist:    "Raging Fist",
	AuraCannon:    "Aura Cannon",
	MeteorStrike:  "Meteor Strike",
	RisingPhoenix: "Rising Phoenix",
	Chakra:        "Chakra",
	RazorGale:     "Razor Gale",
	SoulSpiral:    "Soul Spiral",
	PhantomRush:   "Phantom Rush",
}
