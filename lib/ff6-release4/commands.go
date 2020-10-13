// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
)

type Command uint8

const (
	Fight Command = iota
	ItemCmd
	Magic
	Morph
	Revert
	Steal
	Capture
	Swdtech
	Throw
	Tools
	BlitzCmd
	Runic
	LoreCmd
	Sketch
	Control
	Slot
	Rage
	Leap
	Mimic
	DanceCmd
	Row
	Def
	Jump
	XMagic
	GP_Rain
	Summon
	Health
	Shock
	Possess
	MagiTek
	NoCommand Command = 255
)

var cmdName = map[Command]string{
	Fight:     "Fight",
	ItemCmd:   "Item",
	Magic:     "Magic",
	Morph:     "Morph",
	Revert:    "Revert",
	Steal:     "Steal",
	Capture:   "Capture",
	Swdtech:   "Swdtech",
	Throw:     "Throw",
	Tools:     "Tools",
	BlitzCmd:  "Blitz",
	Runic:     "Runic",
	LoreCmd:   "Lore",
	Sketch:    "Sketch",
	Control:   "Control",
	Slot:      "Slot",
	Rage:      "Rage",
	Leap:      "Leap",
	Mimic:     "Mimic",
	DanceCmd:  "Dance",
	Row:       "Row",
	Def:       "Def",
	Jump:      "Jump",
	XMagic:    "XMagic",
	GP_Rain:   "GP_Rain",
	Summon:    "Summon",
	Health:    "Health",
	Shock:     "Shock",
	Possess:   "Possess",
	MagiTek:   "MagiTek",
	NoCommand: "",
}

func (c Command) MarshalJSON() ([]byte, error) {
	return json.Marshal(cmdName[c])
}

func (c *Command) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range cmdName {
		if v == s {
			*c = k
			return nil
		}
	}
	return fmt.Errorf("\"%s\" isn't a known command", s)
}
