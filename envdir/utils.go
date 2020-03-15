package envdir

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func RunCmd(cmd []string, env map[string]string) (string, int) {
	for key, val := range env {
		err := os.Setenv(key, val)
		if err != nil {
			return "", 1
		}
	}

	cmdStruct := exec.Command(cmd[0], cmd[1:]...)
	var out bytes.Buffer
	cmdStruct.Stdout = &out
	err := cmdStruct.Run()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(out.String()), 0
}
