package horse

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Gallop(path string) (out []byte) {
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
