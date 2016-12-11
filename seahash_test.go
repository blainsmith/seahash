package seahash_test

import (
	"fmt"
	"testing"

	"github.com/blainsmith/seahash"
)

func ExampleHash() {
	// hash some bytes
	hash := seahash.Hash([]byte("to be or not to be"))
	fmt.Println(hash)
	// Output: 1988685042348123509
}

func TestHash(t *testing.T) {
	if seahash.Hash([]byte("to be or not to be")) != 1988685042348123509 {
		t.Fail()
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seahash.Hash([]byte("to be or not to be"))
	}
}

func TestDiffuse(t *testing.T) {
	if seahash.Diffuse(94203824938) != 17289265692384716055 {
		t.Fail()
	}
	if seahash.Diffuse(0xDEADBEEF) != 12110756357096144265 {
		t.Fail()
	}
	if seahash.Diffuse(0) != 0 {
		t.Fail()
	}
	if seahash.Diffuse(1) != 15197155197312260123 {
		t.Fail()
	}
	if seahash.Diffuse(2) != 1571904453004118546 {
		t.Fail()
	}
	if seahash.Diffuse(3) != 16467633989910088880 {
		t.Fail()
	}
}

func BenchmarkDiffuse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seahash.Diffuse(94203824938)
	}
}
