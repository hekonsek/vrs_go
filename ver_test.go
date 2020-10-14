package ver_test

import (
	"github.com/hekonsek/ver"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestReadDefaultVersionIfNoVersionFile(t *testing.T) {
	// Given
	basedir, err := ioutil.TempDir("", "ver-test-*")
	assert.NoError(t, err)
	options, err := ver.NewDefaultReadCurrentOptions()
	assert.NoError(t, err)
	options.Basedir = basedir
	options.GitCommit = false

	// When
	version, err := ver.ReadCurrentVersion(options)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "0.0.0", version)
}

func TestInitialVersionBump(t *testing.T) {
	// Given
	basedir, err := ioutil.TempDir("", "ver-test-*")
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

func TestVersionBump(t *testing.T) {
	// Given
	basedir, err := ioutil.TempDir("", "ver-test-*")
	assert.NoError(t, err)
	options := &ver.BumpOptions{Basedir: basedir, GitCommit: false}
	err = ver.Bump(options)
	assert.NoError(t, err)

	// When
	err = ver.Bump(options)

	// Then
	assert.NoError(t, err)
	readOptions, err := ver.NewDefaultReadCurrentOptions()
	assert.NoError(t, err)
	readOptions.Basedir = basedir
	version, err := ver.ReadCurrentVersion(readOptions)
	assert.NoError(t, err)
	assert.Equal(t, "0.2.0", version)
}

func TestVersionBumpWithCommit(t *testing.T) {
	// Given
	options, err := ver.NewDefaultBumpOptions()
	assert.NoError(t, err)
	basedir, err := ioutil.TempDir("", "ver-test-*")
	assert.NoError(t, err)
	options.Basedir = basedir
	err = exec.Command("git", "init", basedir).Run()
	assert.NoError(t, err)

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
