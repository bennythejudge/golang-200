package utils

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	// AppName is the application's name
	AppName = "todolist"

	// LogStashFormatter is constant used to format logs as logstash format
	LogStashFormatter = "logstash"
	// TextFormatter is constant used to format logs as simple text format
	TextFormatter = "text"
)

// InitLog initializes the logrus logger
func InitLog(logLevel, formatter string) error {

	switch formatter {
	case LogStashFormatter:
		logrus.SetFormatter(&logrustash.LogstashFormatter{
			TimestampFormat: time.RFC3339,
			Type:            AppName,
		})
	default:
		// TODO Set the formatter using the standard logrus TextFormatter with Forced colors and Full timestamp
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
			FullTimestamp: true,
		})
	}

	// TODO set the standard output to os.Stdout
	logrus.SetOutput(os.Stdout)


	// TODO parse the logLevel param
	level, err := logrus.ParseLevel(logLevel)


	// TODO check the parsing error
	if err != nil {
		// TODO if error occurs set the logger level as DebugLevel and return the error
		logrus.SetLevel(logrus.DebugLevel)
		return err
	}
	
	// TODO if no error occurred, set the parsed level as the logger level
	logrus.SetLevel(level)
	return nil
}
