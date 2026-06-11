package keys

import (
	"os"
	"path/filepath"
	"strings"
)

// ListKeys returns a slice of paths to potential SSH private keys in ~/.ssh
func ListKeys() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	sshDir := filepath.Join(home, ".ssh")
	files, err := os.ReadDir(sshDir)
	if err != nil {
		return nil, err
	}

	var keyPaths []string
	for _, file := range files {
		// Basic heuristic: private keys often don't have .pub extension
		if !file.IsDir() && !strings.HasSuffix(file.Name(), ".pub") && !strings.HasSuffix(file.Name(), "known_hosts") && !strings.HasSuffix(file.Name(), "config") {
			keyPaths = append(keyPaths, filepath.Join(sshDir, file.Name()))
		}
	}

	return keyPaths, nil
}
