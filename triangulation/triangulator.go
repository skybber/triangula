// Package triangulation provides utilities for creating Delaunay triangulations.
package triangulation

import (
	"github.com/skybber/Triangula/geom"
	"github.com/skybber/Triangula/normgeom"
	"github.com/fogleman/delaunay"
	"math"
)

// Triangulate creates a Delaunay triangulation from a group of points and returns a set of triangles.
func Triangulate(points normgeom.NormPointGroup, w, h int) []geom.Triangle {

	group := make([]delaunay.Point, len(points))

	width := float64(w)
	height := float64(h)

	for i, p := range points {
		newPoint := delaunay.Point{X: math.Round(p.X * width), Y: math.Round(p.Y * height)}
		group[i] = newPoint
	}

	triangulation, _ := delaunay.Triangulate(group)

	triangles := triangulation.Triangles

	numTris := len(triangles) / 3

	newTriangles := make([]geom.Triangle, numTris)

	for i := 0; i < numTris; i++ {
		in := i * 3
		a := group[triangles[in]]
		b := group[triangles[in+1]]
		c := group[triangles[in+2]]
		tri := geom.NewTriangle(int(a.X), int(a.Y), int(b.X), int(b.Y), int(c.X), int(c.Y))
		newTriangles[i] = tri
	}

	return newTriangles
}

// Area returns the area of a convex polygon.
// Adapted from: https://www.geeksforgeeks.org/area-of-a-polygon-with-given-n-ordered-vertices/.
func Area(polygon []delaunay.Point) int {
	area := 0.

	j := len(polygon) - 1
	for i := 0; i < len(polygon); i++ {
		pI := polygon[i]
		pJ := polygon[j]
		area += (pJ.X + pI.X) * (pJ.Y - pI.Y)
		j = i
	}

	return int(math.Round(math.Abs(area / 2.)))
}
