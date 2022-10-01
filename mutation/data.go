package mutation

import "github.com/skybber/Triangula/normgeom"

// Mutation stores data related to a mutation.
type Mutation struct {
	Old, New normgeom.NormPoint // The point before and after being mutated.
	Index    int                // The index of the point mutated.
}
