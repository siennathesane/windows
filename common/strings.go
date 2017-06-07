// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package common

// this code just enough to do the proper type conversions.

import (
	"github.com/mxplusb/windows"
	"unicode/utf16"
)

const (
	// bit pair values for rune manipulation.
	surr1 = 0xd800
	surr2 = 0xdc00
	surr3 = 0xe000
	surrSelf = 0x10000
	replacementChar = '\uFFFD'
)

// LptStrToString converts the LptStr type to a proper string type.
func LptStrToString(s []windows.LptStr) string {
	a := make([]rune, len(s))
	n := 0
	for i := 0; i < len(s); i++ {
		switch r := s[i]; {
		// if the surrounding bits pairs are valid, we know this bit is a rune.
		case r < surr1, surr3 <= r:
			a[n] = rune(r)
		// basically, we have no idea what it is, so go ahead and decode it.
		case surr1 <= r && r < surr2 && i+1 < len(s) && surr2 <= s[i+1] && s[i+1] < surr3:
			a[n] = utf16.DecodeRune(rune(r), rune(s[i+1]))
			i++
		// if we still can't figure it out, we'll use the default character.
		default:
			a[n] = replacementChar
		}
		n++
	}
	return string(a[:n])
}
