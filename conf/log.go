package conf

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func LogInit() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})
	//log.SetReportCaller(true)
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
}
