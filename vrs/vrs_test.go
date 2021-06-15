package vrs_test

import (
	"github.com/hekonsek/vrs/vrs"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestSerializeVersionOnly(t *testing.T) {
	// Given
	config := &vrs.VrsConfig{Version: "0.0.0"}

	// When
	configBytes, err := yaml.Marshal(config)
	assert.NoError(t, err)

	// Then
	assert.Equal(t, "version: 0.0.0\n", string(configBytes))
}

func TestReadCurrentVersionDetectNoVersioonFile(t *testing.T) {
	// Given
	basedir, err := ioutil.TempDir("", "ver-test-*")
	assert.NoError(t, err)
	options, err := vrs.NewDefaultReadCurrentOptions()
	assert.NoError(t, err)
	options.Basedir = basedir
	options.GitCommit = false
	options.GitPush = false

	// When
	_, err = vrs.ReadCurrentVersion(options)

	// Then
	assert.Equal(t, err, vrs.NoVersioonFileFound)
}

func TestVersionBump(t *testing.T) {
	// Given
	basedir, err := ioutil.TempDir("", "ver-test-*")
	assert.NoError(t, err)
	initOptions, err := vrs.NewDefaultInitOptions()
	assert.NoError(t, err)
	initOptions.GitCommit = false
	initOptions.GitPush = false
	initOptions.Basedir = basedir
	err = vrs.Init(initOptions)
	assert.NoError(t, err)
	options := &vrs.BumpOptions{Basedir: basedir, GitCommit: false}

	// When
	err = vrs.Bump(options)

	// Then
	assert.NoError(t, err)
	readOptions, err := vrs.NewDefaultReadCurrentOptions()
	assert.NoError(t, err)
	readOptions.Basedir = basedir
	version, err := vrs.ReadCurrentVersion(readOptions)
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
	initOptions, err := vrs.NewDefaultInitOptions()
	assert.NoError(t, err)
	initOptions.Basedir = basedir
	initOptions.GitPush = false
	err = vrs.Init(initOptions)
	assert.NoError(t, err)
	options := &vrs.BumpOptions{Basedir: basedir, GitCommit: true, GitPush: false}

	// When
	err = vrs.Bump(options)

	// Then
	assert.NoError(t, err)
	readOptions, err := vrs.NewDefaultReadCurrentOptions()
	assert.NoError(t, err)
	readOptions.Basedir = basedir
	version, err := vrs.ReadCurrentVersion(readOptions)
	assert.NoError(t, err)
	assert.Equal(t, "0.1.0", version)

}
