package profiles

import (
	"os"
	"performance-dashboard/pkg/cli"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSettings(t *testing.T) {
	relativePath := "../../profiles"
	cli.FlagConfigPath = &relativePath
	_, err := os.Stat(relativePath + "/application-default.yaml")
	assert.Nil(t, err)
	settings := GetSettings()
	assert.Equal(t, "jira_project", settings.Schedule.Task[0].ID)
	assert.Equal(t, true, settings.Schedule.Task[0].ExecuteOnStartup)
	assert.Equal(t, time.Duration(0), settings.Schedule.Task[0].DelayedStart)
	assert.Equal(t, time.Duration(time.Hour), settings.Schedule.Task[0].Period)
}