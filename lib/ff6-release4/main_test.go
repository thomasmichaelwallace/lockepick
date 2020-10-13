// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func compare(data, data2 []byte, t *testing.T) {
	if a, b := len(data), len(data2); a != b {
		t.Errorf("lengths differ: %d != %d", a, b)
		return
	}
	if !bytes.Equal(data, data2) {
		t.Error("Not equal..")
		const step = 16
		for i := 0; (i + step) < len(data); i += step {
			a := data[i : i+step]
			b := data2[i : i+step]
			if !bytes.Equal(a, b) {
				both := make([]byte, 2*step)
				for j := 0; j < len(a)/16; j++ {
					aa := a[j*16 : (j+1)*16]
					bb := b[j*16 : (j+1)*16]
					copy(both[j*(2*16):], aa)
					copy(both[j*(2*16)+16:], bb)
				}
				t.Logf("%d\n%s", i, hex.Dump(both))
			}
		}
	}
}
func TestEncode(t *testing.T) {
	s, _ := filepath.Glob("files*.sav.*")
	for _, p := range s {
		t.Log(p)
		data, err := ioutil.ReadFile(p)
		if err != nil {
			t.Error(err)
			continue
		}

		var save SaveFile
		f, err := os.Open(p)
		if err != nil {
			t.Error(err)
			continue
		}
		defer f.Close()

		if err := binary.Read(f, binary.LittleEndian, &save); err != nil {
			t.Error(err)
			continue
		}
		data2 := save.Encode()
		compare(data, data2, t)
		data3 := save.Encode()
		compare(data2, data3, t)
		json.Marshal(save)
		data4 := save.Encode()
		compare(data4, data3, t)

	}
}

func TestJson(t *testing.T) {
	s, _ := filepath.Glob("files*.sav.*")
	for _, p := range s {
		t.Log(p)
		var save SaveFile
		f, err := os.Open(p)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		if err := binary.Read(f, binary.LittleEndian, &save); err != nil {
			t.Error(err)
			continue
		}
		data2, err := json.MarshalIndent(save, "", "\t")
		if err != nil {
			t.Error(err)
			continue
		}
		var save2 SaveFile
		if err := json.Unmarshal(data2, &save2); err != nil {
			t.Error(err)
			continue
		}
		data3, err := json.MarshalIndent(save2, "", "\t")
		if err != nil {
			t.Error(err)
			continue
		}
		a := string(data2)
		b := string(data3)
		if a != b {
			t.Error("not equal")
			break
		}
	}
}
