package horse

import "testing"

func TestGallop(t *testing.T) {

	files := Gallop("testdata")

	if len(files) != 6 {
		t.Error("Expected 6, got ", len(files))
	}
}
