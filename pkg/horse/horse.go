package horse

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Gallop(path string) (out []byte, err error) {

	if _, err := isExist(path); err != nil {
		return nil, fmt.Errorf("There is no %v file or directory", path)
	}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			out = append(out, content...)
		} else if strings.HasPrefix(info.Name(), ".") && len(info.Name()) != 1 {
			return filepath.SkipDir

		}
		return nil
	})

	return
}

func GallopIncludeHidden(path string) (out []byte, err error) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			out = append(out, content...)
		}
		return nil
	})

	return
}

func isExist(path string) (bool, error) {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}
