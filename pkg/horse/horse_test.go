package horse

import "testing"

func TestGallop(t *testing.T) {

	files := Gallop("testdata")

	if len(files) != 6 {
		t.Error("Expected 6, got ", len(files))
	}
}

func TestGallopIncludeHidden(t *testing.T) {

	files := GallopIncludeHidden("testdata")

	if len(files) != 8 {
		t.Error("Expected 8, got ", len(files))
	}
}
