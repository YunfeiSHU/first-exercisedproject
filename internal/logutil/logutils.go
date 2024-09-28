package logutil

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	file, err := os.OpenFile("cmd/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("failed opening file: %v", err)
		panic(err)
	}
	logger.SetOutput(io.MultiWriter(file, os.Stdout))
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	logger.SetReportCaller(true)
	return logger
}
