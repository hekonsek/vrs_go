package semver

import (
	"fmt"
	"strconv"
	"strings"
)

type SemVer struct {
	major       int
	minor       int
	maintenance int
}

func NewSemver(major int, minor int, maintenance int) *SemVer {
	return &SemVer{major: major, minor: minor, maintenance: maintenance}
}

func ParseSemver(versionString string) (*SemVer, error) {
	components := strings.Split(versionString, ".")
	major, err := strconv.Atoi(components[0])
	if err != nil {
		return nil, err
	}
	minor, err := strconv.Atoi(components[1])
	if err != nil {
		return nil, err
	}
	maintenance, err := strconv.Atoi(components[2])
	if err != nil {
		return nil, err
	}
	return &SemVer{major: major, minor: minor, maintenance: maintenance}, nil
}

func (semVer *SemVer) BumpMinor() {
	semVer.minor += 1
	semVer.maintenance = 0
}

func (semVer *SemVer) ToString() string {
	return fmt.Sprintf("%d.%d.%d", semVer.major, semVer.minor, semVer.maintenance)
}
