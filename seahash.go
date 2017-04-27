package seahash

import "bytes"

// Hash will hash a byte array into a uint64 with preset seeds
func Hash(b []uint8) uint64 {
	return HashSeeded(b, 0x16f11fe89b0d677c, 0xb480a793d8e6c86c, 0x6fe2e5aaf078ebc9, 0x14f994a4c5259381)
}

func HashSeeded(b []uint8, k1, k2, k3, k4 uint64) uint64 {
	state := &state{k1, k2, k3, k4}
	buffer := bytes.NewBuffer(b)

	for {
		if buf := buffer.Next(8); len(buf) > 0 {
			state.write(readInt(buf))
		} else {
			break
		}
	}

	return state.finish(len(b))
}

type state struct {
	A uint64
	B uint64
	C uint64
	D uint64
}

func (s *state) write(x uint64) {
	a := s.A
	a = diffuse(a ^ x)

	s.A = s.B
	s.B = s.C
	s.C = s.D
	s.D = a
}

func (s *state) finish(size int) uint64 {
	return diffuse(s.A ^ s.B ^ s.C ^ s.D ^ uint64(size))
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
