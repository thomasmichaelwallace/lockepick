// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

type Status uint8

const (
	Darkness Status = 1 << iota
	Zombie
	Poison
	Magitek
	Invisible
	Imp
	Stone
	Wounded
)
