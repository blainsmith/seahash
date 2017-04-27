// Package seahash implements SeaHash, a non-cryptographic hash function
// created by http://ticki.github.io.
//
// See https://ticki.github.io/blog/seahash-explained.
package seahash

import (
	"bytes"
	"encoding/binary"
	"hash"
)

// Size of a SeaHash checksum in bytes.
const Size = 8

// BlockSize of SeaHash in bytes.
const BlockSize = 8

const (
	chunkSize = 8
	seed1     = 0x16f11fe89b0d677c
	seed2     = 0xb480a793d8e6c86c
	seed3     = 0x6fe2e5aaf078ebc9
	seed4     = 0x14f994a4c5259381
)

type digest struct {
	state  state
	buffer *bytes.Buffer
}

// New creates a new SeaHash hash.Hash
func New() hash.Hash {
	d := &digest{}
	d.Reset()
	return d
}

func (d *digest) Reset() {
	d.state.a = seed1
	d.state.b = seed2
	d.state.c = seed3
	d.state.d = seed4
	d.buffer = bytes.NewBuffer(nil)
}

// Size returns Size constant to satisfy hash.Hash interface
func (d *digest) Size() int { return Size }

// BlockSize returns BlockSize constant to satisfy hash.Hash interface
func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(b []byte) (nn int, err error) {
	d.buffer.Write(b)
	return
}

func (d *digest) Sum(b []byte) []byte {
	d.Write(b)
	return d.checkSum()
}

func (d *digest) checkSum() []byte {
	bl := uint64(d.buffer.Len())
	for {
		if buf := d.buffer.Next(chunkSize); len(buf) > 0 {
			d.state.update(readInt(buf))
		} else {
			break
		}
	}

	r := make([]byte, Size)
	binary.LittleEndian.PutUint64(r, diffuse(d.state.a^d.state.b^d.state.c^d.state.d^bl))
	return r
}

// Sum is a convenience method that returns the checksum of the byte slice
func Sum(b []byte) []byte {
	var d digest
	d.Reset()
	d.Write(b)
	return d.checkSum()
}

type state struct {
	a uint64
	b uint64
	c uint64
	d uint64
}

func (s *state) update(x uint64) {
	a := s.a
	a = diffuse(a ^ x)

	s.a = s.b
	s.b = s.c
	s.c = s.d
	s.d = a
}

func diffuse(x uint64) uint64 {
	x *= 0x6eed0e9da4d94a4f
	a := x >> 32
	b := x >> 60
	x ^= a >> b
	x *= 0x6eed0e9da4d94a4f

	return x
}

func readInt(b []uint8) uint64 {
	var x uint64

	for i := len(b) - 1; i >= 0; i-- {
		x <<= 8
		x |= uint64(b[i])
	}

	return x
}
