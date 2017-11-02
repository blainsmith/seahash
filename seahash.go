// Package seahash implements SeaHash, a non-cryptographic hash function
// created by http://ticki.github.io.
//
// See https://ticki.github.io/blog/seahash-explained.
package seahash

import (
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
	state state
	// Cumulative # of bytes written.
	inputSize int

	// buf[:bufSize] keeps a subword-sized input that was left over from the
	// previous call to Write.
	bufSize int
	buf     [chunkSize]byte
}

// New creates a new SeaHash hash.Hash64
func New() hash.Hash64 {
	d := &digest{}
	d.Reset()
	return d
}

func (d *digest) Reset() {
	d.state.a = seed1
	d.state.b = seed2
	d.state.c = seed3
	d.state.d = seed4
	d.inputSize = 0
	d.bufSize = 0
}

// Size returns Size constant to satisfy hash.Hash interface
func (d *digest) Size() int { return Size }

// BlockSize returns BlockSize constant to satisfy hash.Hash interface
func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(b []byte) (nn int, err error) {
	nn = len(b)
	d.inputSize += len(b)
	if d.bufSize > 0 {
		n := len(d.buf) - d.bufSize
		copy(d.buf[d.bufSize:], b)
		if n > len(b) {
			d.bufSize += len(b)
			return
		}
		d.state.update(readInt(d.buf[:]))
		d.bufSize = 0
		b = b[n:]
	}
	for len(b) >= chunkSize {
		d.state.update(readInt(b[:chunkSize]))
		b = b[chunkSize:]
	}
	if len(b) > 0 {
		d.bufSize = len(b)
		copy(d.buf[:], b)
	}
	return
}

func (d *digest) Sum(b []byte) []byte {
	d.Write(b)
	r := make([]byte, Size)
	binary.LittleEndian.PutUint64(r, d.Sum64())
	return r
}

func (d *digest) Sum64() uint64 {
	if d.bufSize > 0 {
		s := d.state
		s.update(readInt(d.buf[:d.bufSize]))
		return diffuse(s.a ^ s.b ^ s.c ^ s.d ^ uint64(d.inputSize))
	}
	return diffuse(d.state.a ^ d.state.b ^ d.state.c ^ d.state.d ^ uint64(d.inputSize))
}

// Sum is a convenience method that returns the checksum of the byte slice
func Sum(b []byte) []byte {
	var d digest
	d.Reset()
	return d.Sum(b)
}

// Sum64 is a convenience method that returns uint64 checksum of the byte slice
func Sum64(b []byte) uint64 {
	var d digest
	d.Reset()
	d.Write(b)
	return d.Sum64()
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
