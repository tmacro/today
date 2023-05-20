package utils

import (
	"os/user"
	"path/filepath"
	"strings"
)

func ResolvePath(path string) (string, error) {
	path, err := expandTilde(path)
	if err != nil {
		return "", err
	}

	resolved, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return resolved, nil
}

func expandTilde(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	if path == "~" {
		homeDir, err := getHomeDir("")
		if err != nil {
			return "", err
		}
		return homeDir, nil
	}

	// "~user"
	if path[1] != '/' {
		parts := strings.SplitN(path, "/", 2)
		homeDir, err := getHomeDir(parts[0][1:])
		if err != nil {
			return "", err
		}
		return homeDir + "/" + parts[1], nil
	}

	// "~/"
	homeDir, err := getHomeDir("")
	if err != nil {
		return "", err
	}

	return homeDir + path[1:], nil
}

func getHomeDir(name string) (string, error) {
	var usr *user.User
	var err error
	if name != "" {
		usr, err = user.Lookup(name)
	} else {
		usr, err = user.Current()
	}

	if err != nil {
		return "", err
	}

	dir := usr.HomeDir

	return dir, nil
}
