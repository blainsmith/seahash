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
