package seahash

// State is used to store the current state of the hashing as bytes are written to it
type State struct {
	A uint64
	B uint64
	C uint64
	D uint64
}

// Write updates the State with a diffused value and shifts the 4 states around
func (s *State) Write(x uint64) {
	a := s.A
	a = Diffuse(a ^ x)

	s.A = s.B
	s.B = s.C
	s.C = s.D
	s.D = a
}

// Finish performs a final diffuse before returning the final hash
func (s *State) Finish(size int) uint64 {
	return Diffuse(s.A ^ s.B ^ s.C ^ s.D ^ uint64(size))
}
