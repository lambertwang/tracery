package main

import "math"

type pointLight struct {
	center vector
	radius float64
}

func (l pointLight) light(incident vector, normal vector, s scene) (ray, float64) {
	// Compute intersections
	nearestT := math.MaxFloat64
	toRay := lineToRay(incident, l.center)
	for _, shape := range s.shapes {
		t := shape.intersect(toRay)
		// Assume ray cannot cast from inside of the sphere
		if t > 0.01 {
			if t < nearestT {
				nearestT = t
			}
		}
	}

	v := subtractVector(l.center, incident)

	if v.magnitude() > subtractVector(l.center, toRay.incident(nearestT)).magnitude() {
		return toRay, 0.0
	}

	distance := subtractVector(incident, l.center).magnitude()
	attenuation := 1.0 / (1.0 + (2.0/l.radius)*distance + (1.0/math.Pow(l.radius, 2))*math.Pow(distance, 2))

	return toRay, attenuation
}
