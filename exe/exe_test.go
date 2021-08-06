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
	lines, success, err := exec.Run()

	// Then
	assert.NoError(t, err)
	assert.Len(t, lines, 1)
	assert.True(t, success)
}

func TestChangingDir(t *testing.T) {
	// Given
	exec := exe.New("cat resolv.conf").InDirectory("/etc")

	// When
	lines, success, err := exec.Run()

	// Then
	assert.NoError(t, err)
	assert.Greater(t, len(lines), 1)
	assert.True(t, success)
}

func TestNoSuccess(t *testing.T) {
	// Given
	exec := exe.New("cat noSuchFile")

	// When
	lines, success, err := exec.Run()

	// Then
	assert.NoError(t, err)
	assert.False(t, success)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0], "No such file or directory")
}

func TestNoSuccessReport(t *testing.T) {
	// Given
	exec := exe.New("cat noSuchFile")
	_, success, err := exec.Run()

	// When
	report := exec.NoSuccessReport()

	// Then
	assert.NoError(t, err)
	assert.False(t, success)
	assert.Contains(t, report.Error(), "No such file or directory")
}
