package logger

import (
	"fmt"
	"os"

	"cosmossdk.io/log"
	"github.com/cakra-labs/sprune/v7/pkg/app/config"
	"github.com/rs/zerolog"

	cometbftlog "github.com/cometbft/cometbft/libs/log"
	serverlog "github.com/cosmos/cosmos-sdk/server/log"
)

type Logger = cometbftlog.Logger

func NewLogger(cfg *config.Config) (Logger, error) {
	var opts []log.Option
	if cfg.LogLevel != "" {
		logLvl, err := zerolog.ParseLevel(cfg.LogLevel)
		switch {
		case err != nil:
			// If the log level is not a valid zerolog level, then we try to parse it as a key filter.
			filterFunc, err := log.ParseLogLevel(cfg.LogLevel)
			if err != nil {
				return nil, fmt.Errorf("failed to parse log level (%s): %w", cfg.LogLevel, err)
			}

			opts = append(opts, log.FilterOption(filterFunc))
		default:
			opts = append(opts, log.LevelOption(logLvl))
		}
	}

	l := serverlog.CometLoggerWrapper{
		Logger: log.NewLogger(cometbftlog.NewSyncWriter(os.Stdout), opts...).With(log.ModuleKey, cfg.AppName),
	}

	return l, nil
}
