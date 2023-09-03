package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	// multi := zerolog.MultiLevelWriter(consoleWriter, os.Stderr)
	// logger = zerolog.New(multi).With().Timestamp().Logger()
	logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
}

func Info(s string) {
	logger.Info().Msg(s)
}
func Debug(s string) {
	logger.Debug().Msg(s)
}
func Warn(s string) {
	logger.Warn().Msg(s)
}
func Trace(s string) {
	logger.Trace().Msg(s)
}
func Error(s string) {
	logger.Error().Msg(s)
}
func Fatal(err error) {
	logger.Fatal().Err(err).Msg("")
}

func Infof(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	logger.Info().Msg(s)
}
func Debugf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	logger.Debug().Msg(s)
}
func Warnf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	logger.Warn().Msg(s)
}
func Tracef(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	logger.Trace().Msg(s)
}
func Errorf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	logger.Error().Msg(s)
}

func Fatalf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	logger.Fatal().Msg(s)
}
func Printf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	logger.Info().Msg(s)
}
