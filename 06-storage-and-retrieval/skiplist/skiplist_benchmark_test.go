package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateRandomData(size int) (keys, values [][]byte) {
	keys = make([][]byte, size)
	values = make([][]byte, size)

	for i := 0; i < size; i++ {
		keys[i] = []byte(fmt.Sprintf("Key%d", rand.Intn(size)))
		values[i] = []byte(fmt.Sprintf("Value%d", i))
	}

	return keys, values
}

func BenchmarkPut100Items(b *testing.B) {
	keys, values := generateRandomData(100)
	memdb := NewMemDb()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < len(keys); i++ {
			memdb.Put(keys[i], values[i])
		}
	}
}
func BenchmarkPut1000Items(b *testing.B) {
	keys, values := generateRandomData(1000)
	memdb := NewMemDb()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < len(keys); i++ {
			memdb.Put(keys[i], values[i])
		}
	}
}
func BenchmarkPut10000Items(b *testing.B) {
	keys, values := generateRandomData(10000)
	memdb := NewMemDb()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < len(keys); i++ {
			memdb.Put(keys[i], values[i])
		}
	}
}
