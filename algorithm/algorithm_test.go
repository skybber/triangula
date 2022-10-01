package algorithm

import (
	"github.com/skybber/Triangula/algorithm/evaluator"
	"github.com/skybber/Triangula/fitness"
	"github.com/skybber/Triangula/generator"
	imageData "github.com/skybber/Triangula/image"
	"github.com/skybber/Triangula/mutation"
	"github.com/skybber/Triangula/normgeom"
	"github.com/skybber/Triangula/random"
	"image"
	_ "image/jpeg"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func BenchmarkAlgorithm(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	random.Seed(time.Now().UnixNano())

	file, err := os.Open("../imgs/clown.jpg")

	if err != nil {
		panic(err)
	}

	imageFile, _, err := image.Decode(file)

	file.Close()

	if err != nil {
		log.Fatal(err)
	}

	imgData := imageData.ToData(imageFile)

	if err != nil {
		log.Fatal("Arg #2 not an integer")
	}

	pointFactory := func() normgeom.NormPointGroup {

		return (generator.RandomGenerator{}).Generate(1000)
	}
	evaluatorFactory := func(n int) evaluator.Evaluator {
		return evaluator.NewParallel(fitness.TrianglesImageFunctions(imgData, 5, n), 22)
	}

	mutator := mutation.NewGaussianMethod(2/1000, 0.3)

	algo := NewModifiedGenetic(pointFactory, 400, 5, evaluatorFactory, mutator)

	real := func() {
		for i := 0; i < 3000; i++ {
			algo.Step()
		}
	}
	real()
}
