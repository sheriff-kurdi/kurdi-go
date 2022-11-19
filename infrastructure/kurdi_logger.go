package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func Info(obj ...interface{}) {

	// create the logger
	logger := logrus.New()

	logger.SetReportCaller(true)

	// with Json Formatter
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	jsonObj, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	pc, fn, line, _ := runtime.Caller(1)
	fmt.Printf(" %s[%s:%d] %s", runtime.FuncForPC(pc).Name(), fn, line, jsonObj)
}
