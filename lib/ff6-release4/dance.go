// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import "encoding/json"

type (
	Dance        uint8
	DanceMastery uint8
)

const (
	WindRhapsody = 1 << iota
	ForestNocturne
	DesertLullaby
	LoveSerenade
	EarthBlues
	WaterHarmony
	TwilightRequium
	SnowmanRondo
)

func (dm *DanceMastery) UnmarshalJSON(data []byte) error {
	d := make(map[string]bool)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	var i Dance
	for k, v := range dancenames {
		if d[v] {
			i |= k
		}
	}
	*dm = DanceMastery(i)
	return nil
}

func (l DanceMastery) MarshalJSON() ([]byte, error) {
	d := make(map[string]bool)
	for k, v := range dancenames {
		d[v] = Dance(l)&k != 0
	}
	return json.Marshal(d)
}

var dancenames = map[Dance]string{
	WindRhapsody:    "Wind Rhapsody",
	ForestNocturne:  "Forest Nocturne",
	DesertLullaby:   "Desert Lullaby",
	LoveSerenade:    "Love Serenade",
	EarthBlues:      "Earth Blues",
	WaterHarmony:    "Water Harmony",
	TwilightRequium: "Twilight Requium",
	SnowmanRondo:    "Snowman Rondo",
}
