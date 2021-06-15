# vrs - project versioning made easy

`vrs` is a command line tool simplifying versioning of git projects.

## How does it work?

All you have to do is to initialize vrs in the root directory of your project:

```bash
git git@github.com:my_username/my_project.git
cd my_project
vrs init
```

That's it! The command above creates the following `vrs.yml` file in your project:

```
version: 0.0.0
```

During the initialization process:
- `vrs` config file is automatically committed to version control.
- Created commit is tagged with project version from the `vrs` file.
- The commit and the tag are pushed into a remote Git repository.

## Installation

```bash
docker container create --name vrs hekonsek/vrs
sudo docker cp vrs:/vrs /usr/local/bin/
```