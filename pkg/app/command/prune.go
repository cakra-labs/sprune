package command

import (
	"github.com/cakra-labs/sprune/v7/pkg/pruner"
	"github.com/cakra-labs/sprune/v7/tools/logger"
	"github.com/cakra-labs/sprune/v7/types"
	"github.com/spf13/cobra"
)

func pruneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prune",
		Short: "Prune data from the block store and app store",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger, err := logger.NewLogger(appConfig)
			if err != nil {
				return err
			}

			ctx := types.NewContext(
				logger,
			)

			p := pruner.NewPruner(appConfig)

			if appConfig.BlockState {
				if err := p.PruneBlockState(ctx); err != nil {
					return err
				}
			}

			if appConfig.AppState {
				if err := p.PruneAppState(ctx); err != nil {
					return err
				}
			}

			return nil
		},
	}
	return cmd
}
