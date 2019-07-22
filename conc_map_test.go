package gomap_concurrent

import (
	"testing"
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
