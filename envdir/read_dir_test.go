package envdir

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestReadDirWithoutFiles(t *testing.T) {
	envVars := map[string]string{}
	pathToTestDir, err := CreateTempTestDir(envVars)
	if err != nil {
		t.Fatal("ReadDirTest: error of creating test dir", err)
	}

	defer func() {
		err := os.RemoveAll(pathToTestDir)
		if err != nil {
			t.Fatal("RemoveAll: ", err)
		}
	}()

	t.Run("Get env variables if dir is empty", func(t *testing.T) {
		envs, err := ReadDir(pathToTestDir)
		assert.NoError(t, err)
		assert.Equal(t, envs, map[string]string{})
	})

	t.Run("Get env variables if dir without files but has some directories empties and with files", func(t *testing.T) {
		emptyDirName := "empty_dir_name"
		err := os.Mkdir(filepath.Join(pathToTestDir, emptyDirName), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		notEmptyDirName := "not_empty_dir_name"
		pathToNotEmptyDirName := filepath.Join(pathToTestDir, notEmptyDirName)
		err = os.Mkdir(pathToNotEmptyDirName, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		tmpfn := filepath.Join(pathToNotEmptyDirName, "tmpfile")
		if err := ioutil.WriteFile(tmpfn, []byte("temporary file's content"), 0666); err != nil {
			log.Fatal(err)
		}

		envs, err := ReadDir(pathToTestDir)
		assert.NoError(t, err)
		assert.Equal(t, envs, map[string]string{})
	})

}

func TestReadDirWithFiles(t *testing.T) {
	envVars := map[string]string{
		"DJANGO_SETTINGS_MODULE": "mysite.settings",
		"MYSITE_DEBUG":           "true",
		"MYSITE_DEPLOY_DIR":      "/app/mysite",
		"MYSITE_SECRET_KEY":      "Nebuchadnezzar",
	}
	pathToTestDir, err := CreateTempTestDir(envVars)
	if err != nil {
		t.Fatal("ReadDirTest: error of creating test dir", err)
	}

	defer func() {
		err := os.RemoveAll(pathToTestDir)
		if err != nil {
			t.Fatal("RemoveAll: ", err)
		}
	}()

	t.Run("Get env variables", func(t *testing.T) {
		envs, err := ReadDir(pathToTestDir)
		assert.NoError(t, err)
		assert.Equal(t, envs, map[string]string{
			"DJANGO_SETTINGS_MODULE": "mysite.settings",
			"MYSITE_DEBUG":           "true",
			"MYSITE_DEPLOY_DIR":      "/app/mysite",
			"MYSITE_SECRET_KEY":      "Nebuchadnezzar",
		})
	})
}
