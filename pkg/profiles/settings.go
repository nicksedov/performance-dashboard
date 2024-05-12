package profiles

import (
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Settings struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`

	JiraConfig struct {
		BaseURL    string `yaml:"url"`
		ProjectKey string `yaml:"projectKey"`
		BoardID    string `yaml:"boardId"`
		Auth       struct {
			Type     string `yaml:"type"`
			ClientId string `yaml:"clientId"`
			ApiToken string `yaml:"apiToken"`
		} `yaml:"auth"`
	} `yaml:"jira"`

	HttpClientConfig struct {
		RequestTimeout   time.Duration `yaml:"requestTimeout"`
		RequestRateLimit int           `yaml:"requestRateLimit"`
		RetryLimit       int           `yaml:"retryLimit"`
	} `yaml:"httpClient"`

	DbConfig struct {
		DbNode  []Database `yaml:"node"`
	} `yaml:"database"`

	Logger struct {
		File    lumberjack.Logger `yaml:"file"`
		Console struct {
			Mode string `yaml:"mode"`
		}
	} `yaml:"logger"`

	Schedule struct {
		Task []TaskConfig `yaml:"task"`
	} `yaml:"schedule"`
}

type Database struct {
	Host       string `yaml:"host"`
	Port       uint   `yaml:"port"`
	DbName     string `yaml:"db_name"`
	SearchPath string `yaml:"search_path"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	SSLMode    string `yaml:"ssl_mode"`
}

type TaskConfig struct {
	ID               string        `yaml:"id"`
	Period           time.Duration `yaml:"period"`
	ExecuteOnStartup bool          `yaml:"executeOnStartup"`
	DelayedStart     time.Duration `yaml:"delayedStart"`
}