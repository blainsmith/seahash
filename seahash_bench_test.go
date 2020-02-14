package seahash

import (
	"hash/fnv"
	"testing"

	creachadairCity "github.com/creachadair/cityhash"
	dgryskiSpooky "github.com/dgryski/go-spooky"
	dgryskiStadtx "github.com/dgryski/go-stadtx"
	huichenMurmur "github.com/huichen/murmur"
	farmhash "github.com/leemcloughlin/gofarmhash"
	"github.com/pborman/uuid"
	reuseeMurmur "github.com/reusee/mmh3"
	hashlandSpooky "github.com/tildeleb/hashland/spooky"
	zhangMurmur "github.com/zhangxinngang/murmur"
)

var result interface{}

func mkinput(n int) [][]byte {
	rv := make([][]byte, n)
	for i := 0; i < n; i++ {
		rv[i] = uuid.NewRandom()
	}
	return rv
}
func BenchmarkSeahash64(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint64, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		h := New()
		h.Write(input[n])
		output[n] = h.Sum64()
	}
	result = output
}
func BenchmarkFnvHash32(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint32, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		h := fnv.New32()
		h.Write(input[n])
		output[n] = h.Sum32()
	}
	result = output
}
func BenchmarkFnvHash64(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint64, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		h := fnv.New64()
		h.Write(input[n])
		output[n] = h.Sum64()
	}
	result = output
}
func BenchmarkFarmHashHash32(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint32, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = farmhash.Hash32(input[n])
	}
	result = output
}
func BenchmarkFarmHashHash64(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint64, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = farmhash.Hash64(input[n])
	}
	result = output
}
func BenchmarkHuichenMurmur(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint32, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = huichenMurmur.Murmur3(input[n])
	}
	result = output
}
func BenchmarkReuseeMurmur(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint32, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = reuseeMurmur.Sum32(input[n])
	}
	result = output
}
func BenchmarkZhangMurmur(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint32, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = zhangMurmur.Murmur3(input[n])
	}
	result = output
}
func BenchmarkDgryskiSpooky32(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint32, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = dgryskiSpooky.Hash32(input[n])
	}
	result = output
}
func BenchmarkDgryskiSpooky64(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint64, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = dgryskiSpooky.Hash64(input[n])
	}
	result = output
}

func BenchmarkDgryskiStatdx64(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint64, b.N)

	seeds := []uint64{0x0078d65f4a3b926c, 0x87ff56a5543dcb31}
	st := dgryskiStadtx.SeedState(seeds)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = dgryskiStadtx.Hash(&st, input[n])
	}
	result = output
}

func BenchmarkHashlandSpooky32(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint32, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = hashlandSpooky.Hash32(input[n], 0)
	}
	result = output
}
func BenchmarkHashlandSpooky64(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint64, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = hashlandSpooky.Hash64(input[n], 0)
	}
	result = output
}
func BenchmarkCreachadairCity32(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint32, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = creachadairCity.Hash32(input[n])
	}
	result = output
}
func BenchmarkCreachadairCity64(b *testing.B) {
	input := mkinput(b.N)
	output := make([]uint64, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output[n] = creachadairCity.Hash64(input[n])
	}
	result = output
}
