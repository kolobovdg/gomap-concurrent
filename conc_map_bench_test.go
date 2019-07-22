package gomap_concurrent

import (
	"testing"
)

func _SetItems(b *testing.B, shards uint16) {
	m := NewMapWithResolver(baseIntResolver, shards)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(i, "x")
	}
}

func BenchmarkSingleInsert_1_Shard(b *testing.B) {
	runWithCustomShardsNum(_SetItems, b, 1)
}
func BenchmarkSingleInsert_32_Shard(b *testing.B) {
	runWithCustomShardsNum(_SetItems, b, 32)
}
func BenchmarkSingleInsert_256_Shard(b *testing.B) {
	runWithCustomShardsNum(_SetItems, b, 256)
}

func BenchmarkSingleInsert_1024_Shard(b *testing.B) {
	runWithCustomShardsNum(_SetItems, b, 1024)
}

func runWithCustomShardsNum(bench func(b *testing.B, shards uint16), b *testing.B, shardsCount uint16) {
	bench(b, shardsCount)
}
