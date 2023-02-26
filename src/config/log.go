package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

func InitLogger() (*log.Logger, bool) {
	//log.SetLevel(log.DebugLevel)
	//nolint:gosec ==> avoids lint check for gosec lint tool for this line. We can append more linting tools to avoid using comma
	// ex://nolint:govet,errcheck
	logFile, err := os.OpenFile(viper.GetString("LOG_PATH"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) //nolint:gosec
	if err != nil {
		log.Fatal(logFile, err)
		panic("Failed to open log file")
	}
	lvl, _ := log.ParseLevel(viper.GetString("LOG_LEVEL"))
	// close log file
	/*	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Errorf("Given logrus file failed to close: %+v", err)
		}
	}(logFile)*/
	logger := &log.Logger{
		Out:   os.Stderr,
		Hooks: nil,
		/*	Formatter: &log.JSONFormatter{
			TimestampFormat:   "",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			FieldMap: log.FieldMap{
				"FieldKeyTime":  "@timestamp",
				"FieldKeyLevel": "@level",
				"FieldKeyMsg":   "@message",
				"FieldKeyFunc":  "@caller",
			},
			CallerPrettyfier: nil,
			PrettyPrint:      true,
		},*/
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg% \n",
		},
		ReportCaller: true,
		Level:        lvl,
		ExitFunc:     os.Exit,
		BufferPool:   nil,
	}
	logger.SetOutput(logFile)
	return logger, true
}
