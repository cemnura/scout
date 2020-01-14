/*
Copyright The Scout Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package tally

import (
	"io/ioutil"
	"testing"
)

type testpair struct {
	value    []byte
	expected int
}

var caseSensitivePair = []testpair{
	{[]byte("Lorem"), 1},
	{[]byte("nullam"), 2},
	{[]byte("lorem"), 0},
}

var caseInSensitivePair = []testpair{
	{[]byte("Lorem"), 1},
	{[]byte("nullam"), 3},
	{[]byte("lorem"), 1},
}

func TestCaseSensitive(t *testing.T) {
	testfile, _ := ioutil.ReadFile("testdata/lorem.txt")
	for _, pair := range caseSensitivePair {
		count, _ := Tally(testfile, pair.value)
		if count != pair.expected {
			t.Errorf("Didn't find the expected tally of %d, got %d", pair.expected, count)
		}
	}
}

func TestCaseInsensitive(t *testing.T) {
	testfile, _ := ioutil.ReadFile("testdata/lorem.txt")
	for _, pair := range caseInSensitivePair {
		count, _ := TallyCaseSensitive(testfile, pair.value, false)
		if count != pair.expected {
			t.Errorf("Didn't find the expected tally of %d, got %d of %s", pair.expected, count, pair.value)
		}
	}
}
