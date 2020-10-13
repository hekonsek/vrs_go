package ver

import (
	"io/ioutil"
	"os"
	"path"
)

func Bump(basedir string) error {
	version := ""
	versionPath := path.Join(basedir, "version.txt")
	if _, err := os.Stat(versionPath); os.IsNotExist(err) {
		version = "0.0.0"
	} else {
		versionBytes, err := ioutil.ReadFile(versionPath)
		if err != nil {
			return err
		}
	}

	err := ioutil.WriteFile(versionPath, []byte(version), 0644)
	if err != nil {
		return err
	}

	return nil
}