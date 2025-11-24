package codewars

import (
	"math/rand"
	"testing"
)

var health float64
var damage float64

func init() {
	health = rand.ExpFloat64()
	damage = rand.ExpFloat64()
}

func BenchmarkCombatMy(b *testing.B) {
	for b.Loop() {
		_ = CombatMy(health, damage)
	}
}

func BenchmarkCombatMax(b *testing.B) {
	for b.Loop() {
		_ = CombatMax(health, damage)
	}
}
