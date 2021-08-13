package semver_test

import (
	"github.com/hekonsek/vrs/semver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBumpMinor(t *testing.T) {
	// Given
	v, err := semver.ParseSemver("0.0.0")
	assert.NoError(t, err)

	// When
	v.BumpMinor()
	vBumped := v.ToString()

	// Then
	assert.Equal(t, "0.1.0", vBumped)
}
