package main

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	//zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().AnErr("error", errors.New("error message")).Uint64("teste_key", 12).Msg("tesste")
	log.Trace().Msg("trace message")
	log.Debug().Msg("Testing info")
	log.Error().Stack().Err(errors.New("error message")).Msg("")
}
