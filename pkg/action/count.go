package action

import "github.com/cemnura/scout/pkg/horse"
import "github.com/cemnura/scout/pkg/tally"

func Count(value, path string, caseSensitive bool) (int, error) {
	scanned, err := horse.Gallop(path)

	if err != nil {
		return -1, err
	}
	c, _ := tally.TallyCaseSensitive(scanned, []byte(value), caseSensitive)
	return c, nil
}
