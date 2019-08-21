package build

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"os"
	"archive/tar"
	"io"
)

func TestToArchiveWithIgnore(t *testing.T) {
	pluginTar := &Tar{}
	buildCtx, err := pluginTar.ArchiveDirectory("./testdata/workdir/", ".dockerignore")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, buildCtx)

	buf := new(bytes.Buffer)
	buf.ReadFrom(buildCtx)

	dir, err := ioutil.TempDir("", "canary")
	if err != nil {
		t.Error(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile.tar.gz")
	if err := ioutil.WriteFile(tmpfn, buf.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}

	tmpfile, err := os.Open(tmpfn)
	defer tmpfile.Close()

	var fileNames []string
	tr := tar.NewReader(tmpfile)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
		}
		fileNames = append(fileNames, hdr.Name)
	}

	assert.Contains(t, fileNames, "Dockerfile")
	assert.Contains(t, fileNames, "test.txt")
	assert.NotContains(t, fileNames, "exclude.txt")
	assert.NotContains(t, fileNames, "log")
}
