package horse

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Gallop(path string) (out []byte) {
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

func GallopIncludeHidden(path string) (out []byte) {
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
