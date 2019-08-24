package autoimport_test

import (
	"io/ioutil"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/autoimport"
)

func TestSource(t *testing.T) {
	importer := autoimport.New(".")
	src, err := ioutil.ReadFile("testdata/hello.go.txt")
	assert.Nil(t, err)
	newSource, err := importer.Source(src)
	assert.Nil(t, err)
	assert.Contains(t, newSource, []byte("\"fmt\"\n"))
}
