package pruner

import (
	"path/filepath"

	"github.com/cakra-labs/sprune/v7/pkg/app/config"
)

type Pruner struct {
	dbDir string
}

func NewPruner(cfg *config.Config) Pruner {
	dbDir := rootify(cfg.DataDir, cfg.HomePath)

	return Pruner{
		dbDir,
	}
}

func (p Pruner) PruneAppState() {

}

func (p Pruner) PruneBlockState() {

}

func rootify(path, root string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(root, path)
}
