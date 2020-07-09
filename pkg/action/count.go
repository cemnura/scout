package action

import "github.com/cemnura/scout/pkg/horse"
import "github.com/cemnura/scout/pkg/tally"

func Count(value, path string, caseSensitive bool) int {
	scanned := horse.Gallop(path)
	c, _ := tally.TallyCaseSensitive(scanned, []byte(value), caseSensitive)
	return c
}
