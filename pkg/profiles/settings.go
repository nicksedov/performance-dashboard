package profiles

import "gopkg.in/natefinch/lumberjack.v2"

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

	DbConfig struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		DbName   string `yaml:"db_name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"database"`

	Logger struct {
		File lumberjack.Logger `yaml:"file"`
		Console  struct {
			Mode string `yaml:"mode"`
		}
	} `yaml:"logger"`
}
