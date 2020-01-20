package action

import "github.com/cemnura/scout/pkg/horse"
import "github.com/cemnura/scout/pkg/tally"

func Count(value, path string) int {
	scanned := horse.Gallop(path)
	c, _ := tally.Tally(scanned, []byte(value))
	return c
}
