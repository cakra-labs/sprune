package pruner

import (
	"context"
	"path/filepath"

	"github.com/cakra-labs/sprune/v7/pkg/app/config"
	"github.com/cakra-labs/sprune/v7/tools/logger"
	"github.com/cakra-labs/sprune/v7/types"
	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/state"
	cometbftstore "github.com/cometbft/cometbft/store"
	"github.com/neilotoole/errgroup"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Pruner struct {
	cfg   *config.Config
	dbDir string
}

func NewPruner(cfg *config.Config) Pruner {
	dbDir := rootify(cfg.DataDir, cfg.HomePath)

	return Pruner{
		cfg,
		dbDir,
	}
}

func (p Pruner) PruneAppState() {

}

func (p Pruner) PruneBlockState(ctx types.Context) error {
	o := opt.Options{
		DisableSeeksCompaction: true,
	}

	// Get BlockStore
	blockStoreDB, err := db.NewGoLevelDBWithOpts("blockstore", p.dbDir, &o)
	if err != nil {
		return err
	}
	blockStore := cometbftstore.NewBlockStore(blockStoreDB)

	// Get StateStore
	stateDB, err := db.NewGoLevelDBWithOpts("state", p.dbDir, &o)
	if err != nil {
		return err
	}
	stateStore := state.NewStore(stateDB, state.StoreOptions{
		DiscardABCIResponses: false,
	})

	// Define height
	baseHeight := blockStore.Base()
	pruneHeight := blockStore.Height() - int64(p.cfg.Blocks)

	errs, _ := errgroup.WithContext(context.Background())
	errs.Go(func() error {
		p.Logger(ctx).Debug("pruning block store")
		p.cfg.Blocks, err = blockStore.PruneBlocks(pruneHeight)
		if err != nil {
			return err
		}

		p.Logger(ctx).Debug("compacting block store")
		if err := blockStoreDB.DB().CompactRange(util.Range{Start: nil, Limit: nil}); err != nil {
			return err
		}

		return nil
	})

	p.Logger(ctx).Debug("pruning state store")
	err = stateStore.PruneStates(baseHeight, pruneHeight)
	if err != nil {
		return err
	}

	p.Logger(ctx).Debug("compacting state store")
	if err := stateDB.DB().CompactRange(util.Range{Start: nil, Limit: nil}); err != nil {
		return err
	}

	return nil
}

func (p Pruner) Logger(ctx types.Context) logger.Logger {
	return ctx.Logger().With("pkg", "pruner")
}

func rootify(path, root string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(root, path)
}
