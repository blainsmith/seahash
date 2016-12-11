package seahash

import "bytes"

// Hash will hash a byte array into a uint64 with preset seeds
func Hash(b []uint8) uint64 {
	return HashSeeded(b, 0x16f11fe89b0d677c, 0xb480a793d8e6c86c, 0x6fe2e5aaf078ebc9, 0x14f994a4c5259381)
}

// HashSeeded will hash a byte array into a uint64 and allows you to supply your own uint64 seeds
func HashSeeded(b []uint8, k1, k2, k3, k4 uint64) uint64 {
	state := &State{k1, k2, k3, k4}
	buffer := bytes.NewBuffer(b)

	for {
		if buf := buffer.Next(8); len(buf) > 0 {
			state.Write(readInt(buf))
		} else {
			break
		}
	}

	return state.Finish(len(b))
}

// Diffuse helps shift bits around based on the input and isn't meant to be used by itself
func Diffuse(x uint64) uint64 {
	x *= 0x6eed0e9da4d94a4f
	a := x >> 32
	b := x >> 60
	x ^= a >> b
	x *= 0x6eed0e9da4d94a4f

	return x
}

func readInt(b []uint8) uint64 {
	var x uint64 = 0

	for i := len(b) - 1; i >= 0; i-- {
		x <<= 8
		x |= uint64(b[i])
	}

	return x
}
