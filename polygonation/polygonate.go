package polygonation

import (
	"github.com/skybber/Triangula/geom"
	"github.com/skybber/Triangula/normgeom"
	"github.com/skybber/Triangula/triangulation/incrdelaunay"
	"math"
)

func Polygonate(points normgeom.NormPointGroup, w, h int) []geom.Polygon {
	fW, fH := float64(w), float64(h)

	triangulation := incrdelaunay.NewDelaunay(w, h)
	for _, p := range points {
		triangulation.Insert(incrdelaunay.Point{
			X: int32(math.Round(p.X * fW)),
			Y: int32(math.Round(p.Y * fH)),
		})
	}

	var polygons []geom.Polygon

	incrdelaunay.Voronoi(triangulation, func(points []incrdelaunay.FloatPoint) {
		var polygon geom.Polygon

		for _, p := range points {
			new := geom.Point{
				X: int(math.Round(p.X)),
				Y: int(math.Round(p.Y)),
			}

			if len(polygon.Points) == 0 || polygon.Points[len(polygon.Points)-1] != new {
				polygon.Points = append(polygon.Points, new)
			}
		}

		polygons = append(polygons, polygon)
	}, w, h)

	return polygons
}
