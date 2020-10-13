// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type (
	CharId   uint8
	Portrait CharId
)

// 0145
// 76ad
// 839c
const (
	Terra CharId = iota
	Locke
	Cyan
	Shadow
	Edgar
	Sabin
	Celes
	Strago
	Relm
	Setzer
	Mog
	Gau
	Gogo
	Umaro
	Vicks_Wedge
	ImpId
	Leo
	Banon
	EsperTerra    // (Terra)
	MerchantLocke // (Soldier)
	Ghost
	Kefka          //(Terra)
	EmperorGestahl // (Terra)
	SoldierLocke   // [Village Elder] (Terra)
	Empty1         // [Young Man] (Terra) *cannot fight during battle*
	Empty2         // [Interceptor] (Terra) *cannot fight during battle*
	Empty3         // [Opera Celes] (Celes) *cannot fight during battle*
	Empty4         //  [Scholar] (Glitch) *will hang the game when fighting*
)

var cname = map[CharId]string{
	Terra:          "Terra",
	Locke:          "Locke",
	Cyan:           "Cyan",
	Shadow:         "Shadow",
	Edgar:          "Edgar",
	Sabin:          "Sabin",
	Celes:          "Celes",
	Strago:         "Strago",
	Relm:           "Relm",
	Setzer:         "Setzer",
	Mog:            "Mog",
	Gau:            "Gau",
	Gogo:           "Gogo",
	Umaro:          "Umaro",
	Vicks_Wedge:    "Vicks_Wedge",
	ImpId:          "Imp",
	Leo:            "Leo",
	Banon:          "Banon",
	EsperTerra:     "EsperTerra",
	MerchantLocke:  "MerchantLocke",
	Ghost:          "Ghost",
	Kefka:          "Kefka",
	EmperorGestahl: "EmperorGestahl",
	SoldierLocke:   "SoldierLocke",
	Empty1:         "Empty1",
	Empty2:         "Empty2",
	Empty3:         "Empty3",
	Empty4:         "Empty4",
}

func (c *CharId) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range cname {
		if v == s {
			*c = k
			return nil
		}
	}

	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return err
	}
	*c = CharId(i)
	return nil
}

func (c CharId) MarshalJSON() ([]byte, error) {
	v, ok := cname[c]
	if !ok {
		v = fmt.Sprintf("%d", c)
	}
	return json.Marshal(v)
}

func (p *Portrait) UnmarshalJSON(data []byte) error {
	return (*CharId)(p).UnmarshalJSON(data)
}
func (p Portrait) MarshalJSON() ([]byte, error) {
	return CharId(p).MarshalJSON()
}
