package gomap_concurrent

func baseIntResolver(key interface{}, sCount uint16) uint {
	return uint(key.(int)) % uint(sCount)
}
