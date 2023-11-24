package logger

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type AyaLog struct{}

func (s *AyaLog) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:03:04")
	var newLog string
	//HasCaller() 为true才会有调用信息,entry.Caller.Function
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("%s %s %s:%d %s\n", timestamp, entry.Level, fName, entry.Caller.Line, entry.Message)
	} else {
		newLog = fmt.Sprintf("%s %s %s\n", timestamp, entry.Level, entry.Message)
	}
	b.WriteString(newLog)
	return b.Bytes(), nil
}

func NewAyaLog() (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&AyaLog{})
	// logger.SetFormatter(&logrus.JSONFormatter{
	// 	// DisableColors:   true,
	// 	// FullTimestamp:   true,
	// 	TimestampFormat: "2006-01-02 15:03:04",
	// 	CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
	// 		fi := path.Base(f.File)
	// 		return f.Function, fmt.Sprintf("%v %v", fi, f.Line)
	// 	},
	// })
	logger.SetLevel(logrus.InfoLevel)
	return logger, nil
}

type JobLog struct{}

func (s *JobLog) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:03:04")
	newLog := fmt.Sprintf("%s %s\n", timestamp, entry.Message)
	b.WriteString(newLog)
	return b.Bytes(), nil
}
func NewJobLog() (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&JobLog{})
	return logger, nil
}
