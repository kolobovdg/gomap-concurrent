package gomap_concurrent

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMapCreation(t *testing.T) {
	targetShards := uint16(32)
	m := NewMapWithResolver(baseIntResolver, targetShards)
	if m.ShardCount != targetShards {
		t.Error("Incorrect map info")
	}

	if uint16(len(m.shards)) != targetShards {
		t.Error("Incorrect count of shards was created")
	}
}

func TestGMapCRUD(t *testing.T) {
	targetShards := uint16(32)
	roundsNum := 128

	m := NewMapWithResolver(baseIntResolver, targetShards)

	for i := 0; i < roundsNum; i++ {
		m.Set(i, i)
		if x, _ := m.Get(i); x.(int) != i {
			t.Error("Fail while setting/ getting a value")
		}
		if x := m.Len(); x != i+1 {
			t.Error("Where is key set problem")
		}
	}

	for i := 0; i < roundsNum; i++ {
		if value, exists := m.Delete(i); exists && value.(int) == i {
			if (roundsNum - i - 1) != m.Len() {
				t.Error("Count fail on delete")
			}
		} else {
			t.Error("Where is key set problem")
		}
	}
}

func _tryGetSetDel(t *testing.T, resolver func(key interface{}, shards uint16) uint, key interface{}) {
	targetShards := uint16(32)
	m := NewMapWithResolver(resolver, targetShards)

	// Prepare value to simplification different tests
	val := "x"

	m.Set(key, val)

	if l := m.Len(); l != 1 {
		t.Error("Fail set element during test")
	}

	if v, _ := m.Get(key); v.(string) != val {
		t.Error("Value vas corrupted")
	}

	m.Delete(key)
	if l := m.Len(); l != 0 {
		t.Error("Fail del element during test")
	}
}

func TestUintResolver(t *testing.T) {
	_tryGetSetDel(t, baseUint32Resolver, uint32(rand.Uint32()))
}

func TestStringResolver(t *testing.T) {
	_tryGetSetDel(t, baseStringResolver, strconv.Itoa(rand.Int()))
}

func TestStructResolver(t *testing.T) {
	rnd := rand.Uint32()

	_tryGetSetDel(t, exampleTestStructResolver,
		TestStruct{sVar: strconv.Itoa(int(rnd + 1)), uintVar: rnd})
}

func TestMixedResolver(t *testing.T) {
	var varsList []interface{}

	// Try to use different types + time.Time to show default sharder works
	varsList = append(varsList, rand.Int(), rand.Uint32(), strconv.Itoa(rand.Int()),
		TestStruct{sVar: strconv.Itoa(rand.Int()), uintVar: rand.Uint32()}, time.Now())
	for _, v := range varsList {
		_tryGetSetDel(t, exampleTestMixedStructResolver, v)
	}

}
