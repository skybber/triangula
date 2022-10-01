package utils

import (
	"github.com/skybber/Triangula/algorithm"
	"github.com/skybber/Triangula/algorithm/evaluator"
	"github.com/skybber/Triangula/fitness"
	"github.com/skybber/Triangula/generator"
	imageData "github.com/skybber/Triangula/image"
	"github.com/skybber/Triangula/mutation"
	"github.com/skybber/Triangula/normgeom"
	"image"
)

// DefaultAlgorithm returns an algorithm than will be optimal for almost all cases
func DefaultAlgorithm(numPoints int, image image.Image) algorithm.Algorithm {
	img := imageData.ToData(image)

	pointFactory := func() normgeom.NormPointGroup {
		return (generator.RandomGenerator{}).Generate(numPoints)
	}

	evaluatorFactory := func(n int) evaluator.Evaluator {
		return evaluator.NewParallel(fitness.TrianglesImageFunctions(img, 5, n), 22)
	}

	var mutator mutation.Method

	mutator = mutation.DefaultGaussianMethod(numPoints)

	algo := algorithm.NewSimple(pointFactory, 400, 5, evaluatorFactory, mutator)
	return algo
}
