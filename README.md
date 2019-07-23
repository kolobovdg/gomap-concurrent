# gomap-concurrent
[![Build Status](https://api.travis-ci.com/kolobovdg/gomap-concurrent.svg?branch=master)](https://travis-ci.com/kolobovdg/gomap-concurrent)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolobovdg/gomap-concurrent)](https://goreportcard.com/report/github.com/kolobovdg/gomap-concurrent)
[![Sourcegraph](https://sourcegraph.com/github.com/kolobovdg/gomap-concurrent/-/badge.svg)](https://sourcegraph.com/github.com/kolobovdg/gomap-concurrent?badge)
[![GolangCI](https://golangci.com/badges/github.com/kolobovdg/gomap-concurrent.svg)](https://golangci.com)
[![codecov](https://codecov.io/gh/kolobovdg/gomap-concurrent/branch/master/graph/badge.svg)](https://codecov.io/gh/kolobovdg/gomap-concurrent)

Concurrent general-purpose shared map for using with flexible data.

## Import

Import the package:

```go
import (
	gmap "github.com/kolobovdg/gomap-concurrent"
)
```
```bash
go get "github.com/kolobovdg/gomap-concurrent"
```

Running tests:
```bash
go test "github.com/kolobovdg/gomap-concurrent"
```

## Examples

```go
    // Resolver examples avaliable in map_resolvers.go
    func baseIntResolver(key interface{}, sCount uint16) uint {
    	return uint(key.(int)) % uint(sCount)
    }
	targetShards := uint16(32)
	m := NewMapWithResolver(resolver, targetShards)

	value := "x"

	m.Set(key, value)               // Set
	value, exists := m.Get(key)     // Get
	value, exists := m.Delete(key)  // Del
	
	lenOfMap := m.Len()
```

In case of need to store different types of keys:
```go
    // Test struct for make an example of struct resolving in shared map
    type TestStruct struct {
    	sVar    string
    	uintVar uint32
    }
    
    // Example resolver for mixed types of keys
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
    	// Default case if any others didn't fit
    	return 0
    }
```