package bloom

import (
	"github.com/spaolacci/murmur3"
)

type bloomFilter struct {
	m      uint64
	k      uint64
	filter []byte
}

func NewBloomFilter(m, k uint64) *bloomFilter {
	f := make([]byte, m)
	b := bloomFilter{m, k, f}
	return &b
}

func (b *bloomFilter) Add(key string) {
	h := murmur3.New128()
	h.Write([]byte(key))
	ht, he := h.Sum128()
	for i := uint64(1); i <= b.k; i++ {
		idx := (ht + i*he) % b.m
		b.filter[idx] = byte(1)
	}
}

func (b *bloomFilter) Contains(key string) bool {
	h := murmur3.New128()
	h.Write([]byte(key))
	ht, he := h.Sum128()
	for i := uint64(1); i <= b.k; i++ {
		idx := (ht + i*he) % b.m
		if b.filter[idx] == byte(0) {
			return false
		}
	}
	return true
}
