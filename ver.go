package ver

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

const VersioonConfigFileName = "versioon.yml"

type VersioonConfig struct {
	Version string
	Sync    *Sync
}

type Sync struct {
	Files []SyncFile
}

type SyncFile struct {
	Name string
}

var NoVersioonFileFound = errors.New("no versioon file found")

func ParseVersioonConfig(basePath string) (*VersioonConfig, error) {
	versioonConfigPath := path.Join(basePath, VersioonConfigFileName)
	if _, err := os.Stat(versioonConfigPath); err != nil {
		if os.IsNotExist(err) {
			return nil, NoVersioonFileFound
		}
	}

	yml, err := ioutil.ReadFile(versioonConfigPath)
	if err != nil {
		return nil, err
	}

	config := &VersioonConfig{}
	err = yaml.Unmarshal(yml, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (config *VersioonConfig) Write(basePath string) error {
	yml, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(basePath, VersioonConfigFileName), yml, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (config *VersioonConfig) WriteAndCommit(baseDir string, commit bool, commitMessage string) error {
	err := config.Write(baseDir)
	if err != nil {
		return err
	}

	if commit {
		cmd := exec.Command("git", "add", VersioonConfigFileName)
		cmd.Dir = baseDir
		err = cmd.Run()
		if err != nil {
			return err
		}

		cmd = exec.Command("git", "commit", "-m", commitMessage)
		cmd.Dir = baseDir
		err = cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

type InitOptions struct {
	Basedir   string
	GitCommit bool
}

func NewDefaultInitOptions() (*InitOptions, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &InitOptions{
		Basedir:   wd,
		GitCommit: true,
	}, nil
}

func Init(options *InitOptions) error {
	if options == nil {
		o, err := NewDefaultInitOptions()
		if err != nil {
			return err
		}
		options = o
	}
	err := (&VersioonConfig{Version: "0.0.0"}).WriteAndCommit(options.Basedir, options.GitCommit, "Initialized versioon file.")
	if err != nil {
		return err
	}
	return nil
}

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

	config, err := ParseVersioonConfig(options.Basedir)
	if err != nil {
		return err
	}

	oldVersion := config.Version
	versionParts := strings.Split(oldVersion, ".")
	minorVersion, err := strconv.Atoi(versionParts[1])
	if err != nil {
		return err
	}
	config.Version = fmt.Sprintf("%s.%d.%s", versionParts[0], minorVersion+1, versionParts[2])
	err = config.WriteAndCommit(options.Basedir, options.GitCommit, "Version bump.")
	if err != nil {
		return err
	}

	if config.Sync != nil {
		for _, file := range config.Sync.Files {
			err = bumpInFile(options.Basedir, options.GitCommit, file.Name, oldVersion, config.Version)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func bumpInFile(baseDir string, gitCommit bool, file string, oldVersion string, newVersion string) error {
	filePath := path.Join(baseDir, file)
	originalBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	bumpedFile := strings.ReplaceAll(string(originalBytes), oldVersion, newVersion)

	err = ioutil.WriteFile(filePath, []byte(bumpedFile), 0644)
	if err != nil {
		return err
	}

	if gitCommit {
		cmd := exec.Command("git", "add", file)
		cmd.Dir = baseDir
		err = cmd.Run()
		if err != nil {
			return err
		}

		cmd = exec.Command("git", "commit", "-m", "Bumped version.")
		cmd.Dir = baseDir
		err = cmd.Run()
		if err != nil {
			return err
		}
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

	config, err := ParseVersioonConfig(options.Basedir)
	if err != nil {
		return "", err
	}
	return config.Version, nil
}
