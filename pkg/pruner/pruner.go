package pruner

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cakra-labs/sprune/v7/internal/rootmulti"
	"github.com/cakra-labs/sprune/v7/pkg/app/config"
	"github.com/cakra-labs/sprune/v7/tools/logger"
	"github.com/cakra-labs/sprune/v7/types"
	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/state"
	cometbftstore "github.com/cometbft/cometbft/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/neilotoole/errgroup"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Pruner struct {
	cfg   *config.Config
	dbDir string
}

func NewPruner(cfg *config.Config) Pruner {
	dbDir := rootify(cfg.DataDir, cfg.ChainDir)

	return Pruner{
		cfg,
		dbDir,
	}
}

func (p Pruner) PruneAppState(ctx types.Context) error {
	o := opt.Options{
		DisableSeeksCompaction: true,
	}

	// Get BlockStore
	p.Logger(ctx).Debug("pruning application state", "dbDir", p.dbDir)
	appDB, err := db.NewGoLevelDBWithOpts("application", p.dbDir, &o)
	if err != nil {
		return err
	}

	// Load keys
	p.Logger(ctx).Debug("load keys")
	keys := loadKeys(p.cfg.Chain)

	appStore := rootmulti.NewStore(appDB, p.Logger(ctx))
	for _, value := range keys {
		appStore.MountStoreWithDB(value, storetypes.StoreTypeIAVL, nil)
	}

	err = appStore.LoadLatestVersion()
	if err != nil {
		return err
	}

	versions := appStore.GetAllVersions()
	v64 := make([]int64, len(versions))
	for i := 0; i < len(versions); i++ {
		v64[i] = int64(versions[i])
	}

	pruningHeights := v64[:len(v64)-int(p.cfg.BlocksToKeep)]

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		p.Logger(ctx).Info("pruning app store", "versionLenght", len(v64))
		if err := appStore.PruneStores(false, pruningHeights); err != nil {
			return err
		}

		p.Logger(ctx).Debug("compacting application state")
		if err := appDB.DB().CompactRange(util.Range{Start: nil, Limit: nil}); err != nil {
			return err
		}

		p.Logger(ctx).Info(fmt.Sprintf("pruned app stores: %d", len(pruningHeights)))

		return nil
	})

	return g.Wait()
}

func (p Pruner) PruneBlockState(ctx types.Context) error {
	o := opt.Options{
		DisableSeeksCompaction: true,
	}

	// Get BlockStore
	p.Logger(ctx).Debug("pruning block state", "dbDir", p.dbDir)
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
	pruneHeight := blockStore.Height() - int64(p.cfg.BlocksToKeep)

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		p.Logger(ctx).Info("pruning block store", "pruneHeight", pruneHeight)
		prunedBlocks, err := blockStore.PruneBlocks(pruneHeight)
		if err != nil {
			return err
		}

		p.Logger(ctx).Debug("compacting block store")
		if err := blockStoreDB.DB().CompactRange(util.Range{Start: nil, Limit: nil}); err != nil {
			return err
		}

		p.Logger(ctx).Info(fmt.Sprintf("pruned blocks: %d", prunedBlocks))

		return nil
	})

	g.Go(func() error {
		p.Logger(ctx).Info("pruning state store", "baseHeight", baseHeight, "pruneHeight", pruneHeight)
		if err := stateStore.PruneStates(baseHeight, pruneHeight); err != nil {
			return err
		}

		p.Logger(ctx).Debug("compacting state store")
		if err := stateDB.DB().CompactRange(util.Range{Start: nil, Limit: nil}); err != nil {
			return err
		}

		return nil
	})

	return g.Wait()
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
