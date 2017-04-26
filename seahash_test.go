package seahash

import (
	"fmt"
	"testing"
)

func ExampleHash() {
	// hash some bytes
	hash := Hash([]byte("to be or not to be"))
	fmt.Println(hash)
	// Output: 1988685042348123509
}

func TestHash(t *testing.T) {
	if Hash([]byte("to be or not to be")) != 1988685042348123509 {
		t.Fail()
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hash([]byte("to be or not to be"))
	}
}

func TestDiffuse(t *testing.T) {
	if diffuse(94203824938) != 17289265692384716055 {
		t.Fail()
	}
	if diffuse(0xDEADBEEF) != 12110756357096144265 {
		t.Fail()
	}
	if diffuse(0) != 0 {
		t.Fail()
	}
	if diffuse(1) != 15197155197312260123 {
		t.Fail()
	}
	if diffuse(2) != 1571904453004118546 {
		t.Fail()
	}
	if diffuse(3) != 16467633989910088880 {
		t.Fail()
	}
}

func BenchmarkDiffuse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diffuse(94203824938)
	}
}
