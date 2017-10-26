package bloom

import (
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	b1 := NewBloomFilter(1, 1)
	if b1.m != 1 {
		t.Error("b1.m not equal 1")
	}
	if b1.k != 1 {
		t.Error("b1.k not equal 1")
	}
	if len(b1.filter) != 1 {
		t.Error("b1.filter length not equal 1")
	}
	if b1.filter[0] != byte(0) {
		t.Error("b1.filter first element is not equal 0")
	}

	b2 := NewBloomFilter(100, 10)
	if b2.m != 100 {
		t.Error("b2.m not equal 100")
	}
	if b2.k != 10 {
		t.Error("b2.k not equal 10")
	}
	if len(b2.filter) != 100 {
		t.Error("b2.filter length not equal 100")
	}
	if b2.filter[0] != byte(0) {
		t.Error("b2.filter first element is not equal 0")
	}
}

func TestAdd(t *testing.T) {
	b1 := NewBloomFilter(100, 3)
	b1.Add("aaa")
	b1_total := 0
	for _, v := range b1.filter {
		b1_total += int(v)
	}
	if b1_total != 3 {
		t.Error("b1.filter active byte num is not equal 3")
	}

	b2 := NewBloomFilter(10000, 10)
	b2.Add("aaa")
	b2_total := 0
	for _, v := range b2.filter {
		b2_total += int(v)
	}
	if b2_total != 10 {
		t.Error("b2.filter active byte num is not equal 10")
	}
}

func TestContains(t *testing.T) {
	b := NewBloomFilter(1000, 3)
	b.Add("aaa")
	b.Add("bbb")
	b.Add("ccc")

	if b.Contains("aaa") != true {
		t.Error("There is no element expected in bloom filter")
	}
	if b.Contains("ddd") != false {
		t.Error("There is a element not expected in bloom filter")
	}
}
