package main

import (
	"crypto/rand"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"nrvn.cc/go/lol"
)

const (
	defaultTimeout  = 1 * time.Second
	signalChanSize  = 2
	primalityRounds = 1 << 11
)

func main() {
	var (
		signalChan = make(chan os.Signal, signalChanSize)
		done       = make(chan bool, 1)
		debug      = flag.Bool("debug", false, "sets log level to debug")
		ticker     = time.Tick(defaultTimeout)
	)

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		done <- true
	}()

	for {
		select {
		case <-done:
			log.Info().Msg("exiting...go")
			return
		case <-ticker:
			randInt, isPrime, err := lol.Woot(rand.Reader)
			if err != nil {
				log.Fatal().Err(err)
			}

			if isPrime {
				log.Info().Str("int", randInt.String()).Bool("probably prime", isPrime).Send()
			} else {
				log.Debug().Str("int", randInt.String()).Bool("probably prime", isPrime).Send()
			}
		}
	}
}
