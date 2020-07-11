package action

import "github.com/cemnura/scout/pkg/horse"

func Count(value, path string, caseSensitive bool) (int, error) {
	count, err := horse.Gallop(path, value, caseSensitive)

	if err != nil {
		return -1, err
	}
	return count, nil
}
