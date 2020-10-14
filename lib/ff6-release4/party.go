// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import "encoding/json"

type (
	Party             [16]CharacterPosition
	CharacterPosition uint8
	jsonCP            struct {
		DisplayOnWorldMap bool
		Enabled           bool
		Front             bool
		Slot              uint8
		Group             uint8
	}
)


const (
	SlotShift      = 3
	EnabledShift   = 6
	FrontBackShift = 5
	DisplayShift   = 7

	GroupMask     = 3
	SlotMask      = 3 << SlotShift
	DisplayMask   = 1 << DisplayShift
	EnabledMask   = 1 << EnabledShift
	FrontBackMask = 1 << FrontBackShift
	Front         = 0 << FrontBackShift
	Back          = 1 << FrontBackShift
)

func (cp *CharacterPosition) SetEnabled(b bool) {
	if b {
		*cp |= EnabledMask
	} else {
		*cp &^= EnabledMask
	}
}

func (cp CharacterPosition) Enabled() bool {
	return (cp & EnabledMask) != 0
}

func (cp *CharacterPosition) SetSlot(s uint8) {
	*cp &^= SlotMask
	*cp |= CharacterPosition((s << SlotShift) & SlotMask)
}

func (cp CharacterPosition) Slot() uint8 {
	return uint8((cp & SlotMask) >> SlotShift)
}

func (cp *CharacterPosition) SetGroup(g uint8) {
	*cp &^= GroupMask
	*cp |= CharacterPosition(g & GroupMask)
}

func (cp CharacterPosition) Group() uint8 {
	return uint8(cp & 3)
}

func (cp *CharacterPosition) SetDisplayOnWorldMap(b bool) {
	if b {
		*cp |= DisplayMask
	} else {
		*cp &^= DisplayMask
	}
}

func (cp CharacterPosition) DisplayOnWorldMap() bool {
	return (cp & DisplayMask) != 0
}

func (cp CharacterPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonCP{
		DisplayOnWorldMap: cp.DisplayOnWorldMap(),
		Enabled:           cp.Enabled(),
		Front:             (cp & FrontBackMask) == Front,
		Slot:              cp.Slot(),
		Group:             cp.Group(),
	})
}

func (cp *CharacterPosition) UnmarshalJSON(data []byte) error {
	var j jsonCP
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	*cp = 0
	cp.SetDisplayOnWorldMap(j.DisplayOnWorldMap)
	cp.SetEnabled(j.Enabled)
	if !j.Front {
		*cp |= Back
	}
	cp.SetSlot(j.Slot)
	cp.SetGroup(j.Group)
	return nil
}