package transformer

import (
	"errors"
	"github.com/azylman/getl"
	"github.com/azylman/getl/sources/infinite"
	"github.com/azylman/getl/tests"
	"testing"
)

// TestTransformError tests that the upstream Table had all of its data consumed in the case of an
// error from a rowTransformTable.
func TestTransformError(t *testing.T) {
	in := infinite.New()
	out := RowTransform(in, func(row getl.Row) (getl.Row, error) {
		return nil, errors.New("some error")
	})
	// Should receive no rows here because the first response was an error.
	tests.Consumed(t, out)
	// Should receive no rows here because the the transform should have consumed
	// all the rows.
	tests.Consumed(t, in)
}
