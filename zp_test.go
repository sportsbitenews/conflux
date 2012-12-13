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
	"github.com/bmizerany/assert"
	"testing"
)

func zp5(n int) *Zp {
	return NewZp(int64(5), int64(n))
}

func zp7(n int) *Zp {
	return NewZp(int64(7), int64(n))
}

func TestAdd(t *testing.T) {
	a := zp5(1)
	b := zp5(3)
	assert.Equal(t, 0, zp5(4).Z.Cmp(a.Add(b).Z))
}

func TestAddWrap(t *testing.T) {
	a := zp5(1)
	b := zp5(9)
	assert.Equal(t, 0, zp5(0).Z.Cmp(a.Add(b).Z))
}

func TestMinusOne(t *testing.T) {
	a := NewZp(int64(65537), int64(-1))
	assert.Equal(t, int64(65536), a.Z.Int64())
}

func TestMismatchedP(t *testing.T) {
	defer func(){
        r := recover()
		assert.T(t, r != nil)
	}()
	a := zp5(1)
	b := NewZp(int64(65537), int64(9))
	a.Add(b)
	t.Fail()
}
