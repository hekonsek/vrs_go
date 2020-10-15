package ver_test

import (
	"github.com/hekonsek/ver"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestReadCurrentVersionDetectNoVersioonFile(t *testing.T) {
	// Given
	basedir, err := ioutil.TempDir("", "ver-test-*")
	assert.NoError(t, err)
	options, err := ver.NewDefaultReadCurrentOptions()
	assert.NoError(t, err)
	options.Basedir = basedir
	options.GitCommit = false

	// When
	_, err = ver.ReadCurrentVersion(options)

	// Then
	assert.Equal(t, err, ver.NoVersioonFileFound)
}

func TestVersionBump(t *testing.T) {
	// Given
	basedir, err := ioutil.TempDir("", "ver-test-*")
	assert.NoError(t, err)
	initOptions, err := ver.NewDefaultInitOptions()
	assert.NoError(t, err)
	initOptions.GitCommit = false
	initOptions.Basedir = basedir
	err = ver.Init(initOptions)
	assert.NoError(t, err)
	options := &ver.BumpOptions{Basedir: basedir, GitCommit: false}

	// When
	err = ver.Bump(options)

	// Then
	assert.NoError(t, err)
	readOptions, err := ver.NewDefaultReadCurrentOptions()
	assert.NoError(t, err)
	readOptions.Basedir = basedir
	version, err := ver.ReadCurrentVersion(readOptions)
	assert.NoError(t, err)
	assert.Equal(t, "0.1.0", version)

}

func TestVersionBumpWithCommit(t *testing.T) {
	// Given
	// Given
	basedir, err := ioutil.TempDir("", "ver-test-*")
	assert.NoError(t, err)
	err = exec.Command("git", "init", basedir).Run()
	assert.NoError(t, err)
	initOptions, err := ver.NewDefaultInitOptions()
	assert.NoError(t, err)
	initOptions.Basedir = basedir
	err = ver.Init(initOptions)
	assert.NoError(t, err)
	options := &ver.BumpOptions{Basedir: basedir, GitCommit: true}

	// When
	err = ver.Bump(options)

	// Then
	assert.NoError(t, err)
	readOptions, err := ver.NewDefaultReadCurrentOptions()
	assert.NoError(t, err)
	readOptions.Basedir = basedir
	version, err := ver.ReadCurrentVersion(readOptions)
	assert.NoError(t, err)
	assert.Equal(t, "0.1.0", version)

}
