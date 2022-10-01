package rasterize

import "github.com/skybber/Triangula/geom"

func DDAPolygonBlocks(polygon geom.Polygon, blockSize int, line func(x0, x1, y int), block func(x, y int)) {
	polygon.Triangulate(func(triangle geom.Triangle) {
		DDATriangleBlocks(triangle, blockSize, line, block)
	})
}

func DDAPolygon(polygon geom.Polygon, pixel func(x, y int)) {
	polygon.Triangulate(func(triangle geom.Triangle) {
		DDATriangle(triangle, pixel)
	})
}
