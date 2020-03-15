package envdir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunCmd(t *testing.T) {
	envVars := map[string]string{
		"DJANGO_SETTINGS_MODULE": "mysite.settings",
		"MYSITE_DEBUG":           "true",
		"MYSITE_DEPLOY_DIR":      "/app/mysite",
		"MYSITE_SECRET_KEY":      "Nebuchadnezzar",
	}
	t.Run("run cmd", func(t *testing.T) {
		cmd := []string{"printenv", "MYSITE_DEPLOY_DIR"}
		out, code := RunCmd(cmd, envVars)
		assert.Equal(t, code, 0)
		assert.Equal(t, out, "/app/mysite")
	})
}
