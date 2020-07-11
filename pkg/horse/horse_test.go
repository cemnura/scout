package horse

import "testing"

func TestGallop(t *testing.T) {

	occurrances, _ := Gallop("testdata", "1", true)

	if occurrances != 1 {
		t.Error("Expected 1, got ", occurrances)
	}
}

func TestGallopIncludeHidden(t *testing.T) {

	files, _ := GallopIncludeHidden("testdata")

	if len(files) != 8 {
		t.Error("Expected 8, got ", len(files))
	}
}

func TestGallopInvalidFile(t *testing.T) {

	_, err := Gallop("spoon", "1", true)

	if err == nil {
		t.Error("There is a spoon")
	}
}

func BenchmarkGallop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gallop("testdata", "1", true)
	}
}
