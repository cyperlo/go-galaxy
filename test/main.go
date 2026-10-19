package main

import (
	"log"
	"math"
)

const (
	MaxInt32 = math.MaxInt32 // 2147483647

	WarnThreshold     = MaxInt32 * 70 / 100 // 70%
	CriticalThreshold = MaxInt32 * 90 / 100 // 90%
)

func main() {
	log.Printf("%d", MaxInt32)
	log.Printf("%d", WarnThreshold)
	log.Printf("%d", CriticalThreshold)
}