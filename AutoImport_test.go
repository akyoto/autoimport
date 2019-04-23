package autoimport_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/akyoto/autoimport"
	"github.com/stretchr/testify/assert"
)

func TestSource(t *testing.T) {
	importer := autoimport.New(".")
	src, err := ioutil.ReadFile("testdata/hello.go.txt")
	assert.NoError(t, err)
	newSource, err := importer.Source(src)
	assert.NoError(t, err)
	assert.True(t, bytes.Contains(newSource, []byte("fmt\n")))
}
