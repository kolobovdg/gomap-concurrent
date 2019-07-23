package gomap_concurrent

// Example resolver for int keys
func baseIntResolver(key interface{}, sCount uint16) uint {
	return uint(key.(int)) % uint(sCount)
}

// Example resolver for uint32 keys
func baseUint32Resolver(key interface{}, sCount uint16) uint {
	return uint(key.(uint32)) % uint(sCount)
}

// Example resolver for string keys
func baseStringResolver(key interface{}, sCount uint16) uint {
	return uint(fnv32(key.(string))) % uint(sCount)
}

// String to uint (from hash/fnv)
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// Test struct for make an example of struct resolving in shared map
type TestStruct struct {
	sVar    string
	uintVar uint32
}

// Example resolver for custom struct keys
func exampleTestStructResolver(key interface{}, sCount uint16) uint {
	return uint(key.(TestStruct).uintVar) % uint(sCount)
}

// Example resolver for mixed keys
func exampleTestMixedStructResolver(key interface{}, sCount uint16) uint {
	switch key.(type) {
	case int:
		return baseIntResolver(key, sCount)
	case uint32:
		return baseUint32Resolver(key, sCount)
	case string:
		return baseStringResolver(key, sCount)
	case TestStruct:
		return exampleTestStructResolver(key, sCount)
	}
	return 0
}
