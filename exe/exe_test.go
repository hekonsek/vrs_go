package exe_test

import (
	"github.com/hekonsek/vrs/exe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExeEcho(t *testing.T) {
	// Given
	exec := exe.New("echo foo")

	// When
	result := exec.Run()

	// Then
	assert.NoError(t, result.Err())
	assert.Len(t, result.Output(), 1)
	assert.True(t, result.Success())
}

func TestChangingDir(t *testing.T) {
	// Given
	exec := exe.New("cat resolv.conf").InDirectory("/etc")

	// When
	result := exec.Run()

	// Then
	assert.NoError(t, result.Err())
	assert.Greater(t, len(result.Output()), 1)
	assert.True(t, result.Success())
}

func TestNoSuccess(t *testing.T) {
	// Given
	exec := exe.New("cat noSuchFile")

	// When
	result := exec.Run()

	// Then
	assert.NoError(t, result.Err())
	assert.False(t, result.Success())
	assert.Len(t, result.Output(), 1)
	assert.Contains(t, result.Output()[0], "No such file or directory")
}

func TestNoSuccessReport(t *testing.T) {
	// Given
	exec := exe.New("cat noSuchFile")
	result := exec.Run()

	// When
	report := result.NoSuccessReport()

	// Then
	assert.NoError(t, result.Err())
	assert.False(t, result.Success())
	assert.Contains(t, report.Error(), "No such file or directory")
}
