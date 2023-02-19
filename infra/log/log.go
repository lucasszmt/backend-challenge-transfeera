package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
	Warn(msg string)
	Fatal(msg string, err error)
}

func PrettyLogger() logger {
	return logger{
		log.Output(zerolog.ConsoleWriter{Out: os.Stderr}),
	}
}

func NewLogger(level zerolog.Level) logger {
	l := zerolog.Logger{}
	l.WithLevel(level)
	return logger{
		l,
	}
}

type logger struct {
	logger zerolog.Logger
}

func (l *logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *logger) Error(msg string, err error) {
	l.logger.Error().Stack().Err(err).Msgf(msg)
}

func (l *logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *logger) Fatal(msg string, err error) {
	l.logger.Fatal().
		Err(err).
		Msg(msg)
}

type MockLogger struct {
}

func (m MockLogger) Info(msg string) {
	//TODO implement me
}

func (m MockLogger) Error(msg string, err error) {
	//TODO implement me
}

func (m MockLogger) Warn(msg string) {
	//TODO implement me
}

func (m MockLogger) Fatal(msg string, err error) {
	//TODO implement me
}
