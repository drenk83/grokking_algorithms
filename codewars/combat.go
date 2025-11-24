package codewars

import "math"

func CombatMy(health, damage float64) float64 {
	health -= damage

	if health < 0 {
		return 0.0
	}
	return health
}

func CombatMax(health, damage float64) float64 {
	return math.Max(health-damage, 0.0)
}
