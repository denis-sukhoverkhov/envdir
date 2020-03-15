package envdir

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func CreateTempTestDir(envVars map[string]string) (string, error) {

	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range envVars {
		f, err := os.Create(path.Join(dir, key))
		if err != nil {
			return "", err
		}
		_, err = f.WriteString(value)
		if err != nil {
			return "", err
		}
		_ = f.Close()
	}

	return dir, nil
}
