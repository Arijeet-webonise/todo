package logger

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/Sirupsen/logrus"
)

type ILogger interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Initialise()
}

type RealLogger struct {
	log *logrus.Logger
}

func (al *RealLogger) Info(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Info(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Info(args)
	}

}

func (al *RealLogger) Error(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Error(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Error(args)
	}

}

func (al *RealLogger) Debug(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Debug(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Debug(args)
	}
}

func (al *RealLogger) Initialise() {
	al.log = logrus.New()
	al.log.Formatter = new(logrus.TextFormatter)
	al.log.Out = os.Stdout
}
