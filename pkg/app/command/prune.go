package command

import (
	"github.com/neilotoole/errgroup"
	"github.com/spf13/cobra"
)

func pruneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prune [path_to_home]",
		Short: "prune data from the application store and block store",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			errs, _ := errgroup.WithContext(ctx)
			// var err error
			// if tendermint {
			// 	errs.Go(func() error {
			// 		if err = pruneTMData(args[0]); err != nil {
			// 			return err
			// 		}
			// 		return nil
			// 	})
			// }

			// if cosmosSdk {
			// 	err = pruneAppState(args[0])
			// 	if err != nil {
			// 		return err
			// 	}
			// 	return nil

			// }

			return errs.Wait()
		},
	}
	return cmd
}
