package space

import (
    "math"
)

type Planet string

var orbits = map[Planet]float64 {
    "Earth":        31557600,
    "Mercury":      7600543.81992,
    "Venus":        19414149.052176,
    "Mars":         59354032.69008,
    "Jupiter":      374355659.124,
    "Saturn":       929292362.8848,
    "Uranus":       2651370019.3296,
    "Neptune":      5200418560.032,
}

func Age(age float64, planet Planet) float64 {
    return math.Round(age / orbits[planet] * 100) / 100
}
