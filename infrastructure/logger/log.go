package logger

import (
	"errors"
	"fmt"
	"io/ioutil"

	"akawork.io/constant"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var mLog *log.Logger

/**
 * Creates a new Logger
 * See more:
 * https://esc.sh/blog/golang-logging-using-logrus/
 * https://github.com/rifflock/lfshook
 */
func NewLogger(logPath string, logPrefix string) *log.Logger {
	if mLog != nil {
		return mLog
	}

	logPathMap := lfshook.PathMap{
		log.InfoLevel:  logPath + "/" + logPrefix + "_success.log",
		log.TraceLevel: logPath + "/" + logPrefix + "_success.log",
		log.WarnLevel:  logPath + "/" + logPrefix + "_success.log",

		log.DebugLevel: logPath + "/" + logPrefix + "_debug.log",

		log.ErrorLevel: logPath + "/" + logPrefix + "_error.log",
		log.FatalLevel: logPath + "/" + logPrefix + "_error.log",
		log.PanicLevel: logPath + "/" + logPrefix + "_error.log",
	}

	logFormatter := new(log.TextFormatter)
	logFormatter.TimestampFormat = "02-01-2006 15:04:05"
	logFormatter.FullTimestamp = true

	mLog = log.New()
	mLog.Hooks.Add(lfshook.NewHook(
		logPathMap,
		logFormatter,
	))

	if !viper.GetBool(`Debug`) {
		mLog.Out = ioutil.Discard
	}

	return mLog
}

/**
 * Logs trace
 */
func Trace(format string, v ...interface{}) {
	mLog.Tracef(format, v)
}

/**
 * Logs info
 */
func Info(format string, v ...interface{}) {
	mLog.Infof(constant.LogInfoPrefix+format, v...)
}

/**
 * Logs warning
 */
func Warn(format string, v ...interface{}) {
	mLog.Warnf(format, v...)
}

/**
 * Logs debug
 */
func Debug(format string, v ...interface{}) {
	mLog.Debugf(format, v...)
}

/**
 * Logs error
 */
func Error(format string, v ...interface{}) {
	mLog.Errorf(constant.LogErrorPrefix+format, v...)
}

/**
 * Logs fatal
 */
func Fatal(format string, v ...interface{}) {
	go func() {
		msg := fmt.Sprintf(format, v...)
		errors.New(msg)
	}()
	mLog.Fatalf(format, v...)
}

/**
 * Logs panic
 */
func Panic(format string, v ...interface{}) {
	mLog.Panicf(format, v...)
}

/**
 * Logs WorkerInitNewUser info
 */
func WorkerInitNewUser(format string, v ...interface{}) {
	mLog.Infof(constant.LogWorkerInitNewUserInfoPrefix+format, v...)
}

/**
 * Logs WorkerInitNewUser error
 */
func WorkerInitNewUserError(format string, v ...interface{}) {
	mLog.Infof(constant.LogWorkerInitNewUserErrorPrefix+format, v...)
}
