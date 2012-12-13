/*
   conflux - Distributed database synchronization library
	Based on the algorithm described in
		"Set Reconciliation with Nearly Optimal	Communication Complexity",
			Yaron Minsky, Ari Trachtenberg, and Richard Zippel, 2004.

   Copyright (C) 2012  Casey Marshall <casey.marshall@gmail.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package conflux

import (
	"fmt"
	"math/big"
)

var p_128 = big.NewInt(0).SetBytes([]byte{
        0x91,0xae,0x33,0x7c,0x8a,0x19,0xf5,0xb7,
        0xc3,0xca,0xab,0x7f,0xf9,0xf0,0x6d,0x95})

var p_160 = big.NewInt(0).SetBytes([]byte{
        0x81,0xce,0x83,0x36,0x6d,0xb4,0xff,0xb5,
        0xb,0xc9,0x3,0xf7,0x5c,0xd8,0xfd,0x35,
        0xd,0xab,0xa4,0xc3})

var p_256 = big.NewInt(0).SetBytes([]byte{
        0x8c,0x1d,0x8b,0xc4,0xf6,0x99,0xd5,0xc1,
        0x70,0xc0,0xe1,0xfb,0x53,0xc,0xb2,0x6c,
        0x77,0x72,0x81,0x17,0xa,0x99,0x2b,0x72,
        0x4d,0x59,0xc6,0x59,0xa,0x9a,0x54,0xcb})

var p_512 = big.NewInt(0).SetBytes([]byte{
        0xae,0x3f,0x54,0x4f,0xa6,0x5c,0xb4,0xa7,
        0x8a,0x45,0x6d,0x24,0xe7,0x45,0xc9,0xb2,
        0x15,0x94,0x3d,0xd3,0xaf,0x31,0xa0,0xa8,
        0xe4,0xdb,0xba,0x59,0x71,0x44,0xdc,0x9f,
        0x35,0x67,0x88,0xa6,0x35,0xa2,0x4e,0xfa,
        0xcd,0x55,0x54,0x6,0xfa,0x20,0xbe,0xe0,
        0x5e,0x35,0x19,0xb8,0x49,0x21,0x1,0x78,
        0x9a,0x25,0x50,0x5,0xa2,0x56,0x82,0x1d})

// Zp represents a value in the finite field Z(p),
// an integer in which all arithmetic is (mod p).
type Zp struct {
	// The prime bound of the finite field Z(p).
	P *big.Int
	// Some specific Z.
	Z *big.Int
}

func NewZp(p int64, n int64) *Zp {
	zp := &Zp{ P: big.NewInt(p), Z: big.NewInt(n) }
	zp.normalize()
	return zp
}

func (zp *Zp) normalize() {
	zp.Z.Mod(zp.Z, zp.P)
}

func (zp *Zp) Add(values... *Zp) *Zp {
	for _, v := range values {
		assertZp(zp, v)
		zp.Z.Add(zp.Z, v.Z)
		zp.normalize()
	}
	return zp
}

func assertZp(x, y *Zp) {
	if x.P.Cmp(y.P) != 0 {
		panic(fmt.Sprintf("finite field mismatch betwee Z(%v) and Z(%v)", x.P, y.P))
	}
}
