package ver

import (
	"fmt"
	"github.com/hekonsek/osexit"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type BumpOptions struct {
	Basedir   string
	GitCommit bool
}

func NewDefaultBumpOptions() (*BumpOptions, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &BumpOptions{
		Basedir:   wd,
		GitCommit: true,
	}, nil
}

func Bump(options *BumpOptions) error {
	if options == nil {
		o, err := NewDefaultBumpOptions()
		if err != nil {
			return err
		}
		options = o
	}
	version, err := ReadCurrentVersion(&ReadCurrentOptions{
		Basedir:   options.Basedir,
		GitCommit: options.GitCommit,
	})
	if err != nil {
		return err
	}

	versionParts := strings.Split(version, ".")
	minorVersion, err := strconv.Atoi(versionParts[1])
	if err != nil {
		return err
	}
	version = fmt.Sprintf("%s.%d.%s", versionParts[0], minorVersion+1, versionParts[2])

	versionPath := path.Join(options.Basedir, "version.txt")
	err = ioutil.WriteFile(versionPath, []byte(version), 0644)
	if err != nil {
		return err
	}

	if options.GitCommit {
		cmd := exec.Command("git", "add", "version.txt")
		cmd.Dir = options.Basedir
		err = cmd.Run()
		osexit.ExitOnError(err)

		cmd = exec.Command("git", "--git-dir", options.Basedir+"/.git", "commit", "-m", "Version bump.")
		cmd.Dir = options.Basedir
		err = cmd.Run()
		osexit.ExitOnError(err)
	}

	return nil
}

type ReadCurrentOptions struct {
	Basedir   string
	GitCommit bool
}

func NewDefaultReadCurrentOptions() (*ReadCurrentOptions, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &ReadCurrentOptions{
		Basedir:   wd,
		GitCommit: true,
	}, nil
}

func ReadCurrentVersion(options *ReadCurrentOptions) (string, error) {
	if options == nil {
		o, err := NewDefaultReadCurrentOptions()
		if err != nil {
			return "", err
		}
		options = o
	}

	versionPath := path.Join(options.Basedir, "version.txt")
	if _, err := os.Stat(versionPath); os.IsNotExist(err) {
		err = ioutil.WriteFile(versionPath, []byte("0.0.0"), 0644)

		if options.GitCommit {
			cmd := exec.Command("git", "add", "version.txt")
			cmd.Dir = options.Basedir
			err = cmd.Run()
			osexit.ExitOnError(err)

			cmd = exec.Command("git", "--git-dir", options.Basedir+"/.git", "commit", "-m", "Initial version commit.")
			cmd.Dir = options.Basedir
			err = cmd.Run()
			osexit.ExitOnError(err)
		}

		if err != nil {
			return "", err
		}
		return "0.0.0", nil
	} else {
		versionBytes, err := ioutil.ReadFile(versionPath)
		if err != nil {
			return "", err
		}
		return string(versionBytes), nil
	}
}
