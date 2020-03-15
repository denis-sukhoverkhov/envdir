package envdir

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadDir(pathToDir string) (map[string]string, error) {
	envVars := make(map[string]string)

	files, err := ioutil.ReadDir(pathToDir)
	if err != nil {
		return envVars, err
	}

	for _, f := range files {
		if !f.IsDir() {
			name := f.Name()
			pathToFile := filepath.Join(pathToDir, name)
			f, err := os.Open(pathToFile)
			if err != nil {
				return envVars, err
			}
			value := make([]byte, 1000)
			n, err := f.Read(value)
			if err != nil {
				return envVars, err
			}
			envVars[name] = string(value[:n])
			_ = f.Close()
		}
	}

	return envVars, nil
}
