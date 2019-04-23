package autoimport_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/akyoto/autoimport"
)

func TestSource(t *testing.T) {
	importer := autoimport.New("/home/eduard/projects/animenotifier/notify.moe")
	src, _ := ioutil.ReadFile("testdata/activity.go.txt")
	srcWithImports := importer.Source(src)
	src = srcWithImports
	// assert.NotEqual(t, src, srcWithImports)
}

func TestComponents(t *testing.T) {
	importer := autoimport.New("/home/eduard/projects/animenotifier/notify.moe")
	files, _ := ioutil.ReadDir("/home/eduard/projects/animenotifier/notify.moe/components")

	for _, file := range files {
		fmt.Println(file.Name())
		src, _ := ioutil.ReadFile("/home/eduard/projects/animenotifier/notify.moe/components/" + file.Name())
		importer.Source(src)
	}
}
