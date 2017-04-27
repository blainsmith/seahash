package seahash_test

import (
	"fmt"
	"testing"

	"github.com/blainsmith/seahash"
)

func ExampleSum() {
	// hash some bytes
	hash := seahash.Sum([]byte("to be or not to be"))
	fmt.Printf("%x", hash)
	// Output: 75e54a6f823a991b
}

func TestHash(t *testing.T) {
	h := seahash.New()
	h.Write([]byte("to be or "))
	h.Write([]byte("not to be"))
	s := fmt.Sprintf("%x", h.Sum(nil))

	if s != "75e54a6f823a991b" {
		t.Fail()
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seahash.Sum([]byte("to be or not to be"))
	}
}

func TestSizes(t *testing.T) {
	h := seahash.New()

	if h.Size() != 8 {
		t.Fail()
	}

	if h.BlockSize() != 8 {
		t.Fail()
	}
}
