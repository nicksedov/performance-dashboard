package profiles

/*
 * Reading and initializing settings from file on service startup (default name 'settings.yaml')
 * Including:
 *  - PostgreSQL database settings
 *  - Telegram connection settings
 */
import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/go-yaml/yaml"

	"performance-dashboard/pkg/cli"
)

var settings *Settings

func GetSettings() *Settings {
	if settings == nil {
		readSettingsFile()
	}
	return settings
}

func readSettingsFile() {
	settings = &Settings{}

	profilePath := path.Join(*cli.FlagConfigPath, "application-" + *cli.FlagProfile + ".yaml")
	log.Printf("Discovering profile by path: %s", profilePath)
	if strings.TrimSpace(*cli.FlagProfile) != "" {
		ymlFile, ioErr := os.ReadFile(profilePath)
		if ioErr == nil {
			ymlErr := yaml.Unmarshal(ymlFile, &settings)
			if ymlErr != nil {
				log.Fatal(ymlErr)
			}
		} else {
			log.Panic("Error reding profile")
		}
	} else {
		log.Fatalf("Wrong profile name: %s", profilePath)
	}
}
