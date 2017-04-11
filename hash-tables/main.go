package main

import (
	"log"
	"math/rand"
	"time"
)

type Element struct {
	Value interface{}
}

type HashFunc func(e Element) uint32

type SimpleHashTable struct {
	Table []Element
	Hash  HashFunc
	Rand  *rand.Rand
}

func NewHashTable(size int, seed int64) *SimpleHashTable {
	var r *rand.Rand

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	r = rand.New(rand.NewSource(seed))

	hash := makeHashFunc(r, size)

	table := &SimpleHashTable{
		Table: make([]Element, size),
		Hash:  hash,
		Rand:  r,
	}

	return table
}

type HashTable interface {
	Contains(e Element) bool
	Insert(e Element) bool
	Remove(e Element)
}

// Contains returns true if the hash table contains the element
// and false if it does not.
func (sht *SimpleHashTable) Contains(e Element) bool {
	k := sht.Hash(e)

	if sht.Table[k] == e {
		// The element is inside the table
		return true
	}

	// The element does not exist there
	return false
}

// Insert attempts to insert the desired value into the hash table.
// It returns false if a value was not able to be inserted.
func (sht *SimpleHashTable) Insert(e Element) bool {
	if sht.Contains(e) {
		// The element already exists in the table, no need to insert!
		return true
	}

	success, failedE := sht.doInsert(e)
	if !success {
		// The element could not be inserted
		log.Printf("Failed to insert %v into table", failedE)
		return false
	}

	return true
}

// doInsert inserts a value into the table. If a value exists in the location
// of insertion, the original value is displaed and returned, while the
// element to be inserted is added to the table.
func (sht *SimpleHashTable) doInsert(e Element) (bool, Element) {
	var displacedE Element

	for i := 0; i < 10; i++ {
		k := sht.Hash(e)
		if sht.Table[k].Value == nil {
			sht.Table[k] = e
			return true, Element{}
		}

		displacedE = sht.Table[k]
		sht.Table[k] = e
	}

	return false, displacedE
}

// Remove deletes a value from the hash table.
func (sht *SimpleHashTable) Remove(e Element) {
	k := sht.Hash(e)

	if sht.Table[k] == e {
		sht.Table[k] = Element{nil}
		return
	}
}

// log2 returns the number of iterations to reach 0 value
// minus 1.
func log2(size int) uint {
	r := 0
	for size != 0 {
		size >>= 1
		r++
	}

	return uint(r - 1)
}

// makeHashFunc generates a hash function.
// from http://www.keithschwarz.com/interesting/code/cuckoo-hashmap/CuckooHashMap.java.html
func makeHashFunc(r *rand.Rand, size int) HashFunc {
	A := uint(r.Uint32())
	B := uint(r.Uint32())

	return func(e Element) uint32 {
		value := e.Value.(int)
		high := uint(value >> 16)
		low := uint(value & 0x0000FFFF)

		hash := uint32(high*A + low*B)
		return hash >> (32 - log2(size))
	}
}
