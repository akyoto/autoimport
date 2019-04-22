package autoimport_test

import (
	"io/ioutil"
	"testing"

	"github.com/blitzprog/autoimport"
)

func TestSource(t *testing.T) {
	importer := autoimport.New("/home/eduard/projects/animenotifier/notify.moe")
	src, _ := ioutil.ReadFile("testdata/activity.go.txt")
	srcWithImports := importer.Source(src)
	src = srcWithImports
	// assert.NotEqual(t, src, srcWithImports)
}
